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
	item, found := menu[order]
	if found {
		return item.preptime
	} else {
		return 404
	}
}
