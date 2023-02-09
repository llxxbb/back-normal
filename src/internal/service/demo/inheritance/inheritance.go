// 用组合来模拟继承
package inheritance

import "fmt"

// Animal 组合 Mammal 对象来模拟继承
type Animal struct {
	Mammal
}

// 哺乳动物
type Mammal struct {
	Name string
	Age  int
}

func (m *Mammal) Greet() string {
	return fmt.Sprintf("Hello, I'm %s", m.Name)
}
