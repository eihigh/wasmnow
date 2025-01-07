//go:build js

package main

import (
	_ "embed"
	"fmt"
	"syscall/js"
)

//go:embed main.go
var mainGo []byte

func main() {
	fmt.Println("Hello, WebAssembly!")
	fmt.Println(string(mainGo))
	js.Global().Get("document").Get("body").Set("innerHTML", "Hello, WebAssembly!")
}
