package payment

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"

	"github.com/schlucht/myGame/pkg/helpers"
)

type Pay struct {
	Iban     string
	BookedAt time.Time
	Amount   float64
	Balance  float64
	Text     string
	Valuta   time.Time
}

func ReadFile(filename string) (string, error) {
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
			ReadHeader(rec)
			first = false
			continue
		}
		am, _ := strconv.ParseFloat(rec[2], 64)
		bal, _ := strconv.ParseFloat(rec[3], 64)
		t1, _ := time.Parse("2006-01-02", rec[1])
		t2, _ := time.Parse("2006-01-02", rec[5])
		pay = Pay{
			Iban:     rec[0],
			BookedAt: t1,
			Amount:   am,
			Balance:  bal,
			Text:     rec[4],
			Valuta:   t2,
		}
		items = append(items, pay)
		helpers.ErrHandler("Fehler beim lesen der CSV", err)
		if err == io.EOF {
			break
		}
	}
	fmt.Printf("%+v\n", len(items))

	return " ", nil
}

func ReadHeader(firstLine []string) (s []string) {
	s = firstLine
	fmt.Println(s)
	return s
}
