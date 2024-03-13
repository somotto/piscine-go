package piscine

type food struct {
	preptime int
  } 
  var menu = map[string]food{
	"burger": {preptime: 15},
	"chips": {preptime: 10},
	"nuggets": {preptime: 12},
  }
  
  func FoodDeliveryTime(order string) int {
	totalTime := 0
	items := split(order, ",")
	for _, item := range items {
		orderItem,ok := menu[trimSpace(item)]
		if !ok {
			return 0, errorf("404: Menu item \"%s\" notfound",item)
		}
		totalTime +=
	}
  }
  