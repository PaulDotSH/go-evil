package handlers

import (
	"fmt"
	settings "go-evil"
	"go-evil/parser"
	"net/http"
	"path"
)

func RansomwareBuilder(w http.ResponseWriter, r *http.Request) {
	fmt.Println(parser.Parse(path.Join(settings.RansomwarePath, "settings.go")))
}
