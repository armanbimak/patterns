package setter

import (
	"encoding/json"
	"fmt"
)

func NewAgreementCursiveSetter() ISetter {
	return &agreementCursiveSetter{}
}

type agreementCursiveSetter struct{}

func (s *agreementCursiveSetter) SetData(pattern []byte, data json.RawMessage) {
	fmt.Println("Setting cursive data in given pattern.")
}
