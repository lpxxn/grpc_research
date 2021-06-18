module github.com/lpxxn/grpc_research

go 1.15

require (
	github.com/Pallinder/go-randomdata v1.2.0
	github.com/coreos/etcd v3.3.25+incompatible
	github.com/coreos/go-semver v0.3.0 // indirect
	github.com/coreos/go-systemd v0.0.0-00010101000000-000000000000 // indirect
	github.com/coreos/pkg v0.0.0-20180928190104-399ea9e2e55f // indirect
	github.com/gogo/protobuf v1.3.1 // indirect
	github.com/golang/protobuf v1.4.3
	github.com/google/go-cmp v0.5.0 // indirect
	github.com/google/uuid v1.1.2 // indirect
	github.com/grpc-ecosystem/go-grpc-middleware v1.2.0
	github.com/jhump/protoreflect v1.8.2
	github.com/joho/godotenv v1.3.0
	github.com/kelseyhightower/envconfig v1.4.0 // indirect
	go.uber.org/zap v1.16.0 // indirect
	google.golang.org/grpc v1.27.0
	google.golang.org/protobuf v1.25.1-0.20200805231151-a709e31e5d12
)

replace github.com/coreos/go-systemd => github.com/coreos/go-systemd/v22 v22.0.0
