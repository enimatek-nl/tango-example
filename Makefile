build:
	cp $(GOROOT)/misc/wasm/wasm_exec.js ./web/static/wasm_exec.js
	GOARCH=wasm GOOS=js go build -o ./web/static/web.wasm ./web.go

run: build
	go run ./server.go
