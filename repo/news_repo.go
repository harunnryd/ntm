package repo

import (
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
	"ntm/db"
	"ntm/model"
)

type GormNewsRepo struct {}

func NewGormNewsRepo() *GormNewsRepo {
	return new(GormNewsRepo)
}

func (b *GormNewsRepo) orm() *gorm.DB {
	return db.NewORM().Connect().LogMode(true)
}

func (b *GormNewsRepo) create(v interface{}) error {
	return b.orm().Model(model.NewNews()).Create(v).Error
}

func (b *GormNewsRepo) update(id uuid.UUID, v interface{}) error {
	return b.orm().Model(model.NewNews()).Where("id = ?", id).Update(v).Error
}

func (b *GormNewsRepo) delete(id uuid.UUID) error {
	return b.orm().Model(model.NewNews()).Unscoped().Where("id = ?", id).Delete(model.NewNews()).Error
}

func (b *GormNewsRepo) softDelete(id uuid.UUID) error {
	return b.orm().Model(model.NewNews()).Where("id = ?", id).Delete(model.NewNews()).Error
}

type NewsRepo struct { use Repo }

func NewNewsRepo() *NewsRepo {
	return &NewsRepo{use: NewGormNewsRepo()}
}

func (t *NewsRepo) ORM() *gorm.DB {
	return t.use.(*GormNewsRepo).orm()
}

func (t *NewsRepo) Create(v interface{}) error {
	return t.use.create(v)
}

func (t *NewsRepo) Update(id uuid.UUID, v interface{}) error {
	return t.use.update(id, v)
}

func (t *NewsRepo) Delete(id uuid.UUID) error {
	return t.use.delete(id)
}

func (t *NewsRepo) SoftDelete(id uuid.UUID) error {
	return t.use.softDelete(id)
}