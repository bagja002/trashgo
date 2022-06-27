package Models

import "github.com/dgrijalva/jwt-go"

type User struct {
	IDUser         int    `json:"Id_user"`
	Email          string `gorm:"unique"`
	Password       []byte `json:"-"`
	Username       string `json:"username"`
	Address        string `json:"address"`
	Place_of_birth string `json:"Place_of_birth"`
	Gender         string `json:"gender"`
	Foto           string `json:"foto"`
	Title          string `json:"title"`
}

type Aku struct {
	jwt.StandardClaims
	IDUser         int    `json:"Id_user"`
	Email          string `gorm:"unique"`
	Username       string `json:"nama_user"`
	Address        string `json:"alamat"`
	Place_of_birth string `json:"tempat_lahir"`
	Gender         string `json:"kelamin"`
	Foto           string `json:"Foto"`
	Title          string `json:"title"`
}
