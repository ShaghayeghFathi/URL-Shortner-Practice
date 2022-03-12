package http

import (
	"fmt"
	"net/http"

	"github.com/ShaghyeghFathi/URL-Shortner-Practice/internal/db/redis"
	"github.com/labstack/echo/v4"
)

type Handler struct{
	Redisdb redis.Redis
}

type addURLRequest struct{
	Url string `json:"url"`
	ShortnedURL string `json:"shortendurl"`
}

func(h *Handler) addURL(ctx echo.Context) error {
	req:= new(addURLRequest)
	if err := ctx.Bind(req); err != nil {
		return err
	}
	fmt.Println(req.Url, req.ShortnedURL)
	err:= h.Redisdb.Set(ctx.Request().Context() ,req.Url,req.ShortnedURL); if err!= nil{
		return echo.NewHTTPError(http.StatusInternalServerError , err.Error())
	}
	return ctx.NoContent(http.StatusNoContent)
}

func (h Handler)Register(c *echo.Echo){
	c.POST("/link", h.addURL)
	// c.GET("/link" , h.getOriginalURL)
}