package CreekAndContinent

import (
	"testing"
)

func TestUpdatesNormalItemsBeforeSellDate(t *testing.T) {
	item := CreekAndContinent{
		Name:   "normal",
		Price:  10,
		SellIn: 5,
	}
	item.tick()
	assertEquals(t, 9, item.Price)
	assertEquals(t, 4, item.SellIn)
}

func TestUpdatesNormalItemsOnTheSellDate(t *testing.T) {
	item := CreekAndContinent{
		Name:   "normal",
		Price:  10,
		SellIn: 0,
	}
	item.tick()
	assertEquals(t, 8, item.Price)
	assertEquals(t, -1, item.SellIn)
}

func TestUpdatesNormalItemsAfterTheSellDate(t *testing.T) {
	item := CreekAndContinent{
		Name:   "normal",
		Price:  10,
		SellIn: -5,
	}
	item.tick()
	assertEquals(t, 8, item.Price)
	assertEquals(t, -6, item.SellIn)
}

func TestUpdatesNormalItemsWithPriceOf0(t *testing.T) {
	item := CreekAndContinent{
		Name:   "normal",
		Price:  0,
		SellIn: 5,
	}
	item.tick()
	assertEquals(t, 0, item.Price)
	assertEquals(t, 4, item.SellIn)
}

func TestUpdatesJapaneseDenimItemsBeforeTheSellDate(t *testing.T) {
	item := CreekAndContinent{
		Name:   "Japanese Denim",
		Price:  10,
		SellIn: 5,
	}
	item.tick()
	assertEquals(t, 11, item.Price)
	assertEquals(t, 4, item.SellIn)
}

func TestUpdatesJapaneseDenimItemsBeforeTheSellDateWithMaxPrice(t *testing.T) {
	item := CreekAndContinent{
		Name:   "Japanese Denim",
		Price:  50,
		SellIn: 5,
	}
	item.tick()
	assertEquals(t, 50, item.Price)
	assertEquals(t, 4, item.SellIn)
}

func TestUpdatesJapaneseDenimItemsOnTheSellDate(t *testing.T) {
	item := CreekAndContinent{
		Name:   "Japanese Denim",
		Price:  10,
		SellIn: 0,
	}
	item.tick()
	assertEquals(t, 12, item.Price)
	assertEquals(t, -1, item.SellIn)
}

func TestUpdatesJapaneseDenimItemsOnTheSellDateNearMaxPrice(t *testing.T) {
	item := CreekAndContinent{
		Name:   "Japanese Denim",
		Price:  49,
		SellIn: 0,
	}
	item.tick()
	assertEquals(t, 50, item.Price)
	assertEquals(t, -1, item.SellIn)
}

func TestUpdatesJapaneseDenimItemsOnTheSellDateWithMaxPrice(t *testing.T) {
	item := CreekAndContinent{
		Name:   "Japanese Denim",
		Price:  50,
		SellIn: 0,
	}
	item.tick()
	assertEquals(t, 50, item.Price)
	assertEquals(t, -1, item.SellIn)
}

func TestUpdatesJapaneseDenimItemsAfterTheSellDate(t *testing.T) {
	item := CreekAndContinent{
		Name:   "Japanese Denim",
		Price:  10,
		SellIn: -10,
	}
	item.tick()
	assertEquals(t, 12, item.Price)
	assertEquals(t, -11, item.SellIn)
}

func TestUpdatesJapaneseDenimItemsAfterTheSellDateWithMaxPrice(t *testing.T) {
	item := CreekAndContinent{
		Name:   "Japanese Denim",
		Price:  50,
		SellIn: -10,
	}
	item.tick()
	assertEquals(t, 50, item.Price)
	assertEquals(t, -11, item.SellIn)
}

func TestUpdatesWhiteVestTopItemsBeforeTheSellDate(t *testing.T) {
	item := CreekAndContinent{
		Name:   "White Vest Top",
		Price:  10,
		SellIn: 5,
	}
	item.tick()
	assertEquals(t, 10, item.Price)
	assertEquals(t, 5, item.SellIn)
}

func TestUpdatesWhiteVestTopItemsOnTheSellDate(t *testing.T) {
	item := CreekAndContinent{
		Name:   "White Vest Top",
		Price:  10,
		SellIn: 5,
	}
	item.tick()
	assertEquals(t, 10, item.Price)
	assertEquals(t, 5, item.SellIn)
}

func TestUpdatesWhiteVestTopItemsAfterTheSellDate(t *testing.T) {
	item := CreekAndContinent{
		Name:   "White Vest Top",
		Price:  10,
		SellIn: -1,
	}
	item.tick()
	assertEquals(t, 10, item.Price)
	assertEquals(t, -1, item.SellIn)
}

func TestUpdatesFashionShowTicketsItemsBeforeTheSellDate(t *testing.T) {
	item := CreekAndContinent{
		Name:   "Fashion Show Tickets",
		Price:  10,
		SellIn: 11,
	}
	item.tick()
	assertEquals(t, 11, item.Price)
	assertEquals(t, 10, item.SellIn)
}

