package key

import (
	"backend/database"
	"backend/utils"
	"encoding/base64"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
)

var pass = "joensave"

func CreateKey(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)

	var keyData Key
	err := decoder.Decode(&keyData)

	if err != nil {
		panic(err)
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
	db.Create(&keyData)

	results := Message{0, "La llave " + keyData.Name + " se genero exitosamente"}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(results)

	/* ------------------------------ Prueba encriptar ---------------------------- */

	/*publicDecode, _ := base64.StdEncoding.DecodeString(keyData.Public)

	publicKeyEncrypt := utils.BytesToPublicKey(publicDecode)

	ciphertext := utils.EncryptWithPublicKey([]byte("hola mundo"), publicKeyEncrypt)
	fmt.Println("")
	fmt.Println("cadena encriptada: %x", ciphertext)
	fmt.Printf("cadena encriptada: %x\n", ciphertext)
	fmt.Println("cadena encriptada: %x\n", []byte(string(ciphertext)))
	fmt.Println("")

	privateDecode, _ := base64.StdEncoding.DecodeString(keyData.Private)

	privateKeyDecrypt := utils.BytesToPrivateKey(privateDecode, pass)
	plaintext := utils.DecryptWithPrivateKey(ciphertext, privateKeyDecrypt)
	fmt.Println("")
	fmt.Println("cadena desencriptada: %x", plaintext)
	fmt.Println("cadena desencriptada s: %x", string(plaintext))
	fmt.Println("")*/

	/* ------------------------------ Prueba encriptar con Bd ---------------------------- */

	/*var keyDataBd Key

	db.First(&keyDataBd, keyData.ID)

	publicBdDecode, _ := base64.StdEncoding.DecodeString(keyDataBd.Public)
	publicKeyBd := utils.BytesToPublicKey(publicBdDecode)

	ciphertextBd := utils.EncryptWithPublicKey([]byte("Hola mundo"), publicKeyBd)

	fmt.Println("")
	fmt.Println("")
	fmt.Println("_______________________DESDE BD__________________________")
	fmt.Println("cadena encriptada: %x", ciphertextBd)
	fmt.Printf("cadena encriptada: %x\n", ciphertextBd)
	fmt.Println("cadena encriptada: %x\n", []byte(string(ciphertextBd)))
	fmt.Println("")

	privateBdDecode, _ := base64.StdEncoding.DecodeString(keyDataBd.Private)
	privateBdKeyDecrypt := utils.BytesToPrivateKey(privateBdDecode, pass)

	plaintextBd := utils.DecryptWithPrivateKey(ciphertextBd, privateBdKeyDecrypt)
	fmt.Println("")
	fmt.Println("cadena desencriptada: %x", plaintextBd)
	fmt.Println("cadena desencriptada s: %x", string(plaintextBd))
	fmt.Println("")*/

}

//Buscar y listar llaves por nombre /api/key | /api/key?search=1
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

func GetKey(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "id")

	var keyData Key

	db := database.Connection()
	db.Select([]string{"id", "name", "created_at"}).First(&keyData, id)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(keyData)
}

func EncryptText(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)

	var Data DataReceived
	err := decoder.Decode(&Data)

	if err != nil {
		panic(err)
	}

	defer r.Body.Close()

	var keyData Key

	db := database.Connection()
	db.First(&keyData, Data.Id)
	database.CloseConnection(db)

	//publicDecode, _ := base64.StdEncoding.DecodeString(keyData.Public)
	publicKey := utils.BytesToPublicKey([]byte(keyData.Public))
	cipherText := utils.EncryptWithPublicKey([]byte(string(Data.Message)), publicKey)

	results := Message{}
	results.Message = base64.StdEncoding.EncodeToString(cipherText)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(results)
}

func DecryptText(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)

	var Data DataReceived
	err := decoder.Decode(&Data)

	if err != nil {
		panic(err)
	}

	defer r.Body.Close()

	var keyData Key

	db := database.Connection()
	db.First(&keyData, Data.Id)
	database.CloseConnection(db)

	decodedMessage, _ := base64.StdEncoding.DecodeString(Data.Message)

	//privateDecode, _ := base64.StdEncoding.DecodeString(keyData.Private)
	//privateDecode := keyData.Private

	privateKey := utils.BytesToPrivateKey([]byte(keyData.Private), pass)
	plainText := utils.DecryptWithPrivateKey(decodedMessage, privateKey)

	results := Message{}
	results.Message = string(plainText)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(results)
}
