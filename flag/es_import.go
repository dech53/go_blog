package flag

import (
	"context"
	"encoding/json"
	"github.com/elastic/go-elasticsearch/v8/typedapi/core/bulk"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/refresh"
	"os"
	"server/global"
	"server/model/elasticsearch"
	"server/model/other"
	"server/service"
)

func ElasticsearchImport(jsonPath string) (int, error) {
	byteData, err := os.ReadFile(jsonPath)
	if err != nil {
		return 0, err
	}

	var response other.ESIndexResponse
	err = json.Unmarshal(byteData, &response)
	if err != nil {
		return 0, err
	}

	esService := service.ServiceGroupApp.EsService
	indexExists, err := esService.IndexExists(elasticsearch.ArticleIndex())
	if err != nil {
		return 0, err
	}
	if indexExists {
		if err := esService.IndexDelete(elasticsearch.ArticleIndex()); err != nil {
			return 0, err
		}
	}
	err = esService.IndexCreate(elasticsearch.ArticleIndex(), elasticsearch.ArticleMapping())
	if err != nil {
		return 0, err
	}

	var request bulk.Request
	for _, data := range response.Data {
		request = append(request, types.OperationContainer{Index: &types.IndexOperation{Id_: data.ID}})
		request = append(request, data.Doc)
	}

	_, err = global.ESClient.Bulk().
		Request(&request).
		Index(elasticsearch.ArticleIndex()).
		Refresh(refresh.True).
		Do(context.TODO())
	if err != nil {
		return 0, err
	}

	total := len(response.Data)
	return total, nil
}
