package models

type FLOOR struct {
	Buttom int `mapstructure:"buttom" json:"buttom"`
	Top    int `mapstructure:"top" json:"top"`
}

type ConsumptionTime struct {
	Expend int `mapstructure:"expend" json:"expend"`
}

type Config struct {
	Floor           FLOOR           `mapstructure:"floor" json:"floor"`
	ConsumptionTime ConsumptionTime `mapstructure:"consumptionTime" json:"consumptionTime"`
}
