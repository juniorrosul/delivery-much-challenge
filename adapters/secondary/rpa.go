package secondary

import (
	"fmt"

	"github.com/juniorrosul/delivery-much-challenge/application/recipe"
)

type RecipePuppy interface {
	getRecipes(params string) (*recipe.Model, error)
}

type RecipePuppyIntegration struct {
	connector HttpConnector
}

func NewRecipePuppyIntegration(httpConnector HttpConnector) *RecipePuppyIntegration {
	return &RecipePuppyIntegration{
		connector: httpConnector,
	}
}

func (rpi *RecipePuppyIntegration) getRecipes(params string) (*recipe.Model, error) {
	res, err := rpi.connector.DoGet(fmt.Sprintf("/%s", params))

	if err != nil {
		return nil, err
	}

	return res, nil
}
