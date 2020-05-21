package repo

import (
	"ganja/pkg/infra"
)

type base struct {
}

func (base) getById(id string, value interface{}) error {
	return infra.Postgresql().Where("id = ?", id).Limit(1).Find(value).Error
}

func (base) create(value interface{}) error {
	return infra.Postgresql().Create(value).Error
}

func (base) update(value interface{}) error {
	return infra.Postgresql().Updates(value).Error
}

func (base) delete(value interface{}) error {
	return infra.Postgresql().Delete(value).Error
}
