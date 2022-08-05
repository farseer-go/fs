package elasticSearch

import (
	"github.com/olivere/elastic/v7"
)

type esDoc struct {
	Es *elastic.Client
}

// InsertDocToIndex 插入文档
func (esDoc *esDoc) InsertDocToIndex(index string, doc interface{}) (bool, string, error) {
	resp, err := esDoc.Es.Index().Index(index).BodyJson(doc.(map[string]interface{})).Do(ctx)
	if err != nil {
		return false, resp.Result, err
	}
	return true, resp.Result, err
}

// CreateOrUpdateDoc 创建或者更新文档，注意：注意用update更新个别字段时，请用map结构，struct会更新全部的。
func (esDoc *esDoc) CreateOrUpdateDoc(index string, id string, doc interface{}) (bool, string, error) {
	resp, err := esDoc.Es.Update().Index(index).Id(id).Doc(doc).DocAsUpsert(true).Do(ctx)
	if err != nil {
		return false, resp.Result, err
	}
	return true, resp.Result, err
}

// BulkInsertDocs 批量插入
func (esDoc *esDoc) BulkInsertDocs(index string, docs []interface{}) (bool, error) {
	length := len(docs)
	bulkReq := esDoc.Es.Bulk()
	for i := 0; i < length; i++ {
		doc := docs[i]
		req := elastic.NewBulkCreateRequest().Index(index).Type("_doc").Doc(doc)
		bulkReq = bulkReq.Add(req)
	}
	_, err := bulkReq.Do(ctx)
	if err != nil {
		return false, err
	}
	return true, err
}

// DelDoc 删除文档
func (esDoc *esDoc) DelDoc(index string, id string) (bool, string, error) {
	resp, err := esDoc.Es.Delete().Index(index).Id(id).Do(ctx)
	if err != nil {
		return false, resp.Result, err
	}
	return true, resp.Result, err
}

// GetDoc 获取单个文档
func (esDoc *esDoc) GetDoc(index string, id string) (interface{}, error) {
	resp, err := esDoc.Es.Get().Index(index).Id(id).Do(ctx)
	if err != nil {
		return nil, err
	}
	return resp.Source, err
}
