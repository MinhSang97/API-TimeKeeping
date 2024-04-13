package handler

import (
	"app/usecases"
	attendance_dto "app/usecases/dto/attendance_dto"
	"app/usecases/req"
	"app/usecases/res"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"strconv"
)

func AttendanceCheckOut() func(*gin.Context) {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, res.Response{
				StatusCode: http.StatusBadRequest,
				Message:    err.Error(),
				Data:       "Không tìm thấy ID",
			})
			return
		}

		var validate *validator.Validate
		validate = validator.New(validator.WithRequiredStructEnabled())
		req := req.ReqAttendaceCheckOut{}
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

		attendance := attendance_dto.Attendance{
			CheckOutTime:      &req.CheckOutTime,
			LatitudeCheckOut:  req.LatitudeCheckOut,
			LongitudeCheckOut: req.LongitudeCheckOut,
		}

		data := attendance.ToPayload().ToModel()
		uc := usecases.NewAttendanceUseCase()

		err = uc.CheckOut(c.Request.Context(), id, data)

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
			Data:       map[string]interface{}{"Check-out success with ID": id},
		})
	}
}
