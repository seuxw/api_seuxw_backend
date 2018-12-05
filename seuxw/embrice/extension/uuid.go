package extension

import (
	"github.com/satori/go.uuid"
)

// NewUUIDString .
func NewUUIDString() string {
	u, _ := uuid.NewV4()
	return u.String()
}
