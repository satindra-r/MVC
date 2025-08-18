package models

import (
	"database/sql"
	"mvc/pkg/utils"
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

func GetNextSectionOrder() int {
	row := DB.QueryRow(`select max(SectionOrder) from Sections`)
	var sectionOrder int
	var err error

	err = row.Scan(&sectionOrder)

	if err != nil {
		return 1
	}

	return sectionOrder + 1
}

func CreateSection(section Section) error {
	_, err := DB.Exec(`insert into Sections(SectionOrder,SectionName) value (?, ?)`, section.SectionOrder, section.SectionName)
	return err
}

func SwapSections(sectionId1 int, sectionId2 int) error {
	var rows *sql.Rows
	var err error

	var sectionOrder1 int
	var sectionOrder2 int

	tx, err := DB.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	rows, err = tx.Query(`select SectionId, SectionOrder from Sections where SectionId = ? or SectionId = ?`, sectionId1, sectionId2)
	if err != nil {
		return err
	}
	defer rows.Close()

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

	_, err = tx.Exec(`update Sections set sectionOrder = ? where SectionId = ?`, -1, sectionId1)

	if err != nil {
		return err
	}

	_, err = tx.Exec(`update Sections set sectionOrder = ? where SectionId = ?`, sectionOrder1, sectionId2)

	if err != nil {
		return err
	}

	_, err = tx.Exec(`update Sections set sectionOrder = ? where SectionId = ?`, sectionOrder2, sectionId1)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func DeleteSection(sectionId int) error {
	tx, err := DB.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	if sectionId == 1 {
		_, err := tx.Exec(`update Items set SectionId = 2 where SectionId = ?`, sectionId)
		if err != nil {
			return err
		}

	} else {
		_, err := tx.Exec(`update Items set SectionId = 1 where SectionId = ?`, sectionId)
		if err != nil {
			return err
		}

	}
	_, err = tx.Exec(`delete from Sections where SectionId = ?`, sectionId)
	if err != nil {
		return err
	}
	return tx.Commit()
}

func SetUserRole(userId int, role string) error {
	_, err := DB.Exec(`update Users set Role = ? where UserId = ?`, role, userId)
	return err

}

func CreateItem(item Item) error {
	_, err := DB.Exec(`insert into Items (ItemName, SectionId, Price) values (?, ?, ?)`, item.ItemName, item.SectionId, item.Price)
	return err
}

func EditItem(item Item) error {

	if len(item.ItemName) > 0 {
		_, err := DB.Exec(`update Items set ItemName = ? where ItemId = ?`, item.ItemName, item.ItemId)
		if err != nil {
			return err
		}
	}
	if item.SectionId > 0 {
		_, err := DB.Exec(`update Items set SectionId = ? where ItemId = ?`, item.SectionId, item.ItemId)
		if err != nil {
			return err
		}
	}
	if item.Price >= 0 {
		_, err := DB.Exec(`update Items set Price = ? where ItemId = ?`, item.Price, item.ItemId)
		if err != nil {
			return err
		}
	}
	return nil
}

func GetUsers(page int) []User {
	rows, err := DB.Query(`
		select UserId, UserName, Role from Users limit 10 offset ?`, (page-1)*10)
	if utils.LogIfErr(err, "DB error") {
		return nil
	}
	defer rows.Close()

	var users []User

	for rows.Next() {
		var user User
		err = rows.Scan(&user.UserId, &user.UserName, &user.Role)
		if utils.LogIfErr(err, "DB error") {
			return nil
		}
		users = append(users, user)
	}
	return users
}
