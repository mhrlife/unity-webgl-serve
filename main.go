package main

import (
	"flag"
	"github.com/mhrlife/unity-webgl-serve/unitywebgl"
	"log"
	"net/http"
)

func main() {
	var path, port string
	flag.StringVar(&path, "path", ".", "Path to the Unity WebGL game files")
	flag.StringVar(&port, "port", "8080", "Path to the Unity WebGL game files")
	flag.Parse()

	http.Handle("/", unitywebgl.ServeUnityWebGL(path))

	log.Println("Server started on :" + port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Printf("error: %w\n", err)
	}
}
