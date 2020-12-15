module github.com/rhomber/redis/extra/redisotel

go 1.15

replace github.com/rhomber/redis/v8 => ../..

replace github.com/rhomber/redis/extra/rediscmd => ../rediscmd

require (
	github.com/rhomber/redis/extra/rediscmd v0.2.0
	github.com/rhomber/redis/v8 v8.4.2
	go.opentelemetry.io/otel v0.15.0
)
