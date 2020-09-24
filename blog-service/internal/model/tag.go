package model

type Tag struct {
	*Model
    // 标签名称
    Name string  `json:name`
    // 状态 0为禁用、1为启用
    State uint8  `json:state`
}

func (model Tag) TableName() string {
    return "blog_tag"
}