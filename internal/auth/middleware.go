package auth

import (
	"context"
	"fmt"
	"github.com/ingkiller/hackernews/internal/user"
	"log"
	"net/http"
)

var userCtxKey = &contextKey{"user"}

type contextKey struct {
	name string
}

func Middleware() func(http.Handler) http.Handler {

	fmt.Print("Middleware11: %v")
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			header := r.Header.Get("Authorization")

			// Allow unauthenticated users in
			if header == "" {
				next.ServeHTTP(w, r)
				return
			}

			//validate jwt token
			tokenStr := header
			fmt.Print("tokenStr----------------------: %v", tokenStr)
			/*
				username, err := jwt.ParseToken(tokenStr)
				if err != nil {
					http.Error(w, "Invalid token", http.StatusForbidden)
					return
				}
			*/
			if len(tokenStr) < 10 {
				http.Error(w, "Invalid token 10", http.StatusForbidden)
				return
			}
			// create user and check if user exists in db
			user := user.User{Username: "username"}
			id, err := user.GetUserIdByUsername("username")
			if err != nil {
				next.ServeHTTP(w, r)
				return
			}
			user.Id = id
			// put it in context
			ctx := context.WithValue(r.Context(), userCtxKey, &user)
			log.Printf("put context: %v", user)
			// and call the next with our new context
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

func ForContext(ctx context.Context) *user.User {
	raw, _ := ctx.Value(userCtxKey).(*user.User)
	return raw
}
