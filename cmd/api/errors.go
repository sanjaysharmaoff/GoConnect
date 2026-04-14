package main

import (
	"log"
	"net/http"
)

func (app *application) internalServerError(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("msg=%q method=%s path=%s error=%q", "internal server error", r.Method, r.URL.Path, err.Error())
	writeJSONError(w, http.StatusInternalServerError, "the server encountered a problem")
}

func (app *application) badRequestResponse(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("msg=%q method=%s path=%s error=%q", "bad request error", r.Method, r.URL.Path, err.Error())
	writeJSONError(w, http.StatusBadRequest, err.Error())
}

func (app *application) notFoundResponse(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("msg=%q method=%s path=%s error=%q", "bad request error", r.Method, r.URL.Path, err.Error())
	writeJSONError(w, http.StatusNotFound, "Not found")
}

func (app *application) conflictRequestResponse(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("conflict error: method=%s path=%s error=%v", r.Method, r.URL.Path, err)
	writeJSONError(w, http.StatusConflict, "resource conflict")
}

func (app *application) unauthorizedErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("unauthorized  error: method=%s path=%s error=%v", r.Method, r.URL.Path, err)
	writeJSONError(w, http.StatusUnauthorized, "unauthorized")
}

func (app *application) unauthorizedBasicErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	w.Header().Set("WWW-Authenticate", `Basic realm="restricted",charset="UTF-8"`)
	log.Printf("basic unauthorized error: method=%s path=%s error=%v", r.Method, r.URL.Path, err)
	writeJSONError(w, http.StatusUnauthorized, "basic unauthorized")
}
