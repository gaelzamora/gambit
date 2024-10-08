package bd

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"

	"github.com/gaelzamora/gambit/models"
	_ "github.com/go-sql-driver/mysql"
)

func InsertAddress(addr models.Address, User string) error {
	fmt.Println("Comienza el Registro InsertAddress")

	err := DbConnect()

	if err != nil {
		return err
	}

	defer Db.Close()

	sentencia := "INSERT INTO addresses (Add_UserId, Add_Address, Add_City, Add_State, Add_PostalCode, Add_Phone, Add_Title, Add_Name )"
	sentencia += " VALUES ('" + User + "','" + addr.AddAddress + "','" + addr.AddCity + "','" + addr.AddState + "','"
	sentencia += addr.AddPostalCode + "','" + addr.AddPhone + "','" + addr.AddTitle + "','" + addr.AddName + "')"

	_, err = Db.Exec(sentencia)
	if err != nil {
		fmt.Println(err.Error())

		return err
	}

	fmt.Println(sentencia)
	fmt.Println("Insert Address > Ejecucion Exitosa")

	return nil
}

func AddressExists(User string, id int) (error, bool) {
	fmt.Println("Comienza AddressExists")

	err := DbConnect()
	if err != nil {
		return err, false
	}

	defer Db.Close()

	sentence := "SELECT 1 FROM addresses WHERE Add_Id = "+strconv.Itoa(id)+" AND Add_UserId = '"+User+"'"

	fmt.Println(sentence)

	rows, err := Db.Query(sentence)
	if err != nil {
		return err, false
	}

	var valor string

	rows.Next()
	rows.Scan(&valor)

	fmt.Println("AddressExists > Ejecucion Exitosa - valor devuelto "+valor)

	if valor == "1" {
		return nil, true
	}
	
	
	return nil, false
}

func UpdateAddress(addr models.Address) error {
	fmt.Println("Comeinza Update Address")

	err := DbConnect()
	if err != nil {
		return err
	}

	defer Db.Close()

	sentencia := "UPDATE addresses SET "

	if addr.AddAddress != "" {
		sentencia += "Add_Address = '" + addr.AddAddress + "', "
	}
	if addr.AddCity != "" {
		sentencia += "Add_City = '" + addr.AddCity + "', "
	}
	if addr.AddName != "" {
		sentencia += "Add_Name = '" + addr.AddName + "', "
	}
	if addr.AddPhone != "" {
		sentencia += "Add_Phone = '" + addr.AddPhone + "', "
	}
	if addr.AddPostalCode != "" {
		sentencia += "Add_PostalCode = '" + addr.AddPostalCode + "', "
	}
	if addr.AddState != "" {
		sentencia += "Add_State = '" + addr.AddState + "', "
	}
	if addr.AddTitle != "" {
		sentencia += "Add_Title = '" + addr.AddTitle + "', "
	}

	sentencia, _ = strings.CutSuffix(sentencia, ", ")

	sentencia += " WHERE Add_Id = "+strconv.Itoa(addr.AddId)

	_, err = Db.Exec(sentencia)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println(sentencia)
	fmt.Println("Update Address > Ejecución Exitosa")
	return nil
}

func DeleteAddress(id int) error {
	fmt.Println("Comienza el DeleteAddress")

	err := DbConnect()
	if err != nil {
		return err
	}

	defer Db.Close()

	sentence := "DELETE FROM addresses WHERE Add_Id = "+strconv.Itoa(id)

	_, err = Db.Exec(sentence)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println(sentence)
	fmt.Println("Delete Address > Ejecución Exitosa")
	return nil
}

func SelectAddress(User string) ([]models.Address, error) {
	fmt.Println("Comienza SelectAddress")
	Addr := []models.Address{}

	err := DbConnect()
	if err != nil {
		return nil, err
	}

	defer Db.Close()

	sentence := "SELECT Add_Id, Add_Address, Add_City, Add_State, Add_PostalCode, Add_Phone, Add_Title, Add_Name FROM addresses WHERE Add_UserId = '"+User+"'"

	var rows *sql.Rows
	rows, err = Db.Query(sentence)

	if err != nil {
		fmt.Println(err.Error())
		return Addr, err
	}
	defer rows.Close()

	for rows.Next() {
		var a models.Address
		var addId sql.NullInt16
		var addAddress sql.NullString
		var addCity sql.NullString
		var addState sql.NullString
		var addPostalCode sql.NullString
		var addPhone sql.NullString
		var addTitle sql.NullString
		var addName sql.NullString
	
		err := rows.Scan(&addId, &addAddress, &addCity, &addState, &addPostalCode, &addPhone, &addTitle, &addName)
		if err != nil {
			return Addr, err
		}

		a.AddId = int(addId.Int16)
		a.AddAddress = addAddress.String
		a.AddCity = addCity.String
		a.AddState = addState.String
		a.AddPostalCode = addPostalCode.String
		a.AddPhone = addPhone.String
		a.AddTitle = addTitle.String
		a.AddName = addName.String

		Addr = append(Addr, a)
	}

	fmt.Println("Select Addresses > Ejecucion exitosa")

	return Addr, nil
}