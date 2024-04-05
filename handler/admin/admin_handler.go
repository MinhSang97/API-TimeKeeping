package handler

import (
	"app/log"
	"app/model"
	"app/model/req"
	"app/sercurity"
	"app/usecases"
	"app/usecases/dto"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	uuid "github.com/google/uuid"
	"net/http"
)

//func HandleSignUp() func(*gin.Context) {
//	return func(c *gin.Context) {
//		var validate *validator.Validate
//		validate = validator.New(validator.WithRequiredStructEnabled())
//		req := req.ReqSignUp{}
//		if err := c.ShouldBind(&req); err != nil {
//			c.JSON(http.StatusBadRequest, usecases.Response{
//				StatusCode: http.StatusBadRequest,
//				Message:    err.Error(),
//				Data:       nil,
//			})
//			return
//		}
//
//		if err := validate.Struct(req); err != nil {
//			//log.Error(err.Error())
//			c.JSON(http.StatusForbidden, usecases.Response{
//				StatusCode: http.StatusForbidden,
//				Message:    err.Error(),
//				Data:       nil,
//			})
//			return
//		}
//
//		var Data dto.Admin
//
//		if err := c.ShouldBind(&Data); err != nil {
//			c.JSON(http.StatusBadRequest, gin.H{
//				"error": err.Error(),
//			})
//			return
//		}
//
//		err := validate.Struct(Data)
//
//		if err != nil {
//			c.JSON(http.StatusBadRequest, gin.H{
//				"error": err.Error(),
//			})
//			return
//		}
//
//		hash := sercurity.HashAndSalt([]byte(req.PassWord))
//		role := model.ADMIN.String()
//
//		userAdminId, err := uuid.NewUUID()
//
//		if err != nil {
//			log.Error(err.Error())
//			c.JSON(http.StatusForbidden, usecases.Response{
//				StatusCode: http.StatusForbidden,
//				Message:    err.Error(),
//				Data:       nil,
//			})
//			return
//		}
//
//		userAdmin := dto.Admin{
//			UserId:   userAdminId.String(),
//			Name:     req.Name,
//			PassWord: hash,
//			Email:    req.Email,
//			Role:     role,
//			Token:    "",
//		}
//
//		data := userAdmin.ToPayload().ToModel()
//		uc := usecases.NewAdminUseCase()
//
//		err = uc.CreateAdmin(c.Request.Context(), data)
//
//		if err != nil {
//			c.JSON(http.StatusConflict, usecases.Response{
//				StatusCode: http.StatusConflict,
//				Message:    err.Error(),
//				Data:       nil,
//			})
//			return
//		}
//		//gen token
//		token, err := sercurity.GenToken(userAdmin)
//		if err != nil {
//			//log.Error(err.Error())
//			c.JSON(http.StatusInternalServerError, usecases.Response{
//				StatusCode: http.StatusInternalServerError,
//				Message:    err.Error(),
//				Data:       nil,
//			})
//			return
//		}
//		userAdmin.Token = token
//
//		c.JSON(http.StatusOK, gin.H{
//			"data": data.ID,
//		})
//	}
//}

func HandleSignUp() func(*gin.Context) {
	return func(c *gin.Context) {
		var validate *validator.Validate
		validate = validator.New(validator.WithRequiredStructEnabled())
		req := req.ReqSignUp{}
		if err := c.ShouldBind(&req); err != nil {
			c.JSON(http.StatusBadRequest, usecases.Response{
				StatusCode: http.StatusBadRequest,
				Message:    err.Error(),
				Data:       nil,
			})
			return
		}

		if err := validate.Struct(req); err != nil {
			c.JSON(http.StatusForbidden, usecases.Response{
				StatusCode: http.StatusForbidden,
				Message:    err.Error(),
				Data:       nil,
			})
			return
		}

		hash := sercurity.HashAndSalt([]byte(req.PassWord))
		role := model.ADMIN.String()

		userAdminId, err := uuid.NewUUID()

		if err != nil {
			log.Error(err.Error())
			c.JSON(http.StatusForbidden, usecases.Response{
				StatusCode: http.StatusForbidden,
				Message:    err.Error(),
				Data:       nil,
			})
			return
		}

		//gen token
		token, err := sercurity.GenToken(dto.Admin{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, usecases.Response{
				StatusCode: http.StatusInternalServerError,
				Message:    err.Error(),
				Data:       nil,
			})
			return
		}

		//userAdmin := dto.Admin{
		//	UserId:   userAdminId.String(),
		//	Name:     req.Name,
		//	PassWord: hash,
		//	Email:    req.Email,
		//	Role:     role,
		//	Token:    token,
		//}

		userAdmin := dto.Admin{
			UserId:   userAdminId.String(),
			Name:     req.Name,
			PassWord: hash,
			Email:    req.Email,
			Role:     role,
			Token:    token,
		}

		err = validate.Struct(userAdmin)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		data := userAdmin.ToPayload().ToModel()
		uc := usecases.NewAdminUseCase()

		err = uc.CreateAdmin(c.Request.Context(), data)

		if err != nil {
			c.JSON(http.StatusConflict, usecases.Response{
				StatusCode: http.StatusConflict,
				Message:    err.Error(),
				Data:       nil,
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data": data.ID,
		})
	}
}
