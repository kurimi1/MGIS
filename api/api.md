### 1. "获取矿产地基本情况"

1. route definition

- Url: /api/K
- Method: GET
- Request: `Request`
- Response: `Response`

2. request definition



```golang
type Request struct {
	Province string `json:"province"`
	Mine string `json:"mine,optional"`
	Page int `json:"page"`
	Size int `json:"size"`
}
```


3. response definition



```golang
type Response struct {
	Total int `json:"total"`
	Kcdjs []Kcdj `json:"kcdjs"` // 矿产地基本情况
}
```

