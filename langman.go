package langman

import (
	"net/http"
	"strings"

	"github.com/aofei/air"
)

// GasConfig is a set of configurations for the `Gas`.
type GasConfig struct {
	// ParamName is the name of the param used to cover the original
	// Accept-Language header.
	//
	// Default: "accept-language"
	ParamName string

	// CookieName is the name of the cookie used to cover the original
	// Accept-Language header.
	//
	// Default: "accept-language"
	CookieName string

	// CookieMaxAge is the max-age of the cookie used to cover the original
	// Accept-Language header.
	//
	// Default: 0
	CookieMaxAge int

	Skippable func(*air.Request, *air.Response) bool
}

// Gas returns an `air.Gas` that is used to manage the Accept-Language header
// based on the gc.
func Gas(gc GasConfig) air.Gas {
	paramName := gc.ParamName
	if paramName == "" {
		paramName = "accept-language"
	}

	cookieName := gc.CookieName
	if cookieName == "" {
		cookieName = "accept-language"
	}

	return func(next air.Handler) air.Handler {
		return func(req *air.Request, res *air.Response) error {
			if gc.Skippable != nil && gc.Skippable(req, res) {
				return next(req, res)
			}

			param := req.Param(paramName)
			if param == nil {
				cookie := req.Cookie(cookieName)
				if cookie != nil {
					req.Header.Set(
						"Accept-Language",
						cookie.Value,
					)
				}

				return next(req, res)
			}

			als := make([]string, len(param.Values))
			for i, al := range param.Values {
				als[i] = al.String()
			}

			al := strings.Join(als, ",")
			req.Header.Set("Accept-Language", al)
			res.SetCookie(&http.Cookie{
				Name:   cookieName,
				Value:  al,
				Path:   "/",
				MaxAge: gc.CookieMaxAge,
			})

			return next(req, res)
		}
	}
}
