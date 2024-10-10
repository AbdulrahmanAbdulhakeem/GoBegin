package main

import (
  "fmt"
  //"net/http"
  //"log"
  "math/rand"
)

const creditScoreMin = 500
const creditScoreMax = 1000

type credit_rating struct{
	CreditRating int `json:"credit_rating"` 
}

func getCreditScore(w http.ResponseWriter, r *http.Request){
  var creditRating = credit_rating{
    CreditRating:(rand.Intn(creditScoreMax - creditScoreMin) + creditScoreMin)
  }

  w.WriteHeader(http.StatusOK)
  json.NewEncoder(w).Encode(creditRating)
}

func handleRequests(){
  http.Handle("/creditscore" , http.HandlerFunc(getCreditScore))
  log.Fatal(http.ListenAndServe(":8080" , nil))
}

func main(){
  handleRequests()
}
