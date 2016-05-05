# rally-rest-toolkit
Rest Toolkit for Rally in Golang

### Purpose
To proivde an intuitive API in Golang for accessing Rally data through the REST API

### Usage

```sh
go get github.com/comcast/rally-rest-toolkit
```

### Dependencies
```
go get github.com/tools/godep
go get github.com/onsi/gomega
go get github.com/onsi/ginkgo
```

## References
[Rally Web Services Documentation](https://rally1.rallydev.com/slm/doc/webservice)

## Elements Supported
* Hierarchical Requirement
* Defect
* Changeset
* Build Definition
* Build
* Task

## Examples
Examples are located in the [examples folder](examples)

```sh
export API_KEY="insert your api key"
godep go run runme.go
```
