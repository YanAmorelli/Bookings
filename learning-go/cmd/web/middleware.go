package main

import (
	"net/http"
	"github.com/justinas/nosurf"
)

func NoSurf(next http.Handler) http.Handler {
	csfrHandler := nosurf.New(next)

	csfrHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path: "/",
		Secure: false,
		SameSite: http.SameSiteLaxMode,
	})
	return csfrHandler
}