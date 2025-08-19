// 用于演示如何对一个外部类型扩展方法
package extend

import (
	"back/demo/internal/service/demo/inheritance"
	"fmt"
)

// ------------------------------------------------------------------------
// 定义新的类型，注意不能用别名方式 type MyInt = int
type MyInt int

// 为 int 定义扩展方法
func (i MyInt) increment() int {
	return int(i) + 1
}

// ------------------------------------------------------------------------

// 定义新的类型，注意不能用别名方式 type MyAnimal = inheritance.Animal
type MyAnimal inheritance.Animal

// 为 Animal 定义扩展方法
func (animal MyAnimal) hello() string {
	return fmt.Sprintf("%s: hello", animal.Name)
}
