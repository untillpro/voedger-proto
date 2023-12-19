# voedger-proto

- Experimental repository for the voedger product to evaluate some DevOps ideas
- Repository structure is designed after 
  - https://github.com/kubernetes/kubernetes
  - https://github.com/golang-standards/project-layout

## Motivation

- [A&D: Repository structure?](https://dev.heeus.io/launchpad/#!26427)

## Files which every package can have

- errors.go, consts.go
- utils.go // Helpers and potential candidates for goutils module
- [lowered_type1]_string.go, [lowered_type2]_string.go // stringers output
- internal // folder

## Interface package

- Example: [pkg/ibus](pkg/ibus)
- interface.go
- types.go // Public types with methods
  - types_events.go // If there are many types
  - types_schemes.go

## Interface Implementation package
- Example: [pkg/ibusmem](pkg/ibusmem)
  - provide.go
  - New() function
- impl.go
  - impl_types.go // if needed
  - impl_errors.go // if needed
  - impl_myReceiver1.go
  - impl_mySuperReceiver1.go

## Interface and Implementation files

- Interface files + Interface Implementation files

### Just a library

- Example: [cobrau](staging/src/github.com/untillpro/goutils/cobrau)
- Interface and Implementation files // if library provides interface and implementation
- `<package-name>.go` // if library is simple
- `<logical_subpackage1>.go`, `<logical_subpackage2>.go`... // if library is complex


## Package Structure: CLI Tool

- ref [dummytool](cmd/dummytool)
- [main.go](cmd/dummytool/main.go)
  - `execRootCmd()`
    - Use `cobrau.PrepareRootCmd()`
    - Return:
      - `rootCmd.Execute()`
      - or: `cobrau.ExecCommandAndCatchInterrupt(rootCmd)`
- [gorun.sh](cmd/dummytool/gorun.sh) - a helper to run the main func
- <command1>.go, <command2>.go...
- deploy_test.go
  -  `testingu.RunRootTestCases(t, execRootCmd, testCases)`
- Use [internal packages](cmd/dummytool/internal)
