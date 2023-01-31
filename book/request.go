package book

import "encoding/json"

// Catatan jika tipe data integer akan gagal mengconvert karena validator adalah tipe data string,
// jadi int dan string tidak match
// solusinya dengan mengubah tipe data sting menjadi json.Number
// json.Number akan membaca tipe data string menjadi number selama string tersebut sebuah angka contoh "20"
type BookRequest struct{
	Title string `json:"title" binding:"required"`
	Price json.Number `json:"price" binding:"required,number"`
	Description string `json:"description" binding:"required"`
	Rating int32 `json:"rating" binding:"required,number"`
	Discount int32 `json:"discount" binding:"required,number"`
}