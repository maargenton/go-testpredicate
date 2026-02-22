package itertest

import (
	"iter"
	"testing"
)

// VerifySeqCanStopAfterN iterates over seq, breaks after n elements, recovers
// from any panic and reports the error to `t`. A well-behaved iter.Seq must not
// call yield() again after it has returned false.
func VerifySeqCanStopAfterN[T any](t testing.TB, n int, seq iter.Seq[T]) {
	t.Helper()
	if recovered := testSeq(n, seq); recovered != nil {
		t.Errorf("interrupting after %d iterations: %v", n, recovered)
	}
}

func testSeq[T any](n int, seq iter.Seq[T]) (recovered any) {
	defer func() {
		if r := recover(); r != nil {
			recovered = r
		}
	}()

	var i = 0
	for range seq {
		i++
		if i >= n {
			break
		}
	}
	return nil
}

// VerifySeq2CanStopAfterN iterates over seq, breaks after n elements, recovers
// from any panic and reports the error to `t`. A well-behaved iter.Seq2 must not
// call yield() again after it has returned false.
func VerifySeq2CanStopAfterN[K, V any](t testing.TB, n int, seq iter.Seq2[K, V]) {
	t.Helper()
	if recovered := testSeq2(n, seq); recovered != nil {
		t.Errorf("interrupting after %d iterations: %v", n, recovered)
	}
}

func testSeq2[K, V any](n int, seq iter.Seq2[K, V]) (recovered any) {
	defer func() {
		if r := recover(); r != nil {
			recovered = r
		}
	}()

	var i = 0
	for range seq {
		i++
		if i >= n {
			break
		}
	}
	return nil
}
