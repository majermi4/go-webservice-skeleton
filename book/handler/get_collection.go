package handler

import (
	"MyWebService/book/repository"
	"MyWebService/lib/server"
	"github.com/gin-gonic/gin"
	"net/http"
)

type GetBooksItemHandler struct {
	bookRepository *repository.BookRepository
}

func NewGetBooksItemHandler(bookRepository *repository.BookRepository) *GetBooksItemHandler {
	return &GetBooksItemHandler{bookRepository: bookRepository}
}

func (*GetBooksItemHandler) GetConfig() server.HandlerConfig {
	return server.HandlerConfig{
		Path:   "/books",
		Method: http.MethodGet,
	}
}

func (h *GetBooksItemHandler) HandleRequest(c *gin.Context) {
	books, err := h.bookRepository.Find()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})

		return
	}

	c.JSON(http.StatusOK, books)
}
