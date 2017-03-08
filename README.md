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

## 接口
### diff(对比两个文章，先以段为单位，再以子为单位)
- header头：
Authorization: md5(password+当前时间戳) `当前时间戳精确到秒`
- request：
```json
{
    "str1":"nihako",
    "str2":"nihalod"
}
```
### diff/one（对比两文章片段，以字为单位）
同diff
