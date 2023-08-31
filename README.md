# Readme

本项目主要介绍 go 的**实施规范**以及部分语言特性的示例。

可以将本项目作为后端开发的模板。

[go 教程](http://gitlab.cdel.local/arch_job/training/blob/master/doc/tech-stack/go/go-tour.md)

[go 项目规范](http://gitlab.cdel.local/dev-specification/development/blob/master/doc/go-spec.md)

[服务开发规范](http://gitlab.cdel.local/dev-specification/development/tree/master)

运行本项目的所有单元测试

> go test ./src/...

## 生成 mock 代码

使用 [golang/mock](https://github.com/golang/mock), 

安装

```shell
go install github.com/golang/mock/mockgen@v1.6.0
```

在 src 目录下运行

```shell
mockgen -source=".\internal\dao\tmp.go" -destination=".\internal\dao\tmp_mock.go" -package="dao"
```

**注意**：

- 这里生成的 mock 代码没有放到 _test.go 文件里，因为_test.go 中的定义不能跨包使用。这会增加不必要的生产代码，如本示例可执行文件增加了2k。

- 不能跨包使用 test 中的代码问题已有人反馈，但官方拒绝了此提议，原因为改动太大。变通的方法是：哪里使用生成到哪里。

## 使用说明

- 替换 module 名，注意需要批量替换所有相关的 import

- 启动和配置位于 cmd 目录下

- 配置，参考 [go-BaseConfig](https://github.com/llxxbb/go-BaseConfig)：
  
  - 对于不随部署环境变化的配置项放置于：config_default.yml 配置文件中
  
  - 环境配置文件为 config_[env].yml

## Q & A

**Q：编译时遇到 “no required module provides package xxx; to add it:”**

 其中一个场景，在 src 目录下运行 go build ...，会出现此问题。解决方法：

- 编译时明确指定 main.go。 如在根目录下执行：go build ./src/cmd/main.go