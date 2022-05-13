package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func main() {

	rand.Seed(time.Now().UnixNano())

	// Number 1
	printSegitiga(5)

	// Number 2
	result := genPass("Abcdef", "mids")
	fmt.Println(result)

	// Number 3
	result2 := durasifilm(7)
	fmt.Println(result2)
}

func printSegitiga(number int)  {

	if number <= 0 {
		fmt.Println("Terjadi kesalahan")
	} else {
		for brs:= number; brs >= 1; brs--{

			for spasi:= number+1 ; spasi>=brs; spasi--{
				fmt.Printf(" ");
			}
	
			for klm:=1; klm<=brs; klm++{
				fmt.Printf("*");
			}
	
			for klm:=(brs-1); klm>=1; klm--{
				fmt.Printf("*");
			}
			
			fmt.Printf("\n");
		}
	}
}

func genPass(password string, level string) string {
	response := ""
	passwords := strings.ToLower(password)
	if len(passwords) < 6 {
		response += "Panjang password harus lebih dari 6 huruf"
	} else if level == "low" {
		gen1 := strings.ToUpper(passwords[0:1])
		gen2 := strings.ToUpper(passwords[len(passwords)-1 : len(passwords)+0])
		response += gen1 + passwords + gen2
	} else if level == "mid" {
		randAngka := rand.Intn(100)
		chars := "abcdefghijklmnopqrstuvwxyz"
		gen1 := strings.ToUpper(passwords[0:1])
		gen2 := strings.ToUpper(passwords[len(passwords)-1 : len(passwords)+0])
		gen3 := string(chars[rand.Intn(len(chars))])
		response += gen3 + gen1 + passwords + gen2 + strconv.Itoa(randAngka)
	} else if level == "strong" {
		randAngka := rand.Intn(100)
		simbol := "~!@#$%^&*()_+=-{}[];:<>/`?|.,"
		gen1 := strings.ToUpper(passwords[0:1])
		gen2 := strings.ToUpper(passwords[len(passwords)-1 : len(passwords)+0])
		gen3 := string(simbol[rand.Intn(len(simbol))])
		response += gen3 + gen1 + passwords + gen2 + strconv.Itoa(randAngka) + gen3
	} else {
		response += "Terjadi kesalahan"
	}
	
	return response
}

func durasifilm(lamaTerbang int) string {
	
	arrFilm := []int{1, 5, 2, 7, 3, 4, 8, 9}
	temp := 0
	response := ""
	for i := 0; i < len(arrFilm) ; i++ {
		if temp < len(arrFilm)-1 {
			temp += 1
		} else {
			temp += 0
		}
		
		if calculate := arrFilm[i] + arrFilm[temp]; calculate == lamaTerbang{
			response += strconv.Itoa(arrFilm[i]) + " dan " + strconv.Itoa(arrFilm[temp]) + "\n"
		} else {
			response += "Terjadi kesalahan"
		}
	}

	return response
}
