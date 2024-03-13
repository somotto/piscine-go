package piscine

type food struct {
	preptime int
}

var menu = map[string]food{
	"burger":  {preptime: 15},
	"chips":   {preptime: 10},
	"nuggets": {preptime: 12},
}

func FoodDeliveryTime(order string) int {
	totalprepTime := 0
	for _, item := range order {
		foodItem := string(item)
		if menu[foodItem].preptime == 0 {
			return -1
		}
		totalprepTime += menu[foodItem].preptime
	}
	return totalprepTime
}
