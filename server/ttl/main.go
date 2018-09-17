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
	N := 1 * 1024 * 1024
	bin1mb := make([]byte, N)
	rand.Seed(time.Now().UnixNano())
	n, err := rand.Read(bin1mb)
	if err != nil || n != N {
		panic("error")
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/*", func(w http.ResponseWriter, r *http.Request) {
		timeStr, err := time.Now().Add(5 * time.Second).MarshalText()
		if err != nil {
			panic(err)
		}
		w.Header().Set("Chainlightning-Expiry", string(timeStr))
		_, err = io.Copy(w, bytes.NewReader(bin1mb))
		if err != nil {
			panic(err)
		}
	})
	http.ListenAndServe(":13000", r)
}
