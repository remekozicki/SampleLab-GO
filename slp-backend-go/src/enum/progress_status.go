package enum

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

type ProgressStatus int

const (
	ProgressInProgress ProgressStatus = iota
	ProgressDone
)

var statusToString = map[ProgressStatus]string{
	ProgressInProgress: "IN_PROGRESS",
	ProgressDone:       "DONE",
}

var stringToStatus = map[string]ProgressStatus{
	"IN_PROGRESS": ProgressInProgress,
	"DONE":        ProgressDone,
}

// String returns the string representation of the status
func (s ProgressStatus) String() string {
	if val, ok := statusToString[s]; ok {
		return val
	}
	return "UNKNOWN"
}

// MarshalJSON converts the enum to JSON string
func (s ProgressStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}

// UnmarshalJSON parses the JSON string back into the enum
func (s *ProgressStatus) UnmarshalJSON(b []byte) error {
	var str string
	if err := json.Unmarshal(b, &str); err != nil {
		return err
	}
	status, ok := stringToStatus[strings.ToUpper(str)]
	if !ok {
		return fmt.Errorf("nieprawidłowy status: %s", str)
	}
	*s = status
	return nil
}

// Value implements the driver.Valuer interface for GORM
func (s ProgressStatus) Value() (driver.Value, error) {
	return int64(s), nil
}

// Scan implements the sql.Scanner interface for GORM
func (s *ProgressStatus) Scan(value interface{}) error {
	intVal, ok := value.(int64)
	if !ok {
		return errors.New("błąd konwersji statusu z bazy")
	}
	*s = ProgressStatus(intVal)
	return nil
}

func ConvertProgressStatus(input string) (ProgressStatus, error) {
	status, ok := stringToStatus[strings.ToUpper(strings.TrimSpace(input))]
	if !ok {
		return 0, fmt.Errorf("nieprawidłowy status: %s", input)
	}
	return status, nil
}
