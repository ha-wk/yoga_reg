package controllers

import (
	"errors"
	//"fmt"
	//"html/template"
	//"log"
	"net/http"
	//"os"
	"strconv"

	//"strconv"

	"github.com/gorilla/mux"
	"github.com/ha-wk/yoga_reg/apis/auth"
	"github.com/ha-wk/yoga_reg/apis/models"
	"github.com/ha-wk/yoga_reg/apis/responses"
	//razorpay "github.com/razorpay/razorpay-go"
)

type PageVariables struct {
	OrderId uint64 `gorm:"size:255;not null;" json:"id"`
}

/*
	func main() {
		http.HandleFunc("/", App)
		log.Fatal(http.ListenAndServe(":8089", nil))
	}
*/
func (server *Server) App(w http.ResponseWriter, r *http.Request) {

	/*client := razorpay.NewClient("rzp_test_5peFCjW2bKSA1p", "hcX4msA6b3BkcTbLKrri24yB")
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
	}*/

	vars := mux.Vars(r)
	pid, err := strconv.ParseInt(vars["id"], 10, 64) //parsing the id from request body
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	uid, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("UNAUTHORIZED"))
		return
	}
	user := models.User{}
	err = server.DB.Debug().Model(user).Where("id=?", pid).Take(&user).Error
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, err)
		return
	}
	if uid != user.ID { //check if the user attempting to update a post belonging to him or not
		responses.ERROR(w, http.StatusUnauthorized, errors.New("UNAUTHORIZED"))
		return
	}

	/*t, err := template.ParseFiles("app.html")
	if err != nil {
		log.Print("templateparsing error:", err)
	}
	err = t.Execute(w, HomePageVars)
	if err != nil {
		log.Print("template executing error", err)
	}*/
	responses.JSON(w, http.StatusOK, "Successful Payment")
}
