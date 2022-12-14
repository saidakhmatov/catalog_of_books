package handler

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/saidakhmatov/catalog_of_books/helper"

	"github.com/gin-gonic/gin"
	"github.com/saidakhmatov/catalog_of_books/models"
)

// @Summary     Create an author
// @ID          create_author
// @Description Create an author
// @Tags        Author
// @Router      /authors [POST]
// @Accept      json
// @Produce     json
// @Param       author body     models.CreateAuthor true "Author Body"
// @Success     201    {object} models.Response     "Success Response"
// @Response    400    {object} models.Response     "Bad Request Error"
func (h *handler) CreateAuthor(ctx *gin.Context) {
	var ar models.CreateAuthor
	var new_ar models.Author

	if err := ctx.ShouldBindJSON(&ar); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"response": models.Response{
				Error:   err.Error(),
				Data:    nil,
			},
		})
		return
	}

	new_id := helper.UUIDMaker()
	dt := time.Now()

	new_ar.ID = new_id
	new_ar.Firstname = ar.Firstname
	new_ar.Lastname = ar.Lastname
	new_ar.CreatedAt = dt
	new_ar.UpdatedAt = dt

	res, err := h.strg.AuthorRepo().CreateAuthor(new_ar)
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
			Message: "successfully created",
			Data:    res,
		},
	})
}

// @Summary  get all authors
// @ID       get_all_authors_id
// @Router   /authors [get]
// @Tags     Author
// @Produce  json
// @Param    search query    string          false "Search Query"
// @Param    limit  query    string          false "limit"
// @Param    offset query    string          false "offset"
// @Success  200    {object} models.Response "Success Response"
// @Response 404    {object} models.Response "Not found"
func (h *handler) GetAllAuthors(ctx *gin.Context) {
	var qP models.ApplicationQueryParamModel

	offset, offset_exists := ctx.GetQuery("offset")
	if offset_exists {
		res_offset, err := strconv.Atoi(offset)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"response": models.Response{
					Error:   err.Error(),
					Message: "Error while getting all authors",
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
					Message: "Error while getting all authors",
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

	res, err := h.strg.AuthorRepo().GetAllAuthors(qP)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"response": models.Response{
				Error:   err.Error(),
				Message: "Error while getting all authors",
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

// @Summary  get author by ID
// @ID       get_author_id
// @Tags     Author
// @Router   /authors/{id} [get]
// @Produce  json
// @Param    id  path     string          true "Author ID"
// @Success  200 {object} models.Response "Success Response"
// @Response 400 {object} models.Response "Bad Request Error"
// @Response 404    {object} models.Response     "Not found"
func (h *handler) GetAuthor(ctx *gin.Context) {
	id := ctx.Param("id")
	res, err := h.strg.AuthorRepo().GetAuthor(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"response": models.Response{
				Error:   err.Error(),
				Message: "Error while getting author",
				Data:    nil,
			},
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"response": models.Response{
			Error:   "false",
			Message: "success",
			Data:    res,
		},
	})
	return
}

// @Summary  Update Author
// @Tags     Author
// @ID       update_author_id
// @Router   /authors/{id} [put]
// @Accept   json
// @Produce  json
// @Param    id     path     string              true "Author ID"
// @Param    author body     models.UpdateAuthor true "Update Model"
// @Success  200    {object} models.Response     "Success Response"
// @Response 400    {object} models.Response     "Bad Request Error"
// @Response 404 {object} models.Response "Not found"
func (h *handler) UpdateAuthor(ctx *gin.Context) {
	var ar models.UpdateAuthor
	id := ctx.Param("id")

	if err := ctx.ShouldBindJSON(&ar); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"response": models.Response{
				Error:   err.Error(),
				Message: "Error",
				Data:    nil,
			},
		})
		return
	}

	res, err := h.strg.AuthorRepo().UpdateAuthor(ar, id)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"response": models.Response{
				Error:   err.Error(),
				Message: "Error while updating author",
				Data:    nil,
			},
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"response": models.Response{
			Error:   "false",
			Message: "successfully updated",
			Data:    res,
		},
	})
	return
}

// @Summary  delete an author by id
// @Tags     Author
// @Router   /authors/{id} [delete]
// @ID       delete_author_id
// @Param    id  path     string          true "Author ID"
// @Success  200 {object} models.Response "Success Response"
// @Response 400 {object} models.Response "Bad Request Error"
// @Response 404 {object} models.Response "Not found"
func (h *handler) DeleteAuthor(ctx *gin.Context) {
	id := ctx.Param("id")

	res, err := h.strg.AuthorRepo().DeleteAuthor(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"response": models.Response{
				Error:   err.Error(),
				Message: "Error while deleting author",
				Data:    nil,
			},
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"response": models.Response{
			Error:   "false",
			Message: "Successfully Deleted",
			Data:    res,
		},
	})
}
