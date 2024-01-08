package trace

import (
	"fmt"
	"github.com/farseer-go/fs/trace/eumCallType"
	"testing"
)

func TestTrace(t *testing.T) {
	emptyManager := EmptyManager{}
	emptyManager.GetCurTrace()
	emptyManager.TraceDatabase()
	emptyManager.TraceEtcd("", "", 0)
	emptyManager.TraceElasticsearch("", "", "")
	emptyManager.TraceHand("")
	emptyManager.TraceMq("", "", "")
	emptyManager.EntryFSchedule("", 0, 0)
	emptyManager.EntryMqConsumer("", "", "")
	emptyManager.EntryQueueConsumer("", "")
	emptyManager.EntryTask("")
	emptyManager.EntryWatchKey("")
	emptyManager.EntryWebApi("", "", "", "", nil, "", "")
	emptyManager.TraceDatabaseOpen("", "")
	emptyManager.TraceHttp("", "")
	emptyManager.TraceKeyLocation("")
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
	detail.Exception.IsNil()
	detail.SetHttpRequest("", nil, "", "", 0)
	detail.End(fmt.Errorf(""))
	detail.Ignore()
	detail.GetLevel()
	detail.IsIgnore()
	detail.SetSql("", "", "", "", 0)
}
