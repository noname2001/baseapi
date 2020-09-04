package v1

import (
	"baseapi/global"
	"baseapi/global/response"
	"baseapi/models"
	"baseapi/models/request"
	resp "baseapi/models/response"
	"fmt"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

func GetTags(c *gin.Context) {

	var pageInfo request.PageInfo
	pageInfo = request.PageInfo{
		Page:     1,
		PageSize: 20,
	}

	maps := make(map[string]interface{})
	name := c.Query("name")
	if name != "" {
		maps["name"] = name
	}
	err, tagList, total := models.GetTags(pageInfo, maps)

	if err != nil {
		response.FailWithMessage(fmt.Sprintf("获取失败：%v", err), c)
	} else {
		response.OkWithData(resp.PageResult{
			List:     tagList,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, c)
	}
}

func AddTag(c *gin.Context) {
	name := c.Query("name")
	state := com.StrTo(c.DefaultQuery("state", "0")).MustInt()

	valid := validation.Validation{}
	valid.Required(name, "name").Message("名称不能为空")
	valid.MaxSize(name, 100, "name").Message("名称最长为100字符")
	valid.Range(state, 0, 1, "state").Message("状态只允许0或1")
	if !valid.HasErrors() {
		if !models.ExistTagByName(name) {
			tag := models.AddTag(name, &state)
			response.OkWithData(tag, c)
		} else {
			response.FailWithMessage("同名标签已存在", c)
		}
	}
}

func UpdateTag(c *gin.Context) {
	var tag models.Tag
	tagId := c.Param("id")
	global.BA_DB.First(&tag, tagId)

	if tag.ID == 0 {
		response.FailNotExist(c)
		return
	}

	if err := c.ShouldBindJSON(&tag); err != nil {
		fmt.Println(err)
		response.FailWithMessage(err.Error(), c)
		return
	}
	global.BA_DB.Model(&tag).Updates(tag)
	response.OkWithData(tag, c)
}

func DeleteTag(c *gin.Context) {
	var tag models.Tag
	tagId := c.Param("id")
	global.BA_DB.First(&tag, tagId)
	if tag.ID == 0 {
		response.FailNotExist(c)
		return
	}
	global.BA_DB.Delete(&tag)
	response.Ok(c)
}
