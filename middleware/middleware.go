package middleware

import "net/http"

func ContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//Aqui estamos falando que o tipo de retorno será json
		w.Header().Set("Content-type", "aplication/json")
		next.ServeHTTP(w, r)
	})

}
