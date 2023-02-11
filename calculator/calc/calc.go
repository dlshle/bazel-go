package calc
import (
  "fmt"
  "github.com/dlshle/gommon/async"
)

func Calc(a, b int) int {
  return a + b
}

func CalaAsync(e async.Executor, a, b int) {
  e.Execute(func() {
    fmt.Println(a + b)
  })
}
