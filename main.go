package main

import (
	"fmt"
	"time"
)

//func main() {
//	startTime := time.Now() // Thời điểm bắt đầu
//
//	// Lấy thời gian hiện tại theo múi giờ UTC+7
//	currentTime := time.Now().UTC().Add(7 * time.Hour)
//
//	// Đặt đầu tuần là thứ Hai
//	weekStart := currentTime
//	for weekStart.Weekday() != time.Monday {
//		weekStart = weekStart.Add(-24 * time.Hour)
//	}
//
//	// Tính toán cuối tuần (thứ Sáu)
//	weekEnd := weekStart.Add(5 * 24 * time.Hour)
//
//	fmt.Println("Week Start Date:", weekStart.Format("Monday,02-01-2006"))
//	fmt.Println("Week End Date:", weekEnd.Format("Monday, 02-01-2006"))
//
//	endTime := time.Now()                     // Thời điểm kết thúc
//	elapsedTime := endTime.Sub(startTime)     // Thời gian đã trôi qua
//	fmt.Println("Elapsed Time:", elapsedTime) // In ra thời gian đã trôi qua
//}

//func main() {
//	startTime := time.Now() // Thời điểm bắt đầu
//
//	// Lấy thời gian hiện tại theo múi giờ UTC+7
//	currentTime := time.Now().UTC().Add(7 * time.Hour)
//
//	// Tính toán ngày bắt đầu của tuần (thứ Hai)
//	weekday := currentTime.Weekday()
//	daysSinceMonday := int(weekday - time.Monday)
//	if daysSinceMonday < 0 {
//		daysSinceMonday += 7
//	}
//	weekStart := currentTime.AddDate(0, 0, -daysSinceMonday)
//
//	// Tính toán ngày kết thúc của tuần (thứ Sáu)
//	weekEnd := weekStart.Add(4 * 24 * time.Hour)
//
//	// In ra ngày bắt đầu của tuần và ngày kết thúc của tuần với định dạng mong muốn
//	fmt.Println("Week Start Date:", weekStart.Format("Monday, 02-01-2006"))
//	fmt.Println("Week End Date:", weekEnd.Format("Monday, 02-01-2006"))
//
//	endTime := time.Now()                     // Thời điểm kết thúc
//	elapsedTime := endTime.Sub(startTime)     // Thời gian đã trôi qua
//	fmt.Println("Elapsed Time:", elapsedTime) // In ra thời gian đã trôi qua
//}

func main() {
	startTime := time.Now() // Thời điểm bắt đầu

	// Lấy thời gian hiện tại theo múi giờ UTC+7
	currentTime := time.Now().UTC().Add(7 * time.Hour)

	// Tính toán ngày bắt đầu của tuần (thứ Hai)
	weekStart := currentTime.AddDate(0, 0, -int(currentTime.Weekday())+1)

	// Tính toán ngày kết thúc của tuần (thứ Sáu)
	weekEnd := weekStart.Add(6 * 24 * time.Hour)

	// In ra ngày bắt đầu của tuần và ngày kết thúc của tuần với định dạng mong muốn
	fmt.Println("Week Start Date:", weekStart.Format("Monday, 02-01-2006"))
	fmt.Println("Week End Date:", weekEnd.Format("Monday, 02-01-2006"))

	endTime := time.Now()                     // Thời điểm kết thúc
	elapsedTime := endTime.Sub(startTime)     // Thời gian đã trôi qua
	fmt.Println("Elapsed Time:", elapsedTime) // In ra thời gian đã trôi qua
}
