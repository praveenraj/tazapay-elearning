package logic

import (
	"strings"

	"github.com/rs/zerolog/log"

	"tazapay.com/elearning/svc/user/constants"

	"tazapay.com/elearning/common/models"
	"tazapay.com/elearning/common/repository"
	"tazapay.com/elearning/common/request"
	"tazapay.com/elearning/common/responses"
	"tazapay.com/elearning/common/utils"
)

// RegisterUser creates new user in our system
func RegisterUser(payload *request.Payload) *responses.Response {
	var user constants.NewUser
	err := utils.JSONMarshalAndUnmarshal(payload.GetAPI().Body, &user)
	if err != nil {
		log.Error().Msgf("error marshaling create user request: %v", err)
		return responses.InternalError()
	}

	// validate the struct
	isValid, field := utils.IsValidStruct(user)
	if !isValid {
		return responses.BadRequest(utils.MakeResponseError("general", 1100, field)) // value moves to env
	}

	// add new user
	m := models.User{
		FirstName: user.FirstName,
		LastName:  &user.LastName,
		Mobile:    user.Mobile,
		Email:     user.Email,
	}
	err = repository.NewUserDAO().Create(&m)
	if err != nil {
		log.Error().Msgf("error registering new user: %v", err)
		if strings.Contains(err.Error(), "Error 1062") { // mobile or email already registered
			return responses.Conflict()
		}
		return responses.InternalError()
	}
	return responses.Success()
}
