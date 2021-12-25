package struts

import (
	"encoding/csv"
	"io"
	"os"
)

type Recipient struct {
	Email  string
	Param1 string
	Param2 string
}

func (r *Recipient) GetRecipients(path string) ([]Recipient, error) {
	var recipients []Recipient

	f, err := os.Open(path)

	if err != nil {
		return nil, err
	} else {
		defer f.Close()
		csvReader := csv.NewReader(f)
		for {
			r, err := csvReader.Read()

			if err == io.EOF {
				break
			}
			if err != nil {
				return nil, err
			}
			recipient := Recipient{
				Email:  r[0],
				Param1: r[1],
				Param2: r[2],
			}
			recipients = append(recipients, recipient)
		}
	}

	return recipients, nil
}
