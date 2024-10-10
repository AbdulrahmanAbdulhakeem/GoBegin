package main

import (
  "fmt"
  "net/http"
  "log"
  "math/rand"
)

const creditScoreMin = 500
const creditScoreMax = 1000

type credit_rating struct{
	CreditRating int `json:"credit_rating"` 
}
