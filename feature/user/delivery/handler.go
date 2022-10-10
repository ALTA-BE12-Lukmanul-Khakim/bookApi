package delivery

import (
	"bookapi/feature/user/domain"
	"net/http"

	"github.com/labstack/echo/v4"
)

type userHandler struct {
	srv domain.Service
}

func New(e *echo.Echo, srv domain.Service) {
	handler := userHandler{srv: srv}
	e.GET("/users", handler.ShowAllUser())
	e.POST("/users", handler.AddUser())
	// e.GET("/users/:id", handler)
}

func (us *userHandler) AddUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input RegisterFormat
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, FailResponse("cannot bind input"))
		}
		cnv := ToDomain(input)
		res, err := us.srv.AddUser(cnv)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, FailResponse(err.Error()))
		}

		return c.JSON(http.StatusCreated, SuccessResponse("berhasil register", ToResponse(res, "reg")))
	}

}

// func (us *userHandler) UpdateProfile() (domain.Core, error) {

// }
// func (us *userHandler) Profile() (domain.Core, error) {
// 	res, err := us.qry.Get(ID)
// 	if err != nil {
// 		log.Error(err.Error())
// 		if strings.Contains(err.Error(), "table") {
// 			return domain.Core{}, errors.New("database error")
// 		} else if strings.Contains(err.Error(), "found") {
// 			return domain.Core{}, errors.New("no data")
// 		}
// 	}

//		return res, nil
//	}
func (us *userHandler) ShowAllUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		res, err := us.srv.ShowAllUser()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, FailResponse(err.Error()))
		}

		return c.JSON(http.StatusOK, SuccessResponse("success get all user", ToResponse(res, "all")))
	}
}
