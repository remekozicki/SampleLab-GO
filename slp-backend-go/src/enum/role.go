package enum

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

type Role int

const (
	RoleAdmin Role = iota
	RoleWorker
	RoleIntern
)

var roleToString = map[Role]string{
	RoleAdmin:  "ADMIN",
	RoleWorker: "WORKER",
	RoleIntern: "INTERN",
}

var stringToRole = map[string]Role{
	"ADMIN":  RoleAdmin,
	"WORKER": RoleWorker,
	"INTERN": RoleIntern,
}

// String returns the string representation
func (r Role) String() string {
	if val, ok := roleToString[r]; ok {
		return val
	}
	return "UNKNOWN"
}

// JSON marshaling
func (r Role) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.String())
}

func (r *Role) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	role, ok := stringToRole[strings.ToUpper(s)]
	if !ok {
		return fmt.Errorf("nieprawidłowa rola: %s", s)
	}
	*r = role
	return nil
}

// GORM DB support
func (r Role) Value() (driver.Value, error) {
	return int64(r), nil
}

func (r *Role) Scan(value interface{}) error {
	intVal, ok := value.(int64)
	if !ok {
		return errors.New("błąd konwersji roli z bazy")
	}
	*r = Role(intVal)
	return nil
}
