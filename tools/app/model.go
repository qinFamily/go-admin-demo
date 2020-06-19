package app

type Response struct {
	// 代码
	Code int `json:"code" example:"200"`
	// 数据集
	Data interface{} `json:"data"`
	// 消息
	Msg string `json:"msg"`
}

type Page struct {
	List      interface{} `json:"list"`
	Count     int         `json:"count"`
	PageIndex int         `json:"pageIndex"`
	PageSize  int         `json:"pageSize"`
}

type PageResponse struct {
	// 代码
	Code int `json:"code" example:"200"`
	// 数据集
	Data Page `json:"data"`
	// 消息
	Msg string `json:"msg"`
}

func (res *Response) ReturnOK() *Response {
	res.Code = 200
	return res
}

func (res *Response) ReturnError(code int) *Response {
	res.Code = code
	return res
}

func (res *PageResponse) ReturnOK() *PageResponse {
	res.Code = 200
	return res
}

type WorkFlowResponse struct {
	Code  int64 `json:"code" example:"2000"`
	Count int   `json:"count" example:"1"`
	// 消息
	Next     interface{} `json:"next"`
	Previous interface{} `json:"previous"`
	// 另外数据集
	Results interface{} `json:"results"`
}

func (res *WorkFlowResponse) ReturnOK() *WorkFlowResponse {
	res.Code = 20000
	return res
}
