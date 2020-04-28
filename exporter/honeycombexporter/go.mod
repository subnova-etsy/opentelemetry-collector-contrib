module github.com/open-telemetry/opentelemetry-collector-contrib/exporter/honeycombexporter

go 1.14

require (
	github.com/census-instrumentation/opencensus-proto v0.2.1
	github.com/golang/protobuf v1.3.5
	github.com/google/go-cmp v0.4.0
	github.com/honeycombio/opentelemetry-exporter-go v0.3.1
	github.com/klauspost/compress v1.10.3
	github.com/open-telemetry/opentelemetry-collector v0.3.1-0.20200427150635-ca4b8231de7c
	github.com/stretchr/testify v1.5.1
	go.opentelemetry.io/otel v0.4.2
	go.uber.org/zap v1.14.0
)

replace github.com/honeycombio/opentelemetry-exporter-go => github.com/subnova-etsy/opentelemetry-exporter-go v0.99.2
