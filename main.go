package main

import (
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"io/ioutil"

	"github.com/labstack/echo/v4"
)

func main() {
	route := echo.New()
	route.POST("/generate/signature", func(c echo.Context) error {
		jsonData, _ := ioutil.ReadAll(c.Request().Body)
		rawPayload := string(jsonData)
		merchantKey := c.Request().Header.Get("merchant-key")
		fmt.Println(rawPayload, merchantKey)
		h := sha512.New()
		h.Write([]byte(rawPayload + merchantKey))
		bs := h.Sum(nil)
		ourSignature := hex.EncodeToString(bs)

		data := make(map[string]string)
		data["signature"] = ourSignature
		c.JSON(200, data)
		return nil
	})
	route.Start(":8080")
}
