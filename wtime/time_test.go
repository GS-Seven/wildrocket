package wtime

import (
	"fmt"
	"testing"
	"time"
)

func TestWtime_FormatTime(t *testing.T) {
	s := WTtime.FormatTime(time.Now())
	fmt.Println(s)
}

func TestWtime_CurrentDayZero(t *testing.T) {
	dayZero := WTtime.CurrentDayZero()
	fmt.Println(dayZero)
}

func TestWtime_CurrentDayEnd(t *testing.T) {
	end := WTtime.CurrentDayEnd()
	fmt.Println(end)
}

func TestCalculateCurrentTimeAndZeroTime(t *testing.T) {
	d := WTtime.CalculateCurrentTimeAndZeroTime()
	fmt.Println(d.Seconds())
}

func TestPreviousDayStartTime(t *testing.T) {
	o := PreviousDayStartTime()
	fmt.Println(o)
}

func TestPreviousDayEndTime(t *testing.T) {
	e := PreviousDayEndTime()
	fmt.Println(e)
}

func TestPreviousAfterDate(t *testing.T) {
	d := PreviousAfterDate(1)  // 明天
	d2 := PreviousAfterDate(2) // 后天 -- 依次类推
	fmt.Println(d)
	fmt.Println(d2)

	d3 := PreviousAfterDate(-1) // 昨天
	d4 := PreviousAfterDate(-2) // 前天 -- 依次类推
	fmt.Println(d3)
	fmt.Println(d4)
}

func TestPreviousAfterTime(t *testing.T) {
	d := PreviousAfterTime(1)  // 明天当前时间
	d2 := PreviousAfterTime(2) // 后天当前时间 -- 依次类推
	fmt.Println(d)
	fmt.Println(d2)

	d3 := PreviousAfterTime(-1) // 昨天当前时间
	d4 := PreviousAfterTime(-2) // 前天当前时间 -- 依次类推
	fmt.Println(d3)
	fmt.Println(d4)
}
