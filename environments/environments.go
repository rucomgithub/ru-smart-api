package environments

import (
	"strings"
	"time"

	"github.com/spf13/viper"
)

func EnvironmentInit() {
	// ไฟล์ที่จะจัดเก็บตัว Connection string Database
	viper.SetConfigName("environments")
	// ภาษาที่จะใช้ในการ Config
	viper.SetConfigType("yaml")
	// ที่อยู่ของ file config เริ่มค้นหาจาก root ด้านนอกสุด
	viper.AddConfigPath("./environments")
	// แล้วเข้ามาที่ environment folder
	viper.AddConfigPath("environments")

	viper.AutomaticEnv()
	viper.GetViper().SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// เรียก file config.yaml มาใช้
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func TimeZoneInit() {
	ict := time.Now().Local().Location()
	time.Local = ict
}