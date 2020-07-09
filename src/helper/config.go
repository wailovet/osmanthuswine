package helper

var ConfigFile = "config.json"

func SaveConfig(key string, i interface{}) {
	var g interface{}
	JsonByFile(ConfigFile, g)
	if g == nil {
		g = map[string]interface{}{
			key: i,
		}
	} else {
		g.(map[string]interface{})[key] = i
	}
	JsonToFile(ConfigFile, g)
}

func LoadConfig(key string, i interface{}) {
	data := map[string]interface{}{
		key: i,
	}
	JsonByFile(ConfigFile, &data)
}

func GetConfig(key string) interface{} {
	data := map[string]interface{}{}
	JsonByFile(ConfigFile, &data)
	return data[key]
}
