package main

import (
	"bazel-go/calculator/calc"
	"fmt"

	"github.com/dlshle/gommon/async"
)

func main() {
	calc.CalaAsync(async.DirectExecutor, 2, 3)
	fmt.Println(calc.Calc(1, 2))
	fmt.Printf("%v\n", calc.GetCalculatedMessage("test id", 6, 7))
}
