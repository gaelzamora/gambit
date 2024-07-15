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
	