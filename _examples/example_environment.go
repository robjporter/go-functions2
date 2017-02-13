package main

import (
    "fmt"
    "../environment"
)

func main() {
    fmt.Println("COMPILED: ",environment.IsCompiled())
    fmt.Println("COMPILER: ",environment.Compiler())
    fmt.Println("Architecture: ",environment.GOARCH())
    fmt.Println("GO OS: ",environment.GOOS())
    fmt.Println("GO Root: ",environment.GOROOT())
    fmt.Println("GO Version: ",environment.GOVER())
    fmt.Println("GO Path: ",environment.GOPATH())
    fmt.Println("CPU: ",environment.NumCPU())
}