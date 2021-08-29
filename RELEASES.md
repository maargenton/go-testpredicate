# v0.6.2

## Changes

- Remove special case for error formatting. Error values in test context used to
  be handled as a special case and formatted with a simple `%v`; they are now
  formatted just like any other value.


# v0.6.1

## Bug fixes

- Fix issue with `.Contains(...)` failing to match sub-sequence at the very
  beginning of the value.

## Improvements

- Cleanup documentation
- Improve `rake info`, intermediate version number generation

# v0.6.0

## Main Features

- Add `.All()` and `.Any()` collection predicates with `subexpr.Value()`
  sub-predicate to operate om individual elements.
- Add `.Field(keypath)` transformation predicate to extract values from structs
  and maps.

## Improvements

- Update readme badges and fix comment typos
- Add / cleanup documentation comments
- Add release history
- Add workflow to create release on version tags (using rakefile task to precess
  release notes)

# v0.5.0

Significant API change, removing asserter object and using `verify.That()`,
`require.That()` and call chaining to define assertions. v0.4.x API is still
available but should considered deprecated and will be removed by v1.0.0.

## Main Features

- Implement new API based on `verify.That()` / `require.That()` and call
  chaining.
- Switch to go 1.14 minimal requirement, needed for `testing.T.Cleanup()`
- In prettyprint package, remove unused IndentWidth
- Move set operations to `value.Set` and `value.ReflectSet()`
- Move sequence / sub-sequence operations to `value.IndexOfSubsequence()`
- Use code generation and go:generate to collect all predicate functions defined
  in `predicate/impl` and expose them as methods on `builder.Builder`.

## Improvements

- Update documentation
- Remove Travis CI configuration
- Add github action to build and test packages, with coverage on codecov


# v0.4.4

Fix value formatting on assertion failure

# v0.4.3

Fix `value.CompareUnordered()` to not panic with nil interface

# v0.4.2

## Main Features

- Add missing call to `t.Helper()` within nested asserter functions
- Cleanup error messages referring to type issues with quoted type and no value
- Add `time.Time` support in `value.CompareOrdered()`
- Change message format for `p.IsNoError()` and `p.IsError()`

# v0.4.1

Documentation cleanup

# v0.4.0

Significant API change, now using `asserter.New()` to create asserter and
`p.IsTrue()` to refer to predicates.

## Main Features

- Move `predicate` to predicate package and `asserter` to asserter package
- Move all pre-defined predicates to package p
- Move `utils.ValueAs...()` and `utils.Compare...()` to a new value package
- Remove `utils.FormatValue()` and replace remaining uses with
  `prettyprint.FormatValue()`
- Asserter does not abort on failure by default, and details are interpreted as
  key-value pairs with formatted values.

# v0.3.1

## Main Features

- New format for values reporting with a go-like syntax

## Improvements

- Replace `utils.FormatValue()` with new pretty-printer
- Add pkg/prettyprint with new value formatter that display values in a go-like
  syntax, with multi-line formatting, line wrap for long strings, and ellipsis
  for less relevant details.
- Fix collection predicate to only look at PredicateResult to determine success

# v0.3.0

- Fix `IsNil()` and `IsNotNil()` predicates (were only testing for literal nil,
  not nil pointers and such)
- Add `IsTrue()` and `IsFalse()` predicates for convenience

# v0.2.0

## Main Features

- Add predicates to test panic cases
- Add better support for 1.13
- Change `pred.IsError(...)` to be included only with go1.13 or later, using new
  built-in `errors.Is()`, and remove dependency on 'golang.org/x/xerrors'
- Add go1.13 to CI build targets

# v0.1.4

## Main Features

- Change format value limit to 80 characters, with better adaptive formatting
- Add `utils.MaxAbsoluteDifference()` to extract the largest absolute difference
  between equally shaped collections of values
- Change `pred.CloseTo()` to support component-wise comparison of collection
  types

## Improvements

- Cleanup `Asserter` type definitions
- Cleanup `predicate` type definitions
- Update documentation with a summary of the built-in predicates

# v0.1.3

## Main Features

- Add support for `pred.IsError(nil)` for consistency
- Add missing documentation for `pred.IsError()`

# v0.1.2

## Main Features

- Add `pred.Eq()` and `pred.Ne()` alias predicates for IsEqualTo() and
  IsNotEqualTo()
- Add `pred.IsNoError()` and `pred.IsError()` with support for new
  `xerrors.Is()` error wrapping

## Improvements

- Add go 1.12 for Travis CI build
- Update references for renamed account
- Fix typos in documentation examples
- CI: Enable go 1.11 modules and remove go tip version -- too slow

# v0.1.1

Documentation improvements and CI setup

# v0.1.0

Initial public release


# v0.0.0

Overview

## Main Features

## Improvements

## Fixes
