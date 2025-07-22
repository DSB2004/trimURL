package controller

import (
	"url_shortener/pkg/models"
	"url_shortener/pkg/service"
	"url_shortener/pkg/utils"

	"github.com/gofiber/fiber/v2"
)

func HandleCreateShortURL(ctx *fiber.Ctx) error {
	var body models.URLRequestBody
	if err := utils.BodyParser(ctx.Body(), &body); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid JSON format",
		})
	}

	data,err:=body.Validate()

	if(err!=nil){
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	id,err:=service.CreateShortURL(data)
	if(err!=nil){
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message":"Failed to create url",
		})
	}

	endpoint := ctx.Protocol() + "://" + ctx.Hostname()+ "/redirect/" + id
	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message":"URL created",
		"id":id,
		"endpoint":endpoint,
	})

}
func HandleRedirect(ctx *fiber.Ctx)error {
	id:=ctx.Params("id","")
	method:=ctx.Method()
	if len(id)==0{
		return ctx.Status(fiber.StatusNotFound).SendString("Not Found")
	}

	url,err:=service.GetShortURL(id)

	if(err!=nil){
		return ctx.Status(fiber.StatusNotFound).SendString("Not Found")
	}

	found:=false

	for _, m := range url.Method {
		if m == method {
			found = true
			break
		}
	}

	if found==false{
		return ctx.Status(fiber.StatusMethodNotAllowed).SendString("Request Method not Allowed")
	}

	return ctx.Redirect(url.Original)
}