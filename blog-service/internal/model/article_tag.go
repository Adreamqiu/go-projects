package model

type ArticleTag struct {
	*Model
    // 文章ID
    ArticleId uint32  `json:article_id`
    // 标签ID
    TagID uint32  `json:tag_id`
}

func (model ArticleTag) TableName() string {
    return "blog_article_tag"
}