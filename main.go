package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// **** Ham swap*****************************************
func swap(a *Danhsachphong, b *Danhsachphong) {
	var x Danhsachphong

	x = *a
	*a = *b
	*b = x
}

// **** Kieu du lieu phong******************************
type phong struct {
	ID        int `json:"id"`
	Kieuphong int `json:"kieuphong"`
	Price     int `json:"price"`
}

func readFilePhong(fileName string) ([]phong, error) {
	data, err := os.ReadFile(fileName)
	phongData := []phong{}
	if err != nil {

	}
	json.Unmarshal([]byte(data), &phongData)
	return phongData, nil
}

// *****Kieu du lieu khach***********************************
type khach struct {
	Ten         string `json:"Ten khach hang"`
	Email       string `json:"Email"`
	Sdt         string `json:"So dien thoai"`
	Makhachhang string `json:"Ma khach hang"`
	Describe    int    `json:"Describe"`
}

func readFileKhach(fileName string) ([]khach, error) {
	data, err := os.ReadFile(fileName)
	khachData := []khach{}
	if err != nil {

	}
	json.Unmarshal([]byte(data), &khachData)

	return khachData, nil
}

// ****Kieu du lieu danh sach phong********************************
type Danhsachphong struct {
	Ten       string `json:"Ten khach hang"`
	ID        int    `json:"ID phong"`
	Kieuphong int    `json:"Kieu phong"`
	Price     int    `json:"Price"`
}

func readFilePhongthue(fileName string) ([]Danhsachphong, error) {
	data, err := os.ReadFile(fileName)
	danhsachphongData := []Danhsachphong{}
	if err != nil {

	}
	json.Unmarshal([]byte(data), &danhsachphongData)
	return danhsachphongData, nil
}

// *****Ham main****************************************************
func main() {
	fmt.Println("Features:")
	fmt.Println("1. Enter new zoom")
	fmt.Println("2. Enter new customer")
	fmt.Println("3. Book")
	fmt.Println("4. Sort room booking by room type")
	fmt.Println("5. Creat bill for customer")
	fmt.Print("Your choice: ")
	var choice int
	fmt.Scanln(&choice)

	phong0 := []phong{}
	phong0, _ = readFilePhong("Phong.json")

	khach0 := []khach{}
	khach0, _ = readFileKhach("KH.json")
	// //*******************1. Nhap them phong********************************************************************
	if choice == 1 {
		phongnew := phong{}
		fmt.Print("ID: ")
		fmt.Scanln(&phongnew.ID)
		fmt.Print("Type: 1. Single 2. Double 3. VIP?")
		fmt.Scanln(&phongnew.Kieuphong)
		fmt.Print("Price: ")
		fmt.Scanln(&phongnew.Price)
		fmt.Println("Successful!")

		//fmt.Printf("%5d%10d%10d", phong0.ID, phong0.Kieuphong, phong0.Price)
		phong0 = append(phong0, phongnew)
		file, err := json.MarshalIndent(phong0, "", " ")

		if err != nil {
			fmt.Printf("%v", err)
		}

		_ = ioutil.WriteFile("Phong.json", file, 0644)

	}
	// //***************************************************************************************************************

	// // ***************2. Nhap them khach******************************************************************
	if choice == 2 {
		khachnew := khach{}
		fmt.Print("Ten khach hang: ")
		fmt.Scanln(&khachnew.Ten)
		fmt.Print("Email: ")
		fmt.Scanln(&khachnew.Email)
		fmt.Print("So dien thoai: ")
		fmt.Scanln(&khachnew.Sdt)
		fmt.Print("Ma khach hang: ")
		fmt.Scanln(&khachnew.Makhachhang)
		fmt.Print("Mo ta: ")
		fmt.Scanln(&khachnew.Describe)

		khach0 = append(khach0, khachnew)

		file, err := json.MarshalIndent(khach0, "", " ")

		if err != nil {
			fmt.Printf("%v", err)
		}

		_ = ioutil.WriteFile("KH.json", file, 0644)

	}
	// //******************************************************************************************************

	//********3. Danh sach phong thue cho tung khach**************************************************************

	if choice == 3 {
		khach0 := []khach{}
		khach0, _ = readFileKhach("KH.json")

		phong0 := []phong{}
		phong0, _ = readFilePhong("Phong.json")

		check := map[phong]bool{}

		danhsachphong := []Danhsachphong{}

		for i := 0; i < len(khach0); i++ {
			for j := 0; j < len(phong0); j++ {
				if khach0[i].Describe == phong0[j].Kieuphong {
					check[phong0[j]] = false
					m := Danhsachphong{Ten: khach0[i].Ten,
						ID:        phong0[j].ID,
						Kieuphong: phong0[j].Kieuphong,
						Price:     phong0[j].Price}
					danhsachphong = append(danhsachphong, m)
				}

			}
		}
		file, err := json.MarshalIndent(danhsachphong, "", " ")

		if err != nil {
			fmt.Printf("%v", err)
		}

		_ = ioutil.WriteFile("Phongthue.json", file, 0644)

	}

	//***************************************************************************************************************

	//*****4. Sap xep danh sach da luu theo kieu phong***********************************************************************************
	if choice == 4 {
		danhsachphong := []Danhsachphong{}
		danhsachphong, _ = readFilePhongthue("Phongthue.json")
		for i := 0; i < len(danhsachphong)-1; i++ {
			for j := i + 1; j < len(danhsachphong); j++ {
				if danhsachphong[i].Kieuphong > danhsachphong[j].Kieuphong {
					swap(&danhsachphong[i], &danhsachphong[j])
				}
			}
		}
		file, err := json.MarshalIndent(danhsachphong, "", " ")

		if err != nil {
			fmt.Printf("%v", err)
		}

		_ = ioutil.WriteFile("Phongthue.json", file, 0644)

	}
	//******************************************************************************************************************
}
