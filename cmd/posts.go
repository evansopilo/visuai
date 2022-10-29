package main

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/evansopilo/visuai/pkg/data"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

func (app App) CreatePost(c *fiber.Ctx) error {

	ctx, cancel := context.WithTimeout(c.Context(), time.Second*3)
	defer cancel()

	var post data.Post

	if err := c.BodyParser(&post); err != nil {
		app.Logger.Error(err.Error(), map[string]interface{}{"requestid": c.Locals("requestid")})
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "error",
		})
	}

	if err := app.Models.Post.Create(ctx, &post); err != nil {
		app.Logger.Error(err.Error(), map[string]interface{}{"requestid": c.Locals("requestid")})
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "error",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status": "success",
		"data": map[string]string{
			"id": post.ID,
		},
	})
}

func (app App) GetPostByID(c *fiber.Ctx) error {

	ctx, cancel := context.WithTimeout(c.Context(), time.Second*3)
	defer cancel()

	post, err := app.Models.Post.GetByID(ctx, c.Params("post_id"))
	if err != nil {
		switch {
		case errors.Is(err, mongo.ErrNoDocuments):
			app.Logger.Error(err.Error(), map[string]interface{}{"requestid": c.Locals("requestid")})
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"status":  "error",
				"message": fmt.Sprintf("document with id: %v not found", c.Params("post_id")),
			})
		default:
			app.Logger.Error(err.Error(), map[string]interface{}{"requestid": c.Locals("requestid")})
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"status": "error",
			})
		}
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"data":   post,
	})
}

func (app App) GetPostByUserID(c *fiber.Ctx) error {

	ctx, cancel := context.WithTimeout(c.Context(), time.Second*3)
	defer cancel()

	page_num, _ := strconv.Atoi(c.Query("page_num", "1"))

	page_size, _ := strconv.Atoi(c.Query("page_size", "10"))

	skips := page_size * (page_num - 1)

	post, err := app.Models.Post.GetByUserID(ctx, c.Params("user_id"), int64(skips), int64(page_size))
	if err != nil {
		app.Logger.Error(err.Error(), map[string]interface{}{"requestid": c.Locals("requestid")})
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "error",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"data":   post,
	})
}

func (app App) GetPostByCategory(c *fiber.Ctx) error {

	ctx, cancel := context.WithTimeout(c.Context(), time.Second*3)
	defer cancel()

	page_num, _ := strconv.Atoi(c.Query("page_num", "1"))

	page_size, _ := strconv.Atoi(c.Query("page_size", "10"))

	skips := page_size * (page_num - 1)

	post, err := app.Models.Post.GetByCategory(ctx, c.Params("category"), int64(skips), int64(page_size))
	if err != nil {
		app.Logger.Error(err.Error(), map[string]interface{}{"requestid": c.Locals("requestid")})
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "error",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"data":   post,
	})
}

func (app App) GetPostByTags(c *fiber.Ctx) error {

	ctx, cancel := context.WithTimeout(c.Context(), time.Second*3)
	defer cancel()

	page_num, _ := strconv.Atoi(c.Query("page_num", "1"))

	page_size, _ := strconv.Atoi(c.Query("page_size", "10"))

	skips := page_size * (page_num - 1)

	post, err := app.Models.Post.GetByTags(ctx, strings.Split(c.Query("v"), ","), int64(skips), int64(page_size))
	if err != nil {
		app.Logger.Error(err.Error(), map[string]interface{}{"requestid": c.Locals("requestid")})
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "error",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"data":   post,
	})
}

func (app App) GetPost(c *fiber.Ctx) error {

	ctx, cancel := context.WithTimeout(c.Context(), time.Second*3)
	defer cancel()

	page_num, _ := strconv.Atoi(c.Query("page_num", "1"))

	page_size, _ := strconv.Atoi(c.Query("page_size", "10"))

	skips := page_size * (page_num - 1)

	post, err := app.Models.Post.Get(ctx, int64(skips), int64(page_size))
	if err != nil {
		app.Logger.Error(err.Error(), map[string]interface{}{"requestid": c.Locals("requestid")})
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "error",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"data":   post,
	})
}

func (app App) UpdatePost(c *fiber.Ctx) error {

	ctx, cancel := context.WithTimeout(c.Context(), time.Second*3)
	defer cancel()

	var post data.Post

	if err := c.BodyParser(&post); err != nil {
		app.Logger.Error(err.Error(), map[string]interface{}{"requestid": c.Locals("requestid")})
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "error",
		})
	}

	if err := app.Models.Post.UpdateByID(ctx, c.Params("post_id"), &post); err != nil {
		switch {
		case errors.Is(err, data.ErrNoDocument):
			app.Logger.Error(err.Error(), map[string]interface{}{"requestid": c.Locals("requestid")})
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"status":  "error",
				"message": fmt.Sprintf("document with id: %v not found", c.Params("post_id")),
			})
		default:
			app.Logger.Error(err.Error(), map[string]interface{}{"requestid": c.Locals("requestid")})
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"status": "error",
			})
		}
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
	})
}

func (app App) DeletePostByID(c *fiber.Ctx) error {

	ctx, cancel := context.WithTimeout(c.Context(), time.Second*3)
	defer cancel()

	if err := app.Models.Post.DeleteByID(ctx, c.Params("post_id")); err != nil {
		switch {
		case errors.Is(err, data.ErrNoDocument):
			app.Logger.Error(err.Error(), map[string]interface{}{"requestid": c.Locals("requestid")})
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"status":  "error",
				"message": fmt.Sprintf("document with id: %v not found", c.Params("post_id")),
			})
		default:
			app.Logger.Error(err.Error(), map[string]interface{}{"requestid": c.Locals("requestid")})
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"status": "error",
			})
		}
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
	})
}

func (app App) DeletePostByUserID(c *fiber.Ctx) error {

	ctx, cancel := context.WithTimeout(c.Context(), time.Second*3)
	defer cancel()

	if err := app.Models.Post.DeleteByID(ctx, c.Params("user_id")); err != nil {
		switch {
		case errors.Is(err, data.ErrNoDocument):
			app.Logger.Error(err.Error(), map[string]interface{}{"requestid": c.Locals("requestid")})
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"status":  "error",
				"message": fmt.Sprintf("document with user id: %v not found", c.Params("user_id")),
			})
		default:
			app.Logger.Error(err.Error(), map[string]interface{}{"requestid": c.Locals("requestid")})
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"status": "error",
			})
		}
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
	})
}
