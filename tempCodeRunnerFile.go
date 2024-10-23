package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"
	"math"
)

type Trip struct {
	ID           int
	Layanan      string
	Destinasi    string
	Jemput       string
	Jarak        float64
	BeratBarang  float64
	TotalFare    int
}

const MAX = 1000

var dataTrip [MAX]Trip
var index int

func Input(text string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(text)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func main() {
	var choice int
	for {
		fmt.Println("Menu:")
		fmt.Println("1. Tambah Trip")
		fmt.Println("2. Hapus Trip")
		fmt.Println("3. Tampilkan Trip")
		fmt.Println("4. Cari Trip")
		fmt.Println("5. Urutkan Trip")
		fmt.Println("6. Keluar")
		choiceStr := Input("Pilih menu: ")
		choice, _ = strconv.Atoi(choiceStr)

		switch choice {
		case 1:
			tambahTrip()
			fmt.Println("================================================")
		case 2:
			hapusTrip()
			fmt.Println("================================================")
		case 3:
			tampilkanTrip()
			fmt.Println("================================================")
		case 4:
			cariTrip()
			fmt.Println("================================================")
		case 5:
			urutkanTrip()
			fmt.Println("================================================")
		case 6:
			fmt.Print("Anda telah keluar dari program.")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
			fmt.Println("================================================")
		}
	}
}

func tambahTrip() {
	fmt.Println("================================================")
	fmt.Println("Silakan pilih layanan yang ingin Anda gunakan:")
	fmt.Println("1. Motor\n2. Mobil\n3. Paket")
	optionStr := Input("Layanan yang dipilih: ")
	option, _ := strconv.Atoi(optionStr)

	fmt.Println("================================================")

	destinasi := Input("Masukkan Destinasi Tujuan: ")
	jemput := Input("Masukkan Destinasi Jemput: ")

	jarakStr := Input("Masukkan Besar Jarak (Km): ")
	jarak, _ := strconv.ParseFloat(jarakStr, 64)

	var berat float64
	if option == 3 {
		_ = Input("\nMasukkan Jenis Barang: ")

		beratStr := Input("Masukkan Berat Barang (Kg): ")
		berat, _ = strconv.ParseFloat(beratStr, 64)
	
		if berat > 5 {
			fmt.Println("\nBerat lebih dari 5 kg. Anda akan menggunakan layanan mobil antar paket.")
		}
	}

	totalFare := hitungTarif(option, jarak, berat)

	dataTrip[index] = Trip{
		ID:          index + 1,
		Layanan:     panggilLayanan(option),
		Destinasi:   destinasi,
		Jemput:      jemput,
		Jarak:       jarak,
		BeratBarang: berat,
		TotalFare:   totalFare,
	}
	index++
	
	fmt.Println("\nTotal fare: ", totalFare)
	fmt.Println("Data Trip berhasil ditambahkan!")
}

func hapusTrip() {
	idStr := Input("Masukkan ID Trip yang ingin dihapus: ")
	id, _ := strconv.Atoi(idStr)

	index := sequentialSearch(id)
	if index == -1 {
		fmt.Println("Trip tidak ditemukan.")
		return
	}

	for i := index; i < index-1; i++ {
		dataTrip[i] = dataTrip[i+1]
	}
	index--
	fmt.Println("Trip berhasil dihapus.")
}

func tampilkanTrip() {
	for i := 0; i < index; i++ {
		fmt.Printf("ID: %d, Layanan: %s, Destinasi: %s, Jemput: %s, Jarak: %.2f km, Berat Barang: %.2f kg, Total Tarif: %d\n",
			dataTrip[i].ID, dataTrip[i].Layanan, dataTrip[i].Destinasi, dataTrip[i].Jemput,
			dataTrip[i].Jarak, dataTrip[i].BeratBarang, dataTrip[i].TotalFare)
	}
}

func cariTrip() {
	idStr := Input("Masukkan ID Trip yang ingin dicari: ")
	id, _ := strconv.Atoi(idStr)

	index := binarySearch(id)
	if index == -1 {
		fmt.Println("Trip tidak ditemukan.")
		return
	}

	p := dataTrip[index]
	fmt.Printf("ID: %d, Layanan: %s, Destinasi: %s, Jemput: %s, Jarak: %.2f km, Berat Barang: %.2f kg, Total Tarif: %d\n",
		p.ID, p.Layanan, p.Destinasi, p.Jemput, p.Jarak, p.BeratBarang, p.TotalFare)
}

func urutkanTrip() {
	fmt.Println("Pilih kategori pengurutan:")
	fmt.Println("1. Jarak (Ascending)")
	fmt.Println("2. Jarak (Descending)")
	optionStr := Input("Opsi yang dipilih: ")
	option, _ := strconv.Atoi(optionStr)

	switch option {
	case 1:
		selectionSort(true)
	case 2:
		selectionSort(false)
	}
	tampilkanTrip()
}

func hitungTarif(option int, jarak, berat float64) int {
	var totalFare int
	if option == 1 {
		totalFare = int(FareMotor(jarak))
	} else if option == 2 {
		totalFare = int(FareMobil(jarak))
	} else if option == 3 {
		totalFare = int(FarePaket(jarak, berat))
	}
	return totalFare
}

func panggilLayanan(option int) string {
	switch option {
	case 1:
		return "Motor"
	case 2:
		return "Mobil"
	case 3:
		return "Paket"
	default:
		return "Tidak Diketahui"
	}
}

func FareMotor(jarak float64) (tarif float64) {
	const firstKm = 10000
	const nextKmFare = 3000

	tarif = firstKm
	additionalKm := jarak - 1
	if additionalKm > 0 {
		tarif += math.Ceil(additionalKm) * nextKmFare
	}

	return tarif
}

func FareMobil(jarak float64) (tarif float64) {
	const firstKm = 13000
	const nextKmFare = 3500

	tarif = firstKm
	additionalKm := jarak - 1
	if additionalKm > 0 {
		tarif += math.Ceil(additionalKm) * nextKmFare
	}

	return tarif
}

func FarePaket(jarak, berat float64) (tarif float64) {
	const firstKmMotor = 12000
	const firstKmMobil = 26000
	const nextKmFare = 3000
    
	if berat < 5 {
		tarif = firstKmMotor
	} else {
		tarif = firstKmMobil
	}

	additionalKm := jarak - 1
	if additionalKm > 0 {
		tarif += math.Ceil(additionalKm) * nextKmFare
	}

	return tarif
}

func sequentialSearch(id int) int {
	for i := 0; i < index; i++ {
		if dataTrip[i].ID == id {
			return i
		}
	}
	return -1
}

func binarySearch(id int) int {
	left, right := 0, index-1
	for left <= right {
		mid := left + (right-left)/2
		if dataTrip[mid].ID == id {
			return mid
		}
		if dataTrip[mid].ID < id {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

func selectionSort(ascending bool) {
	for i := 0; i < index-1; i++ {
		idx := i
		for j := i + 1; j < index; j++ {
			if ascending {
				if dataTrip[j].Jarak < dataTrip[idx].Jarak {
					idx = j
				}
			} else {
				if dataTrip[j].Jarak > dataTrip[idx].Jarak {
					idx = j
				}
			}
		}
		dataTrip[i], dataTrip[idx] = dataTrip[idx], dataTrip[i]
	}
}