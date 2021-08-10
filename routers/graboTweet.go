package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/vicarkangel/cristianotwitting/bd"
	"github.com/vicarkangel/cristianotwitting/models"
)

/* GraboTweet permite grabar el tweet en la base de datos */
func GraboTweet(w http.ResponseWriter, r *http.Request) {
	var mensaje models.Tweet
	err := json.NewDecoder(r.Body).Decode(&mensaje)

	registro := models.GraboTweet{
		UserID:  IDUsuario,
		Mensaje: mensaje.Mensaje,
		Fecha:   time.Now(),
	}

	/* GraboTweet permite grabar el tweet en la base de datos */
	_, status, err := bd.InsertoTweet(registro)
	if err != nil {
		http.Error(w, "Ocurrio  un error al intentar insertar el registro, reintentar nuevamente "+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "No se ha logrado insertar el registro, reintentar nuevamente ", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
