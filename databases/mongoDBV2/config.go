package mongoDBV2

import "github.com/caarlos0/env/v11"

type Config struct {
	//DATABASE
	//HOST        string `env:"DB_HOST" envDefault:"localhost" json:"db_host"`
	//PORT        string `env:"DB_PORT" envDefault:"27107" json:"db_port"`
	//SSL         string `env:"DB_SSL" envDefault:"disable" json:"db_ssl"`

	NAME        string `env:"DB_NAME" envDefault:"admin" json:"db_name"`
	CLUSTER_URL string `env:"CLUSTER_URL" envDefault:"cluster0.admin.mongodb.net" json:"db_cluster_url"`
	USER        string `env:"DB_USER" json:"db_user"`
	PASS        string `env:"DB_PASS" json:"db_pass"`
	TZ          string `env:"TIME_ZONE" envDefault:"Asia/Jakarta" json:"db_tz"`

	MAX_TIME_OUT_CONNS int `env:"MAX_TIME_OUT_CONNS" envDefault:"300" json:"max_time_out_conns"` //default 300ms
	MIN_POOL_SIZE      int `env:"MIN_POOL_SIZE" envDefault:"10" json:"min_pool_size"`            //default 10
	MAX_POOL_SIZE      int `env:"MAX_POOL_SIZE" envDefault:"100" json:"max_pool_size"`           //default 100

	//ETC
	EngineName string `env:"ENGINE_NAME" json:"engine_name,omitempty"`
}

// Parse for parse env variables to this struct
func (this *Config) Parse() (err error) {
	return env.Parse(this)
}
