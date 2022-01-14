** 限制**

```
    楼层限制:
        1. 最低楼层 BOTTOM_FLOOR : 1
        2. 最高楼层 TOP_FLOOR : 5
```

** 电梯**

```
    属性:
        1. 开关门状态 State : bool
        2. 上升 Up : bool
        3. 当前楼层 Current_Floor : int
        4. 目标楼层 Target_Floor : int
        5. 前进一层楼所使用的固定时间 Consumption_Time : time.Duration

    行为:
        1. 行动到最优的楼层 Advance
            1.1 
```



****

其实还有一些小的逻辑问题要处理, 就是时间有点赶不及, 大致逻辑没问题
