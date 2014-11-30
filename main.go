package main

import (
	"net/http"
	"fmt"
	"encoding/json"
	"strings"
)

// Define a weather data struct
type weatherData struct {
	Name string `json:"name"`
	Main struct {
		Kelvin float64 `json:"temp"`
	} `json:"main"`
}

// Populate query
func query(city string) (weatherData, error) {
	resp, err := http.Get("http://api.openweathermap.org/data/2.5/weather?q=" + city)
	if err != nil {
		return weatherData{}, err
	}

	defer resp.Body.Close()

	var d weatherData

	if err:= json.NewDecoder(resp.Body).Decode(&d); err != nil {
		return weatherData{}, err
	}

	return d, nil
}


func main() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/weather/", func(w http.ResponseWriter, r *http.Request) {
			city:= strings.SplitN(r.URL.Path, "/", 3)[2]
			fmt.Println(r.URL.Path)
			fmt.Println(strings.SplitN(r.URL.Path, "e", 10)[0])
			fmt.Println(strings.SplitN(r.URL.Path, "e", 10)[1])
			fmt.Println(strings.SplitN(r.URL.Path, "e", 10)[2])

			data, err := query(city)

			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			json.NewEncoder(w).Encode(data)
		})
	http.ListenAndServe(":8080", nil)
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello!"))
}
