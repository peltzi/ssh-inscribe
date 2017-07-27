package signapi

import (
	"time"

	"github.com/aakso/ssh-inscribe/pkg/auth"
	"github.com/aakso/ssh-inscribe/pkg/keysigner"
	"github.com/dgrijalva/jwt-go"
)

const (
	TokenLifeSecs = 120
)

type AuthenticatorListEntry struct {
	Authenticator auth.Authenticator
	Default       bool
}

type SignApi struct {
	auth            map[string]auth.Authenticator
	authList        []AuthenticatorListEntry
	defaultAuth     []string
	signer          *keysigner.KeySignerService
	tkey            []byte
	defaultCertLife time.Duration
	maxCertLife     time.Duration
}

func New(
	authList []AuthenticatorListEntry,
	signer *keysigner.KeySignerService,
	tkey []byte,
	defaultlife time.Duration,
	maxlife time.Duration,
) *SignApi {

	authMap := map[string]auth.Authenticator{}
	for _, v := range authList {
		authMap[v.Authenticator.Name()] = v.Authenticator
	}

	return &SignApi{
		auth:            authMap,
		authList:        authList,
		signer:          signer,
		tkey:            tkey,
		defaultCertLife: defaultlife,
		maxCertLife:     maxlife,
	}
}

type SignClaim struct {
	AuthContext *auth.AuthContext
	jwt.StandardClaims
}
