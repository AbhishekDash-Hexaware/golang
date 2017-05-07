
package main
 
import (
  "io"
  "net/http"
  "os"
  "os/signal"
  "syscall"
  "log"
  "fmt"
  "encoding/json"
  "strconv"
)
 
func hello(w http.ResponseWriter, r *http.Request) {
                io.WriteString(w, "Hello world!")
}
 
func hi(w http.ResponseWriter, r *http.Request){
  io.WriteString(w, "Hi!")
}
 
func bus(w http.ResponseWriter, r *http.Request) {
 
                fmt.Fprint(w, "Welcome!\n")
 
                client := &http.Client{}
 
                req, _ := http.NewRequest("GET", "http://api.myjson.com/bins/17y44l", nil)
                req.Header.Add("Accept", "application/json")
                resp, err := client.Do(req)
 
                type BusJSON struct {
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
 
                var result BusJSON
 
                jsonParser := json.NewDecoder(resp.Body)
  if err = jsonParser.Decode(&result); err != nil {
      fmt.Println("parsing config file", err.Error())
  }
 
                for i := range result.BusRoute {
      fmt.Println(result.BusRoute[i].BusNo)
                                                // if(result.BusRoute[i].BusNo == 601){
                                                //            fmt.Println(result.BusRoute[i].Morning.BusRouteText)
                                                // }
  }
 
}
 
 
func busno(w http.ResponseWriter, r *http.Request) {
 
                name := r.URL.Path[len("/bus/route/"):]
                s := string([]byte(name))
                fmt.Println(s)
 
                //fmt.Fprint(w, "Welcome2!\n")
 
                client := &http.Client{}
 
                req, _ := http.NewRequest("GET", "http://api.myjson.com/bins/17y44l", nil)
                req.Header.Add("Accept", "application/json")
                resp, err := client.Do(req)
 
                type BusJSON struct {
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
 
                var result BusJSON
 
                jsonParser := json.NewDecoder(resp.Body)
  if err = jsonParser.Decode(&result); err != nil {
      fmt.Println("parsing config file", err.Error())
  }
 
                for i := range result.BusRoute {
                                                ss, err := strconv.Atoi(s)
                                                if err != nil {
                                                    fmt.Printf("i=%d, type: %T\n", ss, ss)
                                                }
 
                                                if(result.BusRoute[i].BusNo == ss){
                                                                //fmt.Println(result.BusRoute[i].Morning.BusRouteText)
                                                                fmt.Fprint(w, result.BusRoute[i].Morning.BusRouteText)
                                                }
  }
 
}
 
func cleanup() {
    log.Print("Cleaning Logger")
}
 
func main() {
                //http.HandleFunc("/", hello)
  //http.HandleFunc("/hi", hi)
                http.HandleFunc("/bus/number/", bus)
                http.HandleFunc("/bus/route/", busno)
 
  log.Println("\n\r ================================================ \n\r")
 
  fmt.Println("Starting server on http://107.23.138.151:8000")
  log.Println("Starting server\n\r")
  fmt.Println("Press Ctrl+C to stop the server....")
 
  c := make(chan os.Signal, 2)
  signal.Notify(c, os.Interrupt, syscall.SIGTERM)
  go func() {
      <-c
      log.Println("Stopping server\n\r")
      cleanup()
      os.Exit(0)
  }()
 
                http.ListenAndServe(":8000", nil)
 
}
