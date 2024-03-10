package client

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestClientCanHitAPI(t *testing.T) {
	t.Run("happy path - can hit the api and return a pokemon", func(t *testing.T) {
		myClient := NewClient()
		poke, err := myClient.GetPokemonByName(context.Background(), "pikachu")
		assert.NoError(t, err)
		assert.Equal(t, "pikachu", poke.Name)
	})

	t.Run("sad path - return an error when the pokemon does not exist", func(t *testing.T) {
		myClient := NewClient()
		_, err := myClient.GetPokemonByName(context.Background(), "non-existent")
		assert.Error(t, err)
		assert.Equal(t, err, PokemonFetchErr{
			Message:    "non-200 status code from the API",
			StatusCode: 404,
		})
	})

	t.Run("happy path - testing the WithAPIURL option function", func(t *testing.T) {
		myClient := NewClient(WithAPIURL("my-test-url"))
		assert.Equal(t, "my-test-url", myClient.apiURL)
	})

	t.Run("happy path - testing the WithHttpClient option function", func(t *testing.T) {
		myClient := NewClient(
			WithAPIURL("my-test-url"),
			WithHttpClient(&http.Client{
				Timeout: 1 * time.Second,
			}),
		)

		assert.Equal(t, "my-test-url", myClient.apiURL)
		assert.Equal(t, 1*time.Second, myClient.httpClient.Timeout)
	})

	t.Run("happy path - able to hit locally running test server", func(t *testing.T) {
		ts := httptest.NewServer(
			http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprintf(w, `{"name": "pikachu", "height": 10}`)
			}),
		)
		defer ts.Close()

		myClient := NewClient(
			WithAPIURL(ts.URL),
		)
		poke, err := myClient.GetPokemonByName(context.Background(), "pikachu")
		assert.NoError(t, err)
		assert.Equal(t, 10, poke.Height)
		assert.Equal(t, "pikachu", poke.Name)
	})

	t.Run("sad path - able to hit locally running test server", func(t *testing.T) {
		ts := httptest.NewServer(
			http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusInternalServerError)
			}),
		)
		defer ts.Close()

		myClient := NewClient(
			WithAPIURL(ts.URL),
		)
		poke, err := myClient.GetPokemonByName(context.Background(), "pikachu")
		assert.Error(t, err)
		assert.Equal(t, 0, poke.Height)
	})
}
