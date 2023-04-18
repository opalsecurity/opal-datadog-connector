package datadogconnector

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"strings"
)

// Cursor is a struct that contains the next page number
type Cursor struct {
	NextPage int
}

func encodeCursor(cursor *Cursor) (string, error) {
	var buf bytes.Buffer
	encoder := base64.NewEncoder(base64.StdEncoding, &buf)
	err := json.NewEncoder(encoder).Encode(cursor)
	if err != nil {
		return "", err
	}
	encoder.Close()
	return buf.String(), nil
}

func decodeCursor(encodedCursor string) (*Cursor, error) {
	var cursor Cursor
	err := json.NewDecoder(
		base64.NewDecoder(
			base64.StdEncoding,
			strings.NewReader(encodedCursor),
		),
	).Decode(&cursor)
	if err != nil {
		return nil, err
	}
	return &cursor, nil
}
