package main

type kendaraan struct {
	roda      int
	kecepatan int
}

type mobil struct {
	kendaraan
}

func (mobil *mobil) tambahKecepatan(kecepatanBaru int) {
	mobil.kecepatan = mobil.kecepatan + kecepatanBaru
}

func (mobil *mobil) berjalan() {
	mobil.tambahKecepatan(10)
}

func main() {

	mobilCepat := mobil{}

	mobilCepat.berjalan()
	mobilCepat.berjalan()
	mobilCepat.berjalan()

	mobilLamban := mobil{}

	mobilLamban.berjalan()
}