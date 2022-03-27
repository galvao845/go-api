package postMethods

import (
	"context"
	"dataBaseInstance"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const api_url = "https://api.adviceslip.com/advice"

var dbContext = context.Background()

type bodyGetId struct {
	Id string `json:"id"`
}

type bodyGetAdviceData struct {
	Id         string `json:"id"`
	DataAdvice string `json:"dataAdvice"`
}

func throwError(w http.ResponseWriter, err error) {
	if err != nil {
		fmt.Fprintln(w, err.Error())
		//panic(err.Error())
	} else {
		fmt.Fprintln(w, true)
		//panic(err.Error())
	}
}

// API FUNCITIONS
func GetAdviceById(w http.ResponseWriter, r *http.Request) {
	Reqbody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	var ParsedBody bodyGetId
	err = json.Unmarshal(Reqbody, &ParsedBody)
	if err != nil {
		log.Fatal(err)
	}

	if ParsedBody.Id == "" {
		fmt.Fprintln(w, "bodyError")
		return
	}

	resp, err := http.Get(api_url + ParsedBody.Id)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body2, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintln(w, string(body2))
}

// DP FUNCTIONS
func DeleteAdviceDb(w http.ResponseWriter, r *http.Request) {
	Reqbody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	var ParsedBody bodyGetId
	err = json.Unmarshal(Reqbody, &ParsedBody)
	if err != nil {
		log.Fatal(err)
	}

	if ParsedBody.Id == "" {
		fmt.Fprintln(w, "bodyError")
		return
	}

	var conn = dataBaseInstance.ConnectDb()
	defer conn.Close()

	del, err := conn.Query("DELETE FROM ADVICE_TABLE WHERE ID = " + ParsedBody.Id)
	throwError(w, err)
	fmt.Println(del)
}

func InsertAdviceDb(w http.ResponseWriter, r *http.Request) {
	Reqbody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	var ParsedBody bodyGetAdviceData

	err = json.Unmarshal(Reqbody, &ParsedBody)
	if err != nil {
		log.Fatal(err)
	}

	if ParsedBody.Id == "" || ParsedBody.DataAdvice == "" {
		fmt.Fprintln(w, "bodyError")
		return
	}

	var conn = dataBaseInstance.ConnectDb()
	defer conn.Close()

	insert, err := conn.Query("INSERT INTO ADVICE_TABLE (ID, DATA_ADVICE) VALUES (" + ParsedBody.Id + ", '" + ParsedBody.DataAdvice + "')")
	throwError(w, err)
	fmt.Println(insert)
}
