### Common GOROOT Problem in order to fix

```bash

export GO111MODULE=on
#GOPATH MUST BE OUTSIDE OF GOROOT directory!!!
export PATH=$PATH:$GOPATH/bin
export GOROOT=/usr/local/go
export PATH=$PATH:$GOROOT/bin
source ~/.bashrc
```

### Testing sample endpoint (CURL Command)

You can use the following command: 

```bash
curl -i http://localhost:8000/sampleendpoint
```

### Install TableWriter Package

```bash 
go get github.com/olekukonko/tablewriter
```

### Run your test cases

```bash
go test -v
```

## Local Package Import

Use the following code in order to import local package in other modules

```bash 
go mod edit -replace example.com/goroutinesample=../goroutinesample
go mod tidy
go run .
```

## Formatting 

You can use external formatting tool as follows: 

```bash 
gofmt -w goroutinesample/goroutinesamp.go
```



