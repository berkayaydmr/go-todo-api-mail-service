package utils

import (
	"encoding/base64"
	"strings"
)

func DecodeHash(hashed string) (string, error) {
	hashArray, err := base64.StdEncoding.DecodeString(hashed)
	hash := string(hashArray)
	splitMail := strings.TrimRight(hash, "\ufffd\ufffd\ufffdB\ufffd\ufffd\u001c\u0014\ufffd\ufffd\ufffdșo\ufffd$'\ufffdA\ufffdd\ufffd\ufffdL\ufffd\ufffd\ufffd\u001bxR\ufffdU")
	return splitMail, err
}
