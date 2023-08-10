package models

import (
	"go-api/db"
	"net/http"
)

type Hotkeys struct {
	Id     int `json:"id"`
	Name   string `json:"name"`
	Id_kategori int `json:"id_kategori"`
}

func FetchAllHotkeys() (Response, error) {
	var obj Hotkeys
	var arrobj []Hotkeys
	var res Response

	con := db.CreateCon()
	sqlStatement := "SELECT * FROM hotkeys"

	rows, err := con.Query(sqlStatement)
	if err !=nil {
		panic(err)
	}
	defer rows.Close()
	if err != nil{
		return res,err
	}

	for rows.Next(){
		err = rows.Scan(&obj.Id, &obj.Name, &obj.Id_kategori)
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
func FetchHotkeysById(id int) (Response, error) {
	var obj Hotkeys
	var arrobj []Hotkeys
	var res Response

	con := db.CreateCon()
	sqlStatement := "SELECT * FROM hotkeys where id=?"

	rows, err := con.Query(sqlStatement, id)
	if err !=nil {
		panic(err)
	}
	defer rows.Close()
	if err != nil{
		return res,err
	}

	for rows.Next(){
		err = rows.Scan(&obj.Id, &obj.Name, &obj.Id_kategori)
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