package main

import (
	"encoding/csv"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/sadaakisz/go-api/model"
)

var data = [][]string{}

func readCsv(filePath string) [][]string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file: "+filePath, err)
	}
	defer f.Close()
	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for: "+filePath, err)
	}
	return records
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	listaCiudadano := []model.Ciudadano{}
	for _, row := range data {
		x, _ := strconv.Atoi(row[0])
		y, _ := strconv.Atoi(row[1])
		object := model.Ciudadano{
			ESTRATO_SOCIOECONOMICO: x,
			SEGURIDAD_NOCTURNA:     y,
		}
		listaCiudadano = append(listaCiudadano, object)
	}
	listaCiudadano = listaCiudadano[1:]
	//vars := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(listaCiudadano)
}

func main() {
	r := mux.NewRouter()
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
	})
	handler := c.Handler(r)

	data = readCsv("./DatasetSelectivo.csv")
	//fmt.Println(data)

	r.HandleFunc("/", HomeHandler)
	log.Fatal(http.ListenAndServe(":3000", handler))
}
