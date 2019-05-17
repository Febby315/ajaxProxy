# 接口文档

## 开发说明

### 引用架包

#### WEB框架

> go get github.com/julienschmidt/httprouter

## 接口及配置文档

### 配置说明

```ini
[global]
  # 端口配置
  port=8081
```

### 1. 跨域接口 POST /api/post

> 接口参数说明

属性|类型|默认值|允许值|说明
-|-|-|-|-
url|字符串|||需要代理的接口地址
mathod|字符串|"GET"|"GET","POST"|接口的请求方式
header|对象|||请求头(类似键值对)
body|字符串|||请求体

    url和mathod为必填项

#### 使用示例

```javascript
// http://127.0.0.1:8081/api/get?url=encodeURIComponent(url)
var url = "http://img.hb.aicdn.com/d84030573bcad38920ba4b2248cd98cde343c52722674-0PEN0M";
`http://127.0.0.1:8081/api/get?url=${encodeURIComponent(url)}`
```

```javascript
// http://127.0.0.1:8081/api/post
{
  "url":"", //请求地址(必须)
  "mathod":"GET", //请求方式(必须:GET|POST)
  "header":{
    "key":"value"
  } //请求头("可选")键值对
  "body":"",  //请求体("可选")
}
```

### 1. 跨域接口 GET /api/get

> 代理请求静态资源(由于部分今天资源也存在跨域或者限制问题)

#### 使用示例

```javascript
// http://127.0.0.1:8081/api/get?url=encodeURIComponent(url)
var url = "http://img.hb.aicdn.com/d84030573bcad38920ba4b2248cd98cde343c52722674-0PEN0M";
`http://127.0.0.1:8081/api/get?url=${encodeURIComponent(url)}`
```

    同样可传/api/post接口的所有参数,但是代理请求URL将以querystring参数中url参数为准

## 使用说明

1. 双击proxy_xhr.exe
2. 用浏览器打开exampleles目录中实例文件"# ajaxProxy"
