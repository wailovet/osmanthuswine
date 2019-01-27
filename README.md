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
## core.Request 使用说明
### 输入参数 
+ req.GET["参数"] //类型:map[string]string
+ req.POST["参数"] //类型:map[string]string
+ req.REQUEST["参数"] //类型:map[string]string , 为GET以及POST的合并值,当出现值冲突时GET参数会被覆盖
### session与cookie获取
+ req.SESSION["参数"] //类型:map[string]string
+ req.COOKIE["参数"] //类型:map[string]string
### header信息获取
+ req.HEADER["参数"] //类型:map[string]string
### 该取值一般以POST-RAW形式传入原始数据,有可能
+ req.BODY //类型:string
### 获取上传的文件,比较少用到,具体用法懒得写
+ req.FILES //类型:map[string][]*multipart.FileHeader



## core.Responsea 使用说明
### 输出显示
> 如果输出时不处于函数结尾,记得return
+ res.DisplayByData(data interface{})
```
{
    "code":0,
    "data":data,
    "msg":""
}
```
- - -
+ res.DisplayBySuccess(msg string)
```
{
    "code":0,
    "data":null,
    "msg":msg
}
```
- - -
+ res.DisplayByError(msg string, code int)
```
{
    "code":code,
    "data":null,
    "msg":msg
}
```
- - -
+ res.Display(data interface{}, msg string, code int)
```
{
    "code":code,
    "data":data,
    "msg":msg
}
```
- - -
+ res.DisplayByString(data string)
```
data //直接输出data以string形式
```
- - -
+ res.DisplayByRaw(data []byte)
```
data //直接输出data以[]byte形式,可用于直接输出二进制文件
```
### session操作
> 目前session实现基于securecookie,以加密形式储存在cookie中,注意不要存放大量数据,以免超过cookie的最大储存值
+ res.SetSession(name string, value string) //设置session
+ res.DeleteSession(name string) //删除session
+ res.ClearSession() //清空session

### cookie操作
> 尽量以session的形式操作
+ res.SetCookie(name string, value string)