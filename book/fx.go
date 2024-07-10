package book

import (
	"MyWebService/book/handler"
	"MyWebService/book/repository"
	"MyWebService/lib/server"
	"go.uber.org/fx"
)

var Module = fx.Module("book",
	fx.Provide(
		server.AsHandler(handler.NewGetBookItemHandler),
		server.AsHandler(handler.NewGetBooksItemHandler),
		server.AsHandler(handler.NewPostBookItemHandler),
		server.AsHandler(handler.NewPutBookItemHandler),
		server.AsHandler(handler.NewDeleteBookItemHandler),
		repository.NewBookRepository,
	),
)
