## 内部属用基于go-chi的web框架

# 框架引入
> go get -u github.com/wailovet/osmanthuswine

### 目录结构
```
app
 |--index
      |--index.go
html
 |--静态文件....
main.go
config.json
```

# 开始
#### 创建以上目录结构


+ /config.json 配置文件

```json
{
  "port": "8808",
  "host": "0.0.0.0",
  "cross_domain": "*",
  "post_max_memory": 1024000
}
```

+ /main.go文件

```
package main

import (
	"github.com/wailovet/osmanthuswine"
	"github.com/wailovet/osmanthuswine/src/core"
)

func main() {
	//注册index控制器
	core.GetInstanceRouterManage().Registered(&index.Index{})
	//主程序执行
	osmanthuswine.Run()
}
```


+ /app/index/index.go文件

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
## 输入参数
#### 


## core.Responsea 支持方法
### 输出显示
> 如果输出时不处于函数结尾,记得return
+ DisplayByData(data interface{})
```
{
    "code":0,
    "data":data,
    "msg":""
}
```
- - -
+ DisplayBySuccess(msg string)
```
{
    "code":0,
    "data":null,
    "msg":msg
}
```
- - -
+ DisplayByError(msg string, code int)
```
{
    "code":code,
    "data":null,
    "msg":msg
}
```
- - -
+ Display(data interface{}, msg string, code int)
```
{
    "code":code,
    "data":data,
    "msg":msg
}
```
- - -
+ DisplayByString(data string)
```
data //直接输出data以string形式
```
- - -
+ DisplayByRaw(data []byte)
```
data //直接输出data以[]byte形式,可用于直接输出二进制文件
```
### session操作
> 目前session实现基于securecookie,以加密形式储存与cookie中,注意不要存放大量数据,以免超过cookie的最大储存值
+ SetSession(name string, value string) //设置session
+ DeleteSession(name string) //删除session
+ ClearSession() //清空session

### cookie操作
> 尽量以session的形式操作
+ SetCookie(name string, value string)