package main

import (
	"github.com/cockroachdb/apd"
)

type ResultAssignment struct {
	Values []*apd.Decimal
}

type ResultLambdaAssignment string

type ResultTuple struct {
	Values []*apd.Decimal
}

type ResultVariablesTuple struct {
	Variables []string
}
type ResultError string

func NewResultAssignment() ResultAssignment {
	return ResultAssignment{
		Values: make([]*apd.Decimal, 0),
	}
}

func NewResultTuple() ResultTuple {
	return ResultTuple{
		Values: make([]*apd.Decimal, 0),
	}
}

func NewResultVariablesTuple() ResultVariablesTuple {
	return ResultVariablesTuple{
		Variables: make([]string, 0),
	}
}
