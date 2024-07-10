package handler

import (
	"MyWebService/book/repository"
	"MyWebService/lib/server"
	"github.com/gin-gonic/gin"
	"net/http"
)

type PutBookItemHandler struct {
	bookRepository *repository.BookRepository
}

func NewPutBookItemHandler(bookRepository *repository.BookRepository) *PutBookItemHandler {
	return &PutBookItemHandler{bookRepository: bookRepository}
}

func (*PutBookItemHandler) GetConfig() server.HandlerConfig {
	return server.HandlerConfig{
		Path:   "/books/:id",
		Method: http.MethodPut,
	}
}

type BookPutInput struct {
	Title  string   `json:"title"`
	Genres []string `json:"genres"`
}

func (h *PutBookItemHandler) HandleRequest(c *gin.Context) {
	id := c.Param("id")

	book := h.bookRepository.Get(id)
	if book == nil {
		c.Status(http.StatusNotFound)
		return
	}

	var input BookPutInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	book.Title = input.Title
	book.Genres = input.Genres

	if err := h.bookRepository.Update(book); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, book)
}
