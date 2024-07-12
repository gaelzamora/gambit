package bd

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gaelzamora/gambit/models"
	"github.com/gaelzamora/gambit/tools"
)

func UpdateUser(UField models.User, User string) error {
	fmt.Println("Comienza update user")

	err := DbConnect()
	if err != nil {
		return err
	}
	defer Db.Close()

	sentence := "UPDATE users SET "

	coma := ""
	if len(UField.UserFirstName) > 0 {
		coma = ","
		sentence += "User_FirstName = '" + UField.UserFirstName + "'"
	}

	if len(UField.UserLastName) > 0 {
		sentence += coma + "User_LastName = '"+UField.UserLastName+"'"
	}

	sentence += ", User_DataUpg = '"+tools.FechaMySQL()+"' WHERE User_UUID='"+User+"'"
	
	_, err = Db.Exec(sentence)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("Update User > Ejecucion Exitosa")

	return nil

}