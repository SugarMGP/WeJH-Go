package adminController

import (
	"github.com/gin-gonic/gin"
	"wejh-go/app/apiException"
	"wejh-go/app/models"
	"wejh-go/app/services/announcementServices"
	"wejh-go/app/utils"
)

type createAnnouncementForm struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content"`
}
type updateAnnouncementForm struct {
	ID      int    `json:"id" binding:"required"`
	Title   string `json:"title" binding:"required"`
	Content string `json:"content"`
}
type deleteAnnouncementForm struct {
	ID int `json:"id" binding:"required"`
}

func CreateAnnouncement(c *gin.Context) {
	var postForm createAnnouncementForm
	err := c.ShouldBindJSON(&postForm)
	if err != nil {
		apiException.AbortWithException(c, apiException.ParamError, err)
		return
	}

	err = announcementServices.CreateAnnouncement(models.Announcement{Title: postForm.Title, Content: postForm.Content})
	if err != nil {
		apiException.AbortWithError(c, err)
		return
	}
	utils.JsonSuccessResponse(c, nil)
}
func UpdateAnnouncement(c *gin.Context) {
	var postForm updateAnnouncementForm
	err := c.ShouldBindJSON(&postForm)
	if err != nil {
		apiException.AbortWithException(c, apiException.ParamError, err)
		return
	}

	err = announcementServices.UpdateAnnouncement(postForm.ID, models.Announcement{
		Title:   postForm.Title,
		Content: postForm.Content,
	},
	)
	if err != nil {
		apiException.AbortWithException(c, apiException.ParamError, err)
		return
	}
	utils.JsonSuccessResponse(c, nil)

}
func DeleteAnnouncement(c *gin.Context) {
	var postForm deleteAnnouncementForm
	err := c.ShouldBindJSON(&postForm)
	if err != nil {
		apiException.AbortWithException(c, apiException.ParamError, err)
		return
	}

	err = announcementServices.DeleteAnnouncement(postForm.ID)
	if err != nil {
		apiException.AbortWithException(c, apiException.ParamError, err)
		return
	}
	utils.JsonSuccessResponse(c, nil)
}
