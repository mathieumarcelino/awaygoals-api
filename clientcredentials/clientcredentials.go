package clientcredentials

import (
	"encoding/json"
    "net/http"
    "net/url"
    "strconv"
    "strings"
    "log"
    "io/ioutil"
)

type Credentials struct {
	Access_token string
	Expires_in int
	Token_type string
}

func Clientcredentials() string {

	endpoint := "URL"
    data := url.Values{}
    data.Set("grant_type", "client_credentials")

    client := &http.Client{}
    r, err := http.NewRequest("POST", endpoint, strings.NewReader(data.Encode()))
    if err != nil {
        log.Fatal(err)
    }
    r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Authorization", "KEY")
    r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

    res, err := client.Do(r)
    if err != nil {
        log.Fatal(err)
    }
    log.Println(res.Status)
    defer res.Body.Close()
    body, err := ioutil.ReadAll(res.Body)
    if err != nil {
        log.Fatal(err)
    }

	var credentials Credentials
	json.Unmarshal([]byte(string(body)), &credentials)

	final_credentials := credentials.Token_type + " " + credentials.Access_token

	return final_credentials
}


