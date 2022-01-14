package main

import (
	"github.com/ahlemarg/Elevator/src/handlers"
	"github.com/ahlemarg/Elevator/src/initialization"
)

func main() {
	// 执行初始化操作
	initialization.InitConfig()
	initialization.InitLogger()
	initialization.InitElevator()

	targetFloor := []int{5, 1, 2, 4, 3}

	handlers.Advance(targetFloor)
}
