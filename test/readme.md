## 作用

进行模板测试

## api生成
```bash
goctl api go -api desc/test.api -dir . --home ../ --style goZero
```

## model生成
```bash
goctl model mysql ddl --cache -src user.sql -dir model --home ../ --style goZero
```


## 生成swagger
```bash
# 安装
GOPROXY=https://goproxy.cn/,direct go install github.com/zeromicro/goctl-swagger@latest

# 生成
goctl api plugin -plugin goctl-swagger="swagger -filename user.json" -api desc/test.api -dir desc/

# 生成带地址的
goctl api plugin -plugin goctl-swagger="swagger -filename user.json -host 127.0.0.1:9999 -schemes http" -api desc/test.api -dir desc/

# 查看生成的文档
 docker run --rm -p 8083:8080 -e SWAGGER_JSON=/foo/user.json -v $PWD:/foo swaggerapi/swagger-ui
```

## MySQL容器创建

```bash
docker run -d \
  --name test-mysql-01 \
  -p 33069:3306  \
  -e MYSQL_ROOT_PASSWORD=Qwer@1234#$ \
  mysql:5.7 --character-set-server=utf8mb4 --collation-server=utf8mb4_unicode_ci --lower_case_table_names=1
```

## redis容器创建

```bash
docker run -d \
  --name redis-instance-01 \
  -p 36379:6379 \
  redis:7.0.11
```