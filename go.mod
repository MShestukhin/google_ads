module gitlab.somin.ai/analytics/platform/services/google_ads

go 1.16

replace github.com/micro/go-micro/v2 => gitlab.somin.ai/analytics/go-micro/v2 v2.9.2-0.20210714093304-414789c34a58

replace github.com/gocraft/work v0.5.1 => gitlab.somin.ai/analytics/work v0.0.0-20210708121752-22146f643c6f

require (
	github.com/gocraft/work v0.5.1
	github.com/gofrs/uuid v4.0.0+incompatible // indirect
	github.com/golang/protobuf v1.5.2
	github.com/gomodule/redigo v2.0.0+incompatible
	github.com/imdario/mergo v0.3.12 // indirect
	github.com/micro/go-micro/v2 v2.9.1
	github.com/mitchellh/mapstructure v1.4.1
	github.com/pkg/errors v0.9.1
	github.com/urfave/cli/v2 v2.3.0
	gitlab.somin.ai/analytics/platform/pkg/app v0.0.0-20210825145003-68a5b5c80e57
	gitlab.somin.ai/analytics/platform/services/google_api v0.0.0-20210907025915-32f0198c61c5
	golang.org/x/crypto v0.0.0-20210322153248-0c34fe9e7dc2 // indirect
	google.golang.org/protobuf v1.26.0
	gorm.io/gorm v1.21.9
)
