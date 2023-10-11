package model

import (
	"fmt"
	"time"
)

// This is a custom declared scaler
type Timestamp time.Time

func (t *Timestamp) UnmarshalGQL(v any) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("Timestamps must be formatted as strings, not %T", v)
	}

	parsed, err := time.Parse(time.RFC3339, str)
	if err != nil {
		return err
	}

	*t = Timestamp(parsed)
	return nil
}

func (t Timestamp) MarshalGQL() interface{} {
	return time.Time(t).Format(time.RFC3339)
}
