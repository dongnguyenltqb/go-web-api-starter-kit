package repo

import (
	"ganja/pkg/server/entity"
)

type CustomerRepoInterface interface {
	GetById(string) (*entity.Customer, error)
}

type customerRepo struct {
	base
}

func GetCustomerRepo() CustomerRepoInterface {
	return &customerRepo{}
}

func (r *customerRepo) GetById(id string) (*entity.Customer, error) {
	// customerCollection := infra.GetDB().Collection("customers")
	// var cus entity.Customer
	// if err = customerCollection.FindOne(context.Background(), bson.M{
	// 	"customer_id": id,
	// }).Decode(&cus); err != nil {
	// 	return
	// }
	// c = &cus
	// return
	c := &entity.Customer{}
	var err error
	err = r.getById(id, c)
	return c, err
}
