package main

import (
	"log"
	"net/http"
	"net/url"
	"strings"
)

var urls = map[string]string{
	"/":                         "https://linearmouse.app/",
	"/github":                   "https://github.com/linearmouse/linearmouse",
	"/accessibility-permission": "https://github.com/linearmouse/linearmouse/blob/main/ACCESSIBILITY.md",
	"/disable-pointer-acceleration-and-speed": "https://github.com/linearmouse/linearmouse/discussions/201",
	"/bug-report":      "https://github.com/linearmouse/linearmouse/issues/new?template=bug_report.yml&title=%5BBUG%5D+&labels=bug",
	"/feature-request": "https://github.com/linearmouse/linearmouse/issues/new?template=feature_request.yml&title=%5BFeature%5D+&labels=enhancement",
	"/donate":          "https://github.com/sponsors/linearmouse",
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		target, ok := urls[strings.ToLower(r.URL.Path)]
		if !ok {
			http.NotFound(w, r)
			return
		}

		if r.Method != http.MethodGet {
			http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
			return
		}

		targetURL, err := url.Parse(target)
		if err != nil {
			log.Printf("Failed to parse target: %s: %s", target, err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}

		query := targetURL.Query()
		for k, vv := range r.URL.Query() {
			for _, v := range vv {
				query.Add(k, v)
			}
		}
		targetURL.RawQuery = query.Encode()

		target = targetURL.String()

		http.Redirect(w, r, target, http.StatusTemporaryRedirect)
	})

	log.Fatal(http.ListenAndServe(":3000", nil))
}
