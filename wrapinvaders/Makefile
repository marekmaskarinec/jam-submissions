wasm:
	GOARCH=wasm GOOS=js go build -o main.wasm
	mv main.wasm builds/wasm
	mserver 8080 builds/wasm
