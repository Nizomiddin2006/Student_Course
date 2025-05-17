package kurs

import "fmt"


type Kurs struct {
    ID   int
    Nomi string
    Fakultet string
}


func YangiKurs(id int, nomi string, fakultet string) *Kurs {
    return &Kurs{
        ID:       id,
        Nomi:     nomi,
        Fakultet: fakultet,
    }
}

func (k *Kurs) Yoz() {
    fmt.Printf("ID: %d, Nomi: %s, Fakultet: %s\n", k.ID, k.Nomi, k.Fakultet)
}
