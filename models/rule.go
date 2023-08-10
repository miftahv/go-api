package models

import (
	"go-api/db"
	"net/http"
)

type Rule struct {
	Id     int `json:"id"`
	Kode   string `json:"kode"`
	Rule string `json:"rule"`
	
}

func FetchAllRule() (Response, error) {
	var obj Rule
	var arrobj []Rule
	var res Response

	con := db.CreateCon()
	sqlStatement := "SELECT * FROM rule"

	rows, err := con.Query(sqlStatement)
	if err !=nil {
		panic(err)
	}
	defer rows.Close()
	if err != nil{
		return res,err
	}

	for rows.Next(){
		err = rows.Scan(&obj.Id, &obj.Kode,&obj.Rule)
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

func FetchRuleById(id int) (Response, error) {
	var obj Rule
	var arrobj []Rule
	var res Response

	con := db.CreateCon()
	sqlStatement := "SELECT * FROM rule where id=?"

	rows, err := con.Query(sqlStatement, id)
	if err !=nil {
		panic(err)
	}
	defer rows.Close()
	if err != nil{
		return res,err
	}

	for rows.Next(){
		err = rows.Scan(&obj.Id, &obj.Kode, &obj.Rule)
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