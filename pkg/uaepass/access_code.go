package uaepass

import (
	"fmt"
)

// Get the access code first for the next step
// https://docs.uaepass.ae/guides/authentication/web-application/1.-obtaining-the-oauth2-access-code
func RequestAccessCodeURL(url, responseType, redirectURI, clientID, state, scope, acrValues, language string) string {
	requestURL := fmt.Sprintf(
		"%s?response_type=%s&redirect_uri=%s&client_id=%s&state=%s&scope=%s&acr_values=%s&ui_locales=%s",
		url,
		responseType,
		redirectURI,
		clientID,
		state,
		scope,
		acrValues,
		language,
	)
	return requestURL
}
