package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func ParseBody(r *http.Request, x interface{}){
	body, err := ioutil.ReadAll(r.Body)
	if  err != nil{
		fmt.Println(err)
	}
	if err := json.Unmarshal([]byte(body), x); err != nil{ // unmarshal json ni go tiliga o'giradi
		return
	}
}