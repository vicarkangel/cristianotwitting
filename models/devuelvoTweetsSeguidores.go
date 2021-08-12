package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

/* DevuelvoTweets es la estructura con la que devolveremos los Tweests */
type DevuelvoTweetsSeguidores struct {
	ID                primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	UsuarioID         string             `bson:"usuarioid" json:"usuarioId,omitempty"`
	UsuarioRelacionID string             `bson:"usuariorelacionid" json:"usuarioRelationId,omitempty"`
	Tweet             struct {
		Mensaje string    `bson:"mensaje" json:"mensaje,omitempty"`
		Fecha   time.Time `bson:"fecha" json:"fecha,omitempty"`
		ID      string    `bson:"_id" json:"_id,omitempty"`
	}
}
