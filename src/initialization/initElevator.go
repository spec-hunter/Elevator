package initialization

import (
	"time"

	"github.com/ahlemarg/Elevator/src/global"
	"github.com/ahlemarg/Elevator/src/models"
)

func InitElevator() {
	global.Elevator = &models.Elevator{
		State:           false,
		Up:              true,
		CurrentFloor:    1,
		TargetFloor:     1,
		ConsumptionTime: time.Duration(global.Config.ConsumptionTime.Expend),
	}
}
