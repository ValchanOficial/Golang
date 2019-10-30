# Executando Múltiplos Arquivos no Windows

> No caso do Windows o comando abaixo não funciona:

- go run *.go

> Quando for necessário executar múltiplos arquivos no Windows, a forma correta é:

- go run arquivo1.go arquivo2.go
- go run $(ls -1 *.go | grep -v _test.go)