package main

import (
//"io"
"net/http"
  //"os"
  //"syscall"
  //"log"
  "strings"
  "fmt"
  "encoding/json"
  //"strconv"

)
type BusJSON struct {
                   BusServices []struct 
                   { Time string `json:"Time"` 
                   Destination string `json:"Destination"` 
                   BusNo int `json:"Bus_No"` 
                   } `json:"busServices"` 
  }

func busno(w http.ResponseWriter, r *http.Request) {
stopname := r.URL.Path[len("/"):]
s := string([]byte(stopname))
fmt.Println(s)

  var result BusJSON
  client := &http.Client{}
  req, _ := http.NewRequest("GET", "https://api.myjson.com/bins/j46st", nil)
  req.Header.Add("Accept", "application/json")
  resp, err := client.Do(req)
  fmt.Println("I am here error check 1")
   
   jsonParser := json.NewDecoder(resp.Body)
   if err = jsonParser.Decode(&result); err != nil {
      fmt.Println("parsing config file", err.Error())
      fmt.Println("I am here error check 2")                                                
  }
   for i := range result.BusServices {
   			dest := string(result.BusServices[i].Destination)
   			 
   			 if(strings.Contains(dest,stopname)){
   			 	fmt.Println("found the stop",dest,result.BusServices[i].BusNo)

   			 }
   	      
  	}
  
}


func main(){

	fmt.Println("inside main")
	http.HandleFunc("/",busno)
	
	http.ListenAndServe(":8000", nil)
}