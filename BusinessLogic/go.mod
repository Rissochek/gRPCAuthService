module Business

go 1.23.3

require (
	AuthProject v0.0.0
	google.golang.org/grpc v1.69.2
)

require (
	github.com/golang-jwt/jwt v3.2.2+incompatible // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.25.1 // indirect
	github.com/joho/godotenv v1.5.1 // indirect
	golang.org/x/net v0.33.0 // indirect
	golang.org/x/sys v0.28.0 // indirect
	golang.org/x/text v0.21.0 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20241219192143-6b3ec007d9bb // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20241219192143-6b3ec007d9bb // indirect
	google.golang.org/protobuf v1.36.1 // indirect
)

replace AuthProject => ../GoGRPCAuthService
