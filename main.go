package main

import (
	"fmt"
	"net/http"

	"github.com/Knetic/govaluate"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Calculation struct {
	ID         string `json:"id"`
	Expression string `json:"expression"`
	Result     string `json:"result"`
}

type CalculationRequest struct {
	Expression string `json:"expression"`
}

var calculations = []Calculation{} //nessesary to initilization like a slice

func calculateExpression(expression string) (string, error) {
	expr, err := govaluate.NewEvaluableExpression(expression) // 7 + 7 for example
	if err != nil {
		return "", err // mistake like (5 + a)
	}
	result, err := expr.Evaluate(nil)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", result), nil

}

func getCalculations(c echo.Context) error {
	return c.JSON(http.StatusOK, calculations)
}

func postCalculations(c echo.Context) error {
	var request CalculationRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadGateway, map[string]string{"error": "Invalid request"})
	}

	result, err := calculateExpression(request.Expression)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid expression"})
	}

	calc := Calculation{
		ID:         uuid.NewString(),
		Expression: request.Expression,
		Result:     result,
	}

	calculations = append(calculations, calc)
	return c.JSON(http.StatusCreated, calc)
}

func main() {
	e := echo.New()

	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	e.GET("/calculations", getCalculations)
	e.POST("/calculations", postCalculations)
	e.Start("localhost:8080")
}
