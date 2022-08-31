
package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"fmt"
)
//lala
type Makan struct {
	gorm.Model
	Nama  string
	Harga float64
}
type Minuman struct {
	gorm.Model
	Nama  string
	Harga float64
	Topping  string
}
type Staff struct {
	gorm.Model
	Nama  string
	Umur int64
	Tinggal  string
}
func main() {
	// var mk Makan
	var mkA []Makan
	var mkM []Minuman
	var mkS []Staff
	// var mk2 Makan
	var mak Makan


	var nam string
	var har float64
	//minuman
	var namMi string
	var harMi float64
	var topMi string
	//staff
	var namS string
	var umur int64
	var tinggal string
	






	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "jehua:ur9ntmzzWF@tcp(db.kaenova.my.id:2003)/jehua?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	err = db.AutoMigrate(&Makan{}, &Minuman{}, &Staff{})
	if err != nil {
		panic(err.Error())
	}
	var in int = 1
	var masukkan int
	fmt.Println("Apakah anda ingin mengisi database ?")
	fmt.Println("1.Ya")
	fmt.Println("2.tidak")
	fmt.Println("3.tidak")
	fmt.Println("4.tidak")
	fmt.Println("5.tidak")
	fmt.Scanln(&in)
	switch in {
		case 1:
			for in == 1 {

				fmt.Println("masukkan apa yang diinginkan ?")
				fmt.Println("1.Makan")
				fmt.Println("2.Minuman")
				fmt.Println("3.Staff")
				fmt.Scanln(&masukkan)
				switch masukkan {
				case 1:
					fmt.Println("Nama :")
					fmt.Scanln(&nam)
					fmt.Println("Harga :")
					fmt.Scanln(&har)
					// var mak Makan
					mak := Makan{
						Nama:  nam,
						Harga: har,
					}
					result1 := db.Create(&mak)
		
					err1 := result1.Error 
		
					if err1 != nil {
						panic(err1.Error())
					}
				case 2:
					fmt.Println("Nama :")
					fmt.Scanln(&namMi)
					fmt.Println("Harga :")
					fmt.Scanln(&harMi)
					fmt.Println("Topping :")
					fmt.Scanln(&topMi)
					// var minu Minuman
					minu := Minuman{
						Nama:    namMi,
						Harga:   harMi,
						Topping: topMi,
					}
					result2 := db.Create(&minu)
		
					err2 := result2.Error 
		
					if err2 != nil {
						panic(err2.Error())
					}
		
				case 3:
					fmt.Println("Nama :")
					fmt.Scanln(&namS)
					fmt.Println("Umur :")
					fmt.Scanln(&umur)
					fmt.Println("Tempat TInggal :")
					fmt.Scanln(&tinggal)
					// var sta Staff
					sta := Staff{
						Nama:    namS,
						Umur:    umur,
						Tinggal: tinggal,
					}
					result3 := db.Create(&sta)
		
					err3 := result3.Error 
		
					if err3 != nil {
						panic(err3.Error())
					}
				default:
					fmt.Println("masukkan tidak valid")
				}
				
		
		
				fmt.Println("Apakah anda ingin mengisi database lagi ?")
				fmt.Println("1.Ya")
				fmt.Println("2.tidak")
				fmt.Scanln(&in)
		
				
			}
		case 2:
			fmt.Println("selamat Tinggal")
		case 3:
			var lo int
			fmt.Println("Apa yang ingin dilihat ")
			fmt.Println("1.Makanan")
			fmt.Println("2.minuman")
			fmt.Println("3.staff")
			switch lo {
				case 1:
					result5 := db.Find(&mkA)
					err5 := result5.Error 
					if err5 != nil {
						panic(err5.Error())
					}
					fmt.Println(mkA)
				case 2:
					result6 := db.Find(&mkM)
					err6 := result6.Error 
					if err6 != nil {
						panic(err6.Error())
					}
					fmt.Println(mkM)
				case 3:
					result7 := db.Find(&mkS)
					err7 := result7.Error 
					if err7 != nil {
						panic(err7.Error())
					}
					fmt.Println(mkS)

			}
			
		case 4:
			var ki uint
			ki = 5
			fmt.Println("Masukkan ID")
			fmt.Scan(&ki)
			
			mak.ID = ki
			db.Delete(&mak)
		case 5:
		
			var ha string
			var nambaru string
			var harbaru float64
			// lomba := db.First(&mak, ku)
			fmt.Println("masukkan nama makanan")
			fmt.Scan(&ha)
			fmt.Println("masukkan nama makanan baru ")
			fmt.Scan(&nambaru)
			fmt.Println("masukkan harga makanan baru")
			fmt.Scan(&harbaru)

			db.Model(Makan{}).Where("Nama", "kwali").Updates(Makan{Nama: nambaru, Harga: harbaru})
			// db.Model(&lomba).Select("Nama", "harga").Updates(Makan{Nama: ha, Harga: rp})
			// err8 := db.Save(&col).Error 
			// if err8 != nil {
			// 	panic(err8.Error())
			// 	}
			// fmt.Println(&col)



	}
	



	// makan := Makan{Nama: "Soto", Harga: 10000}
	// maka := Makan{Nama: "bakso", Harga: 15000}
	// minum := Minuman{Nama: "Boba", Harga: 15000, Topping: "mangga"}
	// staff := Staff{Nama: "dewa", Umur: 21, Tinggal: "solo"}
	// result5 := db.Create(&makan)
	// result9 := db.Create(&maka)
	// result2 := db.Create(&minum)
	// result3 := db.Create(&staff)

	// err5 := result5.Error 

	// if err5 != nil {
	// 	panic(err5.Error())
	// }


	// err9 := result9.Error 

	// if err9 != nil {
	// 	panic(err9.Error())
	// }





	// err2 := result2.Error 

	// if err2 != nil {
	// 	panic(err2.Error())
	// }
	// err3 := result3.Error 

	// if err3 != nil {
	// 	panic(err3.Error())
	// }
	
	// result4 := db.First(&mk)
	// err4 := result4.Error
	// if err4 != nil {
	// 	panic(err4.Error())
	// }
	// fmt.Println(mk)


	// 	// Get all records
	


	
	// i:= mk.ID
	// result8 := db.First(&mk2, i)

	// err8 := result8.Error 
	// if err5 != nil {
	// 	panic(err8.Error())
	// }
	// fmt.Println(mk2)
	
	// mak.ID = 5
	// db.Delete(&mak)
}
