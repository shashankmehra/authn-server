package meta

import (
	"net/http"

	"github.com/keratin/authn-server/api"
)

func getConfiguration(app *api.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		api.WriteJSON(w, http.StatusOK, map[string]interface{}{
			"issuer":                                app.Config.AuthNURL.String(),
			"response_types_supported":              []string{"id_token"},
			"subject_types_supported":               []string{"public"},
			"id_token_signing_alg_values_supported": []string{"RS256"},
			"claims_supported":                      []string{"iss", "sub", "aud", "exp", "iat", "auth_time"},
			"jwks_uri":                              app.Config.AuthNURL.String() + "/jwks",
		})
	}
}
