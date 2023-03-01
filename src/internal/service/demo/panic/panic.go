package panic

import (
	"fmt"
)

func OutOfIndex() (rtn string) {
	defer func() {
		if err := recover(); err != nil {
			msg := "panic occurred:" + fmt.Sprintf("%v", err)
			rtn = msg
		}
	}()
	arr := [3]int{1, 2, 3}
	idx := 3
	one := arr[idx]
	rtn = fmt.Sprintf("%d", one)
	return
}

func MyPanic() (rtn string) {
	// 仅用于证明 panic 被触发
	defer func() {
		if err := recover(); err != nil {
			rtn = fmt.Sprintf("%v", err)
		}
	}()
	rtn = "ok"
	panic("my panic")
}
