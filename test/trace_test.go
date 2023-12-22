package test

import (
	"fmt"
	"github.com/farseer-go/fs"
	"github.com/farseer-go/fs/container"
	"github.com/farseer-go/fs/modules"
	"github.com/farseer-go/fs/trace"
	"github.com/farseer-go/fs/trace/eumCallType"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTrace(t *testing.T) {
	fs.Initialize[modules.FarseerKernelModule]("unit test")

	assert.Equal(t, "Grpc", eumCallType.Grpc.ToString())
	assert.Equal(t, "Http", eumCallType.Http.ToString())
	assert.Equal(t, "Database", eumCallType.Database.ToString())
	assert.Equal(t, "Redis", eumCallType.Redis.ToString())
	assert.Equal(t, "Mq", eumCallType.Mq.ToString())
	assert.Equal(t, "Elasticsearch", eumCallType.Elasticsearch.ToString())
	assert.Equal(t, "Hand", eumCallType.Hand.ToString())
	assert.Equal(t, "KeyLocation", eumCallType.KeyLocation.ToString())
	assert.Equal(t, "Etcd", eumCallType.Etcd.ToString())
	assert.Equal(t, "", eumCallType.Enum(9).ToString())

	baseTraceDetail := trace.BaseTraceDetail{}
	baseTraceDetail.SetSql("", "", "", 0)
	baseTraceDetail.Ignore()
	assert.Equal(t, true, baseTraceDetail.IsIgnore())
	baseTraceDetail.GetLevel()
	testErr(baseTraceDetail)

	// EmptyManager
	iManager := container.Resolve[trace.IManager]()
	iManager.EntryWebApi("", "", "", "", nil, "", "")
	iManager.EntryFSchedule("", 0, 0)
	iManager.EntryMqConsumer("", "", "")
	iManager.EntryQueueConsumer("")
	iManager.EntryTask("")
	iManager.EntryWatchKey("")
	iManager.TraceMq("", "", "iManager")
	iManager.GetCurTrace()
	iManager.TraceDatabase()
	iManager.TraceDatabaseOpen("", "")
	iManager.TraceElasticsearch("", "", "")
	iManager.TraceEtcd("", "", 0)
	iManager.TraceHand("")
	iManager.TraceHttp("", "")
	iManager.TraceKeyLocation("")
	iManager.TraceMqSend("", "", "", "")
	iManager.TraceRedis("", "", "")

	iManager.TraceHand("").ToString()
	iManager.TraceHand("").GetTraceDetail()
	iManager.TraceHand("").End(nil)
	iManager.TraceHand("").Ignore()
	iManager.TraceHand("").IsIgnore()
	iManager.TraceHand("").GetLevel()
	iManager.TraceHand("").SetSql("", "", "", 0)

	iManager.EntryQueueConsumer("").End()
	iManager.EntryQueueConsumer("").Ignore()
	iManager.EntryQueueConsumer("").GetList()
	iManager.EntryQueueConsumer("").GetTraceId()
	iManager.EntryQueueConsumer("").AddDetail(nil)
	iManager.EntryQueueConsumer("").Error(nil)
	iManager.EntryQueueConsumer("").GetStartTs()
	iManager.EntryQueueConsumer("").SetBody("", 0, "")
}

func testErr(baseTraceDetail trace.BaseTraceDetail) {
	baseTraceDetail.End(fmt.Errorf(""))
}