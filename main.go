package main

import (
	"context"

	"github.com/ridge/must"
)

type X struct {
	c context.Context
}

func main() {
	must.OK(nil)
}
