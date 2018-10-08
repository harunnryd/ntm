package repo

import (
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
	"ntm/db"
	"ntm/model"
)

type GormTopicRepo struct {}

func NewGormTopicRepo() *GormTopicRepo {
	return new(GormTopicRepo)
}

func (b *GormTopicRepo) orm() *gorm.DB {
	return db.NewORM().Connect().LogMode(true)
}

func (b *GormTopicRepo) create(v interface{}) error {
	return b.orm().Model(model.NewTopic()).Create(v).Error
}

func (b *GormTopicRepo) update(id uuid.UUID, v interface{}) error {
	return b.orm().Model(model.NewTopic()).Where("id = ?", id).Update(v).Error
}

func (b *GormTopicRepo) delete(id uuid.UUID) error {
	return b.orm().Model(model.NewTopic()).Unscoped().Where("id = ?", id).Delete(model.NewTopic()).Error
}

func (b *GormTopicRepo) softDelete(id uuid.UUID) error {
	return b.orm().Model(model.NewTopic()).Where("id = ?", id).Delete(model.NewTopic()).Error
}

type TopicRepo struct { use Repo }

func NewTopicRepo() *TopicRepo {
	return &TopicRepo{use: NewGormTopicRepo()}
}

func (t *TopicRepo) ORM() *gorm.DB {
	return t.use.(*GormTopicRepo).orm()
}

func (t *TopicRepo) Create(v interface{}) error {
	return t.use.create(v)
}

func (t *TopicRepo) Update(id uuid.UUID, v interface{}) error {
	return t.use.update(id, v)
}

func (t *TopicRepo) Delete(id uuid.UUID) error {
	return t.use.delete(id)
}

func (t *TopicRepo) SoftDelete(id uuid.UUID) error {
	return t.use.softDelete(id)
}