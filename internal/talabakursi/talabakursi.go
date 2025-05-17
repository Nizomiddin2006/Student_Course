package talabakursi
import (
    "fmt"
    "golang-template/internal/talaba"
    "golang-template/internal/kurs"
)

func Talabakursi(t *talaba.Talaba, k *kurs.Kurs) {
    fmt.Printf("%s talaba %s kursiga yozildi.\n", t.Ism, k.Nomi)
}
