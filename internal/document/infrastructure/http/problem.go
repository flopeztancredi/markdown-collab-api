package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProblemDetails struct {
	Type     string `json:"type"`
	Title    string `json:"title"`
	Status   int    `json:"status"`
	Detail   string `json:"detail,omitempty"`
	Instance string `json:"instance,omitempty"`
}

func writeProblem(c *gin.Context, status int, detail string) {
	title := http.StatusText(status)
	problem := ProblemDetails{
		Type:     "about:blank",
		Title:    title,
		Status:   status,
		Detail:   detail,
		Instance: c.Request.URL.Path,
	}
	if problem.Detail == "" {
		problem.Detail = title
	}

	c.Header("Content-Type", "application/problem+json")
	c.JSON(status, problem)
}
