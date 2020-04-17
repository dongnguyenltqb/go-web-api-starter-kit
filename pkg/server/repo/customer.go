package repo

import (
	"context"
	"ganja/pkg/infra"
	"ganja/pkg/server/entity"

	"go.mongodb.org/mongo-driver/bson"
)

type CustomerRepoInterface interface {
	GetById(string) (*entity.Customer, error)
}

type customerRepo struct {
}

func GetCustomerRepo() CustomerRepoInterface {
	return &customerRepo{}
}

func (*customerRepo) GetById(id string) (c *entity.Customer, err error) {
	customerCollection := infra.GetDB().Collection("customers")
	var cus entity.Customer
	if err = customerCollection.FindOne(context.Background(), bson.M{
		"customer_id": id,
	}).Decode(&cus); err != nil {
		return
	}
	c = &cus
	return
}
