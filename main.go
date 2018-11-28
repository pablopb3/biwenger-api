package main

import (
    "fmt"
    "html"
    "log"
    "encoding/json"
    "net/http"
    "bytes"
    "io/ioutil"
    "github.com/gorilla/mux"
)


func main() {
    router := mux.NewRouter().StrictSlash(true)
    router.HandleFunc("/", Index)
    router.HandleFunc("/getMyPlayers", GetMyPlayers)
    router.HandleFunc("/setLineUp", SetLineUp)
    log.Fatal(http.ListenAndServe(":8080", router))
}

func Index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func login() string {
	user, _ := json.Marshal(User{"pablopb3@gmail.com", "maluco88"})
	response := doRequest("https://biwenger.as.com/api/v2/auth/login", "POST", string(user))
    token := Token{}
    json.Unmarshal([]byte(response), &token)
    return token.Token
}

func doRequest(url string, method string, payload string) string {
	
	var jsonStr = []byte(payload)
    req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonStr))    
    req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	return string(body)

}

func GetMyPlayers(w http.ResponseWriter, r *http.Request) {

    token := login() // or &Foo{}

    url := "https://biwenger.as.com/api/v2/user?fields=*,lineup(type,playersID),players(*,fitness,team,owner),market(*,-userID),offers,-trophies"

    req, err := http.NewRequest("GET", url, nil)    
    req.Header.Set("Content-Type", "application/json")
    req.Header.Set("authorization", "Bearer " + token)
    req.Header.Set("x-league", "757450")
    req.Header.Set("x-version", "560")


	client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
    fmt.Fprintf(w, string(body))
    
}

func SetLineUp(w http.ResponseWriter, r *http.Request) {

    token := login()

    url := "https://biwenger.as.com/api/v2/user?fields=*"

    formation := "4-4-2"
    playerIds := []int{11677,1752,9078,2160,17052,944,9071,1043,10750,10498,15568}
	lineup, _ := json.Marshal(StartingEleven{LineUp{formation, playerIds}})
	var jsonStr = []byte(lineup)
    req, err := http.NewRequest("PUT", url, bytes.NewBuffer(jsonStr))     
    req.Header.Set("Content-Type", "application/json")
    req.Header.Set("authorization", "Bearer " + token)
    req.Header.Set("x-league", "757450")
    req.Header.Set("x-version", "560")

	client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
    fmt.Fprintf(w, string(body))
    
}



