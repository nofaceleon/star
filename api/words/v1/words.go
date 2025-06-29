package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type ProficiencyLevel uint

const (
	ProficiencyLevel1 ProficiencyLevel = iota + 1
	ProficiencyLevel2
	ProficiencyLevel3
	ProficiencyLevel4
	ProficiencyLevel5
)

type CreateReq struct {
	g.Meta             `path:"words" method:"post" sm:"创建" tags:"单词"`
	Word               string           `json:"word" v:"required|length:1,100" dc:"单词" json:"word,omitempty"`
	Definition         string           `json:"definition,omitempty" v:"required|length:1,300" dc:"单词定义"`
	ExampleSentence    string           `json:"exampleSentence,omitempty" v:"required|length:1,300" dc:"例句"`
	ChineseTranslation string           `json:"chineseTranslation,omitempty" v:"required|length:1,300" dc:"中文翻译"`
	Pronunciation      string           `json:"pronunciation,omitempty" v:"required|length:1,100" dc:"发音"`
	ProficiencyLevel   ProficiencyLevel `json:"proficiencyLevel,omitempty" v:"required|between:1,5" dc:"熟练度"`
}

type CreateRes struct {
}

type UpdateReq struct {
	g.Meta             `path:"words" method:"put" sm:"更新" tags:"单词" `
	Id                 uint             `json:"id,omitempty" v:"required"`
	Word               string           `json:"word" v:"required|length:1,100" dc:"单词" json:"word,omitempty"`
	Definition         string           `json:"definition,omitempty" v:"required|length:1,300" dc:"单词定义"`
	ExampleSentence    string           `json:"exampleSentence,omitempty" v:"required|length:1,300" dc:"例句"`
	ChineseTranslation string           `json:"chineseTranslation,omitempty" v:"required|length:1,300" dc:"中文翻译"`
	Pronunciation      string           `json:"pronunciation,omitempty" v:"required|length:1,100" dc:"发音"`
	ProficiencyLevel   ProficiencyLevel `json:"proficiencyLevel,omitempty" v:"required|between:1,5" dc:"熟练度"`
}

type UpdateRes struct {
}

type SingleWord struct {
	Id               uint
	Word             string
	Definition       string
	ProficiencyLevel ProficiencyLevel
}

type ListReq struct {
	g.Meta `path:"words" method:"get" sm:"列表" tags:"单词"`
	Word   string `json:"word" v:"length:1,100" dc:"模糊查询单词"`
	Page   int    `json:"page" v:"min:1" dc:"页码，默认1"`
	Size   int    `json:"size" v:"between:1,100" dc:"每页数量，默认10"`
}

type ListRes struct {
	List  []SingleWord `json:"list"`
	Total int          `json:"total"`
}
