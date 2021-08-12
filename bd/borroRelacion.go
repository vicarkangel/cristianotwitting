package bd

import (
	"context"
	"time"

	"github.com/vicarkangel/cristianotwitting/models"
)

/* BuscoPerfil busca un perfil en la bd */

func BorroRelacion(t models.Relacion) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("cristianotwitting-bd")
	col := db.Collection("relacion")

	_, err := col.DeleteOne(ctx, t)
	if err != nil {
		return false, err
	}
	return false, nil
}
