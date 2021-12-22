package mail

import (
	"encoding/csv"
	"io"
	"log"
	"os"

	"github.com/pullemax/mail-sender/struts"
)

func GetRecipients(path string) []struts.Recipient {
	var recipients []struts.Recipient

	f, err := os.Open(path)
	if err != nil {
		log.Println("Error recovering the recipients. %s", err)
	} else {
		defer f.Close()
		csvReader := csv.NewReader(f)
		for {
			r, err := csvReader.Read()

			if err == io.EOF {
				break
			}
			if err != nil {
				log.Println("Error reading the recipient %s", err)
			}
			recipient := struts.Recipient{
				Email:  r[0],
				Param1: r[1],
				Param2: r[2],
			}
			recipients = append(recipients, recipient)
		}
	}

	return recipients
}
