package handler

import (
	"app/log"
	"app/payload"
	"app/sercurity"
	"app/usecases"
	usersdto "app/usecases/dto/users_dto"
	"app/usecases/req"
	"app/usecases/res"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	uuid "github.com/google/uuid"
	"net/http"
)

func UsersSignUp() func(*gin.Context) {
	return func(c *gin.Context) {
		var validate *validator.Validate
		validate = validator.New(validator.WithRequiredStructEnabled())
		req := req.ReqSignUp{}
		if err := c.ShouldBind(&req); err != nil {
			c.JSON(http.StatusBadRequest, res.Response{
				StatusCode: http.StatusBadRequest,
				Message:    err.Error(),
				Data:       nil,
			})
			return
		}

		if err := validate.Struct(req); err != nil {
			c.JSON(http.StatusForbidden, res.Response{
				StatusCode: http.StatusForbidden,
				Message:    err.Error(),
				Data:       nil,
			})
			return
		}

		PassHash := sercurity.HashAndSalt([]byte(req.PassWord))
		role := payload.EMPLOYEE.String()

		userUsersId, err := uuid.NewUUID()

		if err != nil {
			log.Error(err.Error())
			c.JSON(http.StatusForbidden, res.Response{
				StatusCode: http.StatusForbidden,
				Message:    err.Error(),
				Data:       nil,
			})
			return
		}

		//gen token
		token, err := sercurity.GenTokenUsers(usersdto.Users{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, res.Response{
				StatusCode: http.StatusInternalServerError,
				Message:    err.Error(),
				Data:       nil,
			})
			return
		}

		userUsers := usersdto.Users{
			UserId:   userUsersId.String(),
			Name:     req.Name,
			PassWord: PassHash,
			Email:    req.Email,
			Role:     role,
			Token:    token,
		}

		err = validate.Struct(userUsers)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		data := userUsers.ToPayload().ToModel()
		uc := usecases.NewUsersUseCase()

		err = uc.CreateUsers(c.Request.Context(), data)

		if err != nil {
			c.JSON(http.StatusConflict, res.Response{
				StatusCode: http.StatusConflict,
				Message:    err.Error(),
				Data:       nil,
			})
			return
		}

		c.JSON(http.StatusOK, res.Response{
			StatusCode: http.StatusOK,
			Message:    "Xử lý thành công",
			Data:       data.UserId,
		})
	}
}
