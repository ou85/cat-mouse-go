# Cat and Mouse

## Structure

```plaintext
Documents/
└── cat-mouse-go
        ├── assets
        │   ├── index.html          # Main page
        │   ├── main.wasm           # WebAssembly file
        │   └── wasm_exec.js        # Go runtime for WebAssembly
        ├── server 
        │   └── main.go             # Web server Go
        ├── src
        │   └── main.go             # Source code
        ├── .gitignore
        ├── go.mod
        └── README.md
```

## Build WebAssembly

```bash
GOOS=js GOARCH=wasm go build -o assets/main.wasm src/main.go
```

## Run Server

```bash
cd src
go run main.go
```
