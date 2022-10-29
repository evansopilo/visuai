package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

func (app App) Router() *fiber.App {

	r := fiber.New()

	v1 := r.Group("/v1/api").Use(requestid.New())
	{
		v1.Post("/upload", app.UploadFile)

		v1.Post("/posts", app.CreatePost)

		v1.Get("/posts", app.GetPost)

		v1.Get("/posts/tags", app.GetPostByTags)

		v1.Get("/posts/:post_id", app.GetPostByID)

		v1.Get("/users/:user_id/posts", app.GetPostByUserID)

		v1.Get("/category/:category/posts", app.GetPostByCategory)

		v1.Patch("/posts/:post_id", app.UpdatePost)

		v1.Delete("/posts/:post_id", app.DeletePostByID)

		v1.Delete("/users/:user_id/posts", app.DeletePostByUserID)
	}

	return r
}
