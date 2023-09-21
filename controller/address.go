package controller

import (
	"github.com/go-tron/example/authorize"
	"github.com/go-tron/example/entity"
	"github.com/go-tron/example/service"
	"github.com/go-tron/iris/baseContext"
	"github.com/go-tron/mysql"
	"github.com/kataras/iris/v12"
	"strconv"
)

var AddressController = addressController{}

type addressController struct {
}

func (c *addressController) Route(p iris.Party) {
	p.Post("/list", baseContext.ReturnHandler(c.List))
	p.Post("/get", baseContext.ReturnHandler(c.Get))
	p.Post("/create", baseContext.ReturnHandler(c.Create))
	p.Post("/update", baseContext.ReturnHandler(c.Update))
	p.Post("/remove", baseContext.ReturnHandler(c.Remove))
}

func (c *addressController) List(ctx *baseContext.Context) (interface{}, error) {
	var filters = map[string]interface{}{
		"userId": authorize.GetUserId(ctx),
	}
	return service.AddressService.FindPage(nil, filters)
}

func (c *addressController) Get(ctx *baseContext.Context) (interface{}, error) {
	var params mysql.IdReq
	if err := ctx.JSONReqBody(&params); err != nil {
		return nil, err
	}
	id, err := strconv.Atoi(params.Id.String())
	if err != nil {
		return nil, err
	}
	var filters = map[string]interface{}{
		"userId": authorize.GetUserId(ctx),
	}
	return service.AddressService.FindById(id, filters)
}

func (c *addressController) Create(ctx *baseContext.Context) (interface{}, error) {
	var params entity.Address
	if err := ctx.JSONReqBody(&params); err != nil {
		return nil, err
	}
	params.UserId = authorize.GetUserId(ctx)
	res, err := service.AddressService.Create(&params)
	if err != nil {
		return nil, err
	}
	return ctx.NewSuccess(res).WithMessage("保存成功"), nil
}

func (c *addressController) Update(ctx *baseContext.Context) (interface{}, error) {
	var params entity.Address
	if err := ctx.JSONReqBody(&params); err != nil {
		return nil, err
	}
	var filters = map[string]interface{}{
		"userId": authorize.GetUserId(ctx),
	}
	params.UserId = authorize.GetUserId(ctx)
	res, err := service.AddressService.Update(&params, filters)
	if err != nil {
		return nil, err
	}
	return ctx.NewSuccess(res).WithMessage("保存成功"), nil
}

func (c *addressController) Remove(ctx *baseContext.Context) (interface{}, error) {
	var params entity.Address
	if err := ctx.JSONReqBody(&params); err != nil {
		return nil, err
	}
	var filters = map[string]interface{}{
		"userId": authorize.GetUserId(ctx),
	}
	params.UserId = authorize.GetUserId(ctx)
	return ctx.NewSuccess().WithMessage("删除成功"), service.AddressService.Remove(&params, filters)
}
