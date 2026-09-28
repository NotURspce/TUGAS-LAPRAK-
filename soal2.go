package main
import "fmt"

func main() {
    var Nama, NIM, Kelas string
    fmt.Print("Masukan Nama : ")
    fmt.Scan(&Nama)
    fmt.Print("Masukan NIM : ")
    fmt.Scan(&NIM)
    fmt.Print("Masukan Kelas : ")
    fmt.Scan(&Kelas)
    
    // Menggunakan fmt.Println dan koma
    fmt.Println("Perkenalkan saya adalah", Nama, "salah satu mahasiswa Prodi S1-IF dari kelas", Kelas, "dengan nomor induk mahasiswa", NIM+".")
}
