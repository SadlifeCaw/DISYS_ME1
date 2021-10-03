# DISYS_ME1

works on mac, problems on windows

to generate gRPC, go into DISYS_ME1 directory and "protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative course.proto;"

To fix go gRPC code generation errors: 

  export GOROOT=/usr/local/go
  export GOPATH=$HOME/go
  export GOBIN=$GOPATH/bin
  export PATH=$PATH:$GOROOT:$GOPATH:$GOBIN

Taken from:
https://github.com/golang/protobuf/issues/795#issuecomment-571264398
