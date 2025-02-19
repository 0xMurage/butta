package authn

import (
	"butta/internal/authn/repository"
	"butta/internal/pkg/config"
	"butta/pkg/crypto/argon2id"
	"butta/pkg/crypto/hmac"
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"math/rand/v2"
	"strconv"
	"time"
)

func authWithBasicCredentials(ctx context.Context, cf config.Session, psqlPool *pgxpool.Pool, cred *BasicAuthCredentialsDto) (sessionId string, err error) {
	queries := repository.New(psqlPool)

	authMethod, err := queries.GetUserWithBasicPasswordAuth(ctx, cred.Username)
	if nil != err {
		return "", err
	}

	//compare password hash
	err = argon2id.CompareHashAndPassword(authMethod.SecretHash.String, cred.Password)

	if err != nil {
		return "", err
	}

	return generateSecureSessionId(cf, authMethod.UserProfileID.String(), authMethod.ID.String())
}

func generateSecureSessionId(cf config.Session, profileId, accountId string) (string, error) {

	expiryTimestamp := time.Now().Add(time.Duration(cf.Lifetime) * time.Minute).UnixMilli()

	expiry := strconv.FormatInt(expiryTimestamp, 10)

	sessionId := fmt.Sprintf("%s|%s|%s", accountId, profileId, expiry)

	//sign the data
	mac := hmac.New([]byte(cf.Secret))

	signature, err := mac.Sign(sessionId)

	if err != nil {
		return "", err
	}

	//append the signature on the output
	signedSessionId := fmt.Sprintf("%s|%s|%s|%s", accountId, profileId, expiry, signature)

	return signedSessionId, nil
}

func generatePasswordResetLink(cf config.Config, username string) (string, error) {
	mac := hmac.New([]byte(cf.Session.Secret))

	token := strconv.Itoa(int(rand.Uint32())) //adds just a bit of randomness to the url

	parsedUrl := cf.App.Url.JoinPath("password-reset", token)

	params := parsedUrl.Query()
	params.Set("username", username)

	parsedUrl.RawQuery = params.Encode()

	return mac.SignUrl(*parsedUrl, 15*time.Minute) //expires after 15 minus
}
