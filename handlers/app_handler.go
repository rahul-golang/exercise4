package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"text/template"

	"github.com/rahul-golang/exercise4/models"
	"github.com/rahul-golang/exercise4/services"
)

//AppHandlers handlers
type AppHandlers struct {
	service services.AppService
}

//NewAppHandlers creates dependancy
func NewAppHandlers(service services.AppService) *AppHandlers {
	return &AppHandlers{service: service}
}

//ServeApplication servres html page handlers
func (appHandlers *AppHandlers) ServeApplication(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/layout.html"))
	tmpl.Execute(w, nil)
}

//GetFormatedDateTime handler for crowl application
func (appHandlers *AppHandlers) GetFormatedDateTime(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	dateTime := models.DateTime{}

	err := ReadInput(req.Body, &dateTime)
	if err != nil {
		WriteHTTPError(w, err, http.StatusBadRequest)
		return
	}

	resp, err := appHandlers.service.GetFormatedDateTime(ctx, dateTime)
	if err != nil {
		WriteHTTPError(w, err, http.StatusBadRequest)
		return
	}

	WriteOKResponse(w, resp)

}

//ReadInput from the body
func ReadInput(rBody io.ReadCloser, input interface{}) error {
	decoder := json.NewDecoder(rBody)
	err := decoder.Decode(input)
	return err
}

//WriteHTTPError return HTTP Error Message
func WriteHTTPError(w http.ResponseWriter, err error, statusCode int) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	http.Error(w, err.Error(), statusCode)
}

//WriteOKResponse as a standard JSON response with StatusOK
func WriteOKResponse(w http.ResponseWriter, m interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(m); err != nil {
		WriteHTTPError(w, err, http.StatusInternalServerError)
	}
}
