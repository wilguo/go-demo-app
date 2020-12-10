module github.com/wilguo/go-security-test-app

go 1.15

replace (
	go.opentelemetry.io/contrib/propagators => /Users/wilbeguo/go/src/github.com/open-o11y/opentelemetry-go-contrib-PR/propagators
	go.opentelemetry.io/otel/sdk => /Users/wilbeguo/go/src/github.com/open-telemetry/opentelemetry-go/sdk
)

require (
	github.com/gorilla/mux v1.8.0
	go.opentelemetry.io/contrib/instrumentation/github.com/gorilla/mux/otelmux v0.14.0
	go.opentelemetry.io/contrib/propagators v0.14.0
	go.opentelemetry.io/otel v0.14.0
	go.opentelemetry.io/otel/exporters/otlp v0.14.0
	go.opentelemetry.io/otel/sdk v0.14.0
)
