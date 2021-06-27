package response

import (
	"net/http"

	"../model"
)

type CurrencyResponse struct {

	Message string `json:"message"`
	Status string  `json:"status"`
	StatusCode int `json:"statusCode"`
	CurrencyDetails  []model.CurrencyDetail `json:"currencies"`
	CurrencyDetail  model.CurrencyDetail 

}

func CurrencyDetailsResponse(currencydetails []model.CurrencyDetail) (resposne CurrencyResponse){
	var response CurrencyResponse
	 if !(len(currencydetails)==0){ 
     response =CurrencyResponse{
		 Message :"Data Fetched Successfully ",
		 Status :"OK",
		 StatusCode:http.StatusOK,
		 CurrencyDetails :currencydetails,
	 }
	}else {
		response =CurrencyResponse{
		    Message :"No Data Found",
			Status :"Bad Request",
			StatusCode:http.StatusBadRequest,
			CurrencyDetails :currencydetails,
	}
}
	 return response
	
}

func SingleCurrencyResponse(currencydetail model.CurrencyDetail) (resposne CurrencyResponse){
	var response CurrencyResponse
	 if !(currencydetail==model.CurrencyDetail{}){ 
     response =CurrencyResponse{
		 Message :"Data Fetched Successfully ",
		 Status :"OK",
		 StatusCode:http.StatusOK,
		 CurrencyDetail :currencydetail,
	 }
	}else {
		response =CurrencyResponse{
		    Message :"No Data Found",
			Status :"Bad Request",
			StatusCode:http.StatusBadRequest,
			CurrencyDetail :currencydetail,
	}
}
	 return response
	
}