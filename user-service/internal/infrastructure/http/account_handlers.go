package http

import (
	"net/http"
	"user-service/config"
	"user-service/internal/application"
	"user-service/internal/domain"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type UserHandler struct {
	router  *gin.Engine
	useCase *application.AccountUseCase
}

func NewUserHandler(usecase *application.AccountUseCase) *UserHandler {
	router := gin.Default()
	return &UserHandler{useCase: usecase, router: router}
}

func (uh *UserHandler) Run(conf *config.REST) error {
	return uh.router.Run(":" + conf.Port)
}

func (uh *UserHandler) SetHandler(method, URL string, handlers ...gin.HandlerFunc) {
	switch method {
	case http.MethodGet:
		uh.router.GET(URL, handlers...)
	case http.MethodPost:
		uh.router.POST(URL, handlers...)
	case http.MethodPatch:
		uh.router.PATCH(URL, handlers...)
	case http.MethodPut:
		uh.router.PUT(URL, handlers...)
	case http.MethodDelete:
		uh.router.DELETE(URL, handlers...)
	default:
		logrus.Warn("On this time http-method " + method + " not implemented")
		return
	}
}

func (uh *UserHandler) Register(c *gin.Context) {
	var account domain.Account
	if err := c.BindJSON(&account); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{`error`: `invalid request`})
		return
	}
	if len(account.Login) < 4 || len(account.Password) < 6 {
		c.JSON(http.StatusBadRequest, gin.H{`error`: `invalid login or password`})
		return
	}
	if err := uh.useCase.Register(&account.UUID, account.Login, account.Password); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{`error`: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		`uuid`:  account.UUID,
		`login`: account.Login,
	})
}

func (uh *UserHandler) Login(c *gin.Context) {
	var account domain.Account
	if err := c.BindJSON(&account); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{`error`: err.Error()})
		return
	}
	if len(account.Login) < 4 || len(account.Password) < 6 {
		c.JSON(http.StatusBadRequest, gin.H{`error`: `invalid login or password`})
		return
	}
	newAccount, err := uh.useCase.Login(account.Login, account.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{`error`: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		`uuid`:  newAccount.UUID,
		`login`: newAccount.Login,
	})
}
