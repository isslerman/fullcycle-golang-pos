// To run:
go run main.go wire_gen.go 


DI - Dependency Injection

Ref.: 

- https://github.com/uber-go/fx
- https://github.com/google/wire

Tips: O Google Wire gera o código e não cria uma dependencia no projeto. 

Wire

    // Instala via go install
    // 1. Criamos o arquivo wire.go na raiz do projeto. Pode ser qq nome, mas por o comum usado é wire.go. 
    // 2. rodar wire
    // 3. go:generate para gerar novamente

