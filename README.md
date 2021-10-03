# DISYS_ME1

works on mac, problems on windows

to generate gRPC, go into DISYS_ME1 directory and "protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative course.proto;"

To fix go gRPC code generation errors: https://www.codegrepper.com/search.php?q=protoc-gen-go:%20program%20not%20found%20or%20is%20not%20executable%20please%20specify%20a%20program%20using%20absolute%20path%20or%20make%20sure%20the%20program%20is%20available%20in%20your%20path%20system%20variable%20--go_out:%20protoc-gen-go:%20plugin%20failed%20with%20status%20code%201.
