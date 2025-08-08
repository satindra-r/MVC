package models

import (
	"database/sql"
	"math"
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
	OrderId int     `json:"OrderId"`
	UserId  int     `json:"UserId"`
	Price   float64 `json:"Price"`
	Paid    int     `json:"Paid"`
}

type Dish struct {
	DishId          int    `json:"DishId"`
	ItemId          int    `json:"ItemId"`
	OrderId         int    `json:"OrderId"`
	DishCount       int    `json:"DishCount"`
	SplInstructions string `json:"SplInstructions"`
	Progress        int    `json:"Prepared"`
}

type Section struct {
	SectionId    int    `json:"SectionId"`
	SectionName  string `json:"SectionName"`
	SectionOrder int    `json:"SectionOrder"`
	Colour       int    `json:"Colour"`
}

type SectionedItems struct {
	ItemId      int     `json:"ItemId"`
	ItemName    string  `json:"ItemName"`
	SectionId   int     `json:"SectionId"`
	Price       float64 `json:"Price"`
	SectionName string  `json:"SectionName"`
	Colour      int     `json:"Colour"`
}

type OrderDishes struct {
	OrderId  int        `json:"OrderId"`
	UserId   int        `json:"UserId"`
	Price    float64    `json:"Price"`
	Paid     int        `json:"Paid"`
	Progress float64    `json:"Progress"`
	Dishes   []DishItem `json:"Dishes"`
}

type DishItem struct {
	DishId          int    `json:"DishId"`
	DishCount       int    `json:"DishCount"`
	SplInstructions string `json:"SplInstructions"`
	Prepared        int    `json:"Prepared"`
	ItemName        string `json:"ItemName"`
	ItemPrice       string `json:"ItemPrice"`
	Price           string `json:"Price"`
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

func GetItemPrices() ([]float64, error) {
	var rows *sql.Rows
	var err error

	rows, err = DB.Query(`select Items.Price from Items`)
	if utils.LogIfErr(err, "DB error") {
		return nil, err
	}
	var prices []float64

	for rows.Next() {
		var price float64
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

func GetItems(page int, filters int) []SectionedItems {
	var rows *sql.Rows
	var err error
	if filters > 0 {
		var i = 1
		var filtersList []any
		for filters > 0 {
			if filters%2 == 1 {
				filtersList = append(filtersList, i)
			}
			filters >>= 1
			i++
		}
		var questionMarks = "("
		for i := range len(filtersList) {
			if i > 0 {
				questionMarks += ","
			}
			questionMarks += "?"
		}
		questionMarks += ")"

		rows, err = DB.Query(`select ItemId, ItemName, Items.SectionId, Price,SectionName from Items,Sections where Items.SectionId=Sections.SectionId and Items.SectionId in `+questionMarks+` order by SectionId, ItemId limit 10 offset ?`, append(filtersList, (page-1)*10)...)
		if utils.LogIfErr(err, "DB error") {
			return nil
		}
	} else {
		rows, err = DB.Query(`select ItemId, ItemName, Items.SectionId, Price,SectionName from Items,Sections where Items.SectionId=Sections.SectionId order by SectionId, ItemId limit 10 offset ?`, (page-1)*10)
		if utils.LogIfErr(err, "DB error") {
			return nil
		}
	}

	var items []SectionedItems

	var phiInv = (math.Sqrt(5) - 1) / 2
	for rows.Next() {
		var item SectionedItems
		err = rows.Scan(&item.ItemId, &item.ItemName, &item.SectionId, &item.Price, &item.SectionName)
		_, hue := math.Modf(phiInv * (float64)(item.SectionId))
		item.Colour = (int)(360.0 * hue)
		if utils.LogIfErr(err, "DB error") {
			return nil
		}
		items = append(items, item)
	}
	return items
}

func GetSections() []Section {
	rows, err := DB.Query(`
		select SectionId, SectionName, SectionOrder
			from
			Sections order by SectionOrder
			`)
	if utils.LogIfErr(err, "DB error") {
		return nil
	}
	var sections []Section
	var phiInv = (math.Sqrt(5) - 1) / 2
	for rows.Next() {
		var section Section
		err = rows.Scan(&section.SectionId, &section.SectionName, &section.SectionOrder)
		_, hue := math.Modf(phiInv * (float64)(section.SectionId))
		section.Colour = (int)(360.0 * hue)
		if utils.LogIfErr(err, "DB error") {
			return nil
		}
		sections = append(sections, section)
	}
	return sections
}

func GetUserOrders(userId int, page int) []OrderDishes {
	var orders []OrderDishes
	rows, err := DB.Query(`select Orders.OrderId, Price, Paid,round(100*sum(Prepared*DishCount)/sum(DishCount),2) from Orders,Dishes where Orders.OrderId=Dishes.OrderId and UserId = ? group by OrderId order by OrderId limit 10
                              offset ?`, userId, (page-1)*10)
	if utils.LogIfErr(err, "DB error") {
		return nil
	}

	for rows.Next() {
		var order OrderDishes
		err = rows.Scan(&order.OrderId, &order.Price, &order.Paid, &order.Progress)
		if utils.LogIfErr(err, "DB error") {
			return nil
		}
		dishRows, err := DB.Query(`select DishCount,SplInstructions,Prepared,ItemName from Items,Dishes where Items.ItemId = Dishes.ItemId and OrderId = ?`, order.OrderId)
		if utils.LogIfErr(err, "DB error") {
			return nil
		}

		for dishRows.Next() {
			var dish DishItem
			err = dishRows.Scan(&dish.DishCount, &dish.SplInstructions, &dish.Prepared, &dish.ItemName)
			if utils.LogIfErr(err, "DB error") {
				return nil
			}
			order.Dishes = append(order.Dishes, dish)
		}
		orders = append(orders, order)
	}
	return orders
}

func GetUserOrder(orderId int) (OrderDishes, error) {
	var order OrderDishes
	row := DB.QueryRow(`select
	OrderId,
	UserId,
	Price,
	Paid 
from
	Orders
where Orders.OrderId = ?`, orderId)

	err := row.Scan(&order.OrderId, &order.UserId, &order.Price, &order.Paid)
	if utils.LogIfErr(err, "DB error") {
		return OrderDishes{}, err
	}
	dishRows, err := DB.Query(`select DishCount,SplInstructions,Prepared,ItemName,Price,(Price*DishCount) from Items,Dishes where Items.ItemId = Dishes.ItemId and OrderId = ?`, order.OrderId)
	if utils.LogIfErr(err, "DB error") {
		return OrderDishes{}, err
	}

	for dishRows.Next() {
		var dish DishItem
		err = dishRows.Scan(&dish.DishCount, &dish.SplInstructions, &dish.Prepared, &dish.ItemName, &dish.ItemPrice, &dish.Price)
		if utils.LogIfErr(err, "DB error") {
			return OrderDishes{}, err
		}
		order.Dishes = append(order.Dishes, dish)
	}

	return order, nil
}
