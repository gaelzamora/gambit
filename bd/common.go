package bd

import (
	"database/sql"
	"fmt"
	"os"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gaelzamora/gambit/models"
	"github.com/gaelzamora/gambit/secretm"
)

var SecretModel models.SecretRDSJson
var err error
var Db *sql.DB

func ReadSecret() error {
	SecretModel, err = secretm.GetSecret(os.Getenv("SecretName"))
	return err
}

func DbConnect() error {
	Db, err = sql.Open("mysql", ConStr(SecretModel))
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println(Db.Ping())
	err = Db.Ping()
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("Conexion exitosa de la BD")
	return nil
}

func ConStr(claves models.SecretRDSJson) string {
	var dbUser, authToken, dbEndpoint, dbName string
	dbUser = claves.Username
	authToken = claves.Password
	dbEndpoint = claves.Host
	dbName = "gambit"
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?allowCleartextPasswords=true", dbUser, authToken, dbEndpoint, dbName)
	fmt.Println(dsn)
	return dsn
}

func IsAdmin(userUUID string) (bool, string) {
	fmt.Println("Comienza IsAdmin")

	err := DbConnect()

	if err != nil {
		return false, err.Error()
	}

	defer Db.Close()

	sentence := "SELECT 1 FROM users WHERE User_UUID='" +userUUID+ "' AND User_Status = 0"
	fmt.Println(sentence)

	rows, err := Db.Query(sentence)

	if err != nil {
		return false, err.Error()
	}

	var valor string
	rows.Next()
	rows.Scan(&valor)

	fmt.Println(valor)

	fmt.Println("UserIsAdmin > Ejecucion exitosa - valor devuelto " + valor)
	if valor == "1" {
		return true, ""
	}

	if valor == "0" {
		return false, "No es admin"
	}

	return false, "User is not Admin"
}

func UserExists(UserUUID string) (error, bool) {
	fmt.Println("Comienza UserExists")

	err := DbConnect()
	if err != nil {
		return err, false
	}
	defer Db.Close()

	sentence := "SELECT 1 FROM users WHERE User_UUID='" +UserUUID+ "'"
	fmt.Println(sentence)

	rows, err := Db.Query(sentence)
	if err != nil {
		return err, false
	}

	var valor string
	rows.Next()
	rows.Scan(&valor)

	fmt.Println("UserExists > Ejecucion exitosa - valor devuelto "+valor)

	if valor == "1" {
		return nil, true
	}

	return nil, false
}