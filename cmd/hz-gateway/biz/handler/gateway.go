// Code generated by hertz generator.

package handler

import (
	"TinyMarket/cmd/hz-gateway/biz/errors"
	"TinyMarket/cmd/hz-gateway/biz/types"
	"TinyMarket/kitex_gen/common"
	"bytes"
	"context"
	"fmt"
	"net/http"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/cloudwego/kitex/client/genericclient"
	"github.com/cloudwego/kitex/pkg/generic"
	"github.com/cloudwego/kitex/pkg/kerrors"
)

// Ping .
func Ping(ctx context.Context, c *app.RequestContext) {
	c.JSON(consts.StatusOK, utils.H{
		"message": "pong",
	})
}
type requiredParams struct {
	Method     string `form:"method,required" json:"method"`
	MerchantId string `form:"merchant_id,required" json:"merchant_id"`
	BizParams  string `form:"biz_params,required" json:"biz_params"`
}

var SvcMap = make(map[string]genericclient.Client)

// Gateway handle the request with the query path of prefix `/gateway`.
func Gateway(ctx context.Context, c *app.RequestContext) {
	svcName := c.Param("svc")
	cli, ok := SvcMap[svcName]
	if !ok {
		c.JSON(http.StatusOK, errors.New(common.Err_BadRequest))
		return
	}
	var params requiredParams
	if err := c.BindAndValidate(&params); err != nil {
		hlog.Error(err)
		c.JSON(http.StatusOK, errors.New(common.Err_ServerMethodNotFound))
		return
	}

	req, err := http.NewRequest(http.MethodPost, "", bytes.NewBuffer([]byte(params.BizParams)))
	if err != nil {
		hlog.Warnf("new http request failed: %v", err)
		c.JSON(http.StatusOK, errors.New(common.Err_RequestServerFail))
		return
	}
	req.URL.Path = fmt.Sprintf("/%s/%s", svcName, params.Method)

	customReq, err := generic.FromHTTPRequest(req)
	if err != nil {
		hlog.Errorf("convert request failed: %v", err)
		c.JSON(http.StatusOK, errors.New(common.Err_ServerHandleFail))
		return
	}
	resp, err := cli.GenericCall(ctx, "", customReq)
	respMap := make(map[string]interface{})
	if err != nil {
		hlog.Errorf("GenericCall err:%v", err)
		bizErr, ok := kerrors.FromBizStatusError(err)
		if !ok {
			c.JSON(http.StatusOK, errors.New(common.Err_ServerHandleFail))
			return
		}
		respMap[types.ResponseErrCode] = bizErr.BizStatusCode()
		respMap[types.ResponseErrMessage] = bizErr.BizMessage()
		c.JSON(http.StatusOK, respMap)
		return
	}
	realResp, ok := resp.(*generic.HTTPResponse)
	if !ok {
		c.JSON(http.StatusOK, errors.New(common.Err_ServerHandleFail))
		return
	}
	realResp.Body[types.ResponseErrCode] = 0
	realResp.Body[types.ResponseErrMessage] = "ok"
	c.JSON(http.StatusOK, realResp.Body)
}