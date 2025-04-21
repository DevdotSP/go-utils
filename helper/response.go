package helper

import (
	"github.com/DevdotSP/go-utils/model"

	"github.com/DevdotSP/go-utils/utils"
	"github.com/gofiber/fiber/v3"
)

func JSONResponse(c fiber.Ctx, retCode, retMessage string) error {
	return c.JSON(model.Response{
		ResponseTime: utils.GetResponseTime(c),
		Device:       string(c.RequestCtx().UserAgent()),
		RetCode:      retCode,
		Message:      retMessage,
	})
}

func JSONResponseWithData(c fiber.Ctx, retCode, retMessage string, data interface{}) error {
	return c.JSON(model.Response{
		ResponseTime: utils.GetResponseTime(c),
		Device:       string(c.RequestCtx().UserAgent()),
		RetCode:      retCode,
		Message:      retMessage,
		Data:         data,
	})
}

func JSONResponseWithDataAndToken(c fiber.Ctx, retCode, retMessage string, data interface{}, jwttoken string) error {
	return c.JSON(model.Response{
		ResponseTime: utils.GetResponseTime(c),
		Device:       string(c.RequestCtx().UserAgent()),
		RetCode:      retCode,
		Message:      retMessage,
		Data:         data,
		JwtToken:     jwttoken,
	})
}

func JSONResponseWithDataPageDetails(c fiber.Ctx, retCode, retMessage string, data interface{}, pageDetails *model.PageDetails) error {
	return c.JSON(model.ResponsePageDetails{
		ResponseTime: utils.GetResponseTime(c),
		Device:       string(c.RequestCtx().UserAgent()),
		RetCode:      retCode,
		Message:      retMessage,
		Data:         data,
		PageDetails:  *pageDetails,
	})
}

func JSONResponseWithError(c fiber.Ctx, retCode, retMessage string, err error) error {
	return c.JSON(model.Response{
		ResponseTime: utils.GetResponseTime(c),
		Device:       string(c.RequestCtx().UserAgent()),
		RetCode:      retCode,
		Message:      retMessage,
		Error:        err.Error(),
	})
}

func JSONResponseWithValidationData(c fiber.Ctx, retCode, retMessage string, data interface{}) error {
	return c.JSON(model.Response{
		ResponseTime: utils.GetResponseTime(c),
		Device:       string(c.RequestCtx().UserAgent()),
		RetCode:      retCode,
		Message:      retMessage,
		// Use a clearer field name for validation data, if needed
		ValidationErrors: data, // Or another field name like `Details` depending on your intent
	})
}

func JSONResponseWithValidation(c fiber.Ctx, retCode string, retMessage []string) error {
	return c.JSON(model.Response{
		ResponseTime: utils.GetResponseTime(c),
		Device:       string(c.RequestCtx().UserAgent()),
		RetCode:      retCode,
		Error:        retMessage,
	})
}
