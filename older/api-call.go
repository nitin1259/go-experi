package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// user api ->
/*
{
    "id": 1,
    "name": "Leanne Graham",
    "username": "Bret",
    "email": "Sincere@april.biz",
    "address": {
      "street": "Kulas Light",
      "suite": "Apt. 556",
      "city": "Gwenborough",
      "zipcode": "92998-3874",
      "geo": {
        "lat": "-37.3159",
        "lng": "81.1496"
      }
    },
    "phone": "1-770-736-8031 x56442",
    "website": "hildegard.org",
    "company": {
      "name": "Romaguera-Crona",
      "catchPhrase": "Multi-layered client-server neural-net",
      "bs": "harness real-time e-markets"
    }
  },
*/
// user post
/*
{
    "userId": 1,
    "id": 1,
    "title": "sunt aut facere repellat provident occaecati excepturi optio reprehenderit",
    "body": "quia et suscipit\nsuscipit recusandae consequuntur expedita et cum\nreprehenderit molestiae ut ut quas totam\nnostrum rerum est autem sunt rem eveniet architecto"
  },
*/

// expose -  any  -> [] cerain value

type Address struct {
	Geo Geo `json:"geo"`
}

type Company struct {
	Name string `json:"name"`
}

type Geo struct {
	Lat string `json:"lat"`
	Lng string `json:"lng"`
}

type User struct {
	Id      int     `json:"id"`
	Address Address `json:"address"`
	Company Company `json:"company"`
}

type Posts struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

func api_main() {

	http.HandleFunc("/companydetails", getCompanyDetails)

	log.Fatal(http.ListenAndServe(":8080", nil))

}

func getCompanyDetails(w http.ResponseWriter, req *http.Request) {

	resp1, err := http.Get("https://jsonplaceholder.typicode.com/users")
	if err != nil {
		log.Fatal(err)
	}
	defer resp1.Body.Close()

	if resp1.StatusCode != http.StatusOK {
		log.Fatal("Error while callig api")
	}

	// respData := make(map[string]interface{})
	var users []User
	err = json.NewDecoder(resp1.Body).Decode(&users)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("suers: ", users)

	resp2, err := http.Get("https://jsonplaceholder.typicode.com/posts")
	if err != nil {
		log.Fatal(err)
	}
	defer resp2.Body.Close()

	if resp2.StatusCode != http.StatusOK {
		log.Fatal("Error while callig api")
	}

	var posts []Posts

	err = json.NewDecoder(resp2.Body).Decode(&posts)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("suers: ", posts)
	var num int
	if len(users) < len(posts) {
		num = len(users)
	} else {
		num = len(posts)
	}

	matchrecord := []CompanyDetails{}
	for i := 0; i < num; i++ {

		if users[i].Id == posts[i].Id {
			cd := CompanyDetails{
				Id:    users[i].Id,
				Lat:   users[i].Address.Geo.Lat,
				Lon:   users[i].Address.Geo.Lng,
				Name:  users[i].Company.Name,
				Title: posts[i].Title,
				Body:  posts[i].Body,
			}

			matchrecord = append(matchrecord, cd)
		}
	}

	w.Header().Set("Content-Type", "appication/json")
	json.NewEncoder(w).Encode(matchrecord)

}

type CompanyDetails struct {
	Id    int    `jsong:"id"`
	Lat   string `json:"lat"`
	Lon   string `json:"lon"`
	Name  string `json:"name"`
	Title string `json:"title"`
	Body  string `json:"body"`
}
