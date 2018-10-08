package xhr

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"../../utils"
	"github.com/julienschmidt/httprouter"
)

//Post接口
func Post(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	utils.EnableXDA(w, r)                //跨域请求及日志
	xhr := Xhr{}                         //创建一个XHR对象
	json.NewDecoder(r.Body).Decode(&xhr) //赋值
	// 创建请求客户端
	client := &http.Client{}
	req, _ := http.NewRequest(xhr.Method, xhr.Url, strings.NewReader(xhr.Body))
	// // 伪装请求来源
	// req.Header.Set("Referer", xhr.Url)
	// // 伪装浏览器
	// req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/63.0.3239.132 Safari/537.36")
	//设置用户Header
	for k, v := range xhr.Header {
		req.Header.Set(k, v)
	}
	if res, err := client.Do(req); err == nil {
		log.Printf("%v %v %v\n", xhr.Method, xhr.Url, res.StatusCode)
		defer res.Body.Close()
		//转发服务器header
		for k, v := range res.Header {
			w.Header().Set(k, v[0])
		}
		body, _ := ioutil.ReadAll(res.Body)
		jsonStr, _ := json.Marshal(map[string]interface{}{"status": 1, "msg": nil, "data": string(body)})
		w.Write(jsonStr) //请求数据并返回给客户端
	} else {
		jsonStr, _ := json.Marshal(map[string]interface{}{"status": 0, "msg": string(err.Error()), "data": nil})
		w.Write(jsonStr) //请求数据并返回给客户端
	}
}
