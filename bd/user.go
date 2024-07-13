package bd

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/gaelzamora/gambit/models"
	"github.com/gaelzamora/gambit/tools"
	_ "github.com/go-sql-driver/mysql"
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

	sentence += ", User_DateUpg = '"+tools.FechaMySQL()+"' WHERE User_UUID='"+User+"'"
	
	_, err = Db.Exec(sentence)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("Update User > Ejecucion Exitosa")

	return nil
}

func SelectUser(UserId string) (models.User, error) {
	fmt.Println("Comienza SelectUser")

	user := models.User{}

	err := DbConnect()
	if err != nil {
		return user, err
	}
	defer Db.Close()

	sentence := "SELECT	* FROM users WHERE User_UUID = '"+UserId+"'"

	var rows *sql.Rows
	rows, err = Db.Query(sentence)

	if err != nil {
		fmt.Println(err.Error())
		return user, err
	}
    rows.Next()

	var firstName sql.NullString
	var lastName sql.NullString
	var dateUpg sql.NullTime

	rows.Scan(&user.UserUUID, &user.UserEmail, &firstName, &lastName, &user.UserStatus, &user.UserDateAdd, &dateUpg)

	user.UserFirstName = firstName.String
	user.UserLastName = lastName.String
	user.UserDateUpd = dateUpg.Time.String()

	fmt.Println("Select user > Ejecucion exitosa")

	return user, nil
}

func SelectUsers(Page int) (models.ListUsers, error) {
	fmt.Println("Comienza SelectUsers")

	var lu models.ListUsers
	User := []models.User{}

	err := DbConnect()
	if err != nil {
		return lu, err
	}
	defer Db.Close()

	var offSet int = (Page*10) - 10
	var sentence string
	var sentenceCount string = "SELECT count(*) as registros FROM users"


	sentence = "select * from users LIMIT 10"
	if offSet > 0 {
		sentence += " OFFSET "+strconv.Itoa(offSet)
	} 

	var rowsCount *sql.Rows

	rowsCount, err = Db.Query(sentenceCount)
	if err != nil {
		return lu, err
	}

	defer rowsCount.Close()
	
	rowsCount.Next()
	
	var registros int
	rowsCount.Scan(&registros)
	lu.TotalItems=registros

	var rows *sql.Rows
	rows, err = Db.Query(sentence)

	if err != nil {
		fmt.Println(err.Error())
		return lu, err
	}

	for rows.Next() {
		var u models.User
		var firstName sql.NullString
		var lastName sql.NullString
		var dateUpg sql.NullTime
		
		rows.Scan(&u.UserUUID, &u.UserEmail, &firstName, &lastName, &u.UserStatus, &u.UserDateAdd, &dateUpg)

		u.UserFirstName = firstName.String
		u.UserLastName = lastName.String
		u.UserDateUpd = dateUpg.Time.String()

		User = append(User, u)
	}

	fmt.Println("Select Users > Ejecucion Exitosa")

	lu.Data	= User
	return lu, nil
}