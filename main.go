package main
import (
  "fmt"
  // "bufio"
  // "os"
  // "strings"
  "encoding/json"
  "net/http"
  // "io/ioutil"
  // "time"
)
type Response struct {
  currentObservation struct {
    weather string `json:"weather"`
    tempF string `json:"temp_f"`
  } `json:"current_observation"`
}

// type CurrentObservation struct {
//
// }

func main() {
  var url string = "http://api.wunderground.com/api/f28a93cae85945b6/conditions/q/CA/San_Francisco.json"
  req, err := http.NewRequest("GET", url, nil)
  if err != nil {
    fmt.Println("NewRequest: ", err)
    return
  }

  client := &http.Client{}

  resp, err := client.Do(req)
  if err != nil {
    fmt.Println("Do: ", err)
    return
  }

  defer resp.Body.Close()

  var responseOne Response

  if err := json.NewDecoder(resp.Body).Decode(&responseOne); err != nil {
    fmt.Println(err)
  }

  fmt.Println(responseOne)
  fmt.Println(responseOne.currentObservation)
  fmt.Println(responseOne.currentObservation.weather)
}
