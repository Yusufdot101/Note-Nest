package utilities

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
	w.Header().Add("Content-Type", "application/json")
	// used MarshalIndent instead of Marshal to make it look nice in the terminal
	JSON, err := json.MarshalIndent(message, "", "\t")
	if err != nil {
		return err
	}

	JSON = append(JSON, '\n')
	w.WriteHeader(statusCode)
	_, err = w.Write(JSON)
	return err
}

func ReadJSON(w http.ResponseWriter, r *http.Request, dst any) error {
	// prevent large requests from hogging system resources
	const BodyMaxSizeBytes = 1_000_000
	r.Body = http.MaxBytesReader(w, r.Body, BodyMaxSizeBytes)

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

		case strings.HasPrefix(err.Error(), "json: unknown field "):
			fieldName := strings.TrimPrefix(err.Error(), "json: unknown field ")
			fieldName = strings.Trim(fieldName, "\"")
			return fmt.Errorf("body contains unknown key: %s", fieldName)

		case err.Error() == "http: request body too large":
			BodyMaxSizeMB := float64(BodyMaxSizeBytes) / 1_000_000
			return fmt.Errorf("body cannot be larger than %f.2MB ", BodyMaxSizeMB)

		case errors.Is(err, io.EOF):
			return fmt.Errorf("body cannot be empty")

		// this type of error is due to error from server end, should not happen in regular operations
		case errors.As(err, &invalidUnmarshalErr):
			panic(err)

		default:
			return err
		}
	}

	err = decoder.Decode(&struct{}{})
	if err != io.EOF {
		return fmt.Errorf("body must contain only one JSON")
	}

	return nil
}
