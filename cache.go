package main

import (
    "time"
    "github.com/patrickmn/go-cache"
)

func initCache() (cache.Cache) {
	return *cache.New(5*time.Minute, 10*time.Minute)
}

func addRoutesToCache(c cache.Cache, routes []Route) {
	for _, route := range routes {
		key := "ROUTE_"+route.Method+"_"+route.Path
        c.Set(key, route, 5*time.Minute)
    }
}