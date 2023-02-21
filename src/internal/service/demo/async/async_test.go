package async

import (
	"fmt"
	"testing"
)

// 请已以 debug 方式运行，否则看不到日志输出
func TestAsyncDo(t *testing.T) {
	rtnC := AsyncDo()
	fmt.Println("The async function is called, is waiting process")
	rtnV := <-rtnC
	fmt.Printf("received: %d\n", rtnV)

	// async-await like
	another := <-AsyncDo()
	fmt.Printf("received another: %d\n", another)
}
