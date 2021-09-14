package jwt

import (
	"time"
	"twittor/models"

	jwt "github.com/dgrijalva/jwt-go"
)

func GeneroJWT(t models.Usuario) (string, error){
	//SLICE DE BYTES
	miClave := []byte("MastersdelDesarrollo_grupo")

	payload := jwt.MapClaims{
		"email": t.Email,
		"nombre": t.Nombre,
		"apellidos": t.Apellidos,
		"fecha_nacimiento": t.FechaNacimiento,
		"biografia": t.Biografia,
		"ubicacion": t.Ubicacion,
		"sitioweb": t.SitioWeb,
		"_id": t.ID.Hex(),
		//FORMATO UNIX SE GUARDA DE FORMA EFICIENTE EN UN LONG
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(miClave)
	if err != nil {
		return tokenStr, err
	}

	return tokenStr, nil
}