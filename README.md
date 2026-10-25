# nyow
a dead-simple time tracking CLI

## installation
To install *nyow*, download a binary from [GitHub releases](https://github.com/mifumofu/nyow/releases) and place it somewhere within your $PATH, or execute it directly.

Alternatively, you can install *nyow* with `go install github.com/mifumofu/nyow/cli`, this will install a binary called `cli` to the bin/ directory in your $GOPATH. You should probably rename the binary to `nyow`.

## how to use
*nyow* has a very simplistic and frictionless command line interface. To start a session, you can use `nyow start [CATEGORY]`. To end your session you can use `nyow stop [CATEGORY] [DESCRIPTION]`, a description is a short overview of what you did in this session. *nyow* is designed to allow for multiple trackers to be active at the same time, the only limit being a seperate category per tracker. If you accidentially start a tracker you didn't mean to, you can use `nyow cancel [CATEGORY]` to drop the entry.

When recording a session you may of forgotten to track, you can use the `add` subcommand to manually create an entry. `remove` lets you clean up entries with an ID, which can be obtained from the `list` subcommand. `edit` lets you correct information in the future.

You can get an overview of your recorded sessions via `nyow overview`, which displays a grid - somewhat like GitHub's activity graph - marking how your time was distributed over the past week. To get a raw list, you can use `nyow list` and filter entries out with its flags, like `--category` and `--after`. Commands that show data expose a `--json` flag, intended to be used for scripting when parsed with another tool like `jq` or `python`.
