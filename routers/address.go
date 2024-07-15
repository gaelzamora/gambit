package routers

import (
	"encoding/json"
	"github.com/gaelzamora/gambit/bd"
	"github.com/gaelzamora/gambit/models"
)

func InsertAddress(body string, User string) (int, string) {
	var t models.Address

	err := json.Unmarshal([]byte(body), &t)

	if err != nil {
		return 400, "Error en los datos recibidos: "+err.Error()
	}

	if t.AddAddress == "" {
		return 400, "Debe especificar el Address"
	}

	if t.AddName == "" {
		return 400, "Debe especificar el Name"
	}
	if t.AddTitle == "" {
		return 400, "Debe especificar el Title"
	}
	if t.AddCity == "" {
		return 400, "Debe especificar el City"
	}
	if t.AddPhone == "" {
		return 400, "Debe especificar el Phone"
	}
	if t.AddPostalCode == "" {
		return 400, "Debe especificar el Postal	Code"
	}

	err = bd.InsertAddress(t, User)
	if err != nil {
		return 400, "Ocurrio un error al intentar realizar el registro de Address "+err.Error()
	}

	return 200, "InsertAddress OK"
}

func UpdateAddress(body string, User string, id int) (int, string) {
	var t models.Address

	err := json.Unmarshal([]byte(body), &t)

	if err != nil {
		return 400, "Error en los datos recibidos "+err.Error()
	}

	t.AddId = id

	var encontrado bool
	err, encontrado = bd.AddressExists(User, t.AddId)

	if !encontrado {
		if err != nil {
			return 400, "Error al intentar buscar Address para el usuario "+User+" > "+err.Error()
		}
		return 400, "No se encuentra un registro de ID de Usuario asociado a esa ID de Address"
	}

	err = bd.UpdateAddress(t)

	if err != nil {
		return 400, "Ocurrio un error al intentar realizar la actualizacion del Addres para el ID de Usuario "+User+" > "+err.Error()
	}

	return 200, "UpdateAddress OK"
}

func DeleteAddress(User string, id int) (int, string) {
	err, encontrado := bd.AddressExists(User, id)

	if !encontrado {
		if err != nil {
			return 400, "Error al intentar buscar Address para el usuario "+User+" > "+err.Error()
		}
		return 400, "No se encuentra un registro de ID de Usuario asociado a esa ID de Address"
	}

	err2 := bd.DeleteAddress(id)
	if err2 != nil {
		return 400, "Ocurrio un error al intentar eliminar un Address de "+User
	}

	return 200, "DeleteAddress OK"
}