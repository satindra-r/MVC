package models

func SetPaidOrder(orderId int, paid int) error {
	_, err := DB.Exec(`update Orders
                        set paid = ?
                        where OrderId = ?`, paid, orderId)
	return err
}
