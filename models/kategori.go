package models

import (
	"go-api/db"
	"net/http"
)
type Kategori struct {
	Id     int `json:"id"`
	Name   string `json:"name"`
	Id_family int `json:"id_family"`
}

func FetchAllKategori() (Response, error) {
	var obj Kategori
	var arrobj []Kategori
	var res Response

	con := db.CreateCon()
	sqlStatement := "SELECT * FROM kategori"

	rows, err := con.Query(sqlStatement)
	if err !=nil {
		panic(err)
	}
	defer rows.Close()
	if err != nil{
		return res,err
	}

	for rows.Next(){
		err = rows.Scan(&obj.Id, &obj.Name, &obj.Id_family)
		if err != nil {
			return res,err
		}

		arrobj = append(arrobj, obj)
	}
	res.Status = http.StatusOK
	res.Message = "Succes"
	res.Data = arrobj

	return res,nil

}

func FetchKategoriById(id int) (Response, error) {
	var obj Kategori
	var arrobj []Kategori
	var res Response

	con := db.CreateCon()
	sqlStatement := "SELECT * FROM kategori where id=?"

	rows, err := con.Query(sqlStatement, id)
	if err !=nil {
		panic(err)
	}
	defer rows.Close()
	if err != nil{
		return res,err
	}

	for rows.Next(){
		err = rows.Scan(&obj.Id, &obj.Name, &obj.Id_family)
		if err != nil {
			return res,err
		}

		arrobj = append(arrobj, obj)
	}
	res.Status = http.StatusOK
	res.Message = "Succes"
	res.Data = arrobj

	return res,nil

}