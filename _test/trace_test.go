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

	baseTraceDetail := trace.TraceDetail{}
	baseTraceDetail.SetSql("", "", "", "", 0)
	baseTraceDetail.Ignore()
	assert.Equal(t, true, baseTraceDetail.IsIgnore())
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

	iManager.EntryQueueConsumer("", "").Ignore()
	iManager.EntryQueueConsumer("", "").GetAppInfo()
	iManager.EntryQueueConsumer("", "").AddDetail(&trace.TraceDetail{})
	iManager.EntryQueueConsumer("", "").Error(nil)
	iManager.EntryQueueConsumer("", "").SetBody("", 0, "", nil)
}

func testErr(baseTraceDetail trace.TraceDetail) {
	baseTraceDetail.End(fmt.Errorf(""))
}
