package main

import (
  "encoding/json"
  "net/http"
  "strings"
)

func main() {
  http.HandleFunc("/hello", hello)
  http.HandleFunc("/weather/", func(w http.ResponseWriter, r *http.Request) {
    cityCode := strings.SplitN(r.URL.Path, "/", 2)[1]

    data, err := query(cityCode)
    if err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
      return
    }

    w.Header().Set("Content-Type", "application/json; charset=utf-8")
    json.NewEncoder(w).Encode(data)
  })
  http.ListenAndServe(":8080", nil)
}

func hello(w http.ResponseWriter, r *http.Request){
  w.Write([]byte("hello!"))
}

func query(cityCode string) (weatherData, error) {
  resp, err := http.Get("http://t.weather.sojson.com/api/weather/city/" + cityCode)
  if err != nil {
    return weatherData{}, err
  }

  defer resp.Body.Close()

  var d weatherData
  
  if err := json.NewDecoder(resp.Body).Decode(&d); err != nil {
    return weatherData{}, err
  }

  return d, nil
}

type weatherData struct {
  Time string `json:"time"`
  CityInfo struct {
    City string `json:"city"`
    CityCode string `json:"cityId"`
    Parent string `json:"parent"`
  } `json:"cityInfo"`
  Data struct {
    Shidu string `json:"shidu"`
    Pm25 string `json:"pm25"`
    Pm10 string `json:"pm10"`
    Quality string `json:"quality"`
    Wendu string `json:"wendu"`
    Ganmao string `json:"ganmao"`
  } `json:"data"`
}