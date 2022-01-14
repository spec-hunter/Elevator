package test

import (
	"testing"
	"time"

	"github.com/ahlemarg/Elevator/src/global"
	"github.com/ahlemarg/Elevator/src/handlers"
	"github.com/ahlemarg/Elevator/src/initialization"
	"github.com/ahlemarg/Elevator/src/models"
)

func TestElevatorUP(t *testing.T) {
	global.Elevator = &models.Elevator{
		State:           false,
		Up:              true,
		CurrentFloor:    1,
		TargetFloor:     1,
		ConsumptionTime: time.Duration(global.Config.ConsumptionTime.Expend),
	}

	initialization.InitConfig()
	initialization.InitLogger()

	targetFloor := []int{5, 1, 2, 4, 3}

	handlers.Advance(targetFloor)
}

func TestElevatorDown(t *testing.T) {
	global.Elevator = &models.Elevator{
		State:           false,
		Up:              true,
		CurrentFloor:    5,
		TargetFloor:     5,
		ConsumptionTime: time.Duration(global.Config.ConsumptionTime.Expend),
	}

	initialization.InitConfig()
	initialization.InitLogger()

	targetFloor := []int{5, 1, 2, 4, 3}

	handlers.Advance(targetFloor)
}
