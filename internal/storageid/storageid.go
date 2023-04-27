// Package storageid provides the function for creating a new storage id.
package storageid

import (
	"fmt"
	"math/rand"
	"regexp"
)

// valid id characters (lower alpha and numeric)
// regex must include 1 or more lower alpha or numeric characters
var validId = regexp.MustCompile(`^[a-z0-9]+$`)

// set of valid id characters used for id generation
var validIdCharacters = []rune("abcdefghijklmnopqrstuvwxyz0123456789")

// New creates a new storageid with the specified prefix and length.
func New(prefix string, length uint) (string, error) {

	// validate prefix
	if !validId.MatchString(prefix) {
		return "", fmt.Errorf("storageid.New: invalid prefix %s", prefix)
	}

	// create byte slice, allocating space for remaining id characters
	b := make([]rune, length-uint(len(prefix)))
	for i := range b {
		b[i] = validIdCharacters[rand.Intn(len(validIdCharacters))]
	}

	storageId := prefix + string(b)
	if !validId.MatchString(storageId) || len(storageId) != int(length) {
		return "", fmt.Errorf("storageid.New: invalid id %s", storageId)
	}

	return storageId, nil
}
