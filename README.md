# ginweb  
> use golang to testing gin web frame  

# api test  

## 1. GET 方法  
### api /v1/plain

> 直接访问  
> Query(key) 或者 get 方法的值
> DefaultQuery(key, defaultvalue) 获取 get 方法的值，如没有，则给与默认值    
> GetQueryArray(key) 获取 get 方法的值，值必须是 []string 格式 


```
# curl http://x.x.x./v1/plain
{"msg":{"name":"none","domain":["none"],"token":"none"},"result":true}

# curl 'http://10.189.20.49/v1/plain?name=terry&token=aaa&domain=a&domain=b'
{"msg":{"name":"terry","domain":["a","b"],"token":"aaa"},"result":true}
```

> 假如要获取的数据格式为 []string, int, float64  参考下面方法  

```
var ss []string
c.DefaultQueryVar(key, &ss, []string{"1", "2", "3"})
var i int
c.DefaultQueryVar(key, &i, 3)
var f float64
c.DefaultQueryVar(key, &f, 3.14)
```

## 2.  POST 方法
###  api /v1/plain

> DefaultPostForm(key) 获取 post 中的 data 方法
> PostFormArray(key)  或者 []string 方法

```
# curl -d "name=terry"  -d "token=12312" -d "domain=abc" -d "domain=222"  -X POST  http://10.189.20.49/v1/plain
{"msg":{"name":"terry","domain":["abc","222"],"token":"12312"},"result":true}
```
### api /v1/bind

> 以 json 方式进行参数传递
> 通过 ShouldBindJSON 方式传入参数至预定义 Struct 中

```
# curl -H "Content-Type:application/json" -X POST -d '{"name":"terry","token":"serwer","domain":["a","b"]}'  http://10.189.20.49/v1/bind
{"msg":"user: terry, domain: [a b], token serwer","result":true}
```

### api /v1/config

> 利用 post json 信息， 生成模板配置文件  
> 主要用于客户端自动化生成配置使用  

```
# curl -H "Content-Type:application/json" -X POST -d '{"name":"terry","token":"serwer","domain":["a","b"]}'  http://10.189.20.49/v1/config
Its config template example:
username variable example:

username input exists.

username is : terry


domain variable example:
upstream socks5 127.0.0.1:40408 ".a"
upstream socks5 127.0.0.1:40408 ".b"
upstream socks5 127.0.0.1:40408 "a"
upstream socks5 127.0.0.1:40408 "b"


token variable example:
token is serwer

```
