# 接口文档

## 开发说明

### 引用架包

#### WEB框架 [httprouter](https://github.com/julienschmidt/httprouter)

> go get github.com/julienschmidt/httprouter

## 接口及配置文档

### 配置说明

```ini
[global]
  # 端口配置
  port=8000
```

### 1. 跨域接口 POST /api/xhr

> 接口参数说明

属性|类型|默认值|允许值|说明
-|-|-|-|-
url|字符串|||需要代理的接口地址
mathod|字符串|"GET"|"GET","POST"|接口的请求方式
header|对象|||请求头(类似键值对)
body|字符串|||请求体

    url和mathod为必填项

```javascript
{
  "url":"", //请求地址(必须)
  "mathod":"GET", //请求方式(必须:GET|POST)
  "header":"华高55(8155)" //请求头("可选")
  "body":"",  //请求体("可选")
}
```

## 使用说明

1. 双击proxy_xhr.exe
2. 用浏览器打开exampleles目录中实例文件"# ajaxProxy" 
