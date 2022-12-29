package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	//"os/user"

	"github.com/ha-wk/yoga_reg/apis/auth"
	"github.com/ha-wk/yoga_reg/apis/models"
	"github.com/ha-wk/yoga_reg/apis/responses"
	"golang.org/x/crypto/bcrypt"
)

func (server *Server) Login(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	user := models.User{}

	err = json.Unmarshal(body, &user)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	user.Prepare()
	err = user.Validate("login")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	token, err := server.SignIn(user.Email, user.Password)
	if err != nil {
		//formattedError:=formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	responses.JSON(w, http.StatusOK, token)
}

func (server *Server) SignIn(email, password string) (string, error) {
	var err error
	user := models.User{}

	err = server.DB.Debug().Model(models.User{}).Where("email=?", email).Take(&user).Error
	if err != nil {
		return "", err
	}
	err = models.VerifyPassword(user.Password, password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {

		return "", err
	}
	return auth.CreateToken(user.ID)
}

/*  uid, err := auth.ExtractTokenID(r)
if err != nil {
	responses.ERROR(w, http.StatusUnauthorized, errors.New("UNAUTHORIZED"))
	return
}

if uid != post.AuthorID { //check if the user attempting to update a post belonging to him or not
	responses.ERROR(w, http.StatusUnauthorized, errors.New("UNAUTHORIZED"))
	return
}
*/
