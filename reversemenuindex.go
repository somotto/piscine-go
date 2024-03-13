package piscine

func ReverseMenuIndex(menu []string) []string {
	reversedmenu := make([]string, len(menu))
	for i := range menu {
		reversedmenu[len(menu)-1-i] = menu[i]
	}
	return reversedmenu
}
