package repo

import (
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
	"ntm/db"
	"ntm/model"
)

type TopicRepoInterface interface {
	create(interface{}) error
	update(uuid.UUID, interface{}) error
	delete(uuid.UUID) error
	softDelete(uuid.UUID) error
}

type BaseTopicRepo struct {}

func NewBaseTopicRepo() *BaseTopicRepo {
	return new(BaseTopicRepo)
}

func (b *BaseTopicRepo) orm() *gorm.DB {
	return db.NewORM().Connect()
}

func (b *BaseTopicRepo) create(v interface{}) error {
	return b.orm().Model(model.NewTopic()).Create(v).Error
}

func (b *BaseTopicRepo) update(id uuid.UUID, v interface{}) error {
	return b.orm().Model(model.NewTopic()).Where("id = ?", id).Update(v).Error
}

func (b *BaseTopicRepo) delete(id uuid.UUID) error {
	return b.orm().Model(model.NewTopic()).Unscoped().Where("id = ?", id).Delete(model.NewTopic()).Error
}

func (b *BaseTopicRepo) softDelete(id uuid.UUID) error {
	return b.orm().Model(model.NewTopic()).Where("id = ?", id).Delete(model.NewTopic()).Error
}

type TopicRepo struct { use TopicRepoInterface }

func NewTopicRepo() *TopicRepo {
	return &TopicRepo{use: NewBaseTopicRepo()}
}

func (t *TopicRepo) ORM() *gorm.DB {
	return t.use.(*BaseTopicRepo).orm()
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