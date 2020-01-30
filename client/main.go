package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
	jwt "github.com/dgrijalva/jwt-go"
)

// set MY_JTW_TOKEN=maysupersecretfrase
//var mySingningKey = os.Get(MY_JTW_TOKEN)
var mySigningKey = []byte("maysupersecretfrase")

// GenerateJWT - 
func GenerateJWT() (string, error){
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["authorised"] = true
	claims["user"] = "Danilo Segura" 
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()
	fmt.Println(time.Now().Add(time.Minute).Unix())

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		fmt.Println("Somenthing went wrong: " + err.Error())
		return "", err
	}
	
	return tokenString, nil
}

func homePage(w http.ResponseWriter, r *http.Request){
	validToken, err := GenerateJWT()
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	fmt.Fprintf(w, validToken)
}

func handleRequest(){
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":9001", nil))

}

func main() {
	fmt.Println("My simple Token")
	handleRequest()

	//tokenString, err := GenerateJWT()
	//if err != nil {
	//	fmt.Println("Somenthing went wrong: " + err.Error())
	//}
	//fmt.Println(tokenString)
}
