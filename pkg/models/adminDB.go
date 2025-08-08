package models

import (
	"database/sql"
)

type Item struct {
	ItemId    int     `json:"ItemId"`
	ItemName  string  `json:"ItemName"`
	SectionId int     `json:"SectionId"`
	Price     float64 `json:"Price"`
}

func SetPaidOrder(orderId int, paid int) error {
	_, err := DB.Exec(`update Orders
                        set paid = ?
                        where OrderId = ?`, paid, orderId)
	return err
}

func SwapSections(sectionId1 int, sectionId2 int) error {
	var rows *sql.Rows
	var err error

	var sectionOrder1 int
	var sectionOrder2 int

	rows, err = DB.Query(`select SectionId, SectionOrder from Sections where SectionId = ? or SectionId = ?`, sectionId1, sectionId2)
	if err != nil {
		return err
	}

	for rows.Next() {
		var sectionId int
		var sectionOrder int
		err = rows.Scan(&sectionId, &sectionOrder)
		if err != nil {
			return err
		}

		if sectionId == sectionId1 {
			sectionOrder1 = sectionOrder
		} else {
			sectionOrder2 = sectionOrder
		}
	}

	_, err = DB.Exec(`update Sections set sectionOrder = ? where SectionId = ?`, -1, sectionId1)

	if err != nil {
		return err
	}

	_, err = DB.Exec(`update Sections set sectionOrder = ? where SectionId = ?`, sectionOrder1, sectionId2)

	if err != nil {
		return err
	}

	_, err = DB.Exec(`update Sections set sectionOrder = ? where SectionId = ?`, sectionOrder2, sectionId1)

	return err
}

func SetUserRole(userId int, role string) error {
	_, err := DB.Exec(`update Users set Role = ? where UserId = ?`, role, userId)
	return err

}

func GetNextItemID() int {
	row := DB.QueryRow(`select max(ItemId) from Items`)

	var itemId int
	var err error

	err = row.Scan(&itemId)

	if err != nil {
		return 1
	}

	return itemId + 1

}

func CreateItem(item Item) error {
	_, err := DB.Exec(`insert into Items (ItemId, ItemName, SectionId, Price) values (?, ?, ?, ?)`, item.ItemId, item.ItemName, item.SectionId, item.Price)
	return err
}
