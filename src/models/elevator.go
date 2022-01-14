package models

import (
	"time"
)

type Elevator struct {
	State           bool          // 开关门的状态(false: 关门  true: 开门)
	Up              bool          // 电梯是否是上升状态(true: 上升  false: 下降)
	CurrentFloor    int           // 当前楼层
	TargetFloor     int           // 目标楼层
	ConsumptionTime time.Duration // 行进一层楼固定消耗的时间
}
