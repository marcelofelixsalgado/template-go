package entity_test

import (
	"marcelofelixsalgado/template-go/pkg/domain/customer/entity"
	"testing"
)

func TestCustomerSuccess(t *testing.T) {

	customerId := "8c25d3f4-026a-4435-b8f6-3ed64d7162b4"
	customerName := "John Doe"

	customer, err := entity.NewCustomer(customerId, customerName)

	if err != nil {
		t.Fatalf(err.Error())
	}

	if customer.GetId() != customerId {
		t.Fatalf("Customer Id [%s] is different from expected [%s]", customer.GetId(), customerId)
	}

	if customer.GetName() != customerName {
		t.Fatalf("Customer Name [%s] is different from expected [%s]", customer.GetName(), customerName)
	}
}

func TestCustomerEmptyId(t *testing.T) {

	_, err := entity.NewCustomer("", "John Doe")

	if err == nil {
		t.Fatalf("Should throw error when id is empty")
	}

	if err.Error() != "id is required" {
		t.Fatalf("Error different from expected [%s]", err.Error())
	}
}

func TestCustomerEmptyName(t *testing.T) {

	_, err := entity.NewCustomer("8c25d3f4-026a-4435-b8f6-3ed64d7162b4", "")

	if err == nil {
		t.Fatalf("Should throw error when id is empty")
	}

	if err.Error() != "name is required" {
		t.Fatalf("Error different from expected [%s]", err.Error())
	}
}

func TestCustomerChangeSuccess(t *testing.T) {

	customer, err := entity.NewCustomer("8c25d3f4-026a-4435-b8f6-3ed64d7162b4", "John Doe")

	if err != nil {
		t.Fatalf(err.Error())
	}

	customerNewName := "Changed Name"
	customer.SetName(customerNewName)

	if customer.GetName() != customerNewName {
		t.Fatalf("Customer Name [%s] is different from expected [%s]", customer.GetName(), customerNewName)
	}
}
