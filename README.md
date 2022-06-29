# MGIS

矿产资源查询系统后端-go

## 初始化

> go mod init LGIS

## 使用gorm-gen

<https://github.com/go-gorm/gen>

基于 GORM, 更安全更友好的ORM工具。

### 概览

- 自动生成 CRUD 和 DIY 方法
- 自动根据表结构生成模型（model）代码
- 事务、嵌套事务、保存点、回滚事务点
- 完全兼容 GORM
- 更安全、更友好
- 多种生成代码模式

### 使用gentool生成dao代码

> gentool -dsn "xx:xx@tcp(xx:xx)/xx" -tables "kcdj,kqdz,kttz,mctz,kcmc,kcql,kccl,xksy,kcjs,kckc,kcjj" -outPath "./dao"  
> go mod tidy

### 在model里修改中文字段名的struct

将"数据"字段改为Shuju

```GO
type Kcdj struct {
	KCAAA  string `gorm:"column:KCAAA" json:"KCAAA"`
	KCC    string `gorm:"column:KCC" json:"KCC"`
	JJDAJ  string `gorm:"column:JJDAJ" json:"JJDAJ"`
	JJGLA  string `gorm:"column:JJGLA" json:"JJGLA"`
	DWAAC  string `gorm:"column:DWAAC" json:"DWAAC"`
	DWAAD  string `gorm:"column:DWAAD" json:"DWAAD"`
	KCBA   string `gorm:"column:KCBA" json:"KCBA"`
	KCCA   string `gorm:"column:KCCA" json:"KCCA"`
	KCCB   string `gorm:"column:KCCB" json:"KCCB"`
	PKGKB  string `gorm:"column:PKGKB" json:"PKGKB"`
	KCAOC  string `gorm:"column:KCAOC" json:"KCAOC"`
	KCAOAE string `gorm:"column:KCAOAE" json:"KCAOAE"`
	KCAOAF string `gorm:"column:KCAOAF" json:"KCAOAF"`
	PKD    string `gorm:"column:PKD" json:"PKD"`
	JJDCBF string `gorm:"column:JJDCBF" json:"JJDCBF"`
	Shuju  string `gorm:"column:数据" json:"数据"`
}
```

### 链接数据库

```go
ctx := context.Background()
db, _ := gorm.Open(mysql.Open("xxx:xxx@Jxn@tcp(xxx:xxx)/xxx?charset=utf8mb4&parseTime=True&loc=Local"))
```

## go-zero 编写api服务

### 安装goctl

>go install github.com/zeromicro/go-zero/tools/goctl@latest

#### 配置自动补全

> goctl completion bash > /etc/bash_completion.d/goctl

### 生成api架构代码

> goctl api new api

### 编写api文件

### 生成api代码

> goctl api go -api api.api -dir .

### 生成api文档

> goctl api doc -dir .

### 安装swagger插件生成api json

> go install github.com/zeromicro/goctl-swagger@latest

生成文档, 需要删除"requestBody": {},

> goctl api plugin -plugin goctl-swagger="swagger -filename MGIS.json" -api *.api -dir .
