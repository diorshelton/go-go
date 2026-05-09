# PR Gate

## Code Quality
- [ ] Is the diff scoped to one thing?
- [ ] Have I reviewed my own diff before opening?
- [ ] Is there anything I'd be embarrassed to explain in an interview?

## Documentation
- [ ] Is PHASES.md still accurate — does this move me closer to a "Done when"?
- [ ] Does the PR description explain what changed and why (not just what)?
- [ ] Is the architecture diagram in SPEC.md still accurate?

## Done Criteria
- [ ] Does this PR complete or advance exactly one phase checkpoint?
- [ ] Have I tested the happy path manually?
- [ ] Have I tested at least one edge case (malformed input, empty state, concurrent access)?
- [ ] Does `go vet ./...` pass cleanly?
