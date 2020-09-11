package rpa

import (
	"fmt"
	"testing"

	"github.com/juniorrosul/delivery-much-challenge/application/recipe"
)

func TestNewRecipePuppyIntegration(t *testing.T) {
	t.Run("NO PARAMS", func(t *testing.T) {
		mockedIntegration := NewRecipePuppyIntegration(&ConnectorMock{})

		_, err := mockedIntegration.GetRecipes(&recipe.RequestModel{})

		if err != nil {
			fmt.Println(err)
			t.Error("unexpected error")
		}
	})

	// t.Run("WITH PARAMS", func(t *testing.T) {
	// 	mockedIntegration := NewRecipePuppyIntegration(&ConnectorMock{
	// 		ResponseStatus: 200,
	// 		ResponseBody: `{
	// 			"title": "\nMexican Salsa Recipe\n\n",
	// 			"href": "http:\/\/cookeatshare.com\/recipes\/mexican-salsa-55686",
	// 			"ingredients": "onions, tomato, garlic, jalapeno",
	// 			"thumbnail": "http:\/\/img.recipepuppy.com\/898566.jpg"
	// 		  },
	// 		  {
	// 			"title": "Fast and Low Fat Beans and Tomatoes for a Weeknight",
	// 			"href": "http:\/\/www.recipezaar.com\/Fast-and-Lowfat-Beans-and-Tomatoes-for-a-Weeknight-98962",
	// 			"ingredients": "cannellini beans, tomato, basil, garlic, onions",
	// 			"thumbnail": "http:\/\/img.recipepuppy.com\/555909.jpg"
	// 		  },
	// 		  {
	// 			"title": "Fresh Tomato Basil Sauce",
	// 			"href": "http:\/\/allrecipes.com\/Recipe\/Fresh-Tomato-Basil-Sauce\/Detail.aspx",
	// 			"ingredients": "garlic, olive oil, onions, salt, tomato",
	// 			"thumbnail": "http:\/\/img.recipepuppy.com\/10228.jpg"
	// 		  }`,
	// 	})

	// 	_, err := mockedIntegration.GetRecipes(&recipe.RequestModel{
	// 		Ingredients: []string{"i1", "i2", "i3"},
	// 	})

	// 	if err != nil {
	// 		fmt.Println(err)
	// 		t.Error("unexpected error")
	// 	}
	// })
}
