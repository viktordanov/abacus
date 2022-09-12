BIN = abacus

all: compile run

compile:
	go build -o $(BIN)

recompile-grammar:
	antlr4 -Dlanguage=Go -o parser/ Abacus.g4 -visitor

run:
	./$(BIN)	
