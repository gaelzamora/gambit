package bd

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"

	//	"strconv"
	"github.com/gaelzamora/gambit/models"
	"github.com/gaelzamora/gambit/tools"
	_ "github.com/go-sql-driver/mysql"
	//	"github.com/gaelzamora/gambit/tools"
)

func InsertCategory(c models.Category) (int64, error) {
	fmt.Println("Comienza registro de InsertCategory")

	err := DbConnect()
	if err != nil {
		return 0, err
	}
	defer Db.Close()

	fmt.Println("Hola")

	sentence := "INSERT INTO category (Categ_Name, Categ_Path) VALUES ('" + c.CategName + "','" + c.CategPath + "')"

	var result sql.Result

	result, err = Db.Exec(sentence)
	if err != nil {
		fmt.Println(err.Error())
		return 0, err
	}

	LastInsertId, err2 := result.LastInsertId()

	if err2 != nil {
		return 0, err2
	}

	fmt.Println("Insert Category > Ejecucion exitosa")
	return LastInsertId, nil
}

func UpdateCategory(c models.Category) error {
	fmt.Println("Comienza registro de UpdateCategory")

	err := DbConnect()
	if err != nil {
		return err
	}
	defer Db.Close()

	fmt.Println("Hola")

	sentence := "UPDATE category SET "

	if len(c.CategName)>0 {
		sentence += " Categ_Name = '" + tools.EscapeString(c.CategName) + "'"
	}

	if len(c.CategPath)>0 {
		if !strings.HasSuffix(sentence, "SET ") {
			sentence += ", "
		}
		sentence += "Categ_Path = '" + tools.EscapeString(c.CategPath) + "'"
	}

	sentence += " WHERE Categ_Id = " + strconv.Itoa(c.CategID)

	_, err = Db.Exec(sentence)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("Update Category > Ejecucion Exitosa")
	return nil
}

func DeleteCategory(id int) error {
	fmt.Println("Comienza registro de UpdateCategory")

	err := DbConnect()
	if err != nil {
		return err
	}
	defer Db.Close()

	fmt.Println("Hola")

	sentence := "DELETE FROM category WHERE Categ_Id = " + strconv.Itoa(id)

	_, err = Db.Exec(sentence)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("Delete Category > Ejecucion Exitosa")
	return nil
}

func SelectCategories(CategId int, Slug string) ([]models.Category, error) {
	fmt.Println("Comienza Select Categories")

	var Categ []models.Category

	err := DbConnect()
	if err != nil{
		return Categ, err
	}
	defer Db.Close()

	sentence := "SELECT Categ_Id, Categ_Name, Categ_Path FROM category "

	if CategId > 0 {
		sentence += "WHERE Categ_Id = "+strconv.Itoa(CategId)
	} else {
		if len(Slug) > 0 {
			sentence += "WHERE Categ_Path LIKE '%" + Slug + "%'"
		}
	}

	fmt.Println(sentence)

	var rows *sql.Rows
	rows, _ = Db.Query(sentence)

	for rows.Next() {
		var c models.Category
		var categId sql.NullInt32
		var categName sql.NullString
		var categPath sql.NullString

		err := rows.Scan(&categId, &categName, &categPath)

		if err != nil {
			return Categ, err
		}

		c.CategID = int(categId.Int32)
		c.CategName = categName.String
		c.CategPath = categPath.String

		Categ = append(Categ, c)
	}

	fmt.Println("Select Category > Ejecucion Exitosa")

	return Categ, nil
}