package token

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/macrat/ldapin/config"
)

type OIDCClaims struct {
	jwt.StandardClaims

	Type     string `json:"typ"`
	AuthTime int64  `json:"auth_time,omitempty"`
}

func (claims OIDCClaims) Validate(issuer *config.URL, audience string) error {
	if err := claims.StandardClaims.Valid(); err != nil {
		return err
	}

	if claims.Issuer != issuer.String() {
		return UnexpectedIssuerError
	}

	if claims.Audience != audience {
		return UnexpectedAudienceError
	}

	return nil
}
