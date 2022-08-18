package main

import (
	"flag"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/metall27/ServerAndDB2/internal/app/api"
)

var ( //Подготовка блока локальных переменных
	configPath string
)

func init() { //Подготовка блока инициализации
	//Скажем, что наше приложение будет на этапе запуска получать путь до конфиг файла из внешнего мира
	flag.StringVar(&configPath, "path", "configs/api.toml", "path to config file in .toml format")

}

func main() {
	// В этот момент происходит инициализация переменной configPath со значением
	flag.Parse()
	log.Println("It works")

	//server instance initialization
	config := api.NewConfig() //инстанс с дефолтными значениями

	//Теперь надо попробовать прочитать из .toml/.env, т.к. там может быть новая информация
	_, err := toml.DecodeFile(configPath, config) // Десериализуем toml файл
	if err != nil {
		log.Println("can not find config file, using default values:", err)
	}
	server := api.New(config)

	//api server start
	// if err := server.Start(); err != nil {
	// 	log.Fatal(err)
	// }
	//а можно так запускаться, никакой разницы
	log.Fatal(server.Start())
}
