# Commit Gate

## Code
- [ ] Does the code do what I said it would do?
- [ ] Are there any obvious bugs or edge cases I haven't handled?
- [ ] Did I remove debug output, scratch comments, and TODO stubs?

## Documentation
- [ ] Does the README reflect any changes I made?
- [ ] Did I add a comment anywhere the WHY is non-obvious (a subtle invariant, a workaround, a locking discipline)?
- [ ] Is PHASES.md still accurate?

## Tests
- [ ] Does `go build ./...` succeed?
- [ ] Does `go test ./...` pass?
- [ ] Did I add a test for anything new and testable?
