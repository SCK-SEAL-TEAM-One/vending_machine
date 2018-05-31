package vendingmachine

type vendingmachine struct {
	sumcoin int
	mapcoin map[string]int
	drinks  map[string]int
	coins   []string
}

func (vm *vendingmachine) Insert(coin string) int {

	vm.sumcoin += vm.mapcoin[coin]

	return vm.sumcoin
}

func (vm *vendingmachine) SelectItem(drink string) string {
	if vm.sumcoin >= vm.drinks[drink] {
		vm.sumcoin -= vm.drinks[drink]

	}
	return drink + vm.changes()
}

func (vm *vendingmachine) changes() string {
	var changes string
	for _, coin := range vm.coins {
		for vm.sumcoin >= vm.mapcoin[coin] {
			changes += coin
			vm.sumcoin -= vm.mapcoin[coin]
		}
	}

	return changes
}

func (vm *vendingmachine) Eject() string {
	return vm.changes()
}
