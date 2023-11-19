# flamegraph-test

## How to create flamegraphs
### Unoptimised code flamegraph

Uncooment the code after `// unoptimised.main.go` in `main.go` and comment the code after `// optimised.main.go` in `main.go`.

```zsh
go build main.go
```

```zsh
/bin/bash ./unoptim-launchandprofile.sh
```

To create a SVG flamegraph for the unoptimised code:
```zsh
./FlameGraph/flamegraph.pl unoptimised-output.prof >unoptimised-flamegraph.svg 
```


### Optimised code flamegraph

Uncooment the code after `// optimised.main.go` in `main.go` and comment the code after `// unoptimised.main.go` in `main.go`.

```zsh
go build main.go
```


To create a SVG flamegraph for the optimised code:To create a SVG flamegraph for the optimised code:
```zsh
./FlameGraph/flamegraph.pl optimised-output.prof >optimised-flamegraph.svg 
```