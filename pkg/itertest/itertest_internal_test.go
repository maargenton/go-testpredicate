package itertest

import (
	"fmt"
	"iter"
	"testing"

	"github.com/maargenton/go-testpredicate/pkg/bdd"
	"github.com/maargenton/go-testpredicate/pkg/verify"
)

func Test_testSeq(t *testing.T) {
	bdd.Given(t, "a misbehaving iter.Seq", func(t *bdd.T) {

		var seq = func(n int) iter.Seq[int] {
			return func(yield func(int) bool) {
				for i := range n {
					yield(i)
				}
			}
		}

		t.When("calling testSeq breaking before the end of the sequence", func(t *bdd.T) {
			var recovered = testSeq(3, seq(100))

			t.Then("it should recover from panic and return the value", func(t *bdd.T) {
				verify.That(t, recovered).IsNotNil()
			})
		})

		t.When("calling testSeq breaking after the end of the sequence", func(t *bdd.T) {
			var recovered = testSeq(10, seq(10))

			t.Then("it should return no error", func(t *bdd.T) {
				verify.That(t, recovered).IsNil()
			})
		})

		t.When("calling VerifySeqCanStopAfterN breaking before the end of the sequence", func(t *bdd.T) {
			var tt = &mockT{TB: t}
			VerifySeqCanStopAfterN(tt, 10, seq(100))

			t.Then("Helper() was called", func(t *bdd.T) {
				verify.That(t, tt.helperCalls).Eq(1)
			})

			t.Then("Errorf() was called", func(t *bdd.T) {
				verify.That(t, tt.errors).Length().Eq(1)
			})
		})

		t.When("calling VerifySeqCanStopAfterN breaking after the end of the sequence", func(t *bdd.T) {
			t.Then("it should not report an error", func(t *bdd.T) {
				VerifySeqCanStopAfterN(t, 10, seq(10))
			})
		})
	})
}

func TestSeq2BreakAfterN(t *testing.T) {
	bdd.Given(t, "a misbehaving iter.Seq2", func(t *bdd.T) {

		var seq2 = func(n int) iter.Seq2[int, int] {
			return func(yield func(int, int) bool) {
				for i := range n {
					yield(i, i)
				}
			}
		}

		t.When("calling testSeq2 breaking before the end of the sequence", func(t *bdd.T) {
			var recovered = testSeq2(3, seq2(100))

			t.Then("it should capture the panic and return an error", func(t *bdd.T) {
				verify.That(t, recovered).IsNotNil()
			})
		})

		t.When("calling testSeq2 breaking after the end of the sequence", func(t *bdd.T) {
			var recovered = testSeq2(10, seq2(10))

			t.Then("it should return no error", func(t *bdd.T) {
				verify.That(t, recovered).IsNil()
			})
		})

		t.When("calling VerifySeq2CanStopAfterN breaking before the end of the sequence", func(t *bdd.T) {
			var tt = &mockT{TB: t}
			VerifySeq2CanStopAfterN(tt, 10, seq2(100))

			t.Then("Helper() was called", func(t *bdd.T) {
				verify.That(t, tt.helperCalls).Eq(1)
			})

			t.Then("Errorf() was called", func(t *bdd.T) {
				verify.That(t, tt.errors).Length().Eq(1)
			})
		})

		t.When("calling VerifySeq2CanStopAfterN breaking after the end of the sequence", func(t *bdd.T) {
			t.Then("it should not report an error", func(t *bdd.T) {
				VerifySeq2CanStopAfterN(t, 10, seq2(10))
			})
		})
	})
}

type mockT struct {
	testing.TB
	helperCalls int
	errors      []string
}

func (t *mockT) Helper() {
	t.helperCalls++
}

func (t *mockT) Errorf(format string, args ...any) {
	t.errors = append(t.errors, fmt.Sprintf(format, args...))
}
