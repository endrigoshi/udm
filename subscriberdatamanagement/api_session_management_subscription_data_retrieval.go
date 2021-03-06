/*
 * Nudm_SDM
 *
 * Nudm Subscriber Data Management Service
 *
 * API version: 2.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package subscriberdatamanagement

import (
	"free5gc/lib/http_wrapper"
	"free5gc/src/udm/handler"
	udm_message "free5gc/src/udm/handler/message"
	"github.com/gin-gonic/gin"
)

// GetSmData - retrieve a UE's Session Management Subscription Data
func GetSmData(c *gin.Context) {

	req := http_wrapper.NewRequest(c.Request, nil)
	req.Params["supi"] = c.Param("supi")
	req.Query.Set("plmn-id", c.Query("plmn-id"))

	handleMsg := udm_message.NewHandlerMessage(udm_message.EventGetSmData, req)
	handler.SendMessage(handleMsg)

	rsp := <-handleMsg.ResponseChan
	HTTPResponse := rsp.HTTPResponse

	c.JSON(HTTPResponse.Status, HTTPResponse.Body)

}
