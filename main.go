package main

import (
	"errors"
	"fmt"
)

type Animal struct {
	Name string
	legs int
}

type ToString interface {
	PrintAdr() (string, error)
}

func (n Adresse) PrintAdr() (string, error) {
	if n.Ort == "" {
		return fmt.Sprintf("%v %v\n", n.Name, n.Vorname), nil
	}
	return fmt.Sprintf("%v %v %v\n", n.Name, n.Vorname, n.Ort), nil
}
func (n Animal) PrintAdr() (string, error) {
	if n.Name == "" {
		return fmt.Sprintf("%v\n", n.Name), nil
	}
	return "", errors.New("Kein Name vorhanden")
}

func main() {

	adr := Adresse{
		Name:    "Schmid",
		Vorname: "Lothar",
	}
	fmt.Println(adr.PrintAdr())
	// fmt.Printf("%v %v\n", adr.Name, adr.Vorname)

	var ort string

	fmt.Printf("Bitte Ort eingeben: ")
	fmt.Scan(&ort)
	adr.Ort = ort
	pr, _ := adr.PrintAdr()
	fmt.Printf("%v\n", pr)
	fmt.Printf("%v", adr.String())
}
