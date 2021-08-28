package product

import "database/sql"
import _"fmt"
import "hb-backend-v1/config/database"

type Hasil struct{
	Id_alamat_store int
	Id_tk 			int
	Alamat 			string
	Id_wil 			int
	Latitude 		sql.NullString
	Longitude 		sql.NullString
	
}

func DaftarProduct() ([]Hasil, error){
	connect, err := database.Connect()
	rows, err := connect.Query("SELECT * FROM alamat_store")
	if err != nil{
		return nil, err
	}
	defer rows.Close()
	
	var result []Hasil
	for rows.Next(){
		var each = Hasil{}
		var err = rows.Scan(&each.Id_alamat_store, &each.Id_tk, &each.Alamat, &each.Id_wil, &each.Latitude, &each.Longitude)
		if err != nil{
			return nil, err
		}
		result = append(result, each)
	}
	return result, nil
}