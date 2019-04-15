package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/buger/jsonparser"
	"github.com/encima/openape"
	u "github.com/encima/openape/utils"
)

func main() {
	o := openape.NewServer("config")
	// CUSTOM ROUTES

	// j := u.ParseObject()
	o.AddCustomRoute("/users/login", "POST", func(w http.ResponseWriter, r *http.Request) {
		// vars := mux.Vars(r)
		var res u.JSONResponse
		// TODO validate model here?
		body, err := ioutil.ReadAll(r.Body)

		//TODO handle email login
		bodyUsername, ut, _, err := jsonparser.Get(body, "email")
		bodyPassword, pt, _, err := jsonparser.Get(body, "pwd")
		if ut != jsonparser.NotExist && pt != jsonparser.NotExist || err == nil {
			uname:= string(bodyUsername[:])
			pwd := string(bodyPassword[:])
			getJson := fmt.Sprintf(`{
   "select":{
      "query":[
         {
            "table":"users",
            "fields":[
               "email",
               "id"
            ]
         }
      ]
   },
   "where":[
      {
         "field":"email",
         "table":"users",
         "op":"e",
         "val":"%s",
		 "join": "a"
      },
      {
         "field":"pwd",
         "table":"users",
         "op":"e",
         "val":"%s"
      }
   ]
}`, uname, pwd)
			var jQuery u.JTOS
			err := json.Unmarshal([]byte(getJson), &jQuery)
			if err != nil {
				panic(err)
			}
			query := u.ParseObject(jQuery)
			users := o.DB.GetModels("users", query)
			if string(users.Data) == "null" {
				users.Status = 404
				res = users
				u.SendResponse(w, res)
			} else {
				_, err = jsonparser.ArrayEach(users.Data, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
					jsonparser.EachKey(value, func(idx int, value []byte, vt jsonparser.ValueType, err error) {
						fmt.Println(string(value))
					})
				})
				if err != nil {
					panic(err)
				}

				res = users
				u.SendResponse(w, res)
			}
		} else {
			res.Status = 400
			//res.Data =
			u.SendResponse(w, res)
		}

	})

	o.RunServer()
}
