module github.com/satont/twir/apps/parser

go 1.23.2

toolchain go1.23.3

replace (
	github.com/satont/twir/libs/config => ../../libs/config
	github.com/satont/twir/libs/gomodels => ../../libs/gomodels
	github.com/satont/twir/libs/pubsub => ../../libs/pubsub
	github.com/satont/twir/libs/twitch => ../../libs/twitch
	github.com/satont/twir/libs/types => ../../libs/types
	github.com/twirapp/twir/libs/bus-core => ../../libs/bus-core
	github.com/twirapp/twir/libs/cache => ../../libs/cache
	github.com/twirapp/twir/libs/grpc => ../../libs/grpc
	github.com/twirapp/twir/libs/integrations => ../../libs/integrations
)

require (
	github.com/Masterminds/squirrel v1.5.4
	github.com/getsentry/sentry-go v0.28.1
	github.com/go-redsync/redsync/v4 v4.13.0
	github.com/goccy/go-json v0.10.3
	github.com/google/uuid v1.6.0
	github.com/guregu/null v4.0.0+incompatible
	github.com/hibiken/asynq v0.24.1
	github.com/imroc/req/v3 v3.48.0
	github.com/jackc/pgx/v5 v5.7.1
	github.com/jmoiron/sqlx v1.3.5
	github.com/lib/pq v1.10.9
	github.com/mazznoer/csscolorparser v0.1.3
	github.com/nats-io/nats.go v1.37.0
	github.com/nicklaw5/helix/v2 v2.30.0
	github.com/prometheus/client_golang v1.19.1
	github.com/redis/go-redis/v9 v9.6.1
	github.com/samber/lo v1.47.0
	github.com/satont/twir/libs/config v0.0.0-20241105005614-f68dbf11ade1
	github.com/satont/twir/libs/gomodels v0.0.0-20240225024146-742838c78cea
	github.com/satont/twir/libs/twitch v0.0.0-20231203205548-e635accc6b72
	github.com/satont/twir/libs/types v0.0.0-20241201233432-6c61559e1688
	github.com/satori/go.uuid v1.2.0
	github.com/shkh/lastfm-go v0.0.0-20191215035245-89a801c244e0
	github.com/tidwall/gjson v1.17.0
	github.com/twirapp/twir/libs/bus-core v0.0.0-20240225024146-742838c78cea
	github.com/twirapp/twir/libs/cache v0.0.0-00010101000000-000000000000
	github.com/twirapp/twir/libs/grpc v0.0.0-20241105005614-f68dbf11ade1
	github.com/twirapp/twir/libs/integrations v0.0.0-00010101000000-000000000000
	github.com/twirapp/twir/libs/repositories v0.0.0-20250210025356-9ebc030c5631
	github.com/twirapp/twir/libs/uptrace v0.0.0-20240225000519-a8b08004c468
	github.com/valyala/fasttemplate v1.2.2
	go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc v0.53.0
	go.uber.org/zap v1.27.0
	golang.org/x/exp v0.0.0-20241009180824-f66d83c29e7c
	google.golang.org/grpc v1.67.1
	google.golang.org/protobuf v1.35.1
	gorm.io/driver/postgres v1.5.9
	gorm.io/gorm v1.25.12
)

require (
	github.com/andybalholm/brotli v1.1.1 // indirect
	github.com/avito-tech/go-transaction-manager/drivers/pgxv5/v2 v2.0.0 // indirect
	github.com/avito-tech/go-transaction-manager/trm/v2 v2.0.0 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/cenkalti/backoff/v4 v4.3.0 // indirect
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/cloudflare/circl v1.5.0 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/felixge/httpsnoop v1.0.4 // indirect
	github.com/go-logr/logr v1.4.2 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/go-task/slim-sprig/v3 v3.0.0 // indirect
	github.com/golang-jwt/jwt/v4 v4.5.1 // indirect
	github.com/golang/protobuf v1.5.4 // indirect
	github.com/google/pprof v0.0.0-20241101162523-b92577c0c142 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.21.0 // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20240606120523-5a60cdf6a761 // indirect
	github.com/jackc/puddle/v2 v2.2.2 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/joho/godotenv v1.5.1 // indirect
	github.com/kelseyhightower/envconfig v1.4.0 // indirect
	github.com/klauspost/compress v1.17.11 // indirect
	github.com/lann/builder v0.0.0-20180802200727-47ae307949d0 // indirect
	github.com/lann/ps v0.0.0-20150810152359-62de8c46ede0 // indirect
	github.com/nats-io/nkeys v0.4.7 // indirect
	github.com/nats-io/nuid v1.0.1 // indirect
	github.com/onsi/ginkgo/v2 v2.21.0 // indirect
	github.com/prometheus/client_model v0.5.0 // indirect
	github.com/prometheus/common v0.48.0 // indirect
	github.com/prometheus/procfs v0.12.0 // indirect
	github.com/quic-go/qpack v0.5.1 // indirect
	github.com/quic-go/quic-go v0.48.1 // indirect
	github.com/refraction-networking/utls v1.6.7 // indirect
	github.com/robfig/cron/v3 v3.0.1 // indirect
	github.com/spf13/cast v1.4.1 // indirect
	github.com/tidwall/match v1.1.1 // indirect
	github.com/tidwall/pretty v1.2.1 // indirect
	github.com/uptrace/uptrace-go v1.27.1 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	go.opentelemetry.io/auto/sdk v1.1.0 // indirect
	go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.58.0 // indirect
	go.opentelemetry.io/contrib/instrumentation/runtime v0.53.0 // indirect
	go.opentelemetry.io/otel v1.33.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlplog/otlploghttp v0.4.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetrichttp v1.28.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlptrace v1.28.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp v1.28.0 // indirect
	go.opentelemetry.io/otel/exporters/stdout/stdouttrace v1.28.0 // indirect
	go.opentelemetry.io/otel/log v0.4.0 // indirect
	go.opentelemetry.io/otel/metric v1.33.0 // indirect
	go.opentelemetry.io/otel/sdk v1.28.0 // indirect
	go.opentelemetry.io/otel/sdk/log v0.4.0 // indirect
	go.opentelemetry.io/otel/sdk/metric v1.28.0 // indirect
	go.opentelemetry.io/otel/trace v1.33.0 // indirect
	go.opentelemetry.io/proto/otlp v1.3.1 // indirect
	go.uber.org/dig v1.18.0 // indirect
	go.uber.org/fx v1.22.2 // indirect
	go.uber.org/mock v0.5.0 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	golang.org/x/crypto v0.28.0 // indirect
	golang.org/x/mod v0.21.0 // indirect
	golang.org/x/net v0.30.0 // indirect
	golang.org/x/sync v0.8.0 // indirect
	golang.org/x/sys v0.26.0 // indirect
	golang.org/x/text v0.19.0 // indirect
	golang.org/x/time v0.5.0 // indirect
	golang.org/x/tools v0.26.0 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20240814211410-ddb44dafa142 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20241104194629-dd2ea8efbc28 // indirect
)
