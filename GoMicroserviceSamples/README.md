### Common GOROOT Problem in order to fix

```bash

export GO111MODULE=on
#GOPATH MUST BE OUTSIDE OF GOROOT directory!!!
export PATH=$PATH:$GOPATH/bin
export GOROOT=/usr/local/go
export PATH=$PATH:$GOROOT/bin
source ~/.bashrc
```
