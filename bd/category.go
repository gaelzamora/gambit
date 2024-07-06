package bd

import (
	"database/sql"
	"fmt"
//	"strconv"
//	"strings"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gaelzamora/gambit/models"
//	"github.com/gaelzamora/gambit/tools"
)

func InsertCategory(c models.Category) (int64, error) {
	fmt.Println("Comienza registro de InsertCategory")

	err := DbConnect()
	if err != nil {
		return 0, err
	}
	defer Db.Close()

	sentence := "INSERT INTO category(Categ_Name, Categ_Path) VALUES ('" + c.CategName + "','" + c.CategPath + "')"

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