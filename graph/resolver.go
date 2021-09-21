package graph

//go:generate go run github.com/99designs/gqlgen

import (
	"github.com/zulfaza/gqlgen-todos/graph/model"
)

type Resolver struct{
	todos []*model.Todo
}

