package talaba

import "fmt"

type Talaba struct {
    ID     int
    Ism    string
    Yosh   int
    Fakultet string
}

func YangiTalaba(id int, ism string, yosh int, fakultet string) *Talaba {
    return &Talaba{
        ID:       id,
        Ism:      ism,
        Yosh:     yosh,
        Fakultet: fakultet,
    }
}

func (t *Talaba) Yoz() {
    fmt.Printf("ID: %d, Ism: %s, Yosh: %d, Fakultet: %s\n", t.ID, t.Ism, t.Yosh, t.Fakultet)
}
