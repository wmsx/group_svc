module github.com/wmsx/group_svc

go 1.14

require (
	github.com/golang/protobuf v1.4.2
	github.com/micro/cli/v2 v2.1.2
	github.com/micro/go-micro/v2 v2.9.1
	github.com/wmsx/xconf v0.0.0-20200721142237-75926266fd1a
	google.golang.org/protobuf v1.25.0
	gorm.io/driver/mysql v1.0.1
	gorm.io/gorm v1.20.0
)

// 替换为v1.26.0版本的gRPC库
replace google.golang.org/grpc => google.golang.org/grpc v1.26.0
