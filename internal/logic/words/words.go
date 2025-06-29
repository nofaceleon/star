package words

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	v1 "star2/api/words/v1"
	"star2/internal/dao"
	"star2/internal/model/do"
	"star2/internal/model/entity"
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

// Update 更新单词
func (w *Words) Update(ctx context.Context, id uint, in UpdateInput) error {

	var cls = dao.Words.Columns() //获取这个表中的字段名
	// 用户只能更新自己的单词
	count, err := dao.Words.Ctx(ctx).Where(cls.Uid, in.Uid).Where(cls.Id, id).Count()

	if err != nil {
		return err
	}

	if count == 0 {
		return gerror.New("单词不存在")
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

type ListInput struct {
	Uid  uint
	Word string
	Page int
	Size int
}

func (w *Words) List(ctx context.Context, in ListInput) (list []entity.Words, total int, err error) {
	if in.Page <= 0 {
		in.Page = 1
	}

	if in.Size <= 0 || in.Size > 100 {
		in.Size = 10
	}

	var cls = dao.Words.Columns()
	var orm = dao.Words.Ctx(ctx)

	if in.Uid > 0 {
		orm = orm.Where(cls.Uid, in.Uid)
	}

	if len(in.Word) != 0 {
		orm = orm.WhereLike(cls.Word, "%"+in.Word+"%")
	}

	orm = orm.OrderDesc(cls.Id).OrderDesc(cls.CreatedAt).Page(in.Page, in.Size)
	err = orm.ScanAndCount(&list, &total, true)
	if err != nil {
		return nil, 0, err
	}
	return
}

func (w *Words) Detail(ctx context.Context, id uint, uid uint) (word *entity.Words, err error) {
	var cls = dao.Words.Columns()
	var orm = dao.Words.Ctx(ctx)

	orm = orm.Where(cls.Id, id)
	if uid > 0 {
		orm = orm.Where(cls.Uid, uid)
	}
	err = orm.Scan(&word)
	return
}

func (w *Words) Delete(ctx context.Context, id uint, uid uint) (err error) {
	var cls = dao.Words.Columns()
	var orm = dao.Words.Ctx(ctx)

	orm = orm.Where(cls.Id, id)
	if uid > 0 {
		orm = orm.Where(cls.Uid, uid)
	}
	_, err = orm.Delete()
	return
}

func (w *Words) RandList(ctx context.Context, limit int, uid uint) (list []entity.Words, err error) {
	if limit <= 0 {
		limit = 50
	}
	var cls = dao.Words.Columns()
	var orm = dao.Words.Ctx(ctx)
	if uid > 0 {
		orm = orm.Where(cls.Uid, uid)
	}
	err = orm.Limit(limit).OrderRandom().Scan(&list)
	if err != nil {
		return nil, err
	}
	return
}
