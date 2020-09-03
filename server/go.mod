module cli_chat_tcp/server/main

go 1.15

replace cli_chat_tcp/server/pkg/server => ./pkg/server

require (
	cli_chat_tcp/server/pkg/server v0.0.0-00010101000000-000000000000
	cloud.google.com/go v0.65.0 // indirect
	github.com/chilts/sid v0.0.0-20190607042430-660e94789ec9 // indirect
	github.com/google/uuid v1.1.2 // indirect
	github.com/lithammer/shortuuid v3.0.0+incompatible // indirect
)
