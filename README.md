# Golang实现的文本序列比对

## 安装及运行
```shell
// 安装golang及配置
略
// 安装依赖
go get -u -v github.com/gin-gonic/gin
// 设置配置文件
cd path/to/connfig
mv basic.json/example basic.json
// 运行
cd path/to/main
go run main.go
```

## 请求
- header头：
Authorization: md5(password+当前时间戳) `当前时间戳精确到秒`
- request：
```json
{
    "str1":"nihako",
    "str2":"nihalod"
}
```
