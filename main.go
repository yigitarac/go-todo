package main

import (
	"encoding/json"
	"fmt"
	"math"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

type Gorev struct {
	Isim string
	Sure time.Time
}

var gorevListesi []Gorev

func main() {
	path, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Yol bulunamadı")
		return
	}
	dosyaIcerigi, err := os.ReadFile(filepath.Join(path, ".todo.json"))
	if err != nil {
		fmt.Println("CEYSIN DOSYASI OKUNAMADI")
	} else {
		err = json.Unmarshal(dosyaIcerigi, &gorevListesi)
		if err != nil {
			fmt.Println("CEYSIN DOSYASI TEMİZLENEMEDİ")
		}
	}
	if len(os.Args) == 1 {
		fmt.Println("todo add “Görev Adı“ süre(gün)")
		fmt.Println("todo list")
		fmt.Println("todo remove görev no")
	} else if len(os.Args) >= 2 {
		islem := os.Args[1]
		switch islem {
		case "add":
			if len(os.Args) == 4 {
				alinanGorevAdi := os.Args[2]
				alinanGorevSuresi := os.Args[3]
				gorevEkle(alinanGorevAdi, alinanGorevSuresi, path)
			} else {
				fmt.Println("Eksik ya da hatalı giriş yaptınız!")
				return
			}
		case "list":
			gorevleriListele()
		case "remove":
			if len(os.Args) == 3 {
				silinecekNumara := os.Args[2]
				silinecekIndex, err := strconv.Atoi(silinecekNumara)
				if err != nil {
					fmt.Println("Geçersiz numara")
					return
				}
				gorevSil(silinecekIndex, path)
			} else {
				fmt.Println("Eksik ya da hatalı giriş yaptınız!")
				return
			}
		}
	}
}

func gorevEkle(gorevAdi string, gorevSuresi string, yol string) {
	sayiOlarakGorevSuresi, err := strconv.Atoi(gorevSuresi)
	if err != nil {
		fmt.Println("Hatalı süre girişi")
		return
	}
	zaman := time.Now().Add(24 * time.Hour * time.Duration(sayiOlarakGorevSuresi))
	eklenenGorev := Gorev{
		Isim: gorevAdi,
		Sure: zaman,
	}
	gorevListesi = append(gorevListesi, eklenenGorev)
	fmt.Println("Görev başarıyla eklendi!")
	donusturulenListe, err := json.Marshal(gorevListesi)
	if err != nil {
		fmt.Println("Liste dönüştürülemedi")
		return
	}
	os.WriteFile(filepath.Join(yol, ".todo.json"), donusturulenListe, 0644)
}
func gorevSil(silinecekGorev int, yol string) {
	if len(gorevListesi) == 0 {
		fmt.Println("Silinecek görev yok!")
		return
	}
	index := silinecekGorev - 1
	if silinecekGorev > len(gorevListesi) || silinecekGorev <= 0 {
		fmt.Println("Geçersiz numara!")
		return
	}
	gorevListesi = append(gorevListesi[:index], gorevListesi[index+1:]...)
	fmt.Println("Görev silindi!")
	donusturulenListe, err := json.Marshal(gorevListesi)
	if err != nil {
		fmt.Println("Liste dönüştürülemedi")
		return
	}
	os.WriteFile(filepath.Join(yol, ".todo.json"), donusturulenListe, 0644)
}
func gorevleriListele() {
	if len(gorevListesi) == 0 {
		fmt.Println("HENÜZ BİR GÖREV EKLENMEDİ!")
	}
	for j := range gorevListesi {
		gunselSure := int(math.Round(time.Until(gorevListesi[j].Sure).Hours() / 24))
		fmt.Printf("%d. %s (Süre: %d Gün)\n", j+1, gorevListesi[j].Isim, gunselSure)
	}
}
