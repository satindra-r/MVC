package models

import "mvc/pkg/utils"

func GetUserRole(userId int) (string, error) {
	row := DB.QueryRow(`select Role from Users where UserId = ?`, userId)

	var role string
	var err error

	err = row.Scan(&role)
	return role, err
}

func SetPreparedDish(dishId int, prepared int) error {
	_, err := DB.Exec(`update Dishes
                        set Prepared = ?
                        where DishId = ?`, prepared, dishId)
	return err
}

func GetAllOrders(page int) []OrderDishes {
	var orders []OrderDishes
	rows, err := DB.Query(`select Orders.OrderId, Price, Paid, round(100*sum(Prepared*DishCount)/sum(DishCount),2)
 from Orders,Dishes where Dishes.OrderId = Orders.OrderId group by Orders.OrderId order by OrderId limit 10
                              offset ?`, (page-1)*10)
	if utils.LogIfErr(err, "DB error") {
		return nil
	}

	for rows.Next() {
		var order OrderDishes
		err = rows.Scan(&order.OrderId, &order.Price, &order.Paid, &order.Progress)
		if utils.LogIfErr(err, "DB error") {
			return nil
		}
		dishRows, err := DB.Query(`select DishId,DishCount,SplInstructions,Prepared,ItemName from Items,Dishes where Items.ItemId = Dishes.ItemId and OrderId = ?`, order.OrderId)
		if utils.LogIfErr(err, "DB error") {
			return nil
		}

		for dishRows.Next() {
			var dish DishItem
			err = dishRows.Scan(&dish.DishId, &dish.DishCount, &dish.SplInstructions, &dish.Prepared, &dish.ItemName)
			if utils.LogIfErr(err, "DB error") {
				return nil
			}
			order.Dishes = append(order.Dishes, dish)
		}
		orders = append(orders, order)
	}
	return orders
}
