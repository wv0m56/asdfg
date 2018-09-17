package main

import (
	"bytes"
	"io"
	"math/rand"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	N := 10 * 1024 * 1024
	bin10mb := make([]byte, N)
	rand.Seed(time.Now().UnixNano())
	n, err := rand.Read(bin10mb)
	if err != nil || n != N {
		panic("error")
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/*", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "max-age=5")
		_, err = io.Copy(w, bytes.NewReader(bin10mb))
		if err != nil {
			panic(err)
		}
	})
	http.ListenAndServe(":13000", r)
}
