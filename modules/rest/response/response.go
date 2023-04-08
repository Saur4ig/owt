package response

import (
	"encoding/json"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

const (
	internal = "Internal Server Error"
)

func Failed(c *fiber.Ctx, code int, msg string) error {
	resp := Response{
		Err: &Error{
			Status: code,
			Msg:    msg,
		},
	}
	return sendResponse(c, resp)
}

func Success(c *fiber.Ctx, body any) error {
	resp := Response{
		Success: &Successful{Data: body},
	}

	return sendResponse(c, resp)
}

func sendResponse(c *fiber.Ctx, resp Response) error {
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString(internal)
	}

	return c.Status(http.StatusOK).SendString(string(jsonResp))
}
