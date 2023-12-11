package main

import (
	"fmt"
	"image"
	"image/png"
	"math/rand"
	"net/http"

	"git.sr.ht/~sbinet/gg"
)

func imageGeneration() image.Image {
	dc := gg.NewContext(1000, 200)
	dc.SetRGB(1, 1, 1)
	if err := dc.LoadFontFace("./LatinmodernmathRegular.otf", 100); err != nil {
		panic(err)
	}
	denom := rand.Intn(4) + 2
	prod := rand.Intn(3) + 2
	calc := fmt.Sprintf("%d × (%d / %d) + %d =", rand.Intn(8)+2, denom*prod, denom, rand.Intn(19)+1)
	dc.DrawStringAnchored(calc, 500, 100, 0.5, 0.5)
	return dc.Image()
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Cache-Control", "no-store, no-cache, max-age=0")

	img := imageGeneration()
	png.Encode(w, img)

}

func main() {
	// _ = imageGeneration()
	handler := http.HandlerFunc(handleRequest)
	http.Handle("/", handler)
	http.ListenAndServe(":8080", nil)
}
