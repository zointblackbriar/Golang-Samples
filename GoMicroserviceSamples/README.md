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

