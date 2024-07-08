package bd

import (
	"database/sql"
	"fmt"
	"strconv"
	"github.com/gaelzamora/gambit/models"
	"github.com/gaelzamora/gambit/tools"
	_ "github.com/go-sql-driver/mysql"
)

func InsertProduct(p models.Product) (int64, error) {
	fmt.Println("Comienza Registro")

	err := DbConnect()

	if err != nil {
		return 0, err
	}
	defer Db.Close()

	sentence := "INSERT INTO products (Prod_Title "

	if len(p.ProdDescription) > 0 {
		sentence += ", Prod_Description"
	}
	if p.ProdPrice > 0 {
		sentence += ", Prod_Price"
	}
	if p.ProdCategId > 0 {
		sentence += ", Prod_CategoryId"
	}
	if p.ProdStock > 0 {
		sentence += ", Prod_Stock"
	}
	if len(p.ProdPath) > 0 {
		sentence += ", Prod_Path"
	}

	sentence += ") VALUES ('" + tools.EscapeString(p.ProdTitle) + "'"

	if len(p.ProdDescription) > 0 {
		sentence += ",'" + tools.EscapeString(p.ProdDescription) + "'"
	}
	if p.ProdPrice > 0 {
		sentence += ", " + strconv.FormatFloat(p.ProdPrice, 'e', -1, 64)
	}
	if p.ProdCategId > 0 {
		sentence += ", " + strconv.Itoa(p.ProdCategId)
	}
	if p.ProdStock > 0 {
		sentence += ", " + strconv.Itoa(p.ProdStock)
	}
	if len(p.ProdPath) > 0 {
		sentence += ", '" + tools.EscapeString(p.ProdPath) + "'"
	}

	sentence += ")"	

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

	fmt.Println("Insert Product > Ejecucion exitosa")
	return LastInsertId, nil
}

func UpdateProduct(p models.Product) error {
	fmt.Println("Comienza UPDATE")

	err := DbConnect()
	if err != nil {
		return err
	}

	defer Db.Close()

	sentence := "Update products SET "

	
	sentence = tools.ArmoSentencia(sentence, "Prod_Title", "S", 0, 0, p.ProdTitle)
	sentence = tools.ArmoSentencia(sentence, "Prod_Description", "S", 0, 0, p.ProdDescription)
	sentence = tools.ArmoSentencia(sentence, "Prod_Price", "F", 0, p.ProdPrice, "")
	sentence = tools.ArmoSentencia(sentence, "Prod_CategoryId", "N", p.ProdCategId, 0, "")
	sentence = tools.ArmoSentencia(sentence, "Prod_Stock", "N", p.ProdStock, 0, "")
	sentence = tools.ArmoSentencia(sentence, "Prod_Path", "S", 0, 0, p.ProdPath)

	sentence += " WHERE Prod_Id = "+strconv.Itoa(p.ProdId)

	_, err = Db.Exec(sentence)
	if err != nil {
		fmt.Println(err.Error())
			return err
	}

	fmt.Println("Update Product > Ejecucion Exitosa")

	return nil
}