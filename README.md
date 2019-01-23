## 内部属用基于go-chi的web框架

# 框架引入
> go get -u github.com/wailovet/osmanthuswine

### 目录结构
```
app
 |--index
      |--index.go
config
 |--main.json
html
 |--静态文件....
main.go
```

# 开始
#### 创建以上目录结构
> /main.go文件
```
func main() {
	//注册index控制器
	core.GetInstanceRouterManage().Registered(&index.Index{})
	//主程序执行
	osmanthuswine.Run()
}
```


>/app/index/controller/index.go文件
```
package index

import (
	"github.com/wailovet/osmanthuswine/src/core"
)

type Index struct {
}

func (n *Index) Index(req core.Request, res core.Response) {
	res.DisplayByData(req)
}

```

>/config/main.json
```json
{
  "port": "8808",
  "host": "0.0.0.0",
  "cross_domain": "*",
  "post_max_memory": 1024000
}
```