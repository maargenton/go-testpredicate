## VS Code snippets

In VSCode, open the command palette (⌘⇧P / ^⇧P) and select '_Preferences:
Configure User Snippets_', then '_go.json_' and add any or all of the following
snippets:
- `tfbdd`: inserts the definition of a test function with BDD-style given /when
  / then blocks, using the bifurcated execution context define in the `bdd`
  sub-package.
- `tftcs`: inserts the definition of a test function with data-driven test cases
  and BDD-style given /when / then blocks, using the bifurcated execution
  context define in the `bdd` sub-package.
- `require.That` or `reth`: inserts a placeholder for require.That() statement.
- `verify.That` or `veth`: inserts a placeholder for verify.That() statement.

```json
"test function bdd": {
    "prefix": "tfbdd",
    "body": [
        "func Test$1(t *testing.T) {",
        "\tbdd.Given(t, \"${2:something}\", func(t *bdd.T) {",
        "\t\tt.When(\"${3:doing something}\", func(t *bdd.T) {",
        "\t\t\tt.Then(\"${4:something happens}\", func(t *bdd.T) {",
        "\t\t\t\trequire.That(t, ${5:\"123\"}).Eq(${6:123})",
        "\t\t\t\t$0",
        "\t\t\t})",
        "\t\t})",
        "\t})",
        "}",
    ],
    "description": "Test function with BDD-style given/when/then and require.That"
},
"test function bdd with test cases": {
    "prefix": "tftcs",
    "body": [
        "func Test$1(t *testing.T) {",
        "\tvar tcs = []struct {",
        "\t\tname string",
        "\t}{",
        "\t\t{\"${2:test case}\"},",
        "\t}",
        "\t",
        "\tfor _, tc := range tcs {",
        "\t\tbdd.Given(t, tc.name, func(t *bdd.T) {",
        "\t\t\tt.When(\"${3:doing something}\", func(t *bdd.T) {",
        "\t\t\t\tt.Then(\"${4:something happens}\", func(t *bdd.T) {",
        "\t\t\t\t\trequire.That(t, ${5:\"123\"}).Eq(${6:123})",
        "\t\t\t\t\t$0",
        "\t\t\t\t})",
        "\t\t\t})",
        "\t\t})",
        "\t}",
        "}",
    ],
    "description": "Test function with BDD-style given/when/then, test cases and require.That"
},
"verify": {
    "prefix": "verify.That",
    "body": [
        "verify.That(t, ${0:\"\"}).Eq(\"\")",
    ],
    "description": "verify.That()"
},
"require": {
    "prefix": "require.That",
    "body": [
        "require.That(t, ${0:\"\"}).Eq(\"\")",
    ],
    "description": "require.That()"
},

"given": {
    "prefix": "given",
    "body": [
        "bdd.Given(t, \"$1\", func(t *bdd.T) {$0})",
    ],
    "description": "t.When(...)"
},
"when": {
    "prefix": "when",
    "body": [
        "t.When(\"$1\", func(t *bdd.T) {$0})",
    ],
    "description": "t.When(...)"
},
"with": {
    "prefix": "with",
    "body": [
        "t.With(\"$1\", func(t *bdd.T) {$0})",
    ],
    "description": "t.With(...)"
},
"then": {
    "prefix": "then",
    "body": [
        "t.Then(\"$1\", func(t *bdd.T) {$0})",
    ],
    "description": "t.Then(...)"
},

```
