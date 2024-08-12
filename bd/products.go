package bd

import (
	"database/sql"
	"errors"
	"fmt"
	"strconv"
	"strings"

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

	fmt.Println(sentence)

	_, err = Db.Exec(sentence)
	if err != nil {
		fmt.Println(err.Error())
			return err
	}

	fmt.Println("Update Product > Ejecucion Exitosa")

	return nil
}

func DeleteProduct(id int) error {
	fmt.Println("Comienza Delete Product")

	err := DbConnect()
	if err != nil {
		return err
	}

	defer Db.Close()
	sentence := "DELETE FROM products WHERE Prod_Id = " + strconv.Itoa(id) 
	
	_, err =Db.Exec(sentence)

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("Delete product > Ejecucion exitosa")

	return nil
}

func SelectProduct(p models.Product, choice string, page int, pageSize int, orderType string, orderField string) (models.ProductResp, error) {
	fmt.Println("Comienza SelectProduct")
    var Resp models.ProductResp
    var Prod []models.Product

    err := DbConnect()
    if err != nil {
        return Resp, err
    }
    defer Db.Close()

    var sentence string
    var sentenceCount string
    var where, limit string    

    sentence = "SELECT Prod_Id, Prod_Title, Prod_Description, Prod_CreatedAt, Prod_Updated, Prod_Price, Prod_Path, Prod_CategoryId, Prod_Stock FROM products "
    sentenceCount = "SELECT count(*) as registros FROM products "

    switch choice {
    case "P":
        where = " WHERE Prod_Id = " + strconv.Itoa(p.ProdId)
    case "S":
        where = " WHERE UCASE(CONCAT(Prod_Title, Prod_Description)) LIKE '%" + strings.ToUpper(p.ProdSearch) + "%' "
    case "C":
        where = " WHERE Prod_CategoryId = " + strconv.Itoa(p.ProdCategId)
    case "U":
        where = " WHERE UCASE(Prod_Path) LIKE '%" + strings.ToUpper(p.ProdPath) + "%' "
    case "K":
        join := " JOIN category ON Prod_CategoryId = Categ_Id AND Categ_Path LIKE '%" + strings.ToUpper(p.ProdCategPath) + "%' "
        sentence += join
        sentenceCount += join
    }	

    sentenceCount += where

    fmt.Println("La consulta es " + sentence)
	fmt.Println("El choice es "+choice)
    fmt.Println(sentenceCount)

    rows, err := Db.Query(sentenceCount)
    if err != nil {
        return Resp, err
    }
    defer rows.Close()

    rows.Next()
    var regi sql.NullInt32
    err = rows.Scan(&regi)
    if err != nil {
        return Resp, err
    }

    registros := int(regi.Int32)

    if page > 0 {
        if registros > pageSize {
            limit = " LIMIT " + strconv.Itoa(pageSize)
            if page > 1 {
                offset := pageSize * (page - 1)
                limit += " OFFSET " + strconv.Itoa(offset)
            }
        } else {
            limit = ""
        }
    }

    var orderBy string
    if len(orderField) > 0 {
        switch orderField {
        case "I":
            orderBy = " ORDER BY Prod_Id "
        case "T":
            orderBy = " ORDER BY Prod_Title "
        case "D":
            orderBy = " ORDER BY Prod_Description "
        case "F":
            orderBy = " ORDER BY Prod_CreatedAt "
        case "P":
            orderBy = " ORDER BY Prod_Price "
        case "S":
            orderBy = " ORDER BY Prod_Stock "
        case "C":
            orderBy = " ORDER BY Prod_CategoryId "
        }
        if orderType == "D" {
            orderBy += " DESC"
        }
    }

    sentence += where + orderBy + limit

    fmt.Println(sentence)

    rows, err = Db.Query(sentence)
    if err != nil {
        return Resp, err
    }
    defer rows.Close()

    for rows.Next() {
        var p models.Product
        var ProdId sql.NullInt32
        var ProdTitle sql.NullString
        var ProdDescription sql.NullString
        var ProdCreatedAtBytes []byte
        var ProdUpdatedBytes []byte
        var ProdPrice sql.NullFloat64
        var ProdPath sql.NullString
        var ProdCategoryId sql.NullInt32
        var ProdStock sql.NullInt32

        err := rows.Scan(&ProdId, &ProdTitle, &ProdDescription, &ProdCreatedAtBytes, &ProdUpdatedBytes, &ProdPrice, &ProdPath, &ProdCategoryId, &ProdStock)
        if err != nil {
            fmt.Println("Error en Scan:", err)
            return Resp, err
        }

        p.ProdId = int(ProdId.Int32)
        p.ProdTitle = ProdTitle.String
        p.ProdDescription = ProdDescription.String
        p.ProdCreatedAt = string(ProdCreatedAtBytes)
        p.ProdUpdated = string(ProdUpdatedBytes)
        p.ProdPrice = ProdPrice.Float64
        p.ProdPath = ProdPath.String
        p.ProdCategId = int(ProdCategoryId.Int32)
        p.ProdStock = int(ProdStock.Int32)
        Prod = append(Prod, p)

        fmt.Println("Finaliza")
    }

    Resp.TotalItems = registros
    Resp.Data = Prod

    fmt.Println("Select Product > Ejecucion Exitosa")

    return Resp, nil
}

func UpdateStock(p models.Product) error {
	fmt.Println("Comienza update stock")

	if p.ProdStock==0 {
		return errors.New("{ERROR} debe enviar el Stock a modificar")
	}

	err := DbConnect()
	if err != nil {
		return err
	}

	defer Db.Close()
	sentence := "UPDATE products SET Prod_Stock = Prod_Stock + " + strconv.Itoa(p.ProdStock) + " WHERE Prod_Id = " + strconv.Itoa(p.ProdId) 

	_, err = Db.Exec(sentence)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("Update Stock > Ejecucion exitosa")

	return nil
}