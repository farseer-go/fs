package trace

import (
	"fmt"
	"testing"

	"github.com/farseer-go/fs/trace/eumCallType"
)

func TestTrace(t *testing.T) {
	emptyManager := EmptyManager{}
	emptyManager.GetCurTrace()
	emptyManager.TraceDatabase()
	emptyManager.TraceEtcd("", "", 0)
	emptyManager.TraceElasticsearch("", "", "")
	emptyManager.TraceHand("")
	emptyManager.TraceMq("", "", "")
	emptyManager.EntryFSchedule("", 0, nil)
	emptyManager.EntryMqConsumer("", "", "", "", "")
	emptyManager.EntryQueueConsumer("", "")
	emptyManager.EntryEventConsumer("", "", "")
	emptyManager.EntryTask("")
	emptyManager.EntryWatchKey("")
	emptyManager.EntryWebApi("", "", "", "", nil, "")
	emptyManager.TraceDatabaseOpen("", "")
	emptyManager.TraceHttp("", "")
	emptyManager.TraceMqSend("", "", "", "")
	emptyManager.TraceRedis("", "", "")

	eumCallType.Http.ToString()
	eumCallType.Grpc.ToString()
	eumCallType.Database.ToString()
	eumCallType.Redis.ToString()
	eumCallType.Mq.ToString()
	eumCallType.Elasticsearch.ToString()
	eumCallType.Etcd.ToString()
	eumCallType.Hand.ToString()

	detail := BaseTraceDetail{}
	detail.SetHttpRequest("", nil, nil, "", "", 0)
	detail.End(fmt.Errorf(""))
	detail.Ignore()
	detail.GetLevel()
	detail.IsIgnore()
	detail.SetSql("", "", "", "", 0)

	var traceHand ITraceDetail = &TraceDetailHand{
		BaseTraceDetail: detail,
		Name:            "",
	}
	traceHand.End(nil)
}
