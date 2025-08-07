package models

import (
	"database/sql"
	"mvc/pkg/utils"
)

type User struct {
	UserId   int    `json:"UserId"`
	UserName string `json:"UserName"`
	Role     string `json:"Role"`
	PhoneNo  string `json:"PhoneNo"`
	Address  string `json:"Address"`
	Hash     string `json:"Hash"`
}

type Order struct {
	OrderId int `json:"OrderId"`
	UserId  int `json:"UserId"`
	Price   int `json:"Price"`
	Paid    int `json:"Paid"`
}

type Dish struct {
	DishId          int    `json:"DishId"`
	ItemId          int    `json:"ItemId"`
	OrderId         int    `json:"OrderId"`
	DishCount       int    `json:"DishCount"`
	SplInstructions string `json:"SplInstructions"`
	Prepared        int    `json:"Prepared"`
}

func GetNextUserID() int {
	row := DB.QueryRow(`select max(UserId) from Users`)

	var userId int
	var err error

	err = row.Scan(&userId)

	if err != nil {
		return 1
	}

	return userId + 1

}

func CreateUser(user User) error {
	_, err := DB.Exec(`insert into Users(UserId, UserName, Role, PhoneNo, Address, Hash) value (?, ?, ?, ?, ?, ?)`, user.UserId, user.UserName, user.Role, user.PhoneNo, user.Address, user.Hash)
	return err
}

func GetUserCredentials(userName string) (string, int) {
	var hash string
	var userId int
	var err error
	var row = DB.QueryRow(`select Hash,UserId from Users where UserName = ?`, userName)
	err = row.Scan(&hash, &userId)
	if err != nil {
		return "", -1
	}
	return hash, userId
}

func GetNextDishID() int {
	row := DB.QueryRow(`select max(DishId) from Dishes`)

	var dishId int
	var err error
	err = row.Scan(&dishId)
	if err != nil {
		return 1
	}
	return dishId + 1
}

func GetNextOrderID() int {
	row := DB.QueryRow(`select max(OrderId) from Orders`)

	var orderId int
	var err error
	err = row.Scan(&orderId)
	if err != nil {
		return 1
	}
	return orderId + 1
}

func GetItemPrices() ([]int, error) {
	var rows *sql.Rows
	var err error

	rows, err = DB.Query(`select Items.Price from Items`)
	if utils.LogIfErr(err, "DB error") {
		return nil, err
	}
	var prices []int

	for rows.Next() {
		var price int
		err = rows.Scan(&price)
		if utils.LogIfErr(err, "DB error") {
			return nil, err
		}
		prices = append(prices, price)
	}

	return prices, nil
}

func CreateDish(dish Dish) error {
	_, err := DB.Exec(`insert into Dishes(DishId, ItemId, OrderId, DishCount, SplInstructions, Prepared) value (?, ?, ?, ?, ?, 0)`, dish.DishId, dish.ItemId, dish.OrderId, dish.DishCount, dish.SplInstructions)
	return err
}

func CreateOrder(order Order) error {
	_, err := DB.Exec(`insert into Orders(OrderId, UserId, Price, Paid) value (?, ?, ?, 0)`, order.OrderId, order.UserId, order.Price)
	return err
}
