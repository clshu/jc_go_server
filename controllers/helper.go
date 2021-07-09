package controllers

import (
	"net/http"

	"github.com/gorilla/schema"
)

func parseForm(r *http.Request, dst interface{}) error {
	var decoder = schema.NewDecoder()

	if err := r.ParseForm(); err != nil {
		return err
	}

	if err := decoder.Decode(dst, r.PostForm); err != nil {
		return err
	}

	return nil
}
