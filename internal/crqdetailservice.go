package internal

import (
	"fmt"
  "net/http"
  "encoding/json"
  "io/ioutil"

	"../model"
	"../response"
)



func AllCurrencyDetail() (resposne response.CurrencyResponse){
  var currencyDetails []model.CurrencyDetail
  client := &http.Client{}
 req, err := http.NewRequest("GET", "https://api.hitbtc.com/api/2/public/currency", nil)
 if err != nil {
  fmt.Print(err.Error())
 }
 req.Header.Add("Accept", "application/json")
 req.Header.Add("Content-Type", "application/json")
 resp, err := client.Do(req)
 if err != nil {
  fmt.Print(err.Error())
 }
 bodyBytes, err := ioutil.ReadAll(resp.Body)
 if err != nil {
  fmt.Print(err.Error())
 }

 json.Unmarshal(bodyBytes, &currencyDetails)

 return response.CRQDetailsResponse(currencyDetails)
  
}

func CurrencyDetail(currencyID string) (resposne response.CurrencyResponse){
  var currencyDetails []model.CurrencyDetail
  client := &http.Client{}
 req, err := http.NewRequest("GET", "https://api.hitbtc.com/api/2/public/currency/currencyID", nil)
 if err != nil {
  fmt.Print(err.Error())
 }
 req.Header.Add("Accept", "application/json")
 req.Header.Add("Content-Type", "application/json")
 resp, err := client.Do(req)
 if err != nil {
  fmt.Print(err.Error())
 }
 bodyBytes, err := ioutil.ReadAll(resp.Body)
 if err != nil {
  fmt.Print(err.Error())
 }

 json.Unmarshal(bodyBytes, &currencyDetails)

 return response.CRQDetailsResponse(currencyDetails)
  
}


