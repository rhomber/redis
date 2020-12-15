module github.com/rhomber/redis/example

go 1.14

require (
	github.com/benbjohnson/clock v1.0.3 // indirect
	github.com/rhomber/redis/v8 v8.4.2
	github.com/opentracing/opentracing-go v1.1.1-0.20190913142402-a7454ce5950e // indirect
	go.opentelemetry.io/otel v0.15.0
	go.opentelemetry.io/otel/sdk v0.15.0
	google.golang.org/genproto v0.0.0-20191009194640-548a555dbc03 // indirect
)

replace github.com/rhomber/redis/v8 => ../../
