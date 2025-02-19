package authn

import (
	"butta/internal/authn/repository"
	"butta/internal/pkg/config"
	"butta/pkg/crypto/argon2id"
	"butta/pkg/crypto/hmac"
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
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
