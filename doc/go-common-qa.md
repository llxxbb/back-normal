# java 人员学 go 的一般性问题

## 作用域

go 没有与 java 对等的 public, private, protected,friendly 关键字。要想跨文件使用 go 代码，名字的首字母必须大写。

## 包

一个目录下得所有文件的包名必须使用同一个包名，并保证包名和目录名相同。

引用本模块内部包时，需要以模块名作为前缀+包的完全路径。

引用本地其他模块时，可以在模块名前附加相对路径。

go mod tidy 会自动整理依赖关系，包括[间接依赖](https://blog.csdn.net/juzipidemimi/article/details/104441398)。

参考：[使用go module导入本地包 - 知乎 (zhihu.com)](https://zhuanlan.zhihu.com/p/109828249)

## 错误与异常处理

相对于 java 的异常，go 会将问题分为两种情况进行处理：**意料之中的问题**和**意料之中的问题**。

### 意料之中的问题

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

### 意料之中的问题

[Catch Panics in Go | Delft Stack](https://www.delftstack.com/howto/go/catch-panics-in-golang/)

panic 

打印到文件

recover

## 如何处理继承

## 指针

在使用指针时，需要注意：指针本身不为 nil 但指针的值可能为 nil，这经常会导致空指针 panic!。

## 关闭资源

## 依赖注入

## 隐式接口

## 没有枚举

## 深度复制

## json 处理

## 数据库编程

## 没有 synchronized

## 与 maven 对应的职能如何实现

- 依赖管理

- 打包静态资源

## 测试

单元测试需单独建文件，和要测试的 go文件目录相同，文件名格式为：[待测文件]_test.go

go 自身没有测试用的断言，建议引入依赖 github.com/stretchr/testify

参考[Go的测试框架 - 简书 (jianshu.com)](https://www.jianshu.com/p/fe2f21d4e46d)。包含单元测试、断言、mock，基准测试（性能测试）等

## 比较好的学习资源

官方：

- 交互式学习go: [A Tour of Go](https://go.dev/tour) [中文版](https://tour.go-zh.org)

- [Effective Go - The Go Programming Language (google.cn)](https://golang.google.cn/doc/effective_go)

- [用户手册：涵盖了语言，IDE，测试，数据库等方方面面](https://golang.google.cn/doc/)

民间：

- [Go 语言教程 | 菜鸟教程 (runoob.com)](https://www.runoob.com/go/go-tutorial.html)

- [Go语言入门教程，Golang入门教程（非常详细） (biancheng.net)](http://c.biancheng.net/golang/)
