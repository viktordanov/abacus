package main

import "math/big"

type ResultAssignment struct {
	Values []*big.Float
}

type ResultTuple struct {
	Values []*big.Float
}

type ResultVariablesTuple struct {
	Variables []string
}
type ResultError string

func NewResultAssignment() ResultAssignment {
	return ResultAssignment{
		Values: make([]*big.Float, 0),
	}
}


func NewResultTuple() ResultTuple {
	return ResultTuple{
		Values: make([]*big.Float, 0),
	}
}

func NewResultVariablesTuple() ResultVariablesTuple {
	return ResultVariablesTuple{
		Variables: make([]string, 0),
	}
}