package main

import "fmt"

func main() {
	fmt.Printf("Outro programa em %s!") //%!s(MISSING)
}

// go version
// go help get
// godoc -http=:6060 -> localhost:6060 docs
// go env -> info
// go doc cmd/vet -> verifica código ex.: go vet comandos.go
// go build comandos.go -> gera executável
// ./comandos.exe -> executar
