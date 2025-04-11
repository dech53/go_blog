package service

import (
	"context"
	"encoding/json"
	"errors"
	"server/global"
	"server/model/database"
	"server/model/elasticsearch"
	"server/utils"

	"github.com/elastic/go-elasticsearch/v8/typedapi/core/bulk"
	"github.com/elastic/go-elasticsearch/v8/typedapi/core/search"
	"github.com/elastic/go-elasticsearch/v8/typedapi/core/update"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/refresh"
	"gorm.io/gorm"
)

func (articleService *ArticleService) Create(a *elasticsearch.Article) error {
	_, err := global.ESClient.Index(elasticsearch.ArticleIndex()).Request(a).Refresh(refresh.True).Do(context.TODO())
	return err
}

func (articleService *ArticleService) Delete(ids []string) error {
	var request bulk.Request
	for _, id := range ids {
		request = append(request, types.OperationContainer{Delete: &types.DeleteOperation{Id_: &id}})
	}
	_, err := global.ESClient.Bulk().Request(&request).Index(elasticsearch.ArticleIndex()).Refresh(refresh.True).Do(context.TODO())
	return err
}

func (articleService *ArticleService) Get(id string) (elasticsearch.Article, error) {
	var a elasticsearch.Article
	res, err := global.ESClient.Get(elasticsearch.ArticleIndex(), id).Do(context.TODO())
	if err != nil {
		return elasticsearch.Article{}, err
	}
	if !res.Found {
		return elasticsearch.Article{}, errors.New("document not found")
	}
	err = json.Unmarshal(res.Source_, &a)
	return a, err
}

func (articleService *ArticleService) Update(articleID string, v any) error {
	bytes, err := json.Marshal(v)
	if err != nil {
		return err
	}
	_, err = global.ESClient.Update(elasticsearch.ArticleIndex(), articleID).Request(&update.Request{Doc: bytes}).Refresh(refresh.True).Do(context.TODO())
	return err
}

func (articleService *ArticleService) Exists(title string) (bool, error) {
	req := &search.Request{
		Query: &types.Query{
			Match: map[string]types.MatchQuery{"keyword": {Query: title}},
		},
	}
	res, err := global.ESClient.Search().Index(elasticsearch.ArticleIndex()).Request(req).Size(1).Do(context.TODO())
	if err != nil {
		return false, err
	}
	return res.Hits.Total.Value > 0, nil
}

func (articleService *ArticleService) UpdateCategoryCount(tx *gorm.DB, oldCategory, newCategory string) error {
	if newCategory == oldCategory {
		return nil
	}
	if newCategory != "" {
		var newArticleCategory database.ArticleCategory
		if errors.Is(tx.Where("category = ?", newCategory).First(&newArticleCategory).Error, gorm.ErrRecordNotFound) {
			if err := tx.Create(&database.ArticleCategory{Category: newCategory, Number: 1}).Error; err != nil {
				return err
			}
		} else {
			if err := tx.Model(&newArticleCategory).Update("number", gorm.Expr("number + ?", 1)).Error; err != nil {
				return err
			}
		}
	}

	if oldCategory != "" {
		var oldArticleCategory database.ArticleCategory
		if err := tx.Where("category = ?", oldCategory).First(&oldArticleCategory).Update("number", gorm.Expr("number - ?", 1)).Error; err != nil {
			return err
		}
		if oldArticleCategory.Number == 1 {
			if err := tx.Delete(&oldArticleCategory).Error; err != nil {
				return err
			}
		}
	}

	return nil
}

func (articleService *ArticleService) UpdateTagsCount(tx *gorm.DB, oldTags, newTags []string) error {
	addedTags, removedTags := utils.DiffArrays(oldTags, newTags)

	for _, addedTag := range addedTags {
		var t database.ArticleTag
		if errors.Is(tx.Where("tag = ?", addedTag).First(&t).Error, gorm.ErrRecordNotFound) {
			if err := tx.Create(&database.ArticleTag{Tag: addedTag, Number: 1}).Error; err != nil {
				return err
			}
		} else {
			if err := tx.Model(&t).Update("number", gorm.Expr("number + ?", 1)).Error; err != nil {
				return err
			}
		}
	}

	for _, removedTag := range removedTags {
		var t database.ArticleTag
		if err := tx.Where("tag = ?", removedTag).First(&t).Update("number", gorm.Expr("number - ?", 1)).Error; err != nil {
			return err
		}
		if t.Number == 1 {
			if err := tx.Delete(&t).Error; err != nil {
				return err
			}
		}
	}
	return nil
}
