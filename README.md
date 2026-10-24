# shakedown-fangio

The Fangio demo repository for Those Who Run v2
(those-who-run/paddock#28): a repository that takes the `chassis`
template, to show that it passes CI and that an agent's pull request
waits for the maintainer. The Go content arrives with `copier copy` once
`chassis` has a release (those-who-run/paddock#21).
It is a demo, and is archived once the demo is recorded in paddock.

Issues are off in this repository. Every issue, the roadmap and the decisions
(ADRs) live in [those-who-run/paddock](https://github.com/those-who-run/paddock). A pull request here
names its paddock issue: `Closes those-who-run/paddock#N.` when it meets every
acceptance criterion, `Part of those-who-run/paddock#N` otherwise
([`AGENTS.md`](AGENTS.md)).

## Checks

`make check` runs what CI runs; [`AGENTS.md`](AGENTS.md) says what that is and
what a new clone runs first.
