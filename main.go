//package main
//
//import (
//	"fmt"
//	"io/ioutil"
//	"net/http"
//)
//
//func main() {
//	// Thay thế YOUR_API_KEY bằng API key của bạn
//	apiKey := "AIzaSyAz-hLh34RJEYpzJAxn8YCFlfVPg3iPJWI"
//	// Thay thế YOUR_ADDRESS bằng địa chỉ mà bạn muốn lấy định vị
//	//address := "YOUR_ADDRESS"
//	address := "1600 Amphitheatre Pkwy, Mountain View, CA 94043"
//
//	// Tạo URL cho yêu cầu API
//	url := fmt.Sprintf("https://maps.googleapis.com/maps/api/geocode/json?address=%s&key=%s", address, apiKey)
//
//	// Thực hiện yêu cầu HTTP GET đến API của Google Maps
//	resp, err := http.Get(url)
//	if err != nil {
//		fmt.Println("Lỗi khi gọi API:", err)
//		return
//	}
//	defer resp.Body.Close()
//
//	// Đọc dữ liệu từ phản hồi
//	body, err := ioutil.ReadAll(resp.Body)
//	if err != nil {
//		fmt.Println("Lỗi khi đọc dữ liệu:", err)
//		return
//	}
//
//	// In kết quả
//	fmt.Println(string(body))
//}

// package main
//
// import (
//
//	"context"
//	"fmt"
//	"log"
//
//	"googlemaps.github.io/maps"
//
// )
//
//	func main() {
//		// Tạo một bối cảnh mới
//		ctx := context.Background()
//
//		// Khởi tạo client cho Google Maps API
//		client, err := maps.NewClient(maps.WithAPIKey("AIzaSyAz-hLh34RJEYpzJAxn8YCFlfVPg3iPJWI"))
//		if err != nil {
//			log.Fatalf("maps.NewClient(): %v", err)
//		}
//
//		// Lấy vị trí hiện tại
//		loc, err := client.GetCurrentPlace(ctx)
//		if err != nil {
//			log.Fatalf("client.GetCurrentPlace(): %v", err)
//		}
//
//		// Hiển thị vị trí
//		if len(loc.Places) > 0 {
//			place := loc.Places[0]
//			fmt.Printf("Latitude: %f, Longitude: %f\n", place.Geometry.Location.Lat, place.Geometry.Location.Lng)
//		} else {
//			fmt.Println("No place found")
//		}
//	}
package main

import (
	"context"
	"fmt"
	"log"

	"googlemaps.github.io/maps"
)

func main() {
	// Create a new context
	ctx := context.Background()

	// Initialize a client for the Google Maps API
	client, err := maps.NewClient(maps.WithAPIKey("AIzaSyAz-hLh34RJEYpzJAxn8YCFlfVPg3iPJWI"))
	if err != nil {
		log.Fatalf("maps.NewClient(): %v", err)
	}

	// Geocode an address
	res, err := client.Geocode(ctx, &maps.GeocodingRequest{
		Address: "1600 Amphitheatre Pkwy, Mountain View, CA 94043",
	})
	if err != nil {
		log.Fatalf("client.Geocode(): %v", err)
	}

	// Display the location
	if len(res) > 0 {
		loc := res[0]
		fmt.Printf("Latitude: %f, Longitude: %f\n", loc.Geometry.Location.Lat, loc.Geometry.Location.Lng)
	} else {
		fmt.Println("No place found")
	}
}
