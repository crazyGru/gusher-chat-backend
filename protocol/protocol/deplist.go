package protocol

import (
	"encoding/json"
	"fmt"
)

type DepList []int

// Scan used to convert json into slice of integers
func (d *DepList) Scan(src interface{}) error {
	switch src.(type) {
	case []byte:
		return json.Unmarshal(src.([]byte), &d)
	case nil:
		return nil
	default:
		return fmt.Errorf("incorrect depends_on value in db: %v", src)
	}
}
