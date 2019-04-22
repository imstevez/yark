package configs

var Configs map[string]interface{}

func init() {
	Configs = map[string]interface{}{
		"sql_db": map[string]interface{}{
			"username":       "root",
			"password":       "root",
			"protocol":       "tcp",
			"host":           "127.0.0.1",
			"port":           "3306",
			"database":       "yark",
			"max_life_time":  100,
			"max_open_conns": 16,
			"max_idle_conns": 16,
		},
		"server": map[string]interface{}{
			"port": ":3000",
		},
		"baseInfos": map[string]interface{}{
			"app_name": "Yark v1.0",
			"author":   "StevE.Z",
		},
	}
}
