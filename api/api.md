### 1. "获取矿产地基本情况"

1. route definition

- Url: /api/K
- Method: GET
- Request: `Request`
- Response: `Response`

2. request definition



```golang
type Request struct {
	Province string `form:"province,optional"`
	Mine string `form:"mine,optional"`
	Page int `form:"page"`
	Size int `form:"size"`
}
```


3. response definition



```golang
type Response struct {
	Total int `json:"total"`
	Kcdjs []Kcdj `json:"kcdjs"` // 矿产地基本情况
}
```

### 2. "获取矿区地质情况"

1. route definition

- Url: /api/Kqdz
- Method: GET
- Request: `Request`
- Response: `KqdzRsp`

2. request definition



```golang
type Request struct {
	Province string `form:"province,optional"`
	Mine string `form:"mine,optional"`
	Page int `form:"page"`
	Size int `form:"size"`
}
```


3. response definition



```golang
type KqdzRsp struct {
	Total int `json:"total"`
	Kqdzs []Kqdz `json:"kqdzs"` // 矿区地质情况
}
```

### 3. "获取矿产储量"

1. route definition

- Url: /api/Kccl
- Method: GET
- Request: `Request`
- Response: `KcclRsp`

2. request definition



```golang
type Request struct {
	Province string `form:"province,optional"`
	Mine string `form:"mine,optional"`
	Page int `form:"page"`
	Size int `form:"size"`
}
```


3. response definition



```golang
type KcclRsp struct {
	Total int `json:"total"`
	Kccls []Kccl `json:"kccls"` // 矿产储量
}
```

### 4. "获取矿床技术经济评价"

1. route definition

- Url: /api/Kcjj
- Method: GET
- Request: `Request`
- Response: `KcjjRsp`

2. request definition



```golang
type Request struct {
	Province string `form:"province,optional"`
	Mine string `form:"mine,optional"`
	Page int `form:"page"`
	Size int `form:"size"`
}
```


3. response definition



```golang
type KcjjRsp struct {
	Total int `json:"total"`
	Kcjjs []Kcjj `json:"kcjjs"` // 矿床技术经济评价
}
```

### 5. "获取开采技术条件"

1. route definition

- Url: /api/Kcjs
- Method: GET
- Request: `Request`
- Response: `KcjsRsp`

2. request definition



```golang
type Request struct {
	Province string `form:"province,optional"`
	Mine string `form:"mine,optional"`
	Page int `form:"page"`
	Size int `form:"size"`
}
```


3. response definition



```golang
type KcjsRsp struct {
	Total int `json:"total"`
	Kcjs []Kcj `json:"kcjs"` // 开采技术条件
}
```

### 6. "获取矿产勘查工作概况"

1. route definition

- Url: /api/Kckc
- Method: GET
- Request: `Request`
- Response: `kckcRsp`

2. request definition



```golang
type Request struct {
	Province string `form:"province,optional"`
	Mine string `form:"mine,optional"`
	Page int `form:"page"`
	Size int `form:"size"`
}
```


3. response definition



```golang
type KckcRsp struct {
	Total int `json:"total"`
	Kckcs []Kckc `json:"kckcs"` // 矿产勘查工作概况
}
```

### 7. "获取主要可采煤层特征"

1. route definition

- Url: /api/Kcmc
- Method: GET
- Request: `Request`
- Response: `kcmcRsp`

2. request definition



```golang
type Request struct {
	Province string `form:"province,optional"`
	Mine string `form:"mine,optional"`
	Page int `form:"page"`
	Size int `form:"size"`
}
```


3. response definition



```golang
type KcmcRsp struct {
	Total int `json:"total"`
	Kcmcs []Kcmc `json:"kcmcs"` // 主要可采煤层特征
}
```

### 8. "获取勘查区资源量"

1. route definition

- Url: /api/Kcql
- Method: GET
- Request: `Request`
- Response: `kcqlRsp`

2. request definition



```golang
type Request struct {
	Province string `form:"province,optional"`
	Mine string `form:"mine,optional"`
	Page int `form:"page"`
	Size int `form:"size"`
}
```


3. response definition



```golang
type KcqlRsp struct {
	Total int `json:"total"`
	Kcqls []Kcql `json:"kcqls"` // 勘查区资源量
}
```

### 9. "获取矿体特征"

1. route definition

- Url: /api/Kttz
- Method: GET
- Request: `Request`
- Response: `kttzRsp`

2. request definition



```golang
type Request struct {
	Province string `form:"province,optional"`
	Mine string `form:"mine,optional"`
	Page int `form:"page"`
	Size int `form:"size"`
}
```


3. response definition



```golang
type KttzRsp struct {
	Total int `json:"total"`
	Kttzs []Kttz `json:"kttzs"` // 矿体特征
}
```

### 10. "获取煤矿体特征"

1. route definition

- Url: /api/Mctz
- Method: GET
- Request: `Request`
- Response: `mctzRsp`

2. request definition



```golang
type Request struct {
	Province string `form:"province,optional"`
	Mine string `form:"mine,optional"`
	Page int `form:"page"`
	Size int `form:"size"`
}
```


3. response definition



```golang
type MctzRsp struct {
	Total int `json:"total"`
	Mctzs []Mctz `json:"mctzs"` // 煤矿体特征
}
```

### 11. "获取选矿试验"

1. route definition

- Url: /api/Xksy
- Method: GET
- Request: `Request`
- Response: `xksyRsp`

2. request definition



```golang
type Request struct {
	Province string `form:"province,optional"`
	Mine string `form:"mine,optional"`
	Page int `form:"page"`
	Size int `form:"size"`
}
```


3. response definition



```golang
type XksyRsp struct {
	Total int `json:"total"`
	Xksys []Xksy `json:"xksys"` // 选矿试验
}
```

