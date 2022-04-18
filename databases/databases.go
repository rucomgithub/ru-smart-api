package databases

import (
	"fmt"

	"github.com/go-redis/redis/v8"
	_ "github.com/godror/godror"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
)

type connection struct{}

func NewDatabases() *connection {
	return &connection{}
}

func (c *connection) RedisConnection() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: viper.GetString("redis_cache.addressLocal"),
		// Addr: viper.GetString("redis_cache.address"),
		Password: viper.GetString("redis_cache.password"),
		DB:       viper.GetInt("redis_cache.db-num"),
	})
}

func (c *connection) OracleConnection() (*sqlx.DB, error) {

	dns := fmt.Sprintf("%v", viper.GetString("db.connection"))
	driver := viper.GetString("db.openDriver")

	return sqlx.Open(driver, dns)

}
