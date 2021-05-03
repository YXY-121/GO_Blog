package model
type ArticleTag struct {
		*Model
		TagID  uint32`json:"tag_id"`
		ArticleId uint32 `json:"article_id"`
}

func (a ArticleTag)TableName() string {
	return "blog_article_tag"//这是啥 数据库吗？？
}