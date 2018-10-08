package repo

import "github.com/satori/go.uuid"

type Repo interface {
	create(interface{}) error
	update(uuid.UUID, interface{}) error
	delete(uuid.UUID) error
	softDelete(uuid.UUID) error
}