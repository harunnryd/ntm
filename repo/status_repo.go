package repo

import (
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
	"ntm/db"
	"ntm/model"
)

type GormStatusRepo struct {}

func NewGormStatusRepo() *GormStatusRepo {
	return new(GormStatusRepo)
}

func (b *GormStatusRepo) orm() *gorm.DB {
	return db.NewORM().Connect().LogMode(true)
}

func (b *GormStatusRepo) create(v interface{}) error {
	return b.orm().Model(model.NewStatus()).Create(v).Error
}

func (b *GormStatusRepo) update(id uuid.UUID, v interface{}) error {
	return b.orm().Model(model.NewStatus()).Where("id = ?", id).Update(v).Error
}

func (b *GormStatusRepo) delete(id uuid.UUID) error {
	return b.orm().Model(model.NewStatus()).Unscoped().Where("id = ?", id).Delete(model.NewStatus()).Error
}

func (b *GormStatusRepo) softDelete(id uuid.UUID) error {
	return b.orm().Model(model.NewStatus()).Where("id = ?", id).Delete(model.NewStatus()).Error
}

type StatusRepo struct { use Repo }

func NewStatusRepo() *StatusRepo {
	return &StatusRepo{use: NewGormStatusRepo()}
}

func (t *StatusRepo) ORM() *gorm.DB {
	return t.use.(*GormStatusRepo).orm()
}

func (t *StatusRepo) Create(v interface{}) error {
	return t.use.create(v)
}

func (t *StatusRepo) Update(id uuid.UUID, v interface{}) error {
	return t.use.update(id, v)
}

func (t *StatusRepo) Delete(id uuid.UUID) error {
	return t.use.delete(id)
}

func (t *StatusRepo) SoftDelete(id uuid.UUID) error {
	return t.use.softDelete(id)
}