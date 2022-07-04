package controllers

import (
	"encoding/json"
	"fmt"
	//"math"
	"net/http"
	"strconv"

	"credit-score-system/models"
)

var Criteria = models.Criteria{
	YearsOfOperation : 10,
	LengthOfBusiness: 40,
	AverageMonthlyBusiness: 30,
	HistoryWithGuava: 20,
}

func Score(res http.ResponseWriter , req *http.Request){
	vars := req.URL.Query()
	fmt.Println(vars)
	years := vars.Get("years")
	length := vars.Get("length")
	averageRevenue := vars.Get("averageRevenue") 
	defaulted := vars.Get("defaulted") 
	hasHistory := vars.Get("hasHistory")


	_years,e := strconv.ParseFloat(years,32) 
	_length,e := strconv.ParseFloat(length,32) 
	_averageRevenue,e := strconv.ParseFloat(averageRevenue,32) 
	_defaulted,e := strconv.ParseBool(defaulted) 
	_hasHistory,e := strconv.ParseBool(hasHistory) 

	if(e != nil){
		fmt.Println(e)
	}

	var _score = creditRate(float32(_years), float32(_length), float32(_averageRevenue), _hasHistory , _defaulted);
	json.NewEncoder(res).Encode(models.Response{
		Success : true, Score : _score, Message : "Credit score calculated",
	})
}


func creditRate(y float32,l float32,a float32,h bool,defaulted bool) float32{
	return yearsOfOperation(y) + lengthOfBusiness(l) + avMonthlyBusiness(a) + historyWithGuava(h, defaulted);
}

func yearsOfOperation(years float32) float32{
	var ret float32 = 0;
    
    if(years < 1){
        return Criteria.YearsOfOperation * 0.2;
	}else if (years >= 1 && years <= 3){
        return Criteria.YearsOfOperation * 0.5;
	}else if (years > 3 && years <= 5){
        return Criteria.YearsOfOperation * 0.8;
	}else if (years > 5){
        return Criteria.YearsOfOperation * 1;
	}else{
        return 0;
	}
    return ret;
}

func lengthOfBusiness(months float32) float32{

	if (months < 6){
        return Criteria.LengthOfBusiness * 0.2
	}else if (months >= 6 && months < 12){
        return Criteria.LengthOfBusiness * 0.4;
	}else if (months >= 12 && months <= 36){
        return Criteria.LengthOfBusiness * 0.6;
	}else if (months >= 36 && months <= 60){
        return Criteria.LengthOfBusiness * 0.7;
    }else if (months > 60){
        return Criteria.LengthOfBusiness * 1;
	}
	return Criteria.LengthOfBusiness * 0.2
}

func avMonthlyBusiness(amount float32) float32{
	if (amount < 200000){
        return Criteria.AverageMonthlyBusiness * 0.4;
	}else if (amount >= 200000 && amount <= 500000){
        return Criteria.AverageMonthlyBusiness * 0.7;
	}else if (amount > 500000){
        return Criteria.AverageMonthlyBusiness * 1;
	}else{
        return Criteria.AverageMonthlyBusiness * 0.9;
	}
}

func historyWithGuava(hasHist bool, defaulted bool) float32{
	var _hasHist float32
	if(defaulted){
		_hasHist = -0.2
	}else{
		_hasHist = 1
	}
	if(hasHist){
		return Criteria.HistoryWithGuava * 1 * _hasHist
	}else if (hasHist == false){
        return Criteria.HistoryWithGuava * 0 * _hasHist;
	}else{
        return Criteria.HistoryWithGuava * 0 * _hasHist;
	}
}