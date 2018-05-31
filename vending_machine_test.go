package vendingmachine

import (
	"testing"
)

var mapcoinx = map[string]int{
	"F":  5,
	"T":  10,
	"TW": 2,
	"O":  1,
}

var mapdrinks = map[string]int{
	"SD": 18,
	"CC": 12,
}
var coins = []string{"T", "F", "TW", "O"}

func Test_Insert_T_Should_Display_10(t *testing.T) {
	vm := vendingmachine{mapcoin: mapcoinx, drinks: mapdrinks, coins: coins}
	expected := 10
	amount := vm.Insert("T")
	if expected != amount {
		t.Errorf("Expected %d but %d", expected, amount)
	}
}

func Test_Insert_F_Should_Display_5(t *testing.T) {
	vm := vendingmachine{mapcoin: mapcoinx, drinks: mapdrinks, coins: coins}
	expected := 5
	amount := vm.Insert("F")
	if expected != amount {
		t.Errorf("Expected %d but %d", expected, amount)
	}
}

func Test_Insert_T_F_Should_Display_15(t *testing.T) {
	vm := vendingmachine{mapcoin: mapcoinx, drinks: mapdrinks, coins: coins}
	expected := 15
	vm.Insert("T")
	amount := vm.Insert("F")
	if expected != amount {
		t.Errorf("Expected %d but %d", expected, amount)
	}
}

func Test_Insert_T_F_TW_Should_Display_17(t *testing.T) {
	vm := vendingmachine{mapcoin: mapcoinx, drinks: mapdrinks}
	expected := 17
	vm.Insert("T")
	vm.Insert("F")
	amount := vm.Insert("TW")
	if expected != amount {
		t.Errorf("Expected %d but %d", expected, amount)
	}
}

func Test_Insert_T_F_TW_O_Should_Display_18(t *testing.T) {
	vm := vendingmachine{mapcoin: mapcoinx, drinks: mapdrinks, coins: coins}
	expected := 18
	vm.Insert("T")
	vm.Insert("F")
	vm.Insert("TW")
	amount := vm.Insert("O")
	if expected != amount {
		t.Errorf("Expected %d but %d", expected, amount)
	}
}

func Test_SelectSD_Should_Receive_SD(t *testing.T) {
	vm := vendingmachine{mapcoin: mapcoinx, drinks: mapdrinks, coins: coins}

	vm.Insert("T")
	vm.Insert("F")
	vm.Insert("TW")
	vm.Insert("O")
	expected := "SD"

	drink := vm.SelectItem("SD")
	if expected != drink {
		t.Errorf("Expected %s but %s", expected, drink)
	}
}

func Test_Eject_Should_Return_Coin_T_T_F(t *testing.T) {
	vm := vendingmachine{mapcoin: mapcoinx, drinks: mapdrinks, coins: coins}
	vm.Insert("T")
	vm.Insert("T")
	vm.Insert("F")
	expected := "TTF"
	coins := vm.Eject()
	if expected != coins {
		t.Errorf("Expected %s but %s", expected, coins)
	}
}

func Test_Eject_Should_Return_Coin_T_T(t *testing.T) {
	vm := vendingmachine{mapcoin: mapcoinx, drinks: mapdrinks, coins: coins}
	vm.Insert("T")
	vm.Insert("T")
	expected := "TT"
	coins := vm.Eject()
	if expected != coins {
		t.Errorf("Expected %s but %s", expected, coins)
	}
}

func Test_SelectCC_With_20_Should_Receive_CC_with_changes_F_TW_O(t *testing.T) {
	vm := vendingmachine{mapcoin: mapcoinx, drinks: mapdrinks, coins: coins}

	vm.Insert("T")
	vm.Insert("T")
	expected := "CCFTWO"

	drink := vm.SelectItem("CC")
	if expected != drink {
		t.Errorf("Expected %s but %s", expected, drink)
	}
}

func Test_SelectSD_Should_Receive_SD_And_Eject_Nothing(t *testing.T) {
	vm := vendingmachine{mapcoin: mapcoinx, drinks: mapdrinks, coins: coins}

	vm.Insert("T")
	vm.Insert("F")
	vm.Insert("TW")
	vm.Insert("O")
	expected := ""

	vm.SelectItem("SD")
	coins := vm.Eject()

	if expected != coins {
		t.Errorf("Expected %s but %s", expected, coins)
	}
}
