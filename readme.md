
Web framework：Gin

Orm: gorm

data： sqlite
___
Requirement1:
by typing URL localhost:8080/users to get user information in protobuf format

achieved by function GetUserList

Requirement2:
by typing localhost:8080/users/2 to retrieve user information with id 2

achieved by function GetUserDetail

Security:
using Json web token


## ProtoBuff
use user.proto to generate go file

`protoc --proto_path=pkg --go_out=build/gen pkg/api/user.proto`

## Get jwt token
token can be taken on the website https://jwt.io/ by modifying secret.

secret: abc

# Run
go run main.go