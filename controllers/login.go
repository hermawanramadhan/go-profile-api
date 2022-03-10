package controllers

import (
	"customer-profile/entities"
	"customer-profile/models"
	"customer-profile/utils"
	"encoding/json"
	"fmt"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()

	var user entities.User
	err := d.Decode(&user)
	if err != nil {
		utils.JsonErrorResponse(w, http.StatusBadRequest, fmt.Sprintf("Bad json request: %v", err.Error()))
		return
	}

	if user.Name == "" {
		utils.JsonErrorResponse(w, http.StatusBadRequest, "Bad request: name field required")
		return
	}

	if user.Password == nil {
		utils.JsonErrorResponse(w, http.StatusBadRequest, "Bad request: password field required")
		return
	}

	user, status := models.UserPasswordCheck(user.Name, *user.Password)

	if status {
		utils.JsonSuccessResponse(w, user, "Login success")
	} else {
		utils.JsonErrorResponse(w, http.StatusUnauthorized, "Login failed")
	}
}
