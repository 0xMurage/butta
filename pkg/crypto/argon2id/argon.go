package argon2id

import (
	"butta/pkg/errors"
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"fmt"
	"golang.org/x/crypto/argon2"
	"runtime"
	"strings"
)

type Params struct {
	MemoryCost  uint32
	TimeCost    uint32
	Parallelism uint8
	SaltLength  uint32
	HashLength  uint32
	Version     uint8
}

func NewParams() *Params {
	return &Params{
		MemoryCost:  32,
		TimeCost:    2,
		Parallelism: uint8(runtime.NumCPU()),
		SaltLength:  16,
		HashLength:  32,
		Version:     argon2.Version,
	}
}

func randomCryptoBytes(length uint32) ([]byte, error) {
	byts := make([]byte, length)
	if _, err := rand.Read(byts); err != nil {
		return nil, err
	}
	return byts, nil
}

func HashPassword(password string, params *Params) (string, error) {
	if params == nil {
		params = NewParams()
	}

	salt, err := randomCryptoBytes(params.SaltLength)

	if err != nil {
		return "", errors.Wrap(err, "failed to generate random crypto bytes")
	}

	key := argon2.IDKey([]byte(password), salt, params.TimeCost, params.MemoryCost, params.Parallelism, params.HashLength)

	b64Salt := base64.RawStdEncoding.EncodeToString(salt)
	b64Key := base64.RawStdEncoding.EncodeToString(key)

	output := fmt.Sprintf("$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s", params.Version, params.MemoryCost, params.TimeCost, params.Parallelism, b64Salt, b64Key)
	return output, nil
}

func CompareHashAndPassword(hash, password string) error {
	params, salt, aKey, err := passwordHashDecode(hash)

	if err != nil {
		return errors.Wrap(err, "failed to decode the hash")
	}

	bKey := argon2.IDKey([]byte(password), salt, params.TimeCost, params.MemoryCost, params.Parallelism, params.HashLength)

	if subtle.ConstantTimeCompare(aKey, bKey) == 1 {
		return nil
	}
	return errors.New("crypto/argon2id: hash does not match the password")
}

func passwordHashDecode(hash string) (params *Params, salt, key []byte, err error) {
	parts := strings.Split(hash, "$")

	if len(parts) < 6 {
		return nil, nil, nil, errors.New("crypto/argon2id: invalid hash format")
	}

	if parts[1] != "argon2id" {
		return nil, nil, nil, errors.New("crypto/argon2id: unsupported hash format")
	}

	params = NewParams() //default params

	var version uint8

	_, err = fmt.Sscanf(parts[2], "v=%d", &version)

	if err != nil {
		return nil, nil, nil, errors.New("crypto/argon2id: unsupported hash format")
	}

	if version != params.Version {
		return nil, nil, nil, errors.New("crypto/argon2id: incompatible argon2 hash function version")
	}

	_, err = fmt.Sscanf(parts[3], "m=%d,t=%d,p=%d", &params.MemoryCost, &params.TimeCost, &params.Parallelism)

	if err != nil {
		return nil, nil, nil, errors.New("crypto/argon2id: unsupported hash format")
	}

	salt, err = base64.RawStdEncoding.DecodeString(parts[4])
	if err != nil {
		return nil, nil, nil, errors.Wrap(err, "crypto/argon2id")
	}

	key, err = base64.RawStdEncoding.DecodeString(parts[5])

	if err != nil {
		return nil, nil, nil, err
	}

	return params, salt, key, nil
}
