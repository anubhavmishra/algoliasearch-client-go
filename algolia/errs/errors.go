package errs

import (
	"errors"
	"fmt"
)

var (
	ErrMissingObjectIDErr = errors.New("missing `objectID` field")
	ErrNoMoreHostToTryErr = errors.New("all hosts have been contacted unsuccessfully")
	IndexAlreadyExistsErr = errors.New("destination index already exists, please delete it first as the CopyIndex cannot hold the responsibility of modifying the destination index")
	NoMoreHitsErr         = errors.New("no more hits")
	NoMoreRulesErr        = errors.New("no more rules")
	NoMoreSynonymsErr     = errors.New("no more synonyms")
	SameAppIDErr          = errors.New("indices cannot target the same application ID, please use Client.CopyIndex for same-app index copy instead")
)

func ErrJSONDecode(data []byte, t string) error {
	return fmt.Errorf("cannot decode value %s to type %s", string(data), t)
}
