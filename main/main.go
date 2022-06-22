package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

var Users [6]User

type User struct {
	ID         int
	Email      string
	First_name string
	Last_name  string
	Avatar     string
	Message    string
}

// Response struct
type Response struct {
	Page        int
	Per_page    int
	Total       int
	Total_pages int
	Users       []User `json:"data"`
}

// import user information
func importUsers() {
	// get user info from https://reqres.in/api/users
	res, err := http.Get("https://reqres.in/api/users")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var response Response
	json.Unmarshal(body, &response)

	for i, p := range response.Users {
		p.Message = ""
		Users[i] = p
	}
}

func main() {

	importUsers()

	router := gin.Default()

	// make the map
	m := make(map[string]string)
	for _, p := range Users {
		m[p.Email] = p.First_name
	}

	// put map in the another of router for authentication
	authorized := router.Group("/admin", gin.BasicAuth(m))

	// get request
	authorized.GET("/user", getMessage)

	// post request
	authorized.POST("/user", postMessage)

	router.Run(":8080")
}

/**
save user's message to local folder
*/
func saveFile(u User) {

	userID := strconv.Itoa(u.ID)

	f, err := os.Create(userID + ".txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	_, err2 := f.WriteString(u.Message)

	if err2 != nil {
		log.Fatal(err2)
	}

	fmt.Println("saving message done")
}

func getMessage(c *gin.Context) {
	authUserEmail := c.MustGet(gin.AuthUserKey).(string)

	for _, a := range Users {
		if a.Email == authUserEmail {
			if a.Message == "" {
				c.IndentedJSON(http.StatusNotFound, gin.H{"message": "user never saved a message"})
				return
			} else {
				c.IndentedJSON(http.StatusOK, a.Message)

				// save the message to local filesystem
				saveFile(a)
				return
			}
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "user is not found"})
}

func postMessage(c *gin.Context) {
	authUserEmail := c.MustGet(gin.AuthUserKey).(string)
	var newUser User

	for _, a := range Users {
		if a.Email == authUserEmail {
			if err := c.BindJSON(&newUser); err != nil {
				return
			}
			id := a.ID

			fmt.Printf("i=%d, type: %T\n", id, id)
			Users[id-1].Message = newUser.Message
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "user is not found"})
}
