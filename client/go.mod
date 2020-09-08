module cli_chat_tcp/client/main

go 1.15

replace cli_chat_tcp/client/pkg/chat => ./pkg/chat

replace cli_chat_tcp/client/pkg/render => ./pkg/render

replace cli_chat_tcp/client/pkg/response => ./pkg/response

require (
	cli_chat_tcp/client/pkg/chat v0.0.0-00010101000000-000000000000
	cli_chat_tcp/client/pkg/render v0.0.0-00010101000000-000000000000 // indirect
	cli_chat_tcp/client/pkg/response v0.0.0-00010101000000-000000000000 // indirect
	github.com/inancgumus/screen v0.0.0-20190314163918-06e984b86ed3
	golang.org/x/crypto v0.0.0-20200820211705-5c72a883971a // indirect
)
