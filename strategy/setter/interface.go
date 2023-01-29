package setter

import "encoding/json"

type ISetter interface {
	SetData(pattern []byte, data json.RawMessage)
}
