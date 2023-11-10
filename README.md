# Readme

本项目主要介绍 go 的**实施规范**以及部分语言特性的示例。

可以将本项目作为后端开发的模板。

[go 教程](http://gitlab.cdel.local/arch_job/training/blob/master/doc/tech-stack/go/go-tour.md)

[go 项目规范](http://gitlab.cdel.local/dev-specification/development/blob/master/doc/go-spec.md)

[服务开发规范](http://gitlab.cdel.local/dev-specification/development/tree/master)

## 运行前准备

**修改项目相关信息**：编辑 src/cmd/config_default.yml文件，修改`prj`相关属性。

**模块重命名**：替换 module 名，注意需要批量替换所有相关的 import

**设置环境变量**：PRJ_EVN，用于加载对应的配置文件，如PRJ_ENV=dev。有关配置请参考 [go-BaseConfig](https://github.com/llxxbb/go-BaseConfig)

**配置文件**：位于 src/cmd 目录下

**设置工作目录**：请指定到 src/cmd下，否则IDE环境下可能找不到配置文件

## 启动项目

启动文件位于 src/cmd/main.go

## 编译

> go build ./src/...

## 测试

> go test ./src/...

## 依赖包相关

**注意**：需要进入 src 目录。

更新依赖包

> go get -u ./...

整理依赖包

> go mod tidy

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

## Q & A

**Q：编译时遇到 “no required module provides package xxx; to add it:”**

 其中一个场景，在 src 目录下运行 go build ...，会出现此问题。解决方法：

- 编译时明确指定 main.go。 如在根目录下执行：go build ./src/cmd/main.go