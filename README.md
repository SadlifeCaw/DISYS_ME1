# DISYS_ME1

works on mac, problems on windows

to generate gRPC, go into DISYS_ME1 directory and "protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative course.proto;"

To fix go paths: https://www.codegrepper.com/code-examples/whatever/Please+specify+a+program+using+absolute+path+or+make+sure+the+program+is+available+in+your+PATH+system+variable+--go_out%3A+protoc-gen-go%3A+Plugin+failed+with+status+code+1.
