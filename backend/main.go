package main

import (
	"backend/key"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"

	//"github.com/go-chi/render"
	_ "github.com/lib/pq"
)

func Routes() *chi.Mux {
	router := chi.NewRouter()

	router.Use(
		//render.SetContentType(render.ContentTypeJSON),
		middleware.RequestID,
		middleware.RealIP,
		middleware.Logger,
		//middleware.DefaultCompress,
		//middleware.RedirectSlashes,
		middleware.Recoverer,
		middleware.Timeout(60*time.Second),
	)

	cors := cors.New(cors.Options{
		// AllowedOrigins: []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})

	router.Use(cors.Handler)

	router.Route("/v1", func(r chi.Router) {
		r.Mount("/api/key", key.Routes())
	})

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hi"))
	})

	return router
}

func main() {

	router := Routes()

	/*walkFunc := func(method string, route string, handler htpp.Handler, middlewares ...func(http.Handler) http.Handler) error {
		log.Println("%s %s", method, route)
		return nil
	}

	if err := chi.Walk(router, walkFunc); err != nil{
		log.Panic("")
	}*/

	log.Fatal(http.ListenAndServe(":3333", router))
}
