package exception

import (
	"CptZeru/go-restful-api/helper"
	"CptZeru/go-restful-api/model/web"
	"net/http"

	"github.com/go-playground/validator"
)

func ErrorHandler(w http.ResponseWriter, r *http.Request, err interface{}) {
	if notFoundError(w, r, err) {
		return
	}

	if validationError(w, r, err) {
		return
	}
	internalServerError(w, r, err)
}

func internalServerError(w http.ResponseWriter, r *http.Request, err interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)

	webResponse := web.WebResponse{
		Code:   http.StatusInternalServerError,
		Status: "Internal Server Error",
		Data:   err,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func notFoundError(w http.ResponseWriter, r *http.Request, err interface{}) bool {
	exception, ok := err.(NotFoundError)
	if ok {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)

		webResponse := web.WebResponse{
			Code:   http.StatusNotFound,
			Status: "Not Found",
			Data:   exception.Error,
		}

		helper.WriteToResponseBody(w, webResponse)
		return true
	}
	return false
}

func validationError(w http.ResponseWriter, r *http.Request, err interface{}) bool {
	exception, ok := err.(validator.ValidationErrors)
	if ok {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)

		webResponse := web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Data:   exception.Error(),
		}

		helper.WriteToResponseBody(w, webResponse)
		return true
	}
	return false
}
