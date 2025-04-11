package flag

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"os"
	"server/global"
	"server/model/elasticsearch"
	"server/model/other"
	"time"
)

func ElasticsearchExport() error {
	var response other.ESIndexResponse

	res, err := global.ESClient.Search().
		Index(elasticsearch.ArticleIndex()).
		Scroll("1m").
		Size(1000).
		Query(&types.Query{MatchAll: &types.MatchAllQuery{}}).
		Do(context.TODO())
	if err != nil {
		return err
	}

	for _, hit := range res.Hits.Hits {
		data := other.Data{
			ID:  hit.Id_,
			Doc: hit.Source_,
		}
		response.Data = append(response.Data, data)
	}

	for {
		res, err := global.ESClient.Scroll().ScrollId(*res.ScrollId_).Scroll("1m").Do(context.TODO())
		if err != nil {
			// 如果请求出错，返回错误
			return err
		}

		if len(res.Hits.Hits) == 0 {
			break
		}

		for _, hit := range res.Hits.Hits {
			data := other.Data{
				ID:  hit.Id_,
				Doc: hit.Source_,
			}
			response.Data = append(response.Data, data)
		}
	}

	_, err = global.ESClient.ClearScroll().ScrollId(*res.ScrollId_).Do(context.TODO())
	if err != nil {
		return err
	}

	fileName := fmt.Sprintf("es_%s.json", time.Now().Format("20060102"))

	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	byteData, err := json.Marshal(response)
	if err != nil {
		return err
	}

	_, err = file.Write(byteData)
	if err != nil {
		return err
	}

	return nil
}
