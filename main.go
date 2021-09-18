package main

import (
	"github.com/avtara/mvc-go/config"
	"github.com/avtara/mvc-go/routes"
	"github.com/avtara/mvc-go/utils"
	"github.com/go-playground/validator/v10"
)



func main() {
	config.InitDB()
	e := routes.New()
	e.Validator = &utils.CustomValidator{Validator: validator.New()}
	//e.HTTPErrorHandler = func(err error, c echo.Context) {
	//	report, ok := err.(*echo.HTTPError)
	//	if !ok {
	//		report = echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	//	}
	//
	//	if castedObject, ok := err.(validator.ValidationErrors); ok {
	//		for _, err := range castedObject {
	//			switch err.Tag() {
	//			case "required":
	//				report.Message = fmt.Sprintf("%s is required",
	//					err.Field())
	//			case "email":
	//				report.Message = fmt.Sprintf("%s is not valid email",
	//					err.Field())
	//			case "gte":
	//				report.Message = fmt.Sprintf("%s value must be greater than %s",
	//					err.Field(), err.Param())
	//			case "lte":
	//				report.Message = fmt.Sprintf("%s value must be lower than %s",
	//					err.Field(), err.Param())
	//			}
	//
	//			break
	//		}
	//	}
	//
	//	//c.Logger().Error(report)
	//	c.JSON(report.Code, report)
	//}
	e.Start(":8080")
}