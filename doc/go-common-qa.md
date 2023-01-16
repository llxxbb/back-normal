# java 人员学 go 的一般性问题

这里仅仅是 [A Tour of Go 中文版](https://tour.go-zh.org) 的补充，请优先学期 Tour 中的内容

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

### 访问控制

java: 需要使用 public, private 等关键字，go: 首字母大写便是 public, 否则便是private

### 接口

**隐式接口**

java 需要显式实现接口，而 go 则是隐式的，没有类似于 implement 关键字。只要看起来像鸭子，游起来像鸭子，叫起来像鸭子，那它就是鸭子，不需要显式说明它是鸭子：参考[鸭子类型](https://morven.life/posts/golang-interface-and-composition/#:~:text=Go%20%E6%8E%A5%E5%8F%A3%E4%B8%8E%E7%BB%84%E5%90%88%201%20Go%20%E6%8E%A5%E5%8F%A3%E4%B8%8E%E9%B8%AD%E5%AD%90%E7%B1%BB%E5%9E%8B%20%E4%BB%80%E4%B9%88%E6%98%AF%E2%80%9C%E9%B8%AD%E5%AD%90%E7%B1%BB%E5%9E%8B%E2%80%9D%EF%BC%9F%20...%202)。

如下面示例中 MyError 对 error 接口的实现为隐式接口实现。

示例： src/internal/service/error

**任意类型与类型判断**

在 go 中**空接口**(没有任何方法的接口)有一个广泛而特殊的用法，用于表示任何类型！**因为依据go的隐式接口实现规则，每个自定义类型都会默认实现空接口**。

示例： src/internal/service/anyType

参考：[反射 (google.cn)](https://golang.google.cn/blog/laws-of-reflection)

**接口不要过大**

As [Rob Pike points out](https://go-proverbs.github.io/), "The bigger the interface, the weaker the abstraction."

### 对象比较

字符串间可以直接用 ”==“ 进行比较，不用想 java 那样调用 equal 函数。

### 没有枚举

go 没有枚举定义，可用常量定义来模拟

### 深度复制

值类型的数据，默认全部都是深复制，Array、Int、String、Struct、Float，Bool。

引用类型的数据，默认全部都是浅复制，如指针，Slice，Map。

### 没有 synchronized

go 没有 synchronized， go 认为如果需要重入锁，那么代码是可以优化的。

参考：[Go sync.Mutex - 简书 (jianshu.com)](https://www.jianshu.com/p/9e5554617399)

## 语言高级特性问题

### 如何处理继承

go 没有类，也没有继承的概念，go 用组合来模拟继承。

组合的方式可以把一个类的成员变量声明为另一个类类型，这样就可以使用另一个类的方法和属性，从而实现类似继承的行为。

示例： src/internal/service/inheritance

**要点**：注意匿名字段的使用，可以实现类似于继承的直接调用方式。

### 为已有类型扩展方法

**要点**：方法的接收者和方法必须位于同一个包。解决方法为已有类型定义一个新类型，注意是新类型而不是别名。

示例： src/internal/service/extend 

### 关闭资源

Java 使用 try-with-resources 语句来自动回收资源，go 使用 delay 来释放资源，参考：[Go语言defer（延迟执行语句） (biancheng.net)](http://c.biancheng.net/view/61.html)

示例：src/internal/service/free

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

### 日志

go 内置的 log 库缺少级别和分割能力，本示例使用 [zap](https://github.com/uber-go/zap) + [lumberjack](https://github.com/natefinch/lumberjack)。[Frequently Asked Questions](https://github.com/uber-go/zap/blob/v1.24.0/FAQ.md), [go zap自定义日志输出格式](https://www.jianshu.com/p/fc90ea603ef2)

zap要点：

- 日志有两种输出方式：易用（可将对象直接输出为 json）的 zap.S() 和 高性能的 zap.L()

- 高性能的 logger 只支持结构化的日志输出

- 缺省输出会缓存，所以需要时常落盘： defer zap.L().Sync()

**时间格式化**：格式串必须是**go语言的诞生时间**，**01/02 03:04:05PM ‘06 -0700** ，我们常用的格式为： "2006-01-02 15:04:05.000"。 参考[Golang时间格式化 - 知乎 (zhihu.com)](https://zhuanlan.zhihu.com/p/145009400)

示例： src/tool/logger.go

**已知问题**：输出 json 时只输出 value 不输出 key。解决方法，自行序列化（性能不是很好）,因为 zap 之所以性能好，就是因为 It includes a reflection-free, zero-allocation JSON encoder.

### 配置、环境变量的读取

这里使用 [viper](https://github.com/spf13/viper) ，它提供了下面的功能（来源于 viper）：

- 设置默认值

- 从JSON，TOML，YAML，HCL，Envfile和Java属性properties配置文件读取

- 监控配置文件变更并重新加载（可选）

- 从环境变量读取

- 从远程配置系统（etcd 或 Consul）读取配置

- 从命令行标志读取

- 从缓冲区阅读

示例： src/config

### Web 服务

这里以 https://gin-gonic.com/ 为例进行演示

### 依赖注入

go 基本上不会有依赖注入问题 ，因为**Go 在设计上更倾向于明确的、显式的编程风格**

实际上是利用 go 的代码生成能力

### 注解

Java 支持注解， go 原生支持的不好，一般情况下不建议使用。理由是**Go 在设计上更倾向于明确的、显式的编程风格**。参考：[Go：我有注解，Java：不，你没有！ - 技术颜良 - 博客园 (cnblogs.com)](https://www.cnblogs.com/cheyunhua/p/15409847.html)

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
