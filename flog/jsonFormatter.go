package flog

import "encoding/json"

// JsonFormatter json格式输出
type JsonFormatter struct {
}

func (r *JsonFormatter) Formatter(log *logData) string {
	marshal, _ := json.Marshal(log)
	return string(marshal)
}
