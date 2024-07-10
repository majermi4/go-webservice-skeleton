package handler

import (
	"MyWebService/book/repository"
	"MyWebService/lib/server"
	"github.com/gin-gonic/gin"
	"net/http"
)

type DeleteBookItemHandler struct {
	bookRepository *repository.BookRepository
}

func NewDeleteBookItemHandler(bookRepository *repository.BookRepository) *DeleteBookItemHandler {
	return &DeleteBookItemHandler{bookRepository: bookRepository}
}

func (*DeleteBookItemHandler) GetConfig() server.HandlerConfig {
	return server.HandlerConfig{
		Path:   "/books/:id",
		Method: http.MethodDelete,
	}
}

func (h *DeleteBookItemHandler) HandleRequest(c *gin.Context) {
	id := c.Param("id")

	book := h.bookRepository.Get(id)
	if book == nil {
		c.Status(http.StatusNotFound)

		return
	}

	if err := h.bookRepository.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
