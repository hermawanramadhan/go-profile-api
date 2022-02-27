package controllers

import (
	"customer-profile/entities"
	"customer-profile/models"
	"customer-profile/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {

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

	if user.Age <= 0 {
		utils.JsonErrorResponse(w, http.StatusBadRequest, "Bad request: invalid age")
		return
	}

	if *user.Password == "" {
		utils.JsonResponse(w, http.StatusBadRequest, "Bad request: password field required")
		return
	}

	password_hashed, _ := utils.HashPassword(*user.Password)
	user.Password = &password_hashed

	var riskProfile entities.RiskProfile
	if user.Age >= 30 {
		riskProfile.StockPercent = 72.5
		riskProfile.BondPercent = 21.5
	} else if user.Age >= 20 {
		riskProfile.StockPercent = 54.5
		riskProfile.BondPercent = 25.5
	} else {
		riskProfile.StockPercent = 34.5
		riskProfile.BondPercent = 45.5
	}

	riskProfile.MMPercent = 100 - riskProfile.StockPercent - riskProfile.BondPercent

	user.RiskProfile = &riskProfile
	result := models.CreateUser(&user)
	if result.Error != nil {
		utils.JsonErrorResponse(w, http.StatusInternalServerError, "Internal server error: Can't create a user ")
		return
	}

	//hide password
	user.Password = nil
	utils.JsonSuccessResponse(w, user, "Success create user")

}

func GetUsers(w http.ResponseWriter, r *http.Request) {

	pageInt := 1
	takeInt := 10

	page := r.URL.Query().Get("page")
	take := r.URL.Query().Get("take")

	var err error
	if page != "" {
		pageInt, err = strconv.Atoi(page)
		if err != nil {
			utils.JsonErrorResponse(w, http.StatusBadRequest, "Bad request: invalid query parameter page ")
			return
		}
	}

	if take != "" {
		takeInt, err = strconv.Atoi(take)
		if err != nil {
			utils.JsonErrorResponse(w, http.StatusBadRequest, "Bad request: invalid query parameter take ")
			return
		}
	}

	offset := (pageInt - 1) * takeInt
	limit := takeInt

	users := models.GetUsers(limit, offset)
	utils.JsonSuccessResponse(w, users, "Success get users")
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["id"]

	fmt.Println(vars)

	user, result := models.GetUser(userId)
	if result.RowsAffected == 1 {
		utils.JsonSuccessResponse(w, user, "Success get user")
	} else {
		utils.JsonErrorResponse(w, http.StatusNotFound, "User not found")
	}

}
