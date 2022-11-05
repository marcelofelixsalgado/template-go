package entity

import "errors"

type ICustomer interface {
	GetId() string
	GetName() string
	SetName(name string)
	toString() string
	validate() error
}

type Customer struct {
	id   string
	name string
}

func (customer Customer) GetId() string {
	return customer.id
}

func (customer Customer) GetName() string {
	return customer.name
}

func (customer *Customer) SetName(name string) {
	customer.name = name
}

func (customer Customer) toString() string {
	return "Customer Id: [" + customer.id + "] - Name: [" + customer.name + "]"
}

func (customer Customer) validate() error {

	if customer.id == "" {
		return errors.New("id is required")
	}
	if customer.name == "" {
		return errors.New("name is required")
	}
	return nil
}

func NewCustomer(id string, name string) (ICustomer, error) {
	customer := Customer{
		id:   id,
		name: name,
	}
	err := customer.validate()
	return &customer, err
}
