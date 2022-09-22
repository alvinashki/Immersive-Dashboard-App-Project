package delivery

import (
	"fmt"
	"gp3/config"
	"gp3/features/logs"
	"gp3/middlewares"
	"gp3/utils/helper"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

type Delivery struct {
	logUsecase logs.UsecaseInterface
}

func New(e *echo.Echo, usecase logs.UsecaseInterface) {
	handler := &Delivery{
		logUsecase: usecase,
	}
	e.POST("/feedback", handler.PostLogs, middlewares.JWTMiddleware())
	e.GET("/feedback", handler.GetAllFeedback, middlewares.JWTMiddleware())
}

func (deliv *Delivery) PostLogs(c echo.Context) error {
	idToken, _ := middlewares.ExtractToken(c)

	var dataRequest LogsRequest
	errBind := c.Bind(&dataRequest)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponseHelper("error binding data"))
	}
	dataRequest.UserId = uint(idToken)

	//upload file Image
	imageData, imageInfo, imageErr := c.Request().FormFile("file")
	if imageErr == http.ErrMissingFile || imageErr != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponseHelper("Failed to get image"))
	}

	imageExtension, err_image_extension := helper.CheckFileExtension(imageInfo.Filename, config.ContentImage)
	if err_image_extension != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponseHelper("Image extension error"))
	}

	// check image size
	err_image_size := helper.CheckFileSize(imageInfo.Size, config.ContentImage)
	if err_image_size != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponseHelper("Image size error"))
	}

	// memberikan nama file
	imageName := strconv.Itoa(idToken) + "_" + time.Now().Format("2006-01-02 15:04:05") + "." + imageExtension

	image, errUploadImg := helper.UploadFileToS3(config.EventImages, imageName, config.ContentImage, imageData)

	if errUploadImg != nil {
		fmt.Println(errUploadImg)
		return c.JSON(http.StatusInternalServerError, helper.FailedResponseHelper("failed to upload file"))
	}
	dataRequest.File = image

	// fmt.Println("error =", dataRequest)
	row, err := deliv.logUsecase.CreateData(toCore(dataRequest))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponseHelper("error insert data"))
	}
	if row != 1 {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponseHelper("error insert data"))
	}
	return c.JSON(http.StatusCreated, helper.SuccessResponseHelper("success insert data"))
}

func (deliv *Delivery) GetAllFeedback(c echo.Context) error {
	mentee_id, err := strconv.Atoi(c.QueryParam("mentee_id"))

	if err != nil && mentee_id != 0 {
		return c.JSON(http.StatusBadRequest, helper.FailedResponseHelper("fail converse mentee_id param"))
	}

	dataFeedback, err := deliv.logUsecase.SelectFeedback(mentee_id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponseHelper(err.Error()))
	}

	return c.JSON(http.StatusOK, helper.SuccessDataResponseHelper("get all data mentees success", fromCoreList(dataFeedback)))
}

/*
func (deliv *Delivery) GetAllUser(c echo.Context) error {
	result, err := deliv.logUsecase.GetAllUser()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponseHelper("Failed to get data"))
	}
	return c.JSON(http.StatusOK, helper.SuccessDataResponseHelper("Succes get data", fromCoreList(result)))
}

func (deliv *Delivery) PutUser(c echo.Context) error {
	id := c.Param("id")
	idConv, errConv := strconv.Atoi(id)
	idToken, role := middlewares.ExtractToken(c) //coba

	if errConv != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponseHelper("param must be a number"))
	}

	if idConv == idToken || role == "admin" {
		var dataUpdate UserRequest
		errBind := c.Bind(&dataUpdate)
		if errBind != nil {
			return c.JSON(http.StatusBadRequest, helper.FailedResponseHelper("error binding data"))
		}

		row, err := deliv.logUsecase.PutUser(toCore(dataUpdate), idConv)

		if err != nil || row == 0 {
			return c.JSON(http.StatusInternalServerError, helper.FailedResponseHelper("failed to update data"))
		}

		return c.JSON(http.StatusBadRequest, helper.SuccessResponseHelper("succes update data"))
	}

	return c.JSON(http.StatusBadRequest, helper.FailedResponseHelper("you dont have acces"))

}

func (deliv *Delivery) DeleteAkun(c echo.Context) error {
	id := c.Param("id")
	idConv, errConv := strconv.Atoi(id)
	_, role := middlewares.ExtractToken(c)

	if role != "admin" {
		return c.JSON(http.StatusBadRequest, helper.FailedResponseHelper("you dont have acces"))
	}

	if errConv != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponseHelper("param must be a number"))
	}

	row, err := deliv.logUsecase.DeleteAkun(idConv)
	if err != nil || row == 0 {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponseHelper("failed to delete account"))
	}
	return c.JSON(http.StatusOK, helper.SuccessResponseHelper("succes delete account"))
}

func (deliv *Delivery) SelectUserId(c echo.Context) error {
	id := c.Param("id")
	idConv, errConv := strconv.Atoi(id)

	if errConv != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponseHelper("param must be a number"))
	}

	result, err := deliv.logUsecase.GetUserId(idConv)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponseHelper("failed to get data"))
	}
	return c.JSON(http.StatusOK, helper.SuccessDataResponseHelper("succes get data", fromCore(result)))

}
*/
