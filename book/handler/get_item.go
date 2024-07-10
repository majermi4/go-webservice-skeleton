package handler

import (
	"MyWebService/book/repository"
	"MyWebService/lib/server"
	"github.com/gin-gonic/gin"
	"net/http"
)

type GetBookItemHandler struct {
	bookRepository *repository.BookRepository
}

func NewGetBookItemHandler(bookRepository *repository.BookRepository) *GetBookItemHandler {
	return &GetBookItemHandler{bookRepository: bookRepository}
}

func (*GetBookItemHandler) GetConfig() server.HandlerConfig {
	return server.HandlerConfig{
		Path:   "/books/:id",
		Method: http.MethodGet,
	}
}

func (h *GetBookItemHandler) HandleRequest(c *gin.Context) {
	id := c.Param("id")

	book := h.bookRepository.Get(id)

	if book == nil {
		c.Status(http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, book)
}
