package delivery

import (
	"bookapi/feature/book/domain"
	//"bookapi/feature/book/repository"
	"net/http"

	"github.com/labstack/echo/v4"
)

type bookHandler struct {
	srv domain.Service
}

func New(e *echo.Echo, srv domain.Service) {
	handler := bookHandler{srv: srv}
	e.GET("/books", handler.ShowAllBook())
	e.POST("/books", handler.AddBook())

}

func (bs *bookHandler) AddBook() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input AddBookFormat
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, FailResponse("cannot bind input"))
		}
		cnv := ToDomain(input)
		res, err := bs.srv.AddBook(cnv)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, FailResponse(err.Error()))
		}

		return c.JSON(http.StatusCreated, SuccessResponse("berhasil register", ToResponse(res, "reg")))
	}

}

func (bs *bookHandler) ShowAllBook() echo.HandlerFunc {
	return func(c echo.Context) error {
		res, err := bs.srv.AllBook()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, FailResponse(err.Error()))
		}

		return c.JSON(http.StatusOK, SuccessResponse("success get all user", ToResponse(res, "all")))
	}
}

// func (bs *bookHandler) DeleteBook() echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		var ID uint
// 		res, err :

// 	}
// }
