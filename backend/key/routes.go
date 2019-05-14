package key

import "github.com/go-chi/chi"

func Routes() *chi.Mux {

	router := chi.NewRouter()
	router.Post("/", CreateKey)
	router.Get("/", GetKeys)
	router.Get("/{id}", GetKey)
	router.Post("/encrypt", EncryptText)
	router.Post("/decrypt", DecryptText)
	return router
}
