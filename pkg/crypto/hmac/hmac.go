package hmac

import (
	"butta/pkg/errors"
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
	"net/url"
	"strconv"
	"time"
)

type Signer struct {
	signingKey []byte
}

func (hmacSigner *Signer) hashBytes(bytes []byte) ([]byte, error) {
	hash := hmac.New(sha512.New384, hmacSigner.signingKey)

	_, err := hash.Write(bytes)
	if err != nil {
		return nil, err
	}
	return hash.Sum(nil), nil

}

func (hmacSigner *Signer) Sign(message string) (string, error) {

	hash, err := hmacSigner.hashBytes([]byte(message))

	if err != nil {
		return "", errors.Wrap(err, "hmacSigner.hashBytes")
	}

	return hex.EncodeToString(hash), nil
}

// Verify Returns an error if the message doesn't match the signature
func (hmacSigner *Signer) Verify(message string, signature string) error {
	decodedSignature, err := hex.DecodeString(signature)

	if err != nil {
		return errors.Wrap(err, "unable to decode the signature")
	}

	hash, err := hmacSigner.hashBytes([]byte(message))

	if err != nil {
		return errors.Wrap(err, "unable to hash the message for comparison")
	}

	if hmac.Equal(decodedSignature, hash) {
		return nil
	}
	return errors.New("signature verification failed")
}

func (hmacSigner *Signer) SignUrl(url url.URL, expirationTime time.Duration) (string, error) {

	//this is a workaround to always ensure the query params are sorted
	url.RawQuery = url.Query().Encode()

	if expirationTime > 0 {
		expires := time.Now().Add(expirationTime).UnixMilli()

		query := url.Query()
		query.Set("expires", strconv.FormatInt(expires, 10))

		url.RawQuery = query.Encode()
	}

	signature, err := hmacSigner.Sign(url.String())
	if err != nil {
		return "", err
	}

	query := url.Query()
	query.Set("signature", signature)
	url.RawQuery = query.Encode()

	return url.String(), nil

}

// VerifySignedUrl returns an error if the message doesn't match the signature
func (hmacSigner *Signer) VerifySignedUrl(url url.URL) error {

	if !url.Query().Has("signature") {
		return errors.New("missing signature")
	}

	if url.Query().Has("expires") {
		//verify it has not expired
		expiry := url.Query().Get("expires")
		unixTime, err := strconv.ParseInt(expiry, 10, 64)

		if err != nil {
			return errors.Wrap(err, "unable to parse the expiration time")
		}

		if time.Now().UnixMilli() > unixTime {
			return errors.New("url signature expired")
		}
	}

	//sign without the signature and see if the signature matches
	signature := url.Query().Get("signature")

	query := url.Query()
	query.Del("signature")
	url.RawQuery = query.Encode()

	return hmacSigner.Verify(url.String(), signature)

}

func New(key []byte) *Signer {
	return &Signer{
		signingKey: key,
	}
}
