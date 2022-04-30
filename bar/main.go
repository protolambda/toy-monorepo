package main

import (
	"github.com/protolambda/toy-monorepo/foo"
)

func main() {
	logger := foo.BuildLogger()
	logger.Info("hello")
}
