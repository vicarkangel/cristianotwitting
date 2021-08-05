package routers

import (
	"encoding/json"
	"net/http"

	"github.com/vicarkangel/cristianotwitting/bd"
	"github.com/vicarkangel/cristianotwitting/models"
)

/* Registro funcion para registrar al usuario en la BD */

func Registro(w http.ResponseWriter, r *http.Request) {

	var t models.Usuario
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Error en los datos recibidos"+err.Error(), 400)
	}

	if len(t.Email) == 0 {
		http.Error(w, "El email de usuario es requerido", 400)
	}
	if len(t.Password) < 6 {
		http.Error(w, "Debe especificar una contraceña de al menos 6 caracteres", 400)
		return

	}

	_, encontrado, _ := bd.ChequeoYaExisteUsuario(t.Email)
	if encontrado == true {
		http.Error(w, "Ya existe un usuario registrado con ese email", 400)
		return
	}

	_, status, err := bd.InsertoRegistro(t)
	if err != nil {
		http.Error(w, "Ocurrio un erro al intentar realizar el registo de usuario"+err.Error(), 400)
		return
	}
	if status == false {
		http.Error(w, "No se ha logrado insertar el registro de ususario", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
