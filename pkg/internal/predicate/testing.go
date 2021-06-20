package predicate

// T is a minimal abstracted interface of testing.T for the purpose of //
// predicate construction and evaluation.
type T interface {
	Helper()
	Errorf(format string, args ...interface{})
	FailNow()
	Cleanup(f func())
}
