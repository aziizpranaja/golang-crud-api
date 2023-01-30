package book

// Catatan jika tipe data integer akan gagal mengconvert karena validator adalah tipe data string,
// jadi int dan string tidak match
// solusinya dengan mengubah tipe data sting menjadi json.Number
// json.Number akan membaca tipe data string menjadi number selama string tersebut sebuah angka contoh "20"
type BookInput struct{
	Title string `json:"title" binding:"required"`
	Price interface{} `json:"price" binding:"required,number"`
}