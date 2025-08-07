package models

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
