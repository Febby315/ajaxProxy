package xhr

// Xhr 代理请求参数对象
type Xhr struct {
	URL    string            `json:"url"`    //请求链接
	Method string            `json:"method"` //请求方式
	Header map[string]string `json:"header"` //自定义header
	Body   string            `json:"body"`   //请求数据
}
