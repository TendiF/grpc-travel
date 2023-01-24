package types

import (
	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Claims struct {
	jwt.StandardClaims
	Uid  string
	Role int
}

type User struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	CreatedAt    primitive.DateTime `bson:"created_at,omitempty" json:"created_at"`
	ModifiedAt   primitive.DateTime `bson:"modified_at,omitempty" json:"modified_at"`
	FirstName    string             `bson:"first_name,omitempty" json:"first_name"`
	LastName     string             `bson:"last_name,omitempty" json:"last_name"`
	Gender       string             `bson:"gender,omitempty" json:"gender"`
	Email        string             `bson:"email,omitempty" json:"email"`
	Password     string             `bson:"password,omitempty" json:"password"`
	Address      string             `bson:"address,omitempty" json:"address"`
	Username     string             `bson:"username,omitempty" json:"username"`
	Role         int                `bson:"role,omitempty" json:"role"`
	CodeMerchant string             `bson:"code_merchant,omitempty" json:"code_merchant"`
}

type Customers struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	CreatedAt     primitive.DateTime `bson:"created_at,omitempty" json:"created_at"`
	ModifiedAt    primitive.DateTime `bson:"modified_at,omitempty" json:"modified_at"`
	No            string             `bson:"no,omitempty" json:"no"`
	NIK           string             `bson:"nik,omitempty" json:"nik"`
	Nama          string             `bson:"nama,omitempty" json:"nama"`
	StatusKK      string             `bson:"status_kk,omitempty" json:"status_kk"`
	NoKK          string             `bson:"no_kk,omitempty" json:"no_kk"`
	TanggalLahir  string             `bson:"tanggal_lahir,omitempty" json:"tanggal_lahir"`
	Usia          string             `bson:"usia,omitempty" json:"usia"`
	KotaKab       string             `bson:"kota_kab,omitempty" json:"kota_kab"`
	Kecamatan     string             `bson:"kecamatan,omitempty" json:"kecamatan"`
	DesaKelurahan string             `bson:"desa_kelurahan,omitempty" json:"desa_kelurahan"`
	Kampung       string             `bson:"kampung,omitempty" json:"kampung"`
	RT            string             `bson:"rt,omitempty" json:"rt"`
	RW            string             `bson:"rw,omitempty" json:"rw"`
	Kol           string             `bson:"kol,omitempty" json:"kol"`
	Syahidan      string             `bson:"syahidan,omitempty" json:"syahidan"`
	PJ            string             `bson:"pj,omitempty" json:"pj"`
	Note          string             `bson:"note,omitempty" json:"note"`
	CodeMerchant  string             `bson:"code_merchant,omitempty" json:"code_merchant"`
}

type Reguler struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	CustomerID     primitive.ObjectID `bson:"customer_id,omitempty" json:"customer_id"`
	CreatedAt      primitive.DateTime `bson:"created_at,omitempty" json:"created_at"`
	ModifiedAt     primitive.DateTime `bson:"modified_at,omitempty" json:"modified_at"`
	BulanTahun     string             `bson:"bulantahun,omitempty" json:"bulantahun"`
	Infaq          int                `bson:"infaq,omitempty" json:"infaq"`
	Zakat          int                `bson:"zakat,omitempty" json:"zakat"`
	Shadaqah       int                `bson:"shadaqah,omitempty" json:"shadaqah"`
	Ikhsan         int                `bson:"ikhsan,omitempty" json:"ikhsan"`
	Sabil          int                `bson:"sabil,omitempty" json:"sabil"`
	TabunganFitrah int                `bson:"tabungan_fitrah,omitempty" json:"tabungan_fitrah"`
	TabunganQurban int                `bson:"tabungan_qurban,omitempty" json:"tabungan_qurban"`
	Bina           string             `bson:"bina,omitempty" json:"bina"`
}
