package bd

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/gaelzamora/gambit/models"
)

func InsertOrder(o models.Orders) (int64, error) {
	fmt.Println("Comienza registro de Orders")

	err := DbConnect()
	if err != nil {
		return 0, err
	}

	defer Db.Close()

	sentence := "INSERT INTO orders (Order_UserUUID, Order_Total, Order_AddId) VALUES ('"
	sentence += o.Order_UserUUID + "',"+strconv.FormatFloat(o.Order_Total, 'f', -1, 64) + ","+strconv.Itoa(o.Order_AddId)+")"

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

	for _, od := range o.OrderDetails {
		sentence = "INSERT INTO order_detail (OD_OrderId, OD_ProdId, OD_Quantity, OD_Price) VALUES (" + strconv.Itoa(int(LastInsertId))
		sentence +=	","+strconv.Itoa(od.OD_ProdId)+","+strconv.Itoa(od.OD_Quantity)+","+strconv.FormatFloat(od.OD_Price, 'f', -1, 64)

		fmt.Println(sentence)
		_, err = Db.Exec(sentence)
		if err != nil {
			fmt.Println(err.Error())
			return 0, err
		}
	}

	fmt.Println("Insert Order > Ejecucion Exitosa")

	return LastInsertId, nil
}