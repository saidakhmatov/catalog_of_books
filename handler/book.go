package handler

import (
	"net/http"
	"strconv"
	"time"

	"github.com/saidakhmatov/catalog_of_books/helper"

	"github.com/gin-gonic/gin"
	"github.com/saidakhmatov/catalog_of_books/models"
)

// @Summary     Create a book
// @ID          create_book_id
// @Description has no relation with others
// @Tags        Book
// @Router      /books [post]
// @Accept      json
// @Param       author body models.CreateBook true "Book Body"
// @Produce     json
// @Success     201 {object} models.Response "Success Response"
// @Response    400 {object} models.Response "Bad Request Error"
func (h *handler) CreateBook(ctx *gin.Context) {
	
	var bookCreate models.CreateBook
	
	var book models.Book

	if err := ctx.ShouldBindJSON(&bookCreate); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"response": models.Response{
				Error:   err.Error(),
				Data:    nil,
			},
		})
		return
	}

	dt := time.Now()
	new_id := helper.UUIDMaker()

	book.ID = new_id
	book.CreatedAt = dt
	book.UpdatedAt = dt
	book.CategoryID = bookCreate.CategoryID
	book.AuthorID = bookCreate.AuthorID
	book.BookName = bookCreate.BookName

	res, err := h.strg.BookRepo().CreateBook(book)
	
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"response": models.Response{
				Error:   err.Error(),
				Data:    nil,
			},
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"response": models.Response{
			Error:   "false",
			Message: "Successfully created",
			Data:    res,
		},
	})
}

// @Summary  Get all books
// @ID       get_all_books_id
// @Router   /books [get]
// @Tags     Book
// @Produce  json
// @Param    search query    string          false "search"
// @Param    limit  query    string          false "limit"
// @Param    offset query    string          false "offset"
// @Success  200    {object} models.Response "Success Response"
// @Response 404    {object} models.Response "Some bad request"
func (h *handler) GetAllBooks(ctx *gin.Context) {
	
	var qP models.ApplicationQueryParamModel

	offset, offset_exists := ctx.GetQuery("offset")
	
	if offset_exists {
		res_offset, err := strconv.Atoi(offset)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"response": models.Response{
					Error:   err.Error(),
					Message: "Error while getting all books",
					Data:    nil,
				},
			})
			return
		}

		qP.Offset = res_offset
	}

	limit, limit_exists := ctx.GetQuery("limit")
	
	if limit_exists {
		res_limit, err := strconv.Atoi(limit)	
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"response": models.Response{
					Error:   err.Error(),
					Message: "Error while getting all books",
					Data:    nil,
				},
			})
			return
		}

		qP.Limit = res_limit
	}

	search, search_exists := ctx.GetQuery("search")
	
	if search_exists {
		qP.Search = search
	}

	books, err := h.strg.BookRepo().GetAllBooks(qP)
	
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"response": models.Response{
				Error:   err.Error(),
				Message: "Error while getting all books",
				Data:    nil,
			},
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"response": models.Response{
			Error:   "false",
			Message: "Success",
			Data:    books,
		},
	})
}

// @Summary  Get book by ID
// @ID       get_book_id
// @Tags     Book
// @Router   /books/{id} [get]
// @Produce  json
// @Param    id  path     string          true "Book Category ID"
// @Success  200 {object} models.Response "Success Response"
// @Response 400 {object} models.Response "Bad Request Error"
// @Response 404  {object} models.Response   "Not found"
func (h *handler) GetBook(ctx *gin.Context) {
	
	id := ctx.Param("id")

	res, err := h.strg.BookRepo().GetBook(id)
	
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"response": models.Response{
				Error:   err.Error(),
				Message: "Error while author getting a book",
				Data:    nil,
			},
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"response": models.Response{
			Error:   "false",
			Message: "Success",
			Data:    res,
		},
	})
}

// @Summary  Update book
// @Tags     Book
// @ID       update_book_id
// @Router   /books/{id} [put]
// @Accept   json
// @Produce  json
// @Param    id   path     string            true "Book ID"
// @Param    book body     models.UpdateBook true "Update Model"
// @Success  200  {object} models.Response   "SUccess REsponse"
// @Response 400  {object} models.Response   "Bad Request Error"
// @Response 404 {object} models.Response "Not found"
func (h *handler) UpdateBook(ctx *gin.Context) {
	
	var bookModel models.UpdateBook
	
	id := ctx.Param("id")

	if err := ctx.ShouldBindJSON(&bookModel); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"response": models.Response{
				Error:   err.Error(),
				Data:    nil,
			},
		})
		return
	}

	res, err := h.strg.BookRepo().UpdateBook(bookModel, id)
	
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"response": models.Response{
				Error:   err.Error(),
				Data:    nil,
			},
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"response": models.Response{
			Error:   "false",
			Message: "Success",
			Data:    res,
		},
	})
}

// @Summary  delete an book by id
// @Tags     Book
// @Router   /books/{id} [delete]
// @ID       delete_book_id
// @Param    id  path     string          true "Book ID"
// @Success  200 {object} models.Response "Success Response"
// @Response 400 {object} models.Response "Bad Request Error"
// @Response 404 {object} models.Response "Not found"
func (h *handler) DeleteBook(ctx *gin.Context) {
	
	id := ctx.Param("id")

	res, err := h.strg.BookRepo().DeleteBook(id)
	
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"response": models.Response{
				Error:   err.Error(),
				Message: "Error while deleting book",
				Data:    nil,
			},
		})
		return
	}

	
	ctx.JSON(http.StatusOK, gin.H{
		"response": models.Response{
			Error:   "false",
			Message: "Success",
			Data:    res,
		},
	})
}
