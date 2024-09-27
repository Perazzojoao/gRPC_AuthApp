module mail-service

go 1.22.0

require (
	google.golang.org/grpc v1.67.0
	google.golang.org/protobuf v1.34.2
)

require gopkg.in/alexcesaro/quotedprintable.v3 v3.0.0-20150716171945-2caba252f4dc // indirect

require (
	github.com/joho/godotenv v1.5.1
	golang.org/x/net v0.28.0 // indirect
	golang.org/x/sys v0.24.0 // indirect
	golang.org/x/text v0.17.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240814211410-ddb44dafa142 // indirect
	gopkg.in/gomail.v2 v2.0.0-20160411212932-81ebce5c23df
)
