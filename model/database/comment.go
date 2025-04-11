package database

import (
	"context"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/scriptlanguage"
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
	"server/global"
	"server/model/elasticsearch"
)

type Comment struct {
	global.MODEL
	ArticleID string    `json:"article_id"`
	PID       *uint     `json:"p_id"`
	PComment  *Comment  `json:"-" gorm:"foreignKey:PID"`
	Children  []Comment `json:"children" gorm:"foreignKey:PID"`
	UserUUID  uuid.UUID `json:"user_uuid" gorm:"type:char(36)"`
	User      User      `json:"user" gorm:"foreignKey:UserUUID;references:UUID"`
	Content   string    `json:"content"`
}

func (c *Comment) AfterCreate(_ *gorm.DB) error {
	source := "ctx._source.comments += 1"
	script := types.Script{Source: &source, Lang: &scriptlanguage.Painless}
	_, err := global.ESClient.Update(elasticsearch.ArticleIndex(), c.ArticleID).Script(&script).Do(context.TODO())
	return err
}

func (c *Comment) BeforeDelete(_ *gorm.DB) error {
	var articleID string
	if err := global.DB.Model(&c).Pluck("article_id", &articleID).Error; err != nil {
		return err
	}
	source := "ctx._source.comments -= 1"
	script := types.Script{Source: &source, Lang: &scriptlanguage.Painless}
	_, err := global.ESClient.Update(elasticsearch.ArticleIndex(), articleID).Script(&script).Do(context.TODO())
	return err
}
