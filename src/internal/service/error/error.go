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
