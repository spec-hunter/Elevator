package handlers

import (
	"sort"
	"time"

	"github.com/ahlemarg/Elevator/src/global"
	"github.com/ahlemarg/Elevator/src/models"
	"go.uber.org/zap"
)

func setState(current, target int) {
	// 设置到达对应楼层对电梯状态的设置
	global.Elevator = &models.Elevator{
		State:           false,
		Up:              global.Elevator.Up,
		CurrentFloor:    current,
		TargetFloor:     target,
		ConsumptionTime: time.Duration(1 + global.Config.ConsumptionTime.Expend),
	}
}

func move(floor []int, up bool, current int) {

	index := sort.Search(len(floor), func(i int) bool {
		return floor[i] >= current
	})

	new_index := index
	for i := 0; i < len(floor); i++ {
		if up {
			f := floor[new_index]
			zap.S().Infof("当前已上升至%d楼层, 请稍等.\n", f)
			new_index++
			// 如果到顶了, 则将电梯设置为下降
			if new_index > len(floor)-1 {
				global.Elevator.Up = false
				new_index = 0
			}
			// 设置电梯属性
			setState(f, floor[new_index])
		} else {
			f := floor[new_index]
			zap.S().Infof("当前已下降到楼层%d,请稍等.\n", f)
			new_index--
			// 如果到顶了, 则将电梯设置为下降
			if new_index < 0 {
				global.Elevator.Up = true
				new_index = 0
			}
			// 设置电梯属性
			setState(f, floor[new_index])
		}
	}
}

func Advance(targetFloor []int) {
	// 对targetFloor进行一次排序
	floor := countSort(targetFloor)

	// 根据当前电梯楼层判断行进方向
	if global.Elevator.CurrentFloor == global.Config.Floor.Top {
		global.Elevator.Up = false
	} else if global.Elevator.CurrentFloor == global.Config.Floor.Buttom {
		global.Elevator.Up = true
	}

	// 根据电梯Up属性排序行进的楼层
	move(floor, global.Elevator.Up, global.Elevator.CurrentFloor)
}
