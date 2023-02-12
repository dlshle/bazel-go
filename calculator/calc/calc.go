package calc

import (
	"bazel-go/protos/test"
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

func GetCalculatedMessage(id string, a, b int) test.Event {
	return test.Event{
		Id:      id,
		Payload: fmt.Sprintf("%d", a+b),
	}
}
