package main

import( "fmt"
		"net/http"
		"encoding/json"
		"strings"
		//"io"
		"strconv"
)
//======================================================================//
type StopJSON struct {
                   BusServices []struct 
                   { Time string `json:"Time"` 
                   Destination string `json:"Destination"` 
                   BusNo int `json:"Bus_No"` 
                   } `json:"busServices"` 
  }
type BusNumber struct {
          BusRoute []struct {
                  BusNo   int `json:"Bus_No"`
                  Evening struct {
                              BusRouteText string   `json:"BusRouteText"`
                              BusStops     []string `json:"BusStops"`
                              BusTime      []string `json:"BusTime"`
                        } `json:"evening"`
                        Morning struct {
                              BusRouteText string   `json:"BusRouteText"`
                              BusStops     []string `json:"BusStops"`
                              BusTime      []string `json:"BusTime"`
                  } `json:"morning"`
          } `json:"busRoute"`
  }
  
//======================================================================//
func destination(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "Welcome!\n")

	data := r.URL.Path[len("/destination/"):]
	s := string([]byte(data))
	fmt.Println("data captured is",s)

	var parameters []string =strings.Split(s,"/")
	var loc=strings.ToLower(parameters[0])
	var session=strings.ToLower(parameters[1])
	fmt.Println(loc," ",session)

	var jdata StopJSON
	client := &http.Client{}
	req, _:=http.NewRequest("GET","https://api.myjson.com/bins/j46st",nil)
	req.Header.Add("Accept","application/json")
	resp,err :=client.Do(req)

	jsonParser :=json.NewDecoder(resp.Body)
	if err = jsonParser.Decode(&jdata);err != nil{
	fmt.Println("Error while parsing",err.Error())
	}


	for i := range jdata.BusServices{
			if strings.Contains(strings.ToLower(jdata.BusServices[i].Destination),loc) && strings.Contains(strings.ToLower(jdata.BusServices[i].Time),session){
				fmt.Println("matched",strings.ToLower(jdata.BusServices[i].Destination),strings.ToLower(jdata.BusServices[i].Time),jdata.BusServices[i].BusNo)
			}
		}
		
}


func bus(w http.ResponseWriter, r *http.Request){

	data := r.URL.Path[len("/bus/"):]
	intbusdata, err := strconv.Atoi(data)
	
	
	var jdata BusNumber
	client := &http.Client{}
	req, _:=http.NewRequest("GET","http://api.myjson.com/bins/17y44l",nil)
	req.Header.Add("Accept","application/json")
	resp,err :=client.Do(req)
	fmt.Println("inside bus and executing")
	jsonParser :=json.NewDecoder(resp.Body)
	if err = jsonParser.Decode(&jdata);err != nil{
	fmt.Println("Error while parsing",err.Error())
}
	for i := range jdata.BusRoute{		
			if( intbusdata == jdata.BusRoute[i].BusNo){
			fmt.Println(jdata.BusRoute[i].Morning.BusRouteText)
			fmt.Println(jdata.BusRoute[i].Evening.BusRouteText)
			break
		}
	}

}


func main(){
	fmt.Println("inside main strating execution")

	http.HandleFunc("/destination/",destination)
	http.HandleFunc("/bus/",bus)
	http.ListenAndServe(":3000",nil)
}