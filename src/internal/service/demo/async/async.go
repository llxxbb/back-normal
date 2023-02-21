package async

import (
	"fmt"
	"time"
)

// 异步方法需要返回一个 chan 外部用于接收参数
func AsyncDo() chan int {

	// 4步走

	// 1、创建通道
	r := make(chan int)
	// 2、创建协程
	go func() {
		fmt.Println("processing ...")
		time.Sleep(3 * time.Second)
		fmt.Println("result is ok")
		// 3、将处理结果放入通道
		r <- 1
	}()
	// 4、把通道返回给调用者
	return r

}
