package adresse

import "fmt"

type Adresse struct {
	Name    string
	Vorname string
	Ort     string
}

func (n Adresse) String() string {
	return fmt.Sprintf("%v %v %v\n", n.Name, n.Vorname, n.Ort)
}
