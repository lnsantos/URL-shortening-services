package domain

import (
	"URLshortening/data"
	"encoding/base64"
)

func RegisterNewUrl(
	url string,
) (string, error) {

	encode := base64.StdEncoding.EncodeToString([]byte(url))
	record := data.RegisterUrlDTO{Original: url, Short: encode}

	if url, err := data.SearchUrlByShort(encode); url != "" || err != nil && err.Error() != "mongo: no documents in result" {
		return "", err
	}

	err := data.RegisterUrl(record)

	return encode, err
}
