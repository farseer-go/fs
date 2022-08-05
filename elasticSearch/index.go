package elasticSearch

import (
	"github.com/olivere/elastic/v7"
	"reflect"
)

type esIndex struct {
	Es *elastic.Client
}

// IsIndexExists 检查索引是否存在
func (esIndex *esIndex) IsIndexExists(index string) (bool, error) {
	return esIndex.Es.IndexExists(index).Do(ctx)
}

// GetIndexMappings 获取索引mapping
func (esIndex *esIndex) GetIndexMappings(indices ...string) (map[string]interface{}, error) {
	return esIndex.Es.GetMapping().Index(indices...).Do(ctx)
}

// PutIndexMapping 设置mapping
func (esIndex *esIndex) PutIndexMapping(index string, mapping interface{}) (bool, error) {
	var putResp *elastic.PutMappingResponse
	var err error
	switch reflect.TypeOf(mapping).Kind() {
	case reflect.String:
		putResp, err = esIndex.Es.PutMapping().Index(index).IgnoreUnavailable(true).
			BodyString(mapping.(string)).Do(ctx)
	case reflect.Map:
		putResp, err = esIndex.Es.PutMapping().Index(index).IgnoreUnavailable(true).
			BodyJson(mapping.(map[string]interface{})).Do(ctx)
	}
	if err != nil {
		return false, err
	}
	if putResp.Acknowledged {
		return true, err
	}
	return false, err
}

// CreateIndex 创建索引并设置值
func (esIndex *esIndex) CreateIndex(index string, body interface{}) (bool, error) {
	resp, err := esIndex.Es.CreateIndex(index).BodyJson(body.(map[string]interface{})).Do(ctx)
	if err != nil {
		return false, err
	}
	if resp.Acknowledged {
		return true, err
	}
	return false, err
}

// DeleteIndex 删除索引
func (esIndex *esIndex) DeleteIndex(indices ...string) (bool, error) {
	resp, err := esIndex.Es.DeleteIndex(indices...).Do(ctx)
	if err != nil {
		return false, err
	}
	if resp.Acknowledged {
		return true, err
	}
	return false, err
}
