package error

import (
	"fmt"
	"time"
)

type MyError struct {
	When time.Time
	What string
	Why  string
}

// Implement the built-in error interface
// 这是一种隐式接口实现，不需要显式说明 MyError 实现了 error 接口，只要方法签名一致的就行。
func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s, %s",
		e.When, e.What, e.Why)
}

func DoSomeThing(para string) (string, error) {
	if para == "ok" {
		return para, nil
	}
	return "", &MyError{time.Now(), para, "should give 'ok'"}
}
