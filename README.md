# listVersions
A third-party library for listing the latest versions of application releases. 

Specifically, it gives us the highest patch version of every release between a minimum version and the highest released version. It reads the Github Releases list, uses SemVer for comparison and takes a path to a file as its first argument when executed. It reads this file, which is in the format of:

```
repository,min_version
kubernetes/kubernetes,1.8.0
prometheus/prometheus,2.2.0
```

and it produces output to stdout in the format of:

```
latest versions of kubernetes/kubernetes: [1.13.4 1.12.6 1.11.8 1.10.13 1.9.11 1.8.15]
latest versions of prometheus/prometheus: [2.7.2 2.6.1 2.5.0 2.4.3 2.3.2 2.2.1]
```

## usage

1. run `go build` to build it, obtaining an executable `listVersions.exe`

2. run `listVersions <repo>/<release>,<min_version>`

for example:

`listVersions kubernetes/kubernetes,1.8.0`

## test

test cases are documented in `main_test.go`, which can be run by `go test`
