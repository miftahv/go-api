package models

import (
	"go-api/db"
	"net/http"
)

type Entities struct {
	Id     int `json:"id"`
	Name   string `json:"name"`
	Image string `json:"image"`
	Id_options int `json:"id_options"`
	Kode string `json:"kode"`
	
}

func FetchAllEntities() (Response, error) {
	var obj Entities
	var arrobj []Entities
	var res Response

	con := db.CreateCon()
	sqlStatement := "SELECT * FROM entities"

	rows, err := con.Query(sqlStatement)
	if err !=nil {
		panic(err)
	}
	defer rows.Close()
	if err != nil{
		return res,err
	}

	for rows.Next(){
		err = rows.Scan(&obj.Id, &obj.Name,&obj.Image, &obj.Id_options, &obj.Kode)
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

func FetchEntitiesById(id int) (Response, error) {
	var obj Entities
	var arrobj []Entities
	var res Response

	con := db.CreateCon()
	sqlStatement := "SELECT * FROM entities where id=?"

	rows, err := con.Query(sqlStatement, id)
	if err !=nil {
		panic(err)
	}
	defer rows.Close()
	if err != nil{
		return res,err
	}

	for rows.Next(){
		err = rows.Scan(&obj.Id, &obj.Name, &obj.Image,&obj.Id_options,&obj.Kode)
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