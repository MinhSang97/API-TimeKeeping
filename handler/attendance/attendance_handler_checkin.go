package handler

import (
	"app/usecases"
	attendance_dto "app/usecases/dto/attendance_dto"
	"app/usecases/req"
	"app/usecases/res"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

func AttendanceCheckIn() func(*gin.Context) {
	return func(c *gin.Context) {
		var validate *validator.Validate
		validate = validator.New(validator.WithRequiredStructEnabled())
		req := req.ReqAttendaceCheckIn{}
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
			UserId:           req.UserId,
			Department:       req.Department,
			CheckInTime:      &req.CheckInTime,
			LatitudeCheckIn:  req.LatitudeCheckIn,
			LongitudeCheckIn: req.LongitudeCheckIn,
		}

		data := attendance.ToPayload().ToModel()
		uc := usecases.NewAttendanceUseCase()

		err := uc.CheckIn(c.Request.Context(), data)

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
			Data:       map[string]interface{}{"Check-in success with ID": data.ID},
		})
	}
}
