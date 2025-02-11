package request

import (
	"butta/pkg/errors"
	"encoding/json"
	"io"
)

func JsonDeserializer(payload io.Reader, dest any, AllowUnknownFields bool) error {
	var decoder = json.NewDecoder(payload)

	if AllowUnknownFields != true {
		decoder.DisallowUnknownFields()
	}

	err := decoder.Decode(dest)
	if err == nil {
		return nil
	}

	var unmarshalTypeError *json.UnmarshalTypeError
	var syntaxError *json.SyntaxError

	if errors.As(err, &unmarshalTypeError) {
		return errors.New("data contains invalid value")
	}

	if errors.As(err, &syntaxError) || errors.Is(err, io.ErrUnexpectedEOF) {
		return errors.New("data contains badly-formed JSON")
	}

	if errors.Is(err, io.EOF) {
		return errors.New("data is empty")
	}

	return err
}
