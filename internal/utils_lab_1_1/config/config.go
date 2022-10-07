package config

import (
	"flag"
	"fmt"
	"os"
)

type APMod uint8

const (
	APM_min APMod = iota
	APM_nope
	APM_all
)

type UPMod uint8

const (
	UPM_nope UPMod = iota
	UPM_once
	UPM_everyime
)

type Config struct {
	AutoPerm    APMod
	UserPerm    UPMod
	UserPermVal int
}

var config Config

/*
	--up=nope|once|everyime
	--up-val=[1-9][0-9]*

	--ap=nope|min|all
*/
const (
	uPUsage = `user permutation mode: nope (default)|once|everytime
	nope - Пользователь не выбирает ведущую строку
	once - Пользователь выбирает ведущую строку только один раз
		   При этом строка узказывается с помощью ключа up-val=<col>
		   Отключает автоматическую перестановку на первой итерации
	everytime - Пользователь выбирает ведущую строку на каждой итерации алгоритма
				Ввод происходит в итеративном режиме
				Отключает автоматические перестановки вообще`

	uPValUsage = `user permutation value: номер строки, которая будет выбрана ведущей
	Работает только с ключом -up=once`

	aPUsage = `auto permutation mode: min (default)|nope|all 
	min	- Перестановка тогда, и только тогда, когда главный элемент на 
		   Текущей итреации равен нулю
	all - Перестановка на каждой итерации, главный элемент выбирается макимальный по модулю
	nope - Отключить автоматическую перестановку
		   Если главный элемент равен нулю - выбрасывается ошибка`
)

func Init() {
	upFlag := flag.String("up", "nope", uPUsage)
	apFlag := flag.String("ap", "min", aPUsage)
	flag.IntVar(&config.UserPermVal, "up-val", 0, uPValUsage)

	flag.Parse()

	// check up
	switch *upFlag {
	case "nope":
		config.UserPerm = UPM_nope
	case "once":
		config.UserPerm = UPM_once
	case "everytime":
		config.UserPerm = UPM_everyime
	default:
		fmt.Fprintf(flag.CommandLine.Output(), "up haven't value '%s' see help\n", *apFlag)
		flag.PrintDefaults()
		os.Exit(2)
	}

	// check ap
	switch *apFlag {
	case "min":
		config.AutoPerm = APM_min
	case "nope":
		config.AutoPerm = APM_nope
	case "all":
		config.AutoPerm = APM_all
	default:
		fmt.Fprintf(flag.CommandLine.Output(), "ap haven't value '%s' see help\n", *apFlag)
		flag.PrintDefaults()
		os.Exit(2)
	}
}

func Get() *Config {
	return &config
}
