##内部属用基于go-chi的web框架

#框架引入
> go get -u github.com/wailovet/osmanthuswine

###目录结构
```
app
 |--index
      |--controller
             |--index.go
config
 |--config.json
html
 |--静态文件....
main.go
```

# 开始
#### 创建以上目录结构
> main.go文件
```
func main() {
	core.GetInstanceRouterManage().Registered(&index.Index{})
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