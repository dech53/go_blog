package utils

import (
	"context"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"server/global"
	"server/model/other"
)

func MySQLPagination[T any](model *T, option other.MySQLOption) (list []T, total int64, err error) {
	if option.Page < 1 {
		option.Page = 1
	}
	if option.PageSize < 1 {
		option.PageSize = 10
	}
	if option.Order == "" {
		option.Order = "id desc"
	}

	query := global.DB.Model(model)

	if option.Where != nil {
		query = query.Where(option.Where)
	}

	if err = query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	for _, preload := range option.Preload {
		query = query.Preload(preload)
	}

	err = query.Order(option.Order).
		Limit(option.PageSize).
		Offset((option.Page - 1) * option.PageSize).
		Find(&list).Error

	return list, total, err
}
func EsPagination(ctx context.Context, option other.EsOption) (list []types.Hit, total int64, err error) {
	if option.Page < 1 {
		option.Page = 1
	}
	if option.PageSize < 1 {
		option.PageSize = 10
	}

	from := (option.Page - 1) * option.PageSize
	option.Request.Size = &option.PageSize
	option.Request.From = &from

	res, err := global.ESClient.Search().
		Index(option.Index).
		Request(option.Request).
		SourceIncludes_(option.SourceIncludes...).
		Do(ctx)
	if err != nil {
		return nil, 0, err
	}

	list = res.Hits.Hits
	total = res.Hits.Total.Value
	return list, total, nil
}
