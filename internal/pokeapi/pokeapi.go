package pokeapi

import (
	"github.com/Seeker-09/pokedexCLI/internal/pokeapi/pokecache"
	"net/http"
	"time"
)

const baseURL = "https://pokeapi.co/api/v2"

type Client struct {
	cache      pokecache.Cache
	httpClient http.Client
}

func NewClient(cachInterval time.Duration) Client {
	return Client{
		cache: pokecache.NewCache(cachInterval),
		httpClient: http.Client{
			Timeout: time.Minute,
		},
	}
}
