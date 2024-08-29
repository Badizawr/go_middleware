package endpoint

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

type Service interface {
	DaysLeft() int64
}

type Endpoint struct {
	svc Service
}

func New(svc Service) *Endpoint {
	return &Endpoint{
		svc: svc,
	}
}

func (e *Endpoint) Status(ctx echo.Context) error {
	d := e.svc.DaysLeft()

	s := fmt.Sprintf("Количество дней до Января 2025 осталось: %d", d)

	err := ctx.String(http.StatusOK, s)
	if err != nil {
		return err
	}
	return nil

}
