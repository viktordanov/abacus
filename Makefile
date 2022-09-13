BIN=abacus

all: compile run

test:
	go test -v -covermode=atomic -race ./...

compile:
	go build -o $(BIN) ./cmd/abacus-cli/main.go

recompile-grammar:
	antlr4 -Dlanguage=Go -o parser/ Abacus.g4 -visitor

run:
	./$(BIN)	
