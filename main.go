package main

import (
	"log"
	"net/http"
	"strings"
)

var urls = map[string]string{
	"/":                         "https://linearmouse.app/",
	"/github":                   "https://github.com/linearmouse/linearmouse",
	"/accessibility-permission": "https://github.com/linearmouse/linearmouse/blob/main/ACCESSIBILITY.md",
	"/disable-pointer-acceleration-and-speed": "https://github.com/linearmouse/linearmouse/discussions/201",
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		target, ok := urls[strings.ToLower(r.URL.Path)]
		if !ok {
			http.NotFound(w, r)
			return
		}

		http.Redirect(w, r, target, http.StatusTemporaryRedirect)
	})

	log.Fatal(http.ListenAndServe(":3000", nil))
}
