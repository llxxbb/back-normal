# java 人员学 go 的一般性问题

## 语言基础问题

### 工作区

工作区可以实现多个模块的管理，主要用于多个模块在开发和发布间的依赖问题。

**go官方规范**：工作区由 go.work 文件定义。可覆盖子目录中 go.mod 中的依赖。

参考：[Go 1.18 工作区模式最佳实践 - 知乎 (zhihu.com)](https://zhuanlan.zhihu.com/p/495832968)

### 模块

可作为独立的项目，或者作为`工作区`的子项目。

对依赖进行管理是其核心职能。

**go官方规范**：包由 go.mod 文件定义。

### 包

模块按照按照功能不同进行目录和文件的划分，这些目录便是`包`。

**go官方规范**：一个目录下得所有文件的包名必须**使用同一个包名**，并保证包名和目录名相同，这一点区别于 java 中的一个文件一个包。

引用本模块内部包时，需要以模块名作为前缀+包的完全路径。

引用本地其他模块时，可以在模块名前附加相对路径。

**go mod tidy 会自动整理依赖关系**，包括[间接依赖](https://blog.csdn.net/juzipidemimi/article/details/104441398)。**注意** ，go.mod 中的 // indirect 不要手工编辑， go mod tidy 会自动维护。

参考：[使用go module导入本地包 - 知乎 (zhihu.com)](https://zhuanlan.zhihu.com/p/109828249) 

### 如何处理继承

go 没有类，也没有继承的概念，go 用组合来模拟继承。

组合的方式可以把一个类的成员变量声明为另一个类类型，这样就可以使用另一个类的方法和属性，从而实现类似继承的行为。

示例： src/internal/service/inheritance

**要点**：注意匿名字段的使用，可以实现类似于继承的直接调用方式。

### 为已有类型扩展方法

**要点**：方法的接收者和方法必须位于同一个包。解决方法为已有类型定义一个新类型，注意是新类型而不是别名。

示例： src/internal/service/extend

### 接口

**隐式接口**

**任意类型与类型判断**

go 使用空接口来代表任意类型

```go
interface{} 
```

### 指针

在使用指针时，需要注意：指针本身不为 nil 但指针的值可能为 nil，这经常会导致空指针 panic!。

### 关闭资源

### 对象比较

字符串间可以直接用 ”==“ 进行比较，不用想 java 那样调用 equal 函数。

### 没有枚举

### 深度复制

### 没有 synchronized

### 错误与异常处理

相对于 java 的异常，go 会将问题分为两种情况进行处理：**意料之中的问题**和**意料之中的问题**。

#### 意料之中的问题

对于**意料之中的问题**，go 使用错误机制进行处理，如下面的 ReadFile 会返回两个参数，一个是文件内容，一个是错误。

```go
func main() {
    conent,err:=ioutil.ReadFile("filepath")
    if err !=nil{
        //错误处理
    }else {
        fmt.Println(string(conent))
    }
}
```

如果一个方法返回错误，那么它的**每一层上游调用都应该显式处理 err 对象**。

**建议**：

- 请务必给出明确的错误信息，请不要试图通过调用栈来定位哪里给出的这个错误信息。因为这是预料之中的问题，你应当清楚哪里出了错；因为它只是个字符串，不像 java 的异常有栈信息，给出你调用路径。

- 不建议引用第三方包来增强错误处理以达到类似 java 异常的处理效果，这样就是 java 样式的 go了，反而复杂化了；这样也违背了 go 的设计初衷：意料之中的问题。

- 如果方法在返回错误时信息不足，请上游将上下文信息通过入参传递进来，这样将极大简化错误信息的处理。

- 当然我们可以[自定义一个错误](https://go.dev/tour/methods/19)，来格式化这个错误字符串

示例： src/internal/service/error

#### 意料之外的问题

go 的 panic 函数与 java 的 threw 语句效果差不多，但 go 只有在不可预期的问题上才建议使用 panic, 对于可预期的错误请优先使用错误处理机制，因为错误处理机制会更轻量与简洁。

**要点**：

- panic 会立即退出当前函数

- 要想在 panic 时做一些补救措施，请在 defer 函数中调用 recover() 函数，注意recover() 函数只可以在 defer 函数内使用。

- 如果想在 defer 函数内修改返回值，请对当前函数的**返回值进行命名**，并在defer函数中进行修改。

- 打印调用栈：引入 "runtime/debug"，在代码中调用 debug.PrintStack()。

示例： src/internal/service/panic

## 框架问题

### Web 服务

### 依赖注入

### json 处理

### 数据库编程

### 读取配置文件

### 打包静态资源

### 测试

**go官方规范**：单元测试文件和要测试的 go文件目录相同，后缀为：_test.go

go 自身没有测试用的断言，建议引入依赖 github.com/stretchr/testify

执行项目中的所有单元测试

```shell
# 注意  go test -v ./... 在本项目中不工作
go test -v ./src/...
```

参考[Go的测试框架 - 简书 (jianshu.com)](https://www.jianshu.com/p/fe2f21d4e46d)。包含单元测试、断言、mock，基准测试（性能测试）等

## 比较好的学习资源

官方：

- 交互式学习go: [A Tour of Go](https://go.dev/tour) [中文版](https://tour.go-zh.org)

- [Effective Go - The Go Programming Language (google.cn)](https://golang.google.cn/doc/effective_go)

- [用户手册：涵盖了语言，IDE，测试，数据库等方方面面](https://golang.google.cn/doc/)

民间：

- [Go 语言教程 | 菜鸟教程 (runoob.com)](https://www.runoob.com/go/go-tutorial.html)

- [Go语言入门教程，Golang入门教程（非常详细） (biancheng.net)](http://c.biancheng.net/golang/)
