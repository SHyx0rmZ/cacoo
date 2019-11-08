# github.com/SHyx0rmZ/cacoo/pkg/cacoo
This package implements of parts of [Cacoo](https://nulab.com/products/cacoo/)'s [API](https://developer.nulab.com/docs/cacoo/#cacoo-api-overview). It can be used to retrieve diagrams' contents.

There are two commands included:
- `cmd/cacoo`, which just wraps some of the API calls
- `cmd/cacooviz`, which fetches a sample diagram and exports it as `cacoo.dot` (to be fed to Graphviz)

To build the commands, simply `cd` to either `cmd/cacoo` or `cmd/cacooviz` and run `go build`.
