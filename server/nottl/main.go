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
	N := 100 * 1024 * 1024
	bin100mb := make([]byte, N)
	rand.Seed(time.Now().UnixNano())
	n, err := rand.Read(bin100mb)
	if err != nil || n != N {
		panic("error")
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/*", func(w http.ResponseWriter, r *http.Request) {
		_, err := io.Copy(w, bytes.NewReader(bin100mb))
		if err != nil {
			panic(err)
		}
	})
	http.ListenAndServe(":13000", r)
}
