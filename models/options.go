package models

import (
	"go-api/db"
	"net/http"
)

type Option struct {
	Id     int `json:"id"`
	Name   string `json:"name"`
	Image string `json:"image"`
	Id_hotkeys int `json:"id_hotkeys"`
	Value string `json:"value"`
}

func FetchAllOptions() (Response, error) {
	var obj Option
	var arrobj []Option
	var res Response

	con := db.CreateCon()
	sqlStatement := "SELECT * FROM options"

	rows, err := con.Query(sqlStatement)
	if err !=nil {
		panic(err)
	}
	defer rows.Close()
	if err != nil{
		return res,err
	}

	for rows.Next(){
		err = rows.Scan(&obj.Id, &obj.Name, &obj.Image,&obj.Id_hotkeys, &obj.Value)
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
func FetchOptionById(id int) (Response, error) {
	var obj Option
	var arrobj []Option
	var res Response

	con := db.CreateCon()
	sqlStatement := "SELECT * FROM options where id = ?"

	rows, err := con.Query(sqlStatement, id)
	if err !=nil {
		panic(err)
	}
	defer rows.Close()
	if err != nil{
		return res,err
	}

	for rows.Next(){
		err = rows.Scan(&obj.Id, &obj.Name, &obj.Image,&obj.Id_hotkeys, &obj.Value)
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