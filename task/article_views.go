package task

import (
	"context"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/scriptlanguage"
	"server/global"
	"server/model/elasticsearch"
	"server/service"
	"strconv"
)

func UpdateArticleViewsSyncTask() error {
	articleView := service.ServiceGroupApp.ArticleService.NewArticleView()

	viewsInfo := articleView.GetInfo()
	for id, num := range viewsInfo {
		if num == 0 {
			continue
		}

		source := "ctx._source.views += " + strconv.Itoa(num)
		script := types.Script{Source: &source, Lang: &scriptlanguage.Painless}
		_, err := global.ESClient.Update(elasticsearch.ArticleIndex(), id).Script(&script).Do(context.TODO())
		return err
	}

	articleView.Clear()
	return nil
}
