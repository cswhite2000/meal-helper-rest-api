package main

import (
	"encoding/json"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"io/ioutil"
	"strings"
	"strconv"
)

type Food struct {
	Name string
	Description string
	Calories float64
	Fat float64
	Carbohydrates float64
	Protein float64
}

var food []Food

func GetNutrition(w http.ResponseWriter, r *http.Request) {
	addFood()
	json.NewEncoder(w).Encode(food)
}

func addFood() {

	food = make([]Food, 0)

	files, err := ioutil.ReadDir("./food/")
	if err != nil {
		log.Println(err)
		return
	}

	for _, foodFile := range files {

		data, err := ioutil.ReadFile("./food/" + foodFile.Name())
		if err != nil {
			log.Println(err)
			continue
		}

		dataString := strings.Replace(string(data), "\r", "", 1000);
		dataStrings := strings.Split(dataString, "\n")

		if len(dataStrings) > 5 {
			calories, err := strconv.ParseFloat(dataStrings[1], 64);
			if err != nil {
				log.Println(err)
				continue
			}

			fat, err := strconv.ParseFloat(dataStrings[2], 64);
			if err != nil {
				log.Println(err)
				continue
			}

			carbs, err := strconv.ParseFloat(dataStrings[3], 64);
			if err != nil {
				log.Println(err)
				continue
			}

			protein, err := strconv.ParseFloat(dataStrings[4], 64);
			if err != nil {
				log.Println(err)
				continue
			}

			food = append(food, Food{
				Name:          foodFile.Name(),
				Description:   dataStrings[0],
				Calories:      calories,
				Fat:           fat,
				Carbohydrates: carbs,
				Protein:       protein,
			})

		}
	}
}

func main() {

	addFood()

	router := mux.NewRouter()

	router.HandleFunc("/nutrition/", GetNutrition).Methods("GET")

	log.Fatal(http.ListenAndServe(":80", router))
}