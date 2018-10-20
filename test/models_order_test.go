package test

import (
	"GolangOrdering/models"
	"database/sql"
	"testing"
)

func getdb() *sql.DB {
	db, _ := models.GetDBInstance(
		"joe",
		"U_Xz$M,(3${gpTcZ",
		"185.159.153.22",
		"3306",
		"orders")
	return db
}

func TestModelGetDBInstance(t *testing.T) {
	_, err := models.GetDBInstance(
		"mohamad",
		"123456",
		"localhost",
		"3305",
		"orders")

	ok(t, err)
}

func TestModelUpdateOrder(t *testing.T) {
	order := models.Order{
		ID:     1,
		Status: "taken",
	}
	_, err := order.UpdateOrder(getdb())
	ok(t, err)
}

func TestModelCreateOrder_1(t *testing.T) {
	order := models.Order{
		Status:   "UNASSIGN",
		Distance: 10,
	}

	_, err := order.CreateOrder(getdb())
	ok(t, err)
}

func TestModelGetOrders_3(t *testing.T) {

	_, err := models.GetOrders(getdb(), 1, 10)
	ok(t, err)
}

func TestModelGetOrders_2(t *testing.T) {

	orders, _ := models.GetOrders(getdb(), 10, 0)
	equals(t, 0, len(orders))
}
