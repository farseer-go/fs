package trace

import (
	"strconv"
	"time"

	"github.com/farseer-go/fs/core"
	"github.com/farseer-go/fs/dateTime"
	"github.com/farseer-go/fs/sonyflake"
	"github.com/farseer-go/fs/trace/eumCallType"
	"github.com/farseer-go/fs/trace/eumTraceType"
)

type TraceContext struct {
	TraceId       string            `json:"tid"` // 上下文ID
	AppId         string            `json:"aid"` // 应用ID
	AppName       string            `json:"an"`  // 应用名称
	AppIp         string            `json:"aip"` // 应用IP
	ParentAppName string            `json:"pn"`  // 上游应用
	TraceLevel    int               `json:"tl"`  // 逐层递增（显示上下游顺序）
	StartTs       int64             `json:"st"`  // 调用开始时间戳（微秒）
	EndTs         int64             `json:"et"`  // 调用结束时间戳（微秒）
	UseTs         time.Duration     `json:"ut"`  // 总共使用时间（微秒）
	UseDesc       string            `json:"ud"`  // 总共使用时间（描述）
	TraceType     eumTraceType.Enum `json:"tt"`  // 状态码
	List          []*TraceDetail    `json:"l"`   // 调用的上下文trace.TraceDetail
	TraceCount    int               `json:"tc"`  // 追踪明细数量
	ignore        bool              // 忽略这次的链路追踪
	ignoreDetail  bool              // 忽略链路明细
	Exception     *ExceptionStack   `json:"e"` // 异常信息
	WebContext
	ConsumerContext
	TaskContext
	WatchKeyContext
	CreateAt dateTime.DateTime `json:"ca"` // 请求时间
}

func NewTraceContext() *TraceContext {
	context := &TraceContext{
		AppId:      strconv.FormatInt(core.AppId, 10),
		AppName:    core.AppName,
		AppIp:      core.AppIp,
		TraceId:    strconv.FormatInt(sonyflake.GenerateId(), 10),
		TraceLevel: 0,
		StartTs:    time.Now().UnixMicro(),
		TraceType:  eumTraceType.WebApi,
		CreateAt:   dateTime.Now(),
		List:       make([]*TraceDetail, 0),
	}
	CurTraceContext.Set(context)
	ScopeLevel.Set([]*TraceDetail{})
	return context
}

type WebContext struct {
	WebDomain          string            `json:"wd"`  // 请求域名
	WebPath            string            `json:"wp"`  // 请求地址
	WebMethod          string            `json:"wm"`  // 请求方式
	WebContentType     string            `json:"wct"` // 请求内容类型
	WebStatusCode      int               `json:"wsc"` // 状态码
	WebHeaders         map[string]string `json:"wh"`  // 请求头部
	WebResponseHeaders map[string]string `json:"wh"`  // 响应头部
	WebRequestBody     string            `json:"wrb"` // 请求参数
	WebResponseBody    string            `json:"wpb"` // 输出参数
	WebRequestIp       string            `json:"wip"` // 客户端IP
}

func (receiver WebContext) IsNil() bool {
	return receiver.WebDomain == "" && receiver.WebPath == "" && receiver.WebMethod == "" && receiver.WebContentType == "" && receiver.WebStatusCode == 0
}

type ConsumerContext struct {
	ConsumerServer     string `json:"cs"` // MQ服务器
	ConsumerQueueName  string `json:"cq"` // 队列名称
	ConsumerRoutingKey string `json:"cr"` // 路由KEY
}

func (receiver ConsumerContext) IsNil() bool {
	return receiver.ConsumerServer == "" && receiver.ConsumerQueueName == "" && receiver.ConsumerRoutingKey == ""
}

type TaskContext struct {
	TaskName      string            `json:"tn"`  // 任务名称
	TaskGroupName string            `json:"tgn"` // 任务组ID
	TaskId        int64             `json:"tid"` // 任务ID
	TaskData      map[string]string `json:"td"`  // 任务数据
}

func (receiver TaskContext) IsNil() bool {
	return receiver.TaskName == "" && receiver.TaskGroupName == "" && receiver.TaskId == 0
}

type WatchKeyContext struct {
	WatchKey string `json:"wk"` // KEY
}

func (receiver WatchKeyContext) IsNil() bool {
	return receiver.WatchKey == ""
}

func (receiver *TraceContext) SetBody(requestBody string, statusCode int, responseBody string, responseHead map[string]string) {
	// 限制响应报文最多存储4096长度
	if len(responseBody) > 4096 {
		responseBody = responseBody[:4096]
	}
	receiver.WebContext.WebRequestBody = requestBody
	receiver.WebContext.WebStatusCode = statusCode
	receiver.WebContext.WebResponseBody = responseBody
	receiver.WebContext.WebResponseHeaders = responseHead
}

func (receiver *TraceContext) SetResponseBody(responseBody string) {
	// 限制响应报文最多存储4096长度
	if len(responseBody) > 4096 {
		responseBody = responseBody[:4096]
	}
	receiver.WebContext.WebResponseBody = responseBody
}

// 取消链路
func (receiver *TraceContext) Ignore() {
	receiver.ignore = true
}

func (receiver *TraceContext) IsIgnore() bool {
	return receiver.ignore
}

func (receiver *TraceContext) IgnoreDetail(f func()) {
	defer func() {
		receiver.ignoreDetail = false
	}()

	traceDetail := NewTraceDetail(eumCallType.Hand, "")
	traceDetail.TraceDetailHand.HandName = "忽略明细"
	traceDetail.Comment = "忽略明细"
	traceDetail.Timeline = time.Duration(traceDetail.StartTs-receiver.StartTs) * time.Microsecond
	if len(receiver.List) > 0 {
		traceDetail.UnTraceTs = time.Duration(traceDetail.StartTs-receiver.List[len(receiver.List)-1].EndTs) * time.Microsecond
	} else {
		traceDetail.UnTraceTs = time.Duration(traceDetail.StartTs-receiver.StartTs) * time.Microsecond
	}

	receiver.ignoreDetail = true
	f()
	traceDetail.End(nil)
	receiver.List = append(receiver.List, traceDetail)
}

// AddDetail 添加链路明细
func (receiver *TraceContext) AddDetail(traceDetail *TraceDetail) {
	// 没有忽略明细，才要加入
	if !receiver.ignoreDetail {
		receiver.List = append(receiver.List, traceDetail)
	}
}

func (receiver *TraceContext) Error(err error) {
	if err != nil {
		receiver.Exception = &ExceptionStack{
			ExceptionIsException: true,
			ExceptionMessage:     err.Error(),
			ExceptionDetails:     GetCallerInfo(),
		}
	}
}

func (receiver *TraceContext) GetAppInfo() (string, string, string, string, string) {
	return receiver.TraceId, receiver.AppName, receiver.AppId, receiver.AppIp, receiver.ParentAppName
}
