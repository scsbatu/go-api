package helpers

import (
	"fmt"
	"github.com/google/uuid"
)

func StringToUUIDByte(id string) ([]byte, error) {
	uuID, err := uuid.Parse(id)
	if err != nil {
		return nil, fmt.Errorf("unable to convert the string to UUID")
	}
	idInBinary, err := uuID.MarshalBinary()
	if err != nil {
		return nil, fmt.Errorf("unable to convert the UUID")
	}
	return idInBinary, nil
}
