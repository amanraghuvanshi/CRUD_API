package controllers

import (
	"assignTele/data/request"
	"assignTele/data/response"
	"assignTele/helper"
	"assignTele/service"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type BlogControllers struct {
	BlogsService service.BlogService
}

func NewBlogsControl(blogservice service.BlogService) *BlogControllers {
	return &BlogControllers{BlogsService: blogservice}
}

// Creation
func (controller *BlogControllers) Create(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	blogsCreate := request.BlogCreate{}
	helper.GetReqBody(r, &blogsCreate)

	controller.BlogsService.Create(r.Context(), blogsCreate)
	webResp := response.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   nil,
	}
	helper.WriteResponseBody(w, webResp)
}

// Updation
func (controller *BlogControllers) Update(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	blogUpdate := request.UpdateRequest{}
	helper.GetReqBody(r, &blogUpdate)

	controller.BlogsService.Update(r.Context(), blogUpdate)
	blogID := p.ByName("blogId")
	id, err := strconv.Atoi(blogID)
	helper.PanicIfError(err)
	blogUpdate.ID = id
	webResp := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   nil,
	}
	helper.WriteResponseBody(w, webResp)
}

// Deletion
func (controller *BlogControllers) Delete(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	blogID := p.ByName("blogId")
	id, err := strconv.Atoi(blogID)
	helper.PanicIfError(err)
	controller.BlogsService.Delete(r.Context(), id)

	webResp := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   nil,
	}
	helper.WriteResponseBody(w, webResp)
}

// Get all the data
func (controller *BlogControllers) Find(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	result := controller.BlogsService.FindAll(r.Context())
	webResp := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   result,
	}
	helper.WriteResponseBody(w, webResp)
}

// Getting blogs by ID
func (controller *BlogControllers) FindbyID(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	blogID := p.ByName("blogId")
	id, err := strconv.Atoi(blogID)
	helper.PanicIfError(err)

	result := controller.BlogsService.FindbyID(r.Context(), id)
	webResp := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   result,
	}
	helper.WriteResponseBody(w, webResp)
}
