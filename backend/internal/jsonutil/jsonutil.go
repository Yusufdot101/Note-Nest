package jsonutil

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type Message map[string]any

func WriteJSON(w http.ResponseWriter, message Message, statusCode int) error {
	JSON, err := json.Marshal(message)
	if err != nil {
		return err
	}

	w.WriteHeader(statusCode)
	_, err = w.Write(JSON)
	return err
}

func ReadJSON(w http.ResponseWriter, r *http.Request, dst any) error {
	// prevent large requests from hogging system resources
	const BODY_MAX_BYTES = 1_048_576
	http.MaxBytesReader(w, r.Body, BODY_MAX_BYTES)

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	err := decoder.Decode(dst)
	if err != nil {
		var syntaxErr *json.SyntaxError
		var unmarshalTypeErr *json.UnmarshalTypeError
		var invalidUnmarshalErr *json.InvalidUnmarshalError
		switch {
		case errors.As(err, &syntaxErr):
			return fmt.Errorf("body contains badly formed JSON at character: %d", syntaxErr.Offset)

		case errors.Is(err, io.ErrUnexpectedEOF):
			return fmt.Errorf("body contains badly formed JSON")

		// this is when different variable type is passed to a field, like field expecting int and
		// passed string
		case errors.As(err, &unmarshalTypeErr):
			if unmarshalTypeErr.Field != "" {
				return fmt.Errorf("body contains incorrect type for field: %s", unmarshalTypeErr.Field)
			}
			return fmt.Errorf("body contains incorrect type at character: %d", unmarshalTypeErr.Offset)

		case strings.HasPrefix(err.Error(), "json: unknown "):
			fieldName := strings.TrimPrefix(err.Error(), "json: unknown ")
			return fmt.Errorf("body contains unknown key: %s", fieldName)

		case err.Error() == "http: request body too large":
			return fmt.Errorf("body cannot be larger than %dMB ", BODY_MAX_BYTES)

		case errors.Is(err, io.EOF):
			return fmt.Errorf("body cannot be empty")

		// this type of error is due to error from server end, should not happen in regular operations
		case errors.As(err, &invalidUnmarshalErr):
			panic(err)

		default:
			return err
		}
	}

	return nil
}
