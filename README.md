
## 关于本项目

使用go-zero的template，对其进行了便于自己使用的优化，项目基于1.6.4进行构建
https://go-zero.dev/docs/tutorials/cli/template

## 如何使用

```bash
goctl api --remote "github.com/cplinux98/go-zero-template"
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
