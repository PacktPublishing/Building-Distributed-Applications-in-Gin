package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/mlabouardy/recipes-api/models"
	"github.com/stretchr/testify/assert"
)

func TestListRecipesHandler(t *testing.T) {
	ts := httptest.NewServer(SetupServer())
	defer ts.Close()

	resp, err := http.Get(fmt.Sprintf("%s/recipes", ts.URL))
	defer resp.Body.Close()
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	data, _ := ioutil.ReadAll(resp.Body)

	var recipes []models.Recipe
	json.Unmarshal(data, &recipes)
	assert.Equal(t, len(recipes), 10)
}

func TestUpdateRecipeHandler(t *testing.T) {
	ts := httptest.NewServer(SetupServer())
	defer ts.Close()

	recipe := Recipe{
		ID:   "c0283p3d0cvuglq85log",
		Name: "Oregano Marinated Chicken",
	}

	raw, _ := json.Marshal(recipe)
	resp, err := http.PUT(fmt.Sprintf("%s/recipes/%s", ts.URL, recipe.ID), bytes.NewBuffer(raw))
	defer resp.Body.Close()
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	data, _ := ioutil.ReadAll(resp.Body)

	var payload map[string]string
	json.Unmarshal(data, &payload)

	assert.Equal(t, payload["message"], "Recipe has been updated")
}

func TestDeleteRecipeHandler(t *testing.T) {
	ts := httptest.NewServer(SetupServer())
	defer ts.Close()

	resp, err := http.DELETE(fmt.Sprintf("%s/recipes/c0283p3d0cvuglq85log", ts.URL))
	defer resp.Body.Close()
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	data, _ := ioutil.ReadAll(resp.Body)

	var payload map[string]string
	json.Unmarshal(data, &payload)

	assert.Equal(t, payload["message"], "Recipe has been deleted")
}

func TestFindRecipeHandler(t *testing.T) {
	ts := httptest.NewServer(SetupServer())
	defer ts.Close()

	expectedRecipe := Recipe{
		ID:   "c0283p3d0cvuglq85log",
		Name: "Oregano Marinated Chicken",
		Tags: []string{"main", "chicken"},
	}

	resp, err := http.GET(fmt.Sprintf("%s/recipes/c0283p3d0cvuglq85log", ts.URL))
	defer resp.Body.Close()
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	data, _ := ioutil.ReadAll(resp.Body)

	var actualRecipe Recipe
	json.Unmarshal(data, &actualRecipe)

	assert.Equal(t, expectedRecipe.Name, actualRecipe.Name)
	assert.Equal(t, len(expectedRecipe.Tags), len(actualRecipe.Tags))
}
