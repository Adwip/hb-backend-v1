package product

import "database/sql"
import _"fmt"
import "hb-backend-v1/model"

type Hasil struct{
	Id_alamat_store int
	Id_tk 			int
	Alamat 			string
	Id_wil 			int
	Latitude 		sql.NullString
	Longitude 		sql.NullString
	
}

func ProductList() ([]Hasil, error){
	var result []Hasil
	err := model.Query("SELECT * FROM alamat_store",func(rows *sql.Rows)error {
		for rows.Next(){
			var each = Hasil{}
			var err = rows.Scan(&each.Id_alamat_store, &each.Id_tk, &each.Alamat, &each.Id_wil, &each.Latitude, &each.Longitude)
			if err != nil{
				return err
			}
			result = append(result, each)
		}
		return nil
	})
	if err != nil{
		return nil, err
	}
	return result, nil
}