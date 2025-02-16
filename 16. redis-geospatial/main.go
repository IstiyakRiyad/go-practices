package main

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	goRedis "github.com/redis/go-redis/v9"
)

func main() {
    redisClient := goRedis.NewClient(&goRedis.Options{
        Network: "tcp",
        Addr: ":6379",
        Username: "",
        Password: "",
    })

    ctx, cancel := context.WithTimeout(context.Background(), time.Second * 5)
    defer cancel()


    err := redisClient.GeoAdd(ctx, "geo:1", &goRedis.GeoLocation{
        Name: "b_44",
        Latitude: 23.303,
        Longitude: 90.23,
    }, &goRedis.GeoLocation{
        Name: "a_23",
        Latitude: 23.905,
        Longitude: 90.93,
    }).Err()
    if err != nil {
        slog.Error("Faild to add geo info", slog.String("err", err.Error()))
    }


    result, err := redisClient.GeoSearch(ctx, "geo:1", &goRedis.GeoSearchQuery{
        Latitude: 23.203,
        Longitude: 90.20,
        Radius: 50,
        RadiusUnit: "km",

    }).Result()
    if err != nil {
        slog.Error("Faild to search geo info", slog.String("err", err.Error()))
    }

    fmt.Println(result)
}
