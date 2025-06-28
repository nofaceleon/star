package words

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	v1 "star2/api/words/v1"
	"star2/internal/dao"
	"star2/internal/model/do"
)

type Words struct {
}

func New() *Words {
	return &Words{}
}

type CreateInput struct {
	Uid                uint
	Word               string
	Definition         string
	ExampleSentence    string
	ChineseTranslation string
	Pronunciation      string
	ProficiencyLevel   v1.ProficiencyLevel
}

func (w *Words) Create(ctx context.Context, in CreateInput) error {
	var cls = dao.Words.Columns()

	count, err := dao.Words.Ctx(ctx).Where(cls.Uid, in.Uid).Where(cls.Word, in.Word).Count()
	if err != nil {
		return err
	}
	if count > 0 {
		return gerror.New("单词已存在")
	}

	_, err = dao.Words.Ctx(ctx).Data(do.Words{
		Uid:                in.Uid,
		Word:               in.Word,
		Definition:         in.Definition,
		ExampleSentence:    in.ExampleSentence,
		ChineseTranslation: in.ChineseTranslation,
		Pronunciation:      in.Pronunciation,
		ProficiencyLevel:   in.ProficiencyLevel,
	}).Insert()
	if err != nil {
		return err
	}
	return nil
}

type UpdateInput CreateInput

func (w *Words) Update(ctx context.Context, id int, in UpdateInput) error {

	var cls = dao.Words.Columns()

	count, err := dao.Words.Ctx(ctx).Where(cls.Uid, in.Uid).Where(cls.Word, in.Word).Count()

	if err != nil {
		return err
	}

	if count > 0 {
		return gerror.New("单词已存在")
	}

	_, err = dao.Words.Ctx(ctx).Where(cls.Id, id).Data(do.Words{
		Word:               in.Word,
		Definition:         in.Definition,
		ExampleSentence:    in.ExampleSentence,
		ChineseTranslation: in.ChineseTranslation,
		Pronunciation:      in.Pronunciation,
		ProficiencyLevel:   in.ProficiencyLevel,
	}).Update()

	return err

}
