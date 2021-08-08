package routers

import (
	"encoding/json"
	"net/http"

	"github.com/vicarkangel/cristianotwitting/bd"
	"github.com/vicarkangel/cristianotwitting/models"
)

func ModificarPerfil(w http.ResponseWriter, r *http.Request) {

	var t models.Usuario

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Datos incorrectos "+err.Error(), 400)
		return
	}
	var status bool

	status, err = bd.ModificoRegistro(t, IDUsuario)
	if err != nil {
		http.Error(w, "Ocurió un error al intentar modificar el registo. Reintentar Nuevamente "+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "No se ha logrado modificar el registro del usuario ", 400)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
