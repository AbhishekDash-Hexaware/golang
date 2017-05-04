package main

import (
  "encoding/json"
  "net/http"
)

type Profile struct {
  BusNumber    int
  session    string
  BusStops []string
  BuTime  []string
  session1   string
  BusStops1 []string
  BuTime1  []string
}

func main() {
  http.HandleFunc("/", foo)
  http.ListenAndServe(":3000", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
  profile := Profile{
    601,
    "morning", 
    []string{"Alandur Bus Depot","Trident","Nanganallur Chidambaram Stores","Nanganallur Anjaneyar Temple ","Madipakkam Kumaran Theatre","Madipakkam Ponniamman Temple","Madipakkam Sadhasivam Nagar Via Thoraipakkam Bypass","Siruseri"},
    []string{"6:55:00 AM","7:00:00 AM","7:07:00 AM","7:13:00 AM","7:15:00 AM","8:05:00 AM"},
    "evening",
    []string{"Siruseri","Ponniyamman Temple Bus Stop Via Medavakkam","Jayachandran Nagar","Pallikaranai Bus Stop","Kamakshi Hosptal Junction ","Kaiveli ","Madipakkam Ram Nagar","Madipakkam Sadhasivam Nagar","Madipakkam Ponniamman Temple","Madipakkam Kumaran Theatre","Nanganallur Anjaneyar Temple ","Nanganallur Chidambaram Stores","Trident","Alandur Bus Depot"},
    []string{"00:00:00 AM","0:00:00 AM","0:00:00 AM","0:00:00 AM","0:00:00 AM","0:00:00 AM"},
    }

  js, err := json.Marshal(profile)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
  
  
  w.Header().Set("Content-Type", "application/json")
  w.Write(js)
}
