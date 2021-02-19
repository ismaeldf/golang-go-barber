package middlewares

import (
	"context"
	"github.com/urfave/negroni"
	"ismaeldf/golang-gobarber/modules/users/providers/TokenProvider/models"
	"net/http"
)

const ContextUserKey string = "Id"

type data struct {
	Id string
}

func setContextData(r *http.Request, d *data) (ro *http.Request) {
	ctx := r.Context()
	ctx = context.WithValue(ctx, ContextUserKey, d)
	ro = r.WithContext(ctx)
	return
}

func GetUserIdContext(r *http.Request) string {
	d := *r.Context().Value(ContextUserKey).(*data)
	return d.Id
}

func EnsureAuthenticated(tokenProvider models.ITokenProvider) negroni.Handler {
	return negroni.HandlerFunc(func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		authHeader := r.Header.Get("authorization")

		if authHeader == "" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		id, err := tokenProvider.DecodeToken(authHeader)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		d := data{Id: *id}

		r = setContextData(r, &d)

		next(w, r)
	})
}
