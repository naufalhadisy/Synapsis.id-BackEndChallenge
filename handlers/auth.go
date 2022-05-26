package handlers

import (
	"database/sql"
	"e-commerce/api/models"
	"e-commerce/api/restapi/operations/user"
	"e-commerce/pkg/dao"
	"e-commerce/pkg/utils"
	"fmt"

	"github.com/go-openapi/runtime/middleware"
	"golang.org/x/crypto/bcrypt"
)

type loginImpl struct {
	dbClient *sql.DB
}

func NewUserLoginHandler(db *sql.DB) user.LoginHandler {
	return &loginImpl{
		dbClient: db,
	}
}

func (impl *loginImpl) Handle(params user.LoginParams) middleware.Responder {
	email := params.Login.Email
	userInfo, err := dao.FetchUserDetails(impl.dbClient, *email)
	if err != nil {
		fmt.Println(err.Error())
		return user.NewLoginInternalServerError().WithPayload("Error fetching user details")
	}
	err = bcrypt.CompareHashAndPassword([]byte(*userInfo.Password), []byte(*params.Login.Password))
	if err != nil {
		fmt.Println(err)
		return user.NewRegisterNotFound()
	}
	token, err := utils.GenerateJWT(*email, *userInfo.Fname, userInfo.Lname)
	if err != nil {
		return user.NewLoginInternalServerError().WithPayload("Error defining token")
	}
	return user.NewLoginOK().WithPayload(&models.LoginSuccess{Success: true, Token: token})
}
