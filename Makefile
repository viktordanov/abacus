BIN = abacus

all: compile run

compile:
	go build -o $(BIN)

run:
	./$(BIN)	
