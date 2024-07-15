package clients_manage

import (
	"github.com/gin-gonic/gin"
	"go-oauth/api/oauth_clients"
	"go-oauth/global"
	"strconv"
	"strings"
)

func List(c *gin.Context) {
	userId, _ := c.Get("user_id")
	if userId == "" {
		c.JSON(200, global.BackResp(400, "参数错误", nil))
		return
	}
	page := c.Query("page")
	pageSize := c.Query("pageSize")
	intPage := 0
	intPageSize := -1
	if page != "" && pageSize != "" {
		var err error
		intPage, err = strconv.Atoi(page)
		if err != nil {
			c.JSON(200, global.BackResp(400, "参数错误: 无效的 page 参数", nil))
			return
		}

		intPageSize, err = strconv.Atoi(pageSize)
		if err != nil {
			c.JSON(200, global.BackResp(400, "参数错误: 无效的 pageSize 参数", nil))
			return
		}
	}

	lists, counts := oauth_clients.FindClientsRowList("user_id", userId.(string), (intPage-1)*intPageSize, intPageSize)
	resp := make(map[string]interface{})
	resp["data"] = lists
	resp["total"] = counts
	c.JSON(200, global.BackResp(200, "查询成功", resp))
}
func Delete(c *gin.Context) {
	ids := c.Query("ids")
	if ids == "" {
		c.JSON(200, global.BackResp(400, "参数不能为空", nil))
		return
	}
	idList := strings.Split(ids, ",")
	err := oauth_clients.DeleteRow(idList)
	if err != nil {
		c.JSON(200, global.BackResp(400, "删除失败", nil))
		return
	}
	c.JSON(200, global.BackResp(200, "删除成功", nil))
}
