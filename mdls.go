package mdls

import (
	"fmt"
	"net/http"
)

func HandleJs(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/javascript")
	w.Header().Add("Content-Length", fmt.Sprintf("%v", len(Str_mdls_js)))
	fmt.Fprint(w, Str_mdls_js)
}
func HandleCss(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/css")
	w.Header().Add("Content-Length", fmt.Sprintf("%v", len(Str_mdls_css)))
	fmt.Fprint(w, Str_mdls_css)
}
