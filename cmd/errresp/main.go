package main

import (
	"github.com/mongodb-forks/golinters/errresp"

	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(errresp.Analyzer)
}
