package key

import (
	"backend/database"
	"backend/key/structures"
	"backend/utils"
	"encoding/base64"
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

var pass = "joensave"

//Crear llaves POST /api/key
func CreateKey(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)

	var keyData Key
	err := decoder.Decode(&keyData)

	if err != nil {
		log.Panic(err)
	}

	defer r.Body.Close()

	privateKey, publicKey := utils.GenerateKeyPair(2048)

	public := utils.PublicKeyToBytes(publicKey)
	private := utils.PrivateKeyToBytes(privateKey, pass)

	/*
		keyData.Public = base64.StdEncoding.EncodeToString(public)
		keyData.Private = base64.StdEncoding.EncodeToString(private)
	*/

	keyData.Public = string(public)
	keyData.Private = string(private)

	db := database.Connection()

	if err := db.Create(&keyData).Error; err != nil {
		w.WriteHeader(500)
		return
	}

	results := structures.Message{0, "La llave " + keyData.Name + " se genero exitosamente"}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(results)
}

//Buscar y listar llaves por nombre GET /api/key || /api/key?search=1
func GetKeys(w http.ResponseWriter, r *http.Request) {

	var keys []Key
	search := r.FormValue("search")
	db := database.Connection()

	if search == "" {
		db.Select([]string{"id", "name", "created_at"}).Find(&keys)
	} else {
		db.Select([]string{"id", "name", "created_at"}).Where("name LIKE ?", "%"+search+"%").Find(&keys)
	}

	database.CloseConnection(db)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(keys)
}

//Obtener detalle llave GET /api/key/id
func GetKey(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "id")

	var keyData Key

	db := database.Connection()
	db.Select([]string{"id", "name", "created_at"}).First(&keyData, id)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(keyData)
}

//Encriptar texto plano POST /api/key/encrypt
func EncryptText(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)

	var Data structures.DataReceived
	err := decoder.Decode(&Data)

	if err != nil {
		log.Panic(err)
	}

	defer r.Body.Close()

	var keyData Key

	db := database.Connection()
	db.First(&keyData, Data.Id)
	database.CloseConnection(db)

	//publicDecode, _ := base64.StdEncoding.DecodeString(keyData.Public)

	publicKey := utils.BytesToPublicKey([]byte(keyData.Public))
	cipherText := utils.EncryptWithPublicKey([]byte(string(Data.Message)), publicKey)

	results := structures.Message{}
	results.Message = base64.StdEncoding.EncodeToString(cipherText)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(results)
}

//Desencriptar texto plano POST /api/key/decrypt
func DecryptText(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)

	var Data structures.DataReceived
	err := decoder.Decode(&Data)

	if err != nil {
		log.Panic(err)
	}

	defer r.Body.Close()

	var keyData Key

	db := database.Connection()
	db.First(&keyData, Data.Id)
	database.CloseConnection(db)

	decodedMessage, _ := base64.StdEncoding.DecodeString(Data.Message)

	//privateDecode, _ := base64.StdEncoding.DecodeString(keyData.Private)

	privateKey := utils.BytesToPrivateKey([]byte(keyData.Private), pass)
	plainText := utils.DecryptWithPrivateKey(decodedMessage, privateKey)

	results := structures.Message{}
	results.Message = string(plainText)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(results)
}
