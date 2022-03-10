package databases

import (
	"fmt"

	_ "github.com/godror/godror"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
)

type connection struct{}

func NewDatabases() *connection {
	return &connection{}
}

// func (c *connection) RedisConnection() *redis.Client {
// 	return redis.NewClient(&redis.Options{
// 		Addr: viper.GetString("rdb.addressLocal"),
// 		// Addr: viper.GetString("rdb.address"),
// 		Password: viper.GetString("rdb.password"),
// 		DB:       viper.GetInt("rdb.db-num"),
// 	})
// }

func (c *connection) OracleConnection() (*sqlx.DB, error) {

	dns := fmt.Sprintf("%v", viper.GetString("db.connection"))
	driver := viper.GetString("db.openDriver")

	return sqlx.Open(driver, dns)

}
