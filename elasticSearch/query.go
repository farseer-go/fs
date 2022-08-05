package elasticSearch

import (
	"github.com/olivere/elastic/v7"
	"reflect"
)

// Query 查询结构
type esQuery struct {
	Es *elastic.Client
}

// MustTermOrTermsQuery 精确字段查询-个值或数组
func (esQuery *esQuery) MustTermOrTermsQuery(index string, name string, val interface{}) ([]*elastic.SearchHit, error) {
	query := elastic.NewBoolQuery()
	if reflect.TypeOf(val).Kind() == reflect.Slice {
		query.Must(elastic.NewTermsQuery(name, val))
	} else {
		query.Must(elastic.NewTermsQuery(name, val))
	}
	resp, err := esQuery.Es.Search().Index(index).Query(query).Size(10000).Do(ctx)
	return resp.Hits.Hits, err
}

// MustRangeQuery 范围查询
func (esQuery *esQuery) MustRangeQuery(index string, name string, lte interface{}, gte interface{}) ([]*elastic.SearchHit, error) {
	query := elastic.NewBoolQuery()
	query.Must(elastic.NewRangeQuery(name).Gte(gte).Lte(lte))
	resp, err := esQuery.Es.Search().Index(index).Query(query).Size(10000).Do(ctx)
	return resp.Hits.Hits, err
}

// SizeQuery 指定数量返回
func (esQuery *esQuery) SizeQuery(index string, size int) ([]*elastic.SearchHit, error) {
	resp, err := esQuery.Es.Search().Index(index).Size(size).Do(ctx)
	return resp.Hits.Hits, err
}

// SourceQuery 指定返回字段
func (esQuery *esQuery) SourceQuery(index string, source interface{}) ([]*elastic.SearchHit, error) {
	resp, err := esQuery.Es.Search().Index(index).Source(source).Size(10000).Do(ctx)
	return resp.Hits.Hits, err
}

// SortQuery 指定顺序查询
func (esQuery *esQuery) SortQuery(index string, field string, ascending bool) ([]*elastic.SearchHit, error) {
	resp, err := esQuery.Es.Search().Index(index).Size(10000).Sort(field, ascending).Do(ctx)
	return resp.Hits.Hits, err
}

// AggregationTerms 简单桶聚合
func (esQuery *esQuery) AggregationTerms(index string, name string, field string) (interface{}, error) {
	agg := elastic.NewTermsAggregation().Field(field).Size(10000)
	resp, err := esQuery.Es.Search().
		Index(index).
		Aggregation(name, agg).
		Do(ctx)
	if err != nil {
		return nil, err
	}
	ret, ok := resp.Aggregations.Terms(name)
	if !ok {
		return nil, nil
	}
	bucketKeys := make([]string, len(ret.Buckets))
	for i, _ := range ret.Buckets {
		bucketKeys[i] = ret.Buckets[i].Key.(string)
	}
	return bucketKeys, nil
}

// AggregationTermsWithSubAgg 多级聚合查询
func (esQuery *esQuery) AggregationTermsWithSubAgg(index string, name1 string, field1 string, name2 string, field2 string) (interface{}, error) {
	agg := elastic.NewTermsAggregation().Field(field1).Size(10000)
	resp, err := esQuery.Es.Search().
		Index(index).
		Size(0).
		Aggregation(name1, agg).
		Do(ctx)
	if err != nil {
		return nil, err
	}
	ret, ok := resp.Aggregations.Terms(name1)
	if !ok {
		return nil, nil
	}
	fieldsMap := make(map[string][]interface{})
	for i, _ := range ret.Buckets {
		k := ret.Buckets[i].Key.(string)
		// subAggregation
		agg = elastic.NewTermsAggregation().Field(field2).Size(10000)
		resp, err = esQuery.Es.Search().
			Index(index).
			Size(0).
			Query(elastic.NewBoolQuery().Must(elastic.NewTermQuery(field1, k))).
			Aggregation(name2, agg).
			Do(ctx)
		if err != nil {
			return nil, err
		}
		ret2, _ := resp.Aggregations.Terms(name2)
		for j, _ := range ret2.Buckets {
			fieldsMap[k] = append(fieldsMap[k], ret2.Buckets[j].Key.(interface{}))
		}
	}
	return fieldsMap, nil
}

// AggSub 两级terms usage
func (esQuery *esQuery) AggSub(index string, field1 string, field2 string) (interface{}, error) {
	resp, err := esQuery.Es.Search(index).
		Aggregation(
			"mainAgg",
			elastic.NewTermsAggregation().Field(field1).
				SubAggregation(
					"sub_1", elastic.NewTermsAggregation().Field(field2),
				),
		).
		Size(0).
		Pretty(true).
		Do(ctx)
	if err != nil {
		return nil, err
	}
	ret, _ := resp.Aggregations.Terms("mainAgg")
	return ret, nil
}

// AggregationMax 聚合stats-max
func (esQuery *esQuery) AggregationMax(index string, field string) (interface{}, error) {
	agg := elastic.NewMaxAggregation().Field(field)
	resp, err := esQuery.Es.Search().
		Index(index).
		Aggregation("max_"+field, agg).
		Do(ctx)
	if err != nil {
		return nil, err
	}
	ret, _ := resp.Aggregations.Max("max_" + field)
	ans := *ret.Value
	return ans, nil
}
