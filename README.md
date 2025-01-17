# Cat and Mouse

MVC

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
cd cat-mouse-go
GOOS=js GOARCH=wasm go build -o static/circles.wasm src/circles/main.go
```

## Run Server

```bash
cd src
go run main.go
```
