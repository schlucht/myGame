package payment

import (
	"encoding/csv"
	"io"
	"os"
	"strconv"
	"time"

	"github.com/schlucht/myGame/pkg/helpers"
)

type Pay struct {
	Iban     string    `json:"iban"`
	BookedAt time.Time `json:"booked"`
	Amount   float64   `json:"amount"`
	Balance  float64   `json:"balance"`
	Text     string    `json:"text"`
	Valuta   time.Time `json:"valuta"`
	IsOut    bool      `json:"isOut"`
}

func ReadFile(filename string) ([]Pay, error) {
	f, err := os.Open(filename)
	helpers.ErrHandler("Ein Fehler", err)
	defer f.Close()

	csvReader := csv.NewReader(f)
	csvReader.Comma = ';'
	var first = true
	var items []Pay
	var pay Pay
	for {
		rec, err := csvReader.Read()

		if first {
			// ReadHeader(rec)
			first = false
			continue
		}

		if err == io.EOF {
			break
		}
		if err != nil {
			break
		}
		var isOut bool
		am, _ := strconv.ParseFloat(rec[3], 64)
		bal, _ := strconv.ParseFloat(rec[4], 64)
		t1, _ := time.Parse("2006-01-02", rec[1])
		t2, _ := time.Parse("2006-01-02", rec[5])
		if am < 0 {
			isOut = true
		}
		pay = Pay{
			Iban:     rec[0],
			BookedAt: t1,
			Amount:   am,
			Balance:  bal,
			Text:     rec[2],
			Valuta:   t2,
			IsOut:    isOut,
		}
		items = append(items, pay)
		helpers.ErrHandler("Fehler beim lesen der CSV", err)
	}
	// fmt.Printf("%+v\n", items)

	return items, nil
}

func ReadHeader(firstLine []string) (s []string) {
	s = firstLine
	// fmt.Println(s)
	return s
}
