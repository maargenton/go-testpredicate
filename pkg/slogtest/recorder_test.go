//go:build go1.21

package slogtest_test

import (
	"log"
	"log/slog"
	"testing"
	"time"

	"github.com/maargenton/go-testpredicate/pkg/bdd"
	"github.com/maargenton/go-testpredicate/pkg/require"
	"github.com/maargenton/go-testpredicate/pkg/verify"

	"github.com/maargenton/go-testpredicate/pkg/slogtest"
)

func TestGetFlattenAttrs(t *testing.T) {
	bdd.Given(t, "a slog.Record", func(t *bdd.T) {
		var r = slog.NewRecord(time.Now(), slog.LevelInfo, "message", 0)
		r.Add("foo", 123, "bar", slog.GroupValue(slog.String("baz", "baz-value")))
		t.When("calling GetFlattenAttrs", func(t *bdd.T) {
			t.Then("group attributes are expanded to dot-separated keys", func(t *bdd.T) {
				var attrs = slogtest.GetFlattenAttrs(r)
				require.That(t, attrs).MapKeys().IsEqualSet([]string{
					"foo", "bar.baz",
				})
			})
			t.Then("values are fully resolved, preserving type", func(t *bdd.T) {
				var attrs = slogtest.GetFlattenAttrs(r)
				require.That(t, attrs["foo"]).Eq(123)
				require.That(t, attrs["bar.baz"]).Eq("baz-value")
			})
		})
	})
}

func TestGetAllAttrs(t *testing.T) {
	bdd.Given(t, "a slog.Record", func(t *bdd.T) {
		var r = slog.NewRecord(time.Now(), slog.LevelInfo, "message", 0)
		r.Add("foo", 123, "bar", slog.GroupValue(slog.String("baz", "baz-value")))
		t.When("calling GetAllAttrs", func(t *bdd.T) {
			t.Then("all attributes are stored in resulting map", func(t *bdd.T) {
				var attrs = slogtest.GetAllAttrs(r)
				require.That(t, attrs).MapKeys().IsEqualSet([]string{
					"foo", "bar",
				})
			})
			t.Then("gropup values are stored in nested maps, fully resolved, preserving type", func(t *bdd.T) {
				var attrs = slogtest.GetAllAttrs(r)
				require.That(t, attrs).Field("foo").Eq(123)
				require.That(t, attrs).Field("bar.baz").Eq("baz-value")
			})
		})
	})
}

func TestRecorder(t *testing.T) {
	bdd.Given(t, "a recorder", func(t *bdd.T) {
		var r = &slogtest.Recorder{}
		var l = slog.New(r)

		t.When("calling log.Info()", func(t *bdd.T) {
			l.Info("message", "key", "value")

			t.Then("log is recorded as info, with message and attribute", func(t *bdd.T) {
				var rr = r.GetRecords()
				require.That(t, rr).Length().Eq(1)

				var record = rr[0]
				verify.That(t, record.Level).Eq(slog.LevelInfo)
				verify.That(t, record.Message).Eq("message")

				var attrs = slogtest.GetFlattenAttrs(record)
				verify.That(t, attrs).Field("key").Eq("value")
			})
		})
		t.When("calling log.Info() indirectly from With()", func(t *bdd.T) {
			l.With("kk", "vv").Info("message", "key", "value")

			t.Then("log record include share attributes", func(t *bdd.T) {
				var rr = r.GetRecords()
				require.That(t, rr).Length().Eq(1)

				var attrs = slogtest.GetFlattenAttrs(rr[0])
				verify.That(t, attrs).MapKeys().IsEqualSet([]string{"key", "kk"})
				verify.That(t, attrs).Field("key").Eq("value")
				verify.That(t, attrs).Field("kk").Eq("vv")
			})
		})

		t.When("calling log.Info() indirectly from WithGroup()", func(t *bdd.T) {
			l.WithGroup("group").Info("message", "key", "value")

			t.Then("log record include attributes prefixed by group name", func(t *bdd.T) {
				var rr = r.GetRecords()
				require.That(t, rr).Length().Eq(1)

				var attrs = slogtest.GetFlattenAttrs(rr[0])
				verify.That(t, attrs).MapKeys().IsEqualSet([]string{"group.key"})
				verify.That(t, attrs["group.key"]).Eq("value")
			})
		})

		t.When("calling log.Info() indirectly from mix of With and WithGroup()", func(t *bdd.T) {
			l.With("kk", "vv").WithGroup("group").With("gkk", "gvv").Info("message", "key", "value")

			t.Then("log record include attributes prefixed by group name", func(t *bdd.T) {
				var rr = r.GetRecords()
				require.That(t, rr).Length().Eq(1)

				var attrs = slogtest.GetFlattenAttrs(rr[0])
				verify.That(t, attrs).MapKeys().IsEqualSet([]string{
					"group.key", "group.gkk", "kk",
				})
				verify.That(t, attrs["group.key"]).Eq("value")
				verify.That(t, attrs["group.gkk"]).Eq("gvv")
				verify.That(t, attrs["kk"]).Eq("vv")
			})
		})
	})
}

func TestWithDefaultRecorder(t *testing.T) {
	bdd.Given(t, "something", func(t *bdd.T) {
		log.Printf("before log test1")
		slog.Info("before slog test2")

		var r *slogtest.Recorder
		slogtest.WithSlogRecorder(func(logs *slogtest.Recorder) {
			log.Printf("captured log test3")
			slog.Info("captured slog test4")
			r = logs
		})
		log.Printf("after log test5")
		slog.Info("after sog test6")

		t.When("doing something", func(t *bdd.T) {
			t.Then("something happens", func(t *bdd.T) {
				require.That(t, r.GetRecords()).Length().Eq(2)
			})
		})
	})
}
