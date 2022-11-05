package httputil

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/mayukorin/ebook-merge/app/interfaces/api/httperror"
)

func RespondByJSON(w http.ResponseWriter, r *http.Request, statusCode int, res interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	jres, err := json.Marshal(res)
	if err != nil {
		respondMarshalErrByJSON(w, err)
		return
	}

	_, err = w.Write(jres)
	if err != nil {
		fmt.Println(fmt.Errorf("failed to write jresponse: %w", err))
		return
	}

	w.WriteHeader(statusCode)

	return
}

func RespondErrByJSON(w http.ResponseWriter, r *http.Request, statusCode int, err error) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	httpErr := httperror.HttpError{
		Message: err.Error(),
	}
	jerr, err := json.Marshal(httpErr)
	if err != nil {
		respondMarshalErrByJSON(w, err)
		return
	}

	_, err = w.Write(jerr)
	if err != nil {
		fmt.Println(fmt.Errorf("failed to write jerr: %w", err))
		return
	}

	w.WriteHeader(statusCode)

	return
}

func respondMarshalErrByJSON(w http.ResponseWriter, err error) {

	_, err = w.Write([]byte(err.Error()))
	if err != nil {
		fmt.Println(fmt.Errorf("failed to write marshal err: %w", err))
		return
	}

	w.WriteHeader(http.StatusInternalServerError)

	return
}