func TestUpdatesFashionShowTicketsItemsCloseToTheSellDate(t *testing.T) {
	item := CreekAndContinent{
		Name:   "Fashion Show Tickets",
		Price:  10,
		SellIn: 10,
	}
	item.tick()
	assertEquals(t, 12, item.Price)
	assertEquals(t, 9, item.SellIn)
}

func TestUpdatesFashionShowTicketsItemsBeforeTheSellDateAtMaxPrice(t *testing.T) {
	item := CreekAndContinent{
		Name:   "Fashion Show Tickets",
		Price:  50,
		SellIn: 10,
	}
	item.tick()
	assertEquals(t, 50, item.Price)
	assertEquals(t, 9, item.SellIn)
}

func TestUpdatesFashionShowTicketsItemsVeryCloseToTheSellDate(t *testing.T) {
	item := CreekAndContinent{
		Name:   "Fashion Show Tickets",
		Price:  10,
		SellIn: 5,
	}
	item.tick()
	assertEquals(t, 13, item.Price)
	assertEquals(t, 4, item.SellIn)
}

func TestUpdatesFashionShowTicketsItemsVeryCloseToTheSellDateAtMaxPrice(t *testing.T) {
	item := CreekAndContinent{
		Name:   "Fashion Show Tickets",
		Price:  50,
		SellIn: 5,
	}
	item.tick()
	assertEquals(t, 50, item.Price)
	assertEquals(t, 4, item.SellIn)
}

func TestUpdatesFashionShowTicketsItemsWithOneDayLeftToSell(t *testing.T) {
	item := CreekAndContinent{
		Name:   "Fashion Show Tickets",
		Price:  10,
		SellIn: 1,
	}
	item.tick()
	assertEquals(t, 13, item.Price)
	assertEquals(t, 0, item.SellIn)
}

func TestUpdatesFashionShowTicketsItemsWithOneDayLeftToSellAtMaxPrice(t *testing.T) {
	item := CreekAndContinent{
		Name:   "Fashion Show Tickets",
		Price:  50,
		SellIn: 1,
	}
	item.tick()
	assertEquals(t, 50, item.Price)
	assertEquals(t, 0, item.SellIn)
}

func TestUpdatesFashionShowTicketsItemsOnSellDate(t *testing.T) {
	item := CreekAndContinent{
		Name:   "Fashion Show Tickets",
		Price:  10,
		SellIn: 0,
	}
	item.tick()
	assertEquals(t, 0, item.Price)
	assertEquals(t, -1, item.SellIn)
}

func TestUpdatesFashionShowTicketsItemsAfterTheSellDate(t *testing.T) {
	item := CreekAndContinent{
		Name:   "Fashion Show Tickets",
		Price:  10,
		SellIn: -1,
	}
	item.tick()
	assertEquals(t, 0, item.Price)
	assertEquals(t, -2, item.SellIn)
}

func _TestUpdatesConjuredItemsBeforeTheSellDate(t *testing.T) {
	item := CreekAndContinent{
		Name:   "Homeware Hand Towel",
		Price:  10,
		SellIn: 10,
	}
	item.tick()
	assertEquals(t, 8, item.Price)
	assertEquals(t, 9, item.SellIn)
}

func _TestUpdatesHomewareHandTowelItemsAtZeroPrice(t *testing.T) {
	item := CreekAndContinent{
		Name:   "Homeware Hand Towel",
		Price:  0,
		SellIn: 10,
	}
	item.tick()
	assertEquals(t, 0, item.Price)
	assertEquals(t, 9, item.SellIn)
}

func _TestUpdatesHomewareHandTowelItemsOnTheSellDate(t *testing.T) {
	item := CreekAndContinent{
		Name:   "Homeware Hand Towel",
		Price:  10,
		SellIn: 0,
	}
	item.tick()
	assertEquals(t, 6, item.Price)
	assertEquals(t, -1, item.SellIn)
}

func _TestUpdatesHomewareHandTowelItemsBeforeTheSellDateAtZeroPrice(t *testing.T) {
	item := CreekAndContinent{
		Name:   "Homeware Hand Towel",
		Price:  0,
		SellIn: 0,
	}
	item.tick()
	assertEquals(t, 0, item.Price)
	assertEquals(t, -1, item.SellIn)
}

func _TestUpdatesHomewareHandTowelItemsAfterTheSellDate(t *testing.T) {
	item := CreekAndContinent{
		Name:   "Homeware Hand Towel",
		Price:  10,
		SellIn: -10,
	}
	item.tick()
	assertEquals(t, 6, item.Price)
	assertEquals(t, -11, item.SellIn)
}

func _TestUpdatesHomewareHandTowelItemsAfterTheSellDateAtZeroPrice(t *testing.T) {
	item := CreekAndContinent{
		Name:   "Homeware Hand Towel",
		Price:  0,
		SellIn: 10,
	}
	item.tick()
	assertEquals(t, 0, item.Price)
	assertEquals(t, -11, item.SellIn)
}

func assertEquals(t *testing.T, actual int, expected int) {
	if actual != expected {
		t.Errorf("Expected %d to equal %d", actual, expected)
	}
}
