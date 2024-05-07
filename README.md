
## 关于本项目

使用go-zero的template，对其进行了便于自己使用的优化，项目基于1.6.4进行构建
https://go-zero.dev/docs/tutorials/cli/template

添加的库
- "github.com/Masterminds/squirrel"  用于构建sql语句
- "github.com/pkg/errors" 错误封装


## 如何使用

```bash
# 远程使用
goctl api --remote "github.com/cplinux98/go-zero-template"
# 本地使用
git clone "github.com/cplinux98/go-zero-template" template
goctl api --home template
goctl api go -api desc/test.api -dir . --home template --style goZero
```

## 如何构建自己的

```bash
goctl template init --home "你的路径"

```

## 如何进行测试是否正常运行

```bash
# 构建api
goctl api new . --style goZero --remote "github.com/cplinux98/go-zero-template"

# 构建MySQL的model（带缓存）
goctl model mysql ddl --cache -src user.sql -dir model --home ../ --style goZero

```


## 模板适配自己项目

模板里面有些关于测试项目的包，需要修改成自己的包
- api/handler.tpl
  - test -> 你自己的