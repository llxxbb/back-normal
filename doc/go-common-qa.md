# java 人员学 go 的一般性问题

## 作用域

go 没有与 java 对等的 public, private, protected,friendly 关键字。要想跨文件使用 go 代码，名字的首字母必须大写。

## 包

一个目录下得所有文件的包名必须使用同一个包名，并保证包名和目录名相同。

引用本模块内部包时，需要以模块名作为前缀+包的完全路径。

引用本地其他模块时，可以在模块名前附加相对路径。

参考：[使用go module导入本地包 - 知乎 (zhihu.com)](https://zhuanlan.zhihu.com/p/109828249)

## 错误与异常处理

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

## 比较好的学习资源

官方：

- 交互式学习go: [A Tour of Go](https://go.dev/tour) [中文版](https://tour.go-zh.org)

- [Effective Go - The Go Programming Language (google.cn)](https://golang.google.cn/doc/effective_go)

- [用户手册：涵盖了语言，IDE，测试，数据库等方方面面](https://golang.google.cn/doc/)

民间：

- [Go 语言教程 | 菜鸟教程 (runoob.com)](https://www.runoob.com/go/go-tutorial.html)

- [Go语言入门教程，Golang入门教程（非常详细） (biancheng.net)](http://c.biancheng.net/golang/)
