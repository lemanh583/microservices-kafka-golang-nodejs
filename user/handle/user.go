package handle

import (
	"fmt"
	"net/http"
	"user-services/config"
	"user-services/dto"
	"user-services/service"
	"user-services/util"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type userHandle struct {
	UserService  service.UserService
	MailService  service.MailService
	KafkaService service.KafkaService
}

type UserHandle interface {
	SignUp(c *gin.Context)
	SignIn(c *gin.Context)
}

func NewUserHandle(us service.UserService, ms service.MailService, ks service.KafkaService) UserHandle {
	return &userHandle{
		UserService:  us,
		MailService:  ms,
		KafkaService: ks,
	}
}

func (h *userHandle) SignUp(c *gin.Context) {
	fmt.Println("___h____", h)
	validate := validator.New(validator.WithRequiredStructEnabled())
	var body dto.ReqSignUp
	err := c.ShouldBindJSON(&body)
	if err != nil {
		util.ResponseError(c, http.StatusBadRequest, err.Error())
		return
	}

	err = util.TrimSpace(&body)
	if err != nil {
		util.ResponseError(c, http.StatusBadRequest, err.Error())
		return
	}

	err = validate.Struct(body)
	if err != nil {
		util.ResponseError(c, http.StatusBadRequest, err.Error())
		return
	}

	rs, err := h.UserService.SignUp(c, &body)

	if err != nil {
		util.ResponseError(c, http.StatusBadRequest, err.Error())
		return
	}

	// // send mail
	mailData := dto.MailTransferData{
		To:       []string{rs.Email},
		Subject:  "Verify email",
		Template: fmt.Sprintf("%s/verify/%s", config.Cfg.DomainGateway, rs.ID),
	}
	go h.KafkaService.SendMessage(util.StructToMap(mailData))

	util.ResponseJson(c, http.StatusOK, map[string]interface{}{})
}

func (h *userHandle) SignIn(c *gin.Context) {
	validate := validator.New(validator.WithRequiredStructEnabled())
	var body dto.ReqSignIn
	err := c.ShouldBindJSON(&body)
	if err != nil {
		util.ResponseError(c, http.StatusBadRequest, err.Error())
		return
	}

	err = util.TrimSpace(&body)
	if err != nil {
		util.ResponseError(c, http.StatusBadRequest, err.Error())
		return
	}

	err = validate.Struct(body)
	if err != nil {
		util.ResponseError(c, http.StatusBadRequest, err.Error())
		return
	}

	rs, err := h.UserService.SignIn(c, &body)
	if err != nil {
		util.ResponseError(c, http.StatusBadRequest, err.Error())
		return
	}

	util.ResponseJson(c, http.StatusOK, rs)
}
