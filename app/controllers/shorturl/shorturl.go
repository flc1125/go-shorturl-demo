package shorturl

import (
	"context"
	"crypto/md5"
	"fmt"
	"net/http"
	"shorturl/app/utils/url"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/labstack/echo/v4"
)

const LONG_PREFIX string = "LONG:"
const SHORT_PREFIX string = "SHORT:"

var redisClient *redis.Client
var ctx = context.Background()

func init() {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
}

func getGlobalId() int {
	num, _ := redisClient.Incr(ctx, "SHORT:URL:GLOBAL").Result()

	return int(num)
}

// 长链接转义
func longurlEncode(longurl string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(longurl)))
}

func Create(c echo.Context) error {
	longurl, expire := c.FormValue("longurl"), c.FormValue("expire")
	longurlHash := longurlEncode(longurl)
	expireInt, _ := strconv.Atoi(expire)

	if expireInt <= 0 {
		expireInt = 7 * 24 * 60 // 7天
	}

	shorturl, err := redisClient.Get(ctx, LONG_PREFIX+longurlHash).Result()
	if err == nil {
		return c.JSON(
			http.StatusOK, echo.Map{
				"status": 1, "msg": "SUCCESS", "data": echo.Map{
					"longurl": longurl, "shorturl": shorturl}})
	}

	shorturl = url.Encode(getGlobalId())
	redisClient.SetEX(ctx, LONG_PREFIX+longurlHash, shorturl, time.Minute*time.Duration(expireInt))
	redisClient.SetEX(ctx, SHORT_PREFIX+shorturl, longurl, time.Minute*time.Duration(expireInt))

	return c.JSON(
		http.StatusOK, echo.Map{
			"status": 1, "msg": "SUCCESS", "data": echo.Map{
				"longurl": longurl, "shorturl": shorturl}})
}

func Redirect(c echo.Context) error {
	shorturl := c.Param("url")

	longurl, err := redisClient.Get(ctx, SHORT_PREFIX+shorturl).Result()
	if err != nil {
		return c.HTML(http.StatusFound, c.Param("url")+"：Not Found")
	}

	return c.Redirect(http.StatusFound, longurl)
}
