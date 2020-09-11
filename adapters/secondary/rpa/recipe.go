package rpa

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/juniorrosul/delivery-much-challenge/adapters/secondary"
	"github.com/juniorrosul/delivery-much-challenge/application/recipepuppy"
)

// RecipePuppy - Interface
type RecipePuppy interface {
	getRecipes(params string) (*recipepuppy.ResponseModel, error)
}

// RecipePuppyIntegration - Base struct for API integration
type RecipePuppyIntegration struct {
	conn secondary.HTTPConnector
}

// NewRecipePuppyIntegration - Integration for API
func NewRecipePuppyIntegration(httpConnector secondary.HTTPConnector) *RecipePuppyIntegration {
	return &RecipePuppyIntegration{
		conn: httpConnector,
	}
}

// GetRecipes - Get recipes from external API
func (rpi *RecipePuppyIntegration) GetRecipes(request *recipepuppy.RequestModel) (*recipepuppy.ResponseModel, error) {
	var rpr *recipepuppy.ResponseModel
	params := strings.Join(request.Ingredients, ",")
	res, err := rpi.conn.DoGet(fmt.Sprintf("/?i=%s", params))
	if err != nil || params == "" {
		return nil, err
	}
	body, err := secondary.GetBodyResponse(res)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &rpr); err != nil {
		return nil, err
	}
	return rpr, nil
}
