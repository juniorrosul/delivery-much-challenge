package secondary

import (
	"encoding/json"
	"fmt"

	"github.com/juniorrosul/delivery-much-challenge/application/recipe"
)

// RecipePuppy - Interface
type RecipePuppy interface {
	getRecipes(params string) (*recipe.RecipePuppyResponseModel, error)
}

// RecipePuppyIntegration - Base struct for API integration
type RecipePuppyIntegration struct {
	conn HTTPConnector
}

// NewRecipePuppyIntegration - Integration for API
func NewRecipePuppyIntegration(httpConnector HTTPConnector) *RecipePuppyIntegration {
	return &RecipePuppyIntegration{
		conn: httpConnector,
	}
}

// GetRecipes - Get recipes from external API
func (rpi *RecipePuppyIntegration) GetRecipes(params string) (*recipe.RecipePuppyResponseModel, error) {
	res, err := rpi.conn.DoGet(fmt.Sprintf("/%s", params))

	if err != nil {
		return nil, err
	}

	body, err := GetBodyResponse(res)
	body = []byte(`{
		"title": "\nMexican Salsa Recipe\n\n",
		"href": "http:\/\/cookeatshare.com\/recipes\/mexican-salsa-55686",
		"ingredients": "onions, tomato, garlic, jalapeno",
		"thumbnail": "http:\/\/img.recipepuppy.com\/898566.jpg"
	  },
	  {
		"title": "Fast and Low Fat Beans and Tomatoes for a Weeknight",
		"href": "http:\/\/www.recipezaar.com\/Fast-and-Lowfat-Beans-and-Tomatoes-for-a-Weeknight-98962",
		"ingredients": "cannellini beans, tomato, basil, garlic, onions",
		"thumbnail": "http:\/\/img.recipepuppy.com\/555909.jpg"
	  },
	  {
		"title": "Fresh Tomato Basil Sauce",
		"href": "http:\/\/allrecipes.com\/Recipe\/Fresh-Tomato-Basil-Sauce\/Detail.aspx",
		"ingredients": "garlic, olive oil, onions, salt, tomato",
		"thumbnail": "http:\/\/img.recipepuppy.com\/10228.jpg"
	  }`)

	var rpr *recipe.RecipePuppyResponseModel

	if err = json.Unmarshal(body, &rpr); err != nil {
		return nil, err
	}

	return rpr, nil
}
