//go:build go1.21

// Package slogtest provides support for recording structured log slog messages
// during tests and verifying that the expected messages have been emitted.
//
// `slogtest.Recorder` is a full-featured `slog.Handler` that records every
// structured log message and supports both shared attributes with `With()` and
// nested groups with `WithGroup()`.
//
// `slogtest.GetAllAttrs` and `slogtest.GetFlattenAttrs` are helper functions to
// retrieve all the attributes of a record as a map, with fully resolved values.
//
// `slogtest.WithSlogRecorder()` evaluate a function with a temporary
// `slogtest.Recorder` installed as the default logger of the slog packages.
package slogtest

import (
	"context"
	"log"
	"log/slog"
	"os"
	"sync"
)

// Recorder is an implementation of slog.Handler that records each log entry for
// later verification during unit tests.
type Recorder struct {
	records []slog.Record
	mu      sync.Mutex
}

var _ slog.Handler = (*Recorder)(nil)

func (h *Recorder) Enabled(context.Context, slog.Level) bool { return true }
func (h *Recorder) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &nestedRecorder{parent: h, attrs: attrs}
}

func (h *Recorder) WithGroup(name string) slog.Handler {
	return &nestedRecorder{parent: h, group: name}
}

func (h *Recorder) Handle(ctx context.Context, r slog.Record) error {
	h.mu.Lock()
	defer h.mu.Unlock()

	h.records = append(h.records, r)
	return nil
}

func (h *Recorder) GetRecords() []slog.Record {
	h.mu.Lock()
	defer h.mu.Unlock()
	return h.records
}

// GetFlattenAttrs returns all the attributes of a record as a single-level map,
// where keys of nested groups are expanded to a dot-separated syntax, and all
// the values are fully resolved.
func GetFlattenAttrs(r slog.Record) map[string]any {
	var attrs = make(map[string]any)
	var attrList []slog.Attr

	r.Attrs(func(attr slog.Attr) bool {
		attrList = append(attrList, attr)
		return true
	})
	flattenAttrs(attrs, "", attrList)
	return attrs
}

func flattenAttrs(attrs map[string]any, prefix string, group []slog.Attr) {
	for _, attr := range group {
		var key = prefix + attr.Key
		if attr.Value.Kind() == slog.KindGroup {
			flattenAttrs(attrs, key+".", attr.Value.Group())
		} else {
			attrs[key] = attr.Value.Resolve().Any()
		}
	}
}

// GetAllAttrs returns all the attributes of a record as a map, preserving the
// hierarchy of groups, with fully resolved values.
func GetAllAttrs(r slog.Record) map[string]any {
	var attrs = make(map[string]any)
	r.Attrs(func(attr slog.Attr) bool {
		attrs[attr.Key] = expandAttrValue(attr.Value)
		return true
	})
	return attrs
}

func expandAttrValue(v slog.Value) any {
	if v.Kind() == slog.KindGroup {
		var attrs = make(map[string]any)
		for _, attr := range v.Group() {
			attrs[attr.Key] = expandAttrValue(attr.Value)
		}
		return attrs
	} else {
		return v.Resolve().Any()
	}
}

// WithSlogRecorder creates and install a new log recorder as the global
// slog.Default(), and passes the recorder to the nested function. Upon return,
// it restores the prior slog.Default() and attempts to restore the state of the
// log package as well -- if a log.SetOutput() was installed after a call to
// slog.SetDefault(), the output will not be restored. Note that this function
// has global side-effects and may not behave as expected with concurrently
// executing tests.
func WithSlogRecorder(f func(logs *Recorder)) {
	var r = &Recorder{}
	var defaultLogger = slog.Default()
	var defaultLogFlags = log.Flags()

	slog.SetDefault(slog.New(r))
	defer func() {
		// Restore log output to os.Stderr before restoring prior slog; this is
		// intended to cover 2 cases:
		//
		// - the restored slog is the default logger in which case the log
		//   package is not updated but should be restored to operating
		//   independently and log to stderr by default.
		// - the restored slog is a custom logger, in which case slog
		//   reconfigure the log package logger to use the custom logger as
		//   output.
		log.SetFlags(defaultLogFlags)
		log.SetOutput(os.Stderr)
		slog.SetDefault(defaultLogger)
	}()

	f(r)
}

type nestedRecorder struct {
	parent slog.Handler
	attrs  []slog.Attr
	group  string
}

var _ slog.Handler = (*nestedRecorder)(nil)

func (h *nestedRecorder) Enabled(ctx context.Context, l slog.Level) bool {
	return h.parent.Enabled(ctx, l)
}

func (h *nestedRecorder) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &nestedRecorder{parent: h, attrs: attrs}
}

func (h *nestedRecorder) WithGroup(name string) slog.Handler {
	return &nestedRecorder{parent: h, group: name}
}

func (h *nestedRecorder) Handle(ctx context.Context, r slog.Record) error {
	if h.group != "" || len(h.attrs) > 0 {
		var rr = slog.NewRecord(r.Time, r.Level, r.Message, r.PC)
		if len(h.attrs) > 0 {
			rr.AddAttrs(h.attrs...)
			r.Attrs(func(a slog.Attr) bool {
				rr.AddAttrs(a)
				return true
			})
		} else if h.group != "" {
			var attrs []any
			r.Attrs(func(a slog.Attr) bool {
				attrs = append(attrs, a)
				return true
			})
			rr.AddAttrs(slog.Group(h.group, attrs...))
		}
		r = rr
	}
	return h.parent.Handle(ctx, r)
}
