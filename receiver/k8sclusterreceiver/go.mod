module github.com/open-telemetry/opentelemetry-collector-contrib/receiver/k8sclusterreceiver

go 1.14

require (
	github.com/census-instrumentation/opencensus-proto v0.2.1
	github.com/iancoleman/strcase v0.0.0-20171129010253-3de563c3dc08
	github.com/open-telemetry/opentelemetry-collector v0.3.1-0.20200427150635-ca4b8231de7c
	github.com/stretchr/testify v1.5.1
	go.uber.org/zap v1.14.1
	k8s.io/api v0.17.4
	k8s.io/apimachinery v0.17.4
	k8s.io/client-go v0.0.0-20190620085101-78d2af792bab
)
