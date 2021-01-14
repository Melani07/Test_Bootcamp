package main

import "fmt"

func rerata(kel []int) int {
	jumlah := 0
	for i := 0; i < len(kel); i++ {
		jumlah = jumlah + kel[i]
	}
	return jumlah / len(kel)
}

func total(kel []int) int {
	jumlah := 0
	for i := 0; i < len(kel); i++ {
		jumlah = jumlah + kel[i]
	}
	return jumlah
}

func sorting(kel []int) []int {
	var a int
	for i := 0; i < len(kel); i++ {
		for j := i + 1; j < len(kel); j++ {
			if kel[i] < kel[j] {
				a = kel[i]
				kel[i] = kel[j]
				kel[j] = a
			}
		}
	}

	return kel
}

func main() {
	// Soal Algoritma Pertama 10 Point
	fmt.Print("Soal Algoritma Pertama \n")
	for i := 1; i <= 7; i++ {
		for j := i; j <= 7; j++ {
			if i == j || j == 1 {
				fmt.Print(j)
			} else {
				fmt.Print("0")
			}
		}
		fmt.Print("\n")
	}
	fmt.Print("=========================\n")

	//Soal Algoritma Kedua
	fmt.Print("\nSoal Algoritma Kedua \n")
	var bilangan = [32]int{
		1, 2, 33, 44, 55, 33, 44, 22, 44, 33, 2,
		77, 66, 1, 2, 3, 4, 5, 6, 7, 89, 3,
		3, 8, 9, 75, 4, 3, 2, 2, 1, 3}

	// fmt.Print(len(bilangan))
	var A []int = bilangan[0:11]
	var B []int = bilangan[11:22]
	var C []int = bilangan[22:32]

	//Sorting array
	print("\n>>>sorting array \n")
	kelA := sorting(A)
	fmt.Print("Kelompok A : ", kelA, "\n")
	kelB := sorting(B)
	fmt.Print("Kelompok B : ", kelB, "\n")
	kelC := sorting(C)
	fmt.Print("Kelompok C : ", kelC)

	//total array
	print("\n\n>>>total array")
	kelompokA := total(A)
	fmt.Print("\nTotal Kelompok A : ", kelompokA)
	kelompokB := total(B)
	fmt.Print("\nTotal Kelompok B : ", kelompokB)
	kelompokC := total(C)
	fmt.Print("\nTotal Kelompok C : ", kelompokC)

	print("\n\n>>>Rata-rata array")
	rata2A := rerata(A)
	fmt.Print("\nRata-rata Kelompok A : ", rata2A)
	rata2B := rerata(B)
	fmt.Print("\nRata-rata Kelompok B : ", rata2B)
	rata2C := rerata(C)
	fmt.Print("\nRata-rata Kelompok C : ", rata2C)

	print("\n\n>>>Max array \n")
	fmt.Print("Nilai Max Kelompok A : ", kelA[0:1], "\n")
	fmt.Print("Nilai Max Kelompok B : ", kelB[0:1], "\n")
	fmt.Print("Nilai Max Kelompok C : ", kelC[0:1], "\n")

	print("\n\n>>>Min array \n")
	fmt.Print("Nilai Min Kelompok A : ", kelA[10:11], "\n")
	fmt.Print("Nilai Min Kelompok B : ", kelB[10:11], "\n")
	fmt.Print("Nilai Min Kelompok C : ", kelC[9:10], "\n")
}
