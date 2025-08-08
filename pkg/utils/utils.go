package utils

import (
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Middleware func(w http.ResponseWriter, r *http.Request) *http.Request

func LogIfErr(e error, msg string) bool {
	if e != nil {
		log.Printf(msg+": %v\n", e)
	}
	return (e != nil)
}

func PanicIfErr(e error, msg string) {
	if e != nil {
		log.Panicf(msg+": %v\n", e)
	}
}

func QuitIfErr(e error, msg string) {
	if e != nil && !errors.Is(e, http.ErrServerClosed) {
		log.Fatalf(msg+": %v\n", e)
	}
}

func RespondFailure(w http.ResponseWriter, statusCode int, msg string) {
	w.WriteHeader(statusCode)
	var e error
	_, e = fmt.Fprint(w, msg)
	_ = LogIfErr(e, "Connection Error")
}

func RespondSuccess(w http.ResponseWriter, statusCode int, msg string) {
	w.WriteHeader(statusCode)
	var e error
	_, e = fmt.Fprint(w, msg)
	_ = LogIfErr(e, "Connection Error")
}

func ReflectAndLogErr(w http.ResponseWriter, statusCode int, e error, msg string) bool {
	if e != nil {
		_ = LogIfErr(e, msg)
		w.WriteHeader(statusCode)
		_, e = fmt.Fprint(w, msg)
		_ = LogIfErr(e, "Connection Error")
		return true
	}
	return false
}

func GetOrReflect(w http.ResponseWriter, r *http.Request, item string) (string, bool) {
	var val string
	val = r.Header.Get(item)
	var e error
	if len(val) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = fmt.Fprint(w, "No "+item+" Provided")
		_ = LogIfErr(e, "Connection Error")
		return "", true
	}
	return val, false
}

func SetMiddlewares(middlewares []Middleware) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		for i := range middlewares {
			middleware := middlewares[i]
			r = middleware(w, r)
			if r == nil {
				return
			}
		}
	}
}

func SetRoute(router *mux.Router, method string, route string, middlewares ...Middleware) {
	router.HandleFunc(route, SetMiddlewares(middlewares)).Methods(method)
}
