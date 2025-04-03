package http

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/mojtabamovahedi/url-shorter/internal/model"
	repo "github.com/mojtabamovahedi/url-shorter/internal/repository"
	"github.com/mojtabamovahedi/url-shorter/internal/service"
	"log"
	"net/http"
)

func Save(service service.LinkService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			ctx = c.Request.Context()
		)
		req := RequestBody{}
		if err := c.ShouldBindBodyWith(&req, binding.JSON); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "wrong body",
			})
			return
		}
		if len(req.Url) == 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "url required",
			})
			return
		}

		url, err := service.FindLinkByUrl(ctx, model.MainUrl(req.Url))
		if err != nil && !errors.Is(err, repo.LinkNotFoundErr) {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "database error",
			})
			return
		}

		if err == nil {
			c.JSON(http.StatusOK, gin.H{
				"message": "success",
				"short":   url.Short,
			})
			return
		}

		link := model.Link{
			Url: model.MainUrl(req.Url),
		}

		link.ID, err = service.CreateLink(ctx, link)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "database error",
			})
			return
		}

		err = link.MakeShort()
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "can't make short",
			})
			return
		}

		_, err = service.UpdateLink(ctx, link)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "database error",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "success",
			"short":   string(link.Short),
		})
	}
}

func Redirect(service service.LinkService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx = c.Request.Context()
		short := c.Param("short")
		if len(short) == 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "short url required",
			})
			return
		}

		link := model.Link{
			Short: model.ShortUrl(short),
		}

		if !link.ShortUrlValidation() {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "short url in not valid",
			})
			return
		}

		url, err := service.FindLinkByShortUrl(ctx, link.Short)
		if err != nil && !errors.Is(err, repo.LinkNotFoundErr) {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "database error",
			})
			return
		}

		if errors.Is(err, repo.LinkNotFoundErr) {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "short url not found",
			})
			return
		}

		c.Redirect(http.StatusFound, string(url.Url))
	}
}

type RequestBody struct {
	Url string `json:"url" binding:"required"`
}
