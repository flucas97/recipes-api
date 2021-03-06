package recipe

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIngredientsToSortedSlice(t *testing.T) {
	sr := SearchResult{
		Results: []FromRecipepuppy{
			{
				Title:       "Bolo de Banana",
				Link:        "https://www.delivery-much-test.com.br/recipes/bolo-de-banana",
				Ingredients: "ovo, caramelo, banana",
			},
		},
	}

	result := sr.IngredientsToSortedSlice()
	assert.Equal(t, []string{"banana", "caramelo", "ovo"}, result[0].Ingredients, "failed to cast or sort ingredients. recipe.IngredientsToSlice")
	assert.Contains(t, result[0].Ingredients, "caramelo", "ovo", "banana")
}
