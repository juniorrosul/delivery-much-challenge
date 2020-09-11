package secondary

import (
	"encoding/json"
	"fmt"

	"github.com/juniorrosul/delivery-much-challenge/application/recipe"
)

type RecipePuppy interface {
	getRecipes(params string) (*recipe.RecipePuppyResponse, error)
}

type RecipePuppyIntegration struct {
	conn HttpConnector
}

func NewRecipePuppyIntegration(httpConnector HttpConnector) *RecipePuppyIntegration {
	return &RecipePuppyIntegration{
		conn: httpConnector,
	}
}

func (rpi *RecipePuppyIntegration) getRecipes(params string) (*recipe.RecipePuppyResponse, error) {
	res, err := rpi.conn.DoGet(fmt.Sprintf("/%s", params))

	if err != nil {
		return nil, err
	}

	body, err := getBodyResponse(res)

	var rpr *recipe.RecipePuppyResponse

	if err = json.Unmarshal(body, &rpr); err != nil {
		return nil, err
	}

	return rpr, nil
}
