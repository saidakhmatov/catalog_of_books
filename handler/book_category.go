package handler

import (
	"net/http"
	"strconv"
	"time"

	"github.com/saidakhmatov/catalog_of_books/helper"

	"github.com/gin-gonic/gin"
	"github.com/saidakhmatov/catalog_of_books/models"
)

// @Summary     Create a book category
// @ID          create_book_category_id
// @Description has no relation with others
// @Tags        BookCategory
// @Router      /book_category [POST]
// @Accept      json
// @Param       author body models.CreateBookCategory true "Author Body"
// @Produce     json
// @Success     201 {object} models.Response "Success Response"
// @Response    400 {object} models.Response "Bad Request Error"
func (h *handler) CreateBookCategory(ctx *gin.Context) {
	var bookCatCreate *models.CreateBookCategory
	var bookCat models.BookCategory

	if err := ctx.ShouldBindJSON(&bookCatCreate); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"response": models.Response{
				Error:   err.Error(),
				Data:    nil,
			},
		})
		return
	}

	dt := time.Now()

	bookCat.ID = helper.UUIDMaker()
	bookCat.CategoryName = bookCatCreate.CategoryName
	bookCat.CreatedAt = dt
	bookCat.UpdatedAt = dt

	res, err := h.strg.BookCategoryRepo().CreateBookCategory(bookCat)
	
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
			Message: "Successfully Created",
			Data:    res,
		},
	})
}

// @Summary  Get all book categories
// @ID       get_all_book_categories
// @Router   /book_category [GET]
// @Tags     BookCategory
// @Produce  json
// @Param    search query    string          false "Search Query"
// @Param    limit  query    string          false "limit"
// @Param    offset query    string          false "offset"
// @Success  200    {object} models.Response "Success Response"
// @Response 404    {object} models.Response "Bad Request Error"
func (h *handler) GetAllBookCategories(ctx *gin.Context) {
	var qP models.ApplicationQueryParamModel

	offset, offset_exists := ctx.GetQuery("offset")

	if offset_exists {
		res_offset, err := strconv.Atoi(offset)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"response": models.Response{
					Error:   err.Error(),
					Message: "Error while getting all categories",
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
					Message: "Error while getting all categories",
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

	bookCats, err := h.strg.BookCategoryRepo().GetAllBookCategories(qP)

	if err != nil {

		ctx.JSON(http.StatusBadRequest, gin.H{
			"response": models.Response{
				Error:   err.Error(),
				Message: "Error while getting all categories",
				Data:    nil,
			},
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{

		"response": models.Response{
			Error:   "false",
			Message: "Success",
			Data:    bookCats,
		},
	})
}

// @Summary  Get book category by ID
// @ID       get_book_category_id
// @Tags     BookCategory
// @Router   /book_category/{id} [get]
// @Produce  json
// @Param    id  path     string          true "Book Category ID"
// @Success  200 {object} models.Response "Success Response"
// @Response 400 {object} models.Response "Bad Request Error"
// @Response 404    {object} models.Response           "Not found"
func (h *handler) GetBookCategory(ctx *gin.Context) {
	id := ctx.Param("id")

	res, err := h.strg.BookCategoryRepo().GetBookCategory(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"response": models.Response{
				Error:   err.Error(),
				Message: "Error while getting book category",
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

// @Summary  Update book category
// @Tags     BookCategory
// @ID       update_author_id
// @Router   /book_category/{id} [put]
// @Accept   json
// @Produce  json
// @Param    id     path     string                    true "Book Category ID"
// @Param    author body     models.UpdateBookCategory true "Update Model"
// @Success  200    {object} models.Response           "Success Response"
// @Response 400    {object} models.Response           "Bad Request Error"
// @Response 404 {object} models.Response "Not found"
func (h *handler) UpdateBookCategory(ctx *gin.Context) {
	var bookCatModel *models.UpdateBookCategory
	id := ctx.Param("id")

	if err := ctx.ShouldBindJSON(&bookCatModel); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"response": models.Response{
				Error:   err.Error(),
				Data:    nil,
			},
		})
		return
	}

	res, err := h.strg.BookCategoryRepo().UpdateBookCategory(bookCatModel, id)
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

// @Summary  delete an book category by id
// @Tags     BookCategory
// @Router   /book_category/{id} [delete]
// @ID       delete_book_category_id
// @Param    id  path     string          true "Book Category ID"
// @Success  200 {object} models.Response "Success Response"
// @Response 400 {object} models.Response "Bad Request Error"
// @Response 404 {object} models.Response "Not found"
func (h *handler) DeleteBookCategory(ctx *gin.Context) {
	id := ctx.Param("id")

	res, err := h.strg.BookCategoryRepo().DeleteBookCategory(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"response": models.Response{
				Error:   err.Error(),
				Message: "Error while deleting category",
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
