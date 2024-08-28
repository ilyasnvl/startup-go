package handler

import (
	"net/http"
	"startup/campaign"
	"startup/helper"
	"startup/user"
	"strconv"

	"github.com/gin-gonic/gin"
)

//tangkap parameter di handler
//handler ke service
//service yg menentukan repository mana yg di panggil
//repository: getAll, getByUserID
//db

type campaignHandler struct {
	service campaign.Service
}

func NewCampaignHandler(service campaign.Service) *campaignHandler {
	return &campaignHandler{service}
}

func (h *campaignHandler) GetCampaigns(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Query("user_id"))

	campaigns, err := h.service.GetCampaigns(userID)
	if err != nil {
		response := helper.APIResponse("Error to get campaigns", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("List of campaigns", http.StatusOK, "success", campaign.Formatcampaigns(campaigns))
	c.JSON(http.StatusOK, response)
}

func (h *campaignHandler) GetCampaign(c *gin.Context) {
	//api/v1/campaigns/2
	//handler : mapping id yang di url ke struct input => service, call formatter
	//service : inputnya struct input => menangkap id di url, manggil repo
	//repository : get campaign by id

	var input campaign.GetCampaignDetailInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse("gagal to get detail of campaign", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	campaignDetail, err := h.service.GetCampaignByID(input)
	if err != nil {
		response := helper.APIResponse("Failed to get detail of campaign. err:"+err.Error(), http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("campaign detail", http.StatusOK, "success", campaign.FormatCampaignDetail(campaignDetail))
	c.JSON(http.StatusOK, response)
}

//tangkap paramter dari user ke input struct
//ambil current user dari jwt
//panggil service, paramternya input struct (dan juga buat slug)
//panggil repository, untuk simpan data campaign terbaru

func (h *campaignHandler) CreateCampaign(c *gin.Context) {
	var input campaign.CreateCampaignInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to create campaign", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)

	input.User = currentUser

	//panggil service
	newCampaign, err := h.service.CreateCampaign(input)
	if err != nil {
		response := helper.APIResponse("Failed to create campaign", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to create campaign", http.StatusOK, "success", campaign.FormatCampaign(newCampaign))
	c.JSON(http.StatusOK, response)
}

//user masukkan input
//handler
//mapping dari input ke input struct
//input dari user, dan juga input yang ada di uri (passing ke service)
//service(find campaign by id, tangkap parameter)
//repository update campaign

func (h *campaignHandler) UpdateCampaign(c *gin.Context) {
	var inputID campaign.GetCampaignDetailInput

	err := c.ShouldBindUri(&inputID)
	if err != nil {
		response := helper.APIResponse("Failed to update campaign", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var inputData campaign.CreateCampaignInput

	err = c.ShouldBindJSON(&inputData)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to update campaign", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)

	inputData.User = currentUser

	updateCampaign, err := h.service.UpdateCampaign(inputID, inputData)
	if err != nil {
		response := helper.APIResponse("Failed to update campaign", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return

	}

	response := helper.APIResponse("Success to update campaign", http.StatusOK, "success", campaign.FormatCampaign(updateCampaign))
	c.JSON(http.StatusOK, response)

}
