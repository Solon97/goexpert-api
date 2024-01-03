package common

import (
	"fmt"
	"time"
)

type ISODate time.Time

func (d *ISODate) UnmarshalJSON(data []byte) error {
	// Remova as aspas do JSON
	strDate := string(data)
	strDate = strDate[1 : len(strDate)-1] // Remove as aspas

	// Converta a string para time.Time
	t, err := time.Parse("2006-01-02", strDate)
	if err != nil {
		return fmt.Errorf("invalid date format: %v", err)
	}

	// Atualize o valor do ISODate
	*d = ISODate(t)
	return nil
}

func (d ISODate) String() string {
	return time.Time(d).Format("2006-01-02")
}
