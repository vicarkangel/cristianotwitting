package models

type Relacion struct {
	UsuarioID         string `json:"usuarioid" json:"usuarioId"`
	UsuarioRelacionID string `bson:"usuariorelacionid" json:"usuariorelacionId"`
}
