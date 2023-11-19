# flamegraph-test

## How to create flamegraphs
```zsh
go build main.go
```

```zsh
/bin/bash ./launchandprofile.sh
```

To create a SVG flamegraph for the unoptimised code:
```zsh
./FlameGraph/flamegraph.pl output.prof >unoptimised-flamegraph.svg 
```

To create a SVG flamegraph for the optimised code:
```zsh
./FlameGraph/flamegraph.pl output-optimised.prof >optimised-flamegraph.svg 
```