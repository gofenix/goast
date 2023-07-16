package demo

import (
	"context"
	"fmt"
)

type Foo struct {
	i int
}

type Bar interface {
	Do(ctx context.Context) error
}

func demo() {
	a := sum(1, 2)
	fmt.Println(a)
}
