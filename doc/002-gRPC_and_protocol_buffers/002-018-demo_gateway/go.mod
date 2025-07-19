module demo_gateway

go 1.24.0

require (
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.27.1
	google.golang.org/genproto/googleapis/api v0.0.0-20250707201910-8d1bb00bc6a7
	google.golang.org/grpc v1.73.0
	google.golang.org/protobuf v1.36.6
)

require (
	go.opentelemetry.io/otel v1.36.0 // indirect
	go.opentelemetry.io/otel/sdk/metric v1.36.0 // indirect
	golang.org/x/net v0.41.0 // indirect
	golang.org/x/sys v0.33.0 // indirect
	golang.org/x/text v0.26.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250715232539-7130f93afb79 // indirect
)

replace google.golang.org/genproto => google.golang.org/genproto v0.0.0-20250715232539-7130f93afb79
