package piscine

type food struct {
	preptime int
}

func FoodDeliveryTime(order string) int {
	burger := food{15}
	chips := food{10}
	nuggets := food{12}
	if order == "burger" {
		return burger.preptime
	} else if order == "chips" {
		return chips.preptime
	} else if order == "nuggets" {
		return nuggets.preptime
	}
	return 404
}
