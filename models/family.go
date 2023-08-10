package models

import (
	"go-api/db"
	"net/http"
)

type Family struct {
	Id     int `json:"id"`
	Name   string `json:"name"`
	Id_sec int `json:"id_sec"`
	Ischecked bool `json:"ischecked"`
}

func FetchAllFamily() (Response, error) {
	var obj Family
	var arrobj []Family
	var res Response

	con := db.CreateCon()
	sqlStatement := "SELECT * FROM families"

	rows, err := con.Query(sqlStatement)
	if err !=nil {
		panic(err)
	}
	defer rows.Close()
	if err != nil{
		return res,err
	}

	for rows.Next(){
		err = rows.Scan(&obj.Id, &obj.Name, &obj.Id_sec, &obj.Ischecked)
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