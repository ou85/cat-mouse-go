# Cat and Mouse

MVC

## Structure

```plaintext
Documents/
└── cat-mouse-go
        ├── server 
        │   └── main.go             
        ├── src
        │   ├── mouse                 # Cat and mouse
        │   │   ├── controller.go
        │   │   ├── entities.go
        │   │   ├── main.go             
        │   │   └── render.go        
        │   └── random                # Random symbols
        │       └── main.go          
        ├── static
        │   └── style.css        
        ├── go.mod
        ├── index.html                # Main page
        ├── mouse.wasm                # WebAssembly file
        ├── random.html          
        ├── README.md
        └── wasm_exec.js              # Go runtime for WebAssembly
```

## Build WebAssembly

```bash
cd cat-mouse-go
GOOS=js GOARCH=wasm go build -o mouse.wasm ./src/mouse
```

## Run Server

```bash
cd server
go run main.go
```
