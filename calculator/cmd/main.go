package main
import (
  "fmt"
  "bazel-go/calculator/calc"
  "github.com/dlshle/gommon/async"
)

func main() {
  calc.CalaAsync(async.DirectExecutor, 2, 3)
  fmt.Println(calc.Calc(1, 2))
}
