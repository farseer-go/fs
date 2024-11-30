package test

import (
	"fmt"
	"testing"

	"github.com/farseer-go/fs/trace"
	"github.com/farseer-go/fs/trace/eumCallType"
	"github.com/stretchr/testify/assert"
)

func TestTrace(t *testing.T) {
	assert.Equal(t, "Grpc", eumCallType.Grpc.ToString())
	assert.Equal(t, "Http", eumCallType.Http.ToString())
	assert.Equal(t, "Database", eumCallType.Database.ToString())
	assert.Equal(t, "Redis", eumCallType.Redis.ToString())
	assert.Equal(t, "Mq", eumCallType.Mq.ToString())
	assert.Equal(t, "Elasticsearch", eumCallType.Elasticsearch.ToString())
	assert.Equal(t, "Hand", eumCallType.Hand.ToString())
	assert.Equal(t, "Etcd", eumCallType.Etcd.ToString())
	assert.Equal(t, "", eumCallType.Enum(9).ToString())

	baseTraceDetail := trace.BaseTraceDetail{}
	baseTraceDetail.SetSql("", "", "", "", 0)
	baseTraceDetail.Ignore()
	assert.Equal(t, true, baseTraceDetail.IsIgnore())
	baseTraceDetail.GetLevel()
	testErr(baseTraceDetail)

	// EmptyManager
	iManager := trace.EmptyManager{}
	iManager.EntryWebApi("", "", "", "", nil, "")
	iManager.EntryFSchedule("", 0, nil)
	iManager.EntryTaskGroup("", "", 0)
	iManager.EntryMqConsumer("", "", "", "", "")
	iManager.EntryQueueConsumer("", "")
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
	iManager.TraceMqSend("", "", "", "")
	iManager.TraceRedis("", "", "")

	iManager.TraceHand("").ToString()
	iManager.TraceHand("").GetTraceDetail()
	iManager.TraceHand("").End(nil)
	iManager.TraceHand("").Ignore()
	iManager.TraceHand("").IsIgnore()
	iManager.TraceHand("").GetLevel()
	iManager.TraceHand("").SetSql("", "", "", "", 0)
	iManager.TraceHand("").SetHttpRequest("", nil, nil, "", "", 0)

	iManager.EntryQueueConsumer("", "").Ignore()
	iManager.EntryQueueConsumer("", "").GetAppInfo()
	iManager.EntryQueueConsumer("", "").AddDetail(nil)
	iManager.EntryQueueConsumer("", "").Error(nil)
	iManager.EntryQueueConsumer("", "").SetBody("", 0, "")
}

func testErr(baseTraceDetail trace.BaseTraceDetail) {
	baseTraceDetail.End(fmt.Errorf(""))
}
