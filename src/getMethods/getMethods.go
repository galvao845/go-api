package getMethods

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

//ID SERIAL PRIMARY KEY NOT NULL,
//DATA_ADVICE VARCHAR(500) NOT NULL
type DataAdviceStruct struct {
	id         int
	DataAdvice string
}

type bodyGetId struct {
	Id string `json:"id"`
}

type bodyGetAdviceData struct {
	Id         string `json:"id"`
	DataAdvice string `json:"dataAdvice"`
}

func throwError(err error) {
	if err != nil {
		panic(err.Error())
	}
}

// API FUNCITIONS
func GetAdvice(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get(api_url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintln(w, string(body))
}

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

	resp, err := http.Get(api_url + "/" + string(ParsedBody.Id))
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
func GetAdviceFromDb(w http.ResponseWriter, r *http.Request) {
	var conn = dataBaseInstance.ConnectDb()
	defer conn.Close()

	sqlStatment, err := conn.Query("select * from ADVICE_TABLE")
	throwError(err)

	for sqlStatment.Next() {
		var dataAdviceStruct DataAdviceStruct
		err := sqlStatment.Scan(&dataAdviceStruct.id, &dataAdviceStruct.DataAdvice)
		throwError(err)
		fmt.Fprintln(w, "{"+fmt.Sprint(dataAdviceStruct.id)+":"+dataAdviceStruct.DataAdvice+"}")
	}
}
