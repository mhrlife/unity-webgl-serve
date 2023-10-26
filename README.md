# Unity WebGL Serve
`unity-webgl-serve` is a simple command-line tool designed to serve Unity's compressed WebGL games. It helps developers in the development stage by setting up a server tailored for Unity WebGL content.
## Features
- Automatically sets appropriate `Content-Encoding` headers based on file extensions:
  - `.gz` files: Content-Encoding: `gzip`
  - `.br` files: Content-Encoding: `br`
- Sets the `Content-Type` header for:
  - `.wasm`, `.wasm.br`, and `.wasm.gz` files: `Content-Type: application/wasm`
## Installation
To install unity-webgl-serve, use the go install command:
```bash
 go install github.com/mhrlife/unity-webgl-serve@latest
```
Make sure your `$GOPATH/bin` directory is in your system's PATH.
```bash
export PATH="$HOME/go/bin:$PATH"
```
By default, this is `C:\Users\<YourUsername>\go\bin` on Windows or `$HOME/go/bin` on Linux/Mac.
## Usage
To use the tool, navigate to the directory where your Unity WebGL game files are located and run:
```bash
unity-webgl-serve --path=./path_to_your_webgl_files --port 8081
```
Replace `./path_to_your_webgl_files` with the path to your Unity WebGL game files. The server will start and use the provided path to serve the WebGL content.








