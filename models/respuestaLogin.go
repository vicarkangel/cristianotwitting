package models

/* Respuesta del token que devuelve con el login */
type RespuestaLogin struct {
	Token string `json:"token,omitemty"`
}
