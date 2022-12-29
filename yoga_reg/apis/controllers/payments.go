package controllers

import (
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/ha-wk/yoga_reg/apis/auth"
	"github.com/ha-wk/yoga_reg/apis/models"
	"github.com/ha-wk/yoga_reg/apis/responses"
	razorpay "github.com/razorpay/razorpay-go"
)

type PageVariables struct {
	OrderId string
}

/*
	func main() {
		http.HandleFunc("/", App)
		log.Fatal(http.ListenAndServe(":8089", nil))
	}
*/
func App(w http.ResponseWriter, r *http.Request) {

	client := razorpay.NewClient("key_id", "secret_key")
	data := map[string]interface{}{
		"amount":   50000, //IN PAISE.../ BY 100
		"currency": "INR",
		"receipt":  "some_receipt_id",
	}
	body, err := client.Order.Create(data, nil)

	if err != nil {
		fmt.Printf("PROBLEM GETTING REPO INFO %V\n", err)
		os.Exit(1)
	}

	value := body["id"]
	str := value.(string)
	HomePageVars := PageVariables{
		OrderId: str,
	}

	uid, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("UNAUTHORIZED"))
		return
	}
	user := models.User{}
	if uid != user.ID { //check if the user attempting to update a post belonging to him or not
		responses.ERROR(w, http.StatusUnauthorized, errors.New("UNAUTHORIZED"))
		return
	}

	t, err := template.ParseFiles("app.html")
	if err != nil {
		log.Print("templateparsing error:", err)
	}
	err = t.Execute(w, HomePageVars)
	if err != nil {
		log.Print("template executing error", err)
	}
}
