package handler

import (
	"MyWebService/book/data"
	"MyWebService/book/repository"
	"MyWebService/lib/server"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type PostBookItemHandler struct {
	bookRepository *repository.BookRepository
}

func NewPostBookItemHandler(bookRepository *repository.BookRepository) *PostBookItemHandler {
	return &PostBookItemHandler{bookRepository: bookRepository}
}

func (*PostBookItemHandler) GetConfig() server.HandlerConfig {
	return server.HandlerConfig{
		Path:   "/books",
		Method: http.MethodPost,
	}
}

type BookPostInput struct {
	Title  string   `json:"title"`
	Genres []string `json:"genres"`
}

func (h *PostBookItemHandler) HandleRequest(c *gin.Context) {
	var input BookPostInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	book := &data.Book{
		Title:  input.Title,
		Genres: input.Genres,
	}

	err := h.bookRepository.Create(book)
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, book)
}
