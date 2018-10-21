//Author Mohammad Naser Abbasanadi
//Creating Date 2018-10-20
// order.go is to connect to database and manipuate data accordingly

package models

import (
	"GolangOrdering/logger"
	"database/sql"
)

//Order is represent the structure of order in database
type Order struct {
	ID       int    `json:"id"`
	Distance int    `json:"distance"`
	Status   string `json:"status"`
}

//UpdateOrder is allow to update the status of orderids
func (p *Order) UpdateOrder(db *sql.DB) (int, error) {

	cnt := -1
	err := db.QueryRow("select  count(iOrderid) from orderinfo where iOrderid=?", p.ID).Scan(&cnt)
	if cnt == 0 {
		return -1, nil
	}

	if err != nil {
		return 0, err
	}

	stmt, err := db.Prepare("UPDATE orderinfo SET vStatus=? WHERE iOrderid=? and vStatus=?")
	if err != nil {
		logger.Log.Println("Cannot run update statement", err)
		return 0, err
	}

	res, err := stmt.Exec(p.Status, p.ID, "UNASSIGN")
	if err != nil {
		logger.Log.Println("Cannot run update statement", err)
		return 0, err
	}

	affect, err := res.RowsAffected()
	if err != nil {
		logger.Log.Println("Cannot run update statement", err)
		return 0, err
	}

	return int(affect), nil
}

//CreateOrder is responsable to insert new order to database
func (p *Order) CreateOrder(db *sql.DB) (int, error) {
	stmt, err := db.Prepare("INSERT INTO orderinfo(vStatus, iDistance) VALUES(?, ?)")
	if err != nil {
		logger.Log.Println("Cannot prepare DB statement", err)
		return 0, err
	}

	res, err := stmt.Exec(p.Status, p.Distance)
	if err != nil {
		logger.Log.Println("Cannot run insert statement", err)
		return 0, err
	}

	id, _ := res.LastInsertId()

	return int(id), nil

}

//GetOrders is responsable to retrieve all orders based on offest provided by user
func GetOrders(db *sql.DB, page, limit int) ([]Order, error) {
	rows, err := db.Query(
		"SELECT iOrderid , vStatus , iDistance  FROM orderinfo LIMIT ? , ?",
		page, limit)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	orders := []Order{}

	for rows.Next() {
		var p Order
		if err := rows.Scan(&p.ID, &p.Status, &p.Distance); err != nil {
			return nil, err
		}
		orders = append(orders, p)
	}

	return orders, nil
}
