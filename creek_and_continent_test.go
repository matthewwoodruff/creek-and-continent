package CreekAndContinent

import (
	"testing"
)

func TestUpdatesNormalItemsBeforeSellDate(t *testing.T) {
	item := CreekAndContinent{
		Name:    "normal",
		Quality: 10,
		SellIn:  5,
	}
	item.tick()
	assertEquals(t, 9, item.Quality)
	assertEquals(t, 4, item.SellIn)
}

func TestUpdatesNormalItemsOnTheSellDate(t *testing.T) {
	item := CreekAndContinent{
		Name:    "normal",
		Quality: 10,
		SellIn:  0,
	}
	item.tick()
	assertEquals(t, 8, item.Quality)
	assertEquals(t, -1, item.SellIn)
}

func TestUpdatesNormalItemsAfterTheSellDate(t *testing.T) {
	item := CreekAndContinent{
		Name:    "normal",
		Quality: 10,
		SellIn:  -5,
	}
	item.tick()
	assertEquals(t, 8, item.Quality)
	assertEquals(t, -6, item.SellIn)
}

func TestUpdatesNormalItemsWithQualityOf0(t *testing.T) {
	item := CreekAndContinent{
		Name:    "normal",
		Quality: 0,
		SellIn:  5,
	}
	item.tick()
	assertEquals(t, 0, item.Quality)
	assertEquals(t, 4, item.SellIn)
}

func TestUpdatesJapaneseDenimItemsBeforeTheSellDate(t *testing.T) {
	item := CreekAndContinent{
		Name:    "Japanese Demin",
		Quality: 10,
		SellIn:  5,
	}
	item.tick()
	assertEquals(t, 11, item.Quality)
	assertEquals(t, 4, item.SellIn)
}

func TestUpdatesJapaneseDenimItemsBeforeTheSellDateWithMaxQuality(t *testing.T) {
	item := CreekAndContinent{
		Name:    "Japanese Demin",
		Quality: 50,
		SellIn:  5,
	}
	item.tick()
	assertEquals(t, 50, item.Quality)
	assertEquals(t, 4, item.SellIn)
}

func TestUpdatesJapaneseDenimItemsOnTheSellDate(t *testing.T) {
	item := CreekAndContinent{
		Name:    "Japanese Demin",
		Quality: 10,
		SellIn:  0,
	}
	item.tick()
	assertEquals(t, 12, item.Quality)
	assertEquals(t, -1, item.SellIn)
}

func TestUpdatesJapaneseDenimItemsOnTheSellDateNearMaxQuality(t *testing.T) {
	item := CreekAndContinent{
		Name:    "Japanese Demin",
		Quality: 49,
		SellIn:  0,
	}
	item.tick()
	assertEquals(t, 50, item.Quality)
	assertEquals(t, -1, item.SellIn)
}

func TestUpdatesJapaneseDenimItemsOnTheSellDateWithMaxQuality(t *testing.T) {
	item := CreekAndContinent{
		Name:    "Japanese Demin",
		Quality: 50,
		SellIn:  0,
	}
	item.tick()
	assertEquals(t, 50, item.Quality)
	assertEquals(t, -1, item.SellIn)
}

func TestUpdatesJapaneseDenimItemsAfterTheSellDate(t *testing.T) {
	item := CreekAndContinent{
		Name:    "Japanese Demin",
		Quality: 10,
		SellIn:  -10,
	}
	item.tick()
	assertEquals(t, 12, item.Quality)
	assertEquals(t, -11, item.SellIn)
}

func TestUpdatesJapaneseDenimItemsAfterTheSellDateWithMaxQuality(t *testing.T) {
	item := CreekAndContinent{
		Name:    "Japanese Demin",
		Quality: 50,
		SellIn:  -10,
	}
	item.tick()
	assertEquals(t, 50, item.Quality)
	assertEquals(t, -11, item.SellIn)
}

func TestUpdatesWhiteVestTopItemsBeforeTheSellDate(t *testing.T) {
	item := CreekAndContinent{
		Name:    "White Vest Top",
		Quality: 10,
		SellIn:  5,
	}
	item.tick()
	assertEquals(t, 10, item.Quality)
	assertEquals(t, 5, item.SellIn)
}

func TestUpdatesWhiteVestTopItemsOnTheSellDate(t *testing.T) {
	item := CreekAndContinent{
		Name:    "White Vest Top",
		Quality: 10,
		SellIn:  5,
	}
	item.tick()
	assertEquals(t, 10, item.Quality)
	assertEquals(t, 5, item.SellIn)
}

func TestUpdatesWhiteVestTopItemsAfterTheSellDate(t *testing.T) {
	item := CreekAndContinent{
		Name:    "White Vest Top",
		Quality: 10,
		SellIn:  -1,
	}
	item.tick()
	assertEquals(t, 10, item.Quality)
	assertEquals(t, -1, item.SellIn)
}

func TestUpdatesFashionShowTicketsItemsBeforeTheSellDate(t *testing.T) {
	item := CreekAndContinent{
		Name:    "Fashion Show Tickets",
		Quality: 10,
		SellIn:  11,
	}
	item.tick()
	assertEquals(t, 11, item.Quality)
	assertEquals(t, 10, item.SellIn)
}

func TestUpdatesFashionShowTicketsItemsCloseToTheSellDate(t *testing.T) {
	item := CreekAndContinent{
		Name:    "Fashion Show Tickets",
		Quality: 10,
		SellIn:  10,
	}
	item.tick()
	assertEquals(t, 12, item.Quality)
	assertEquals(t, 9, item.SellIn)
}

func TestUpdatesFashionShowTicketsItemsBeforeTheSellDateAtMaxQuality(t *testing.T) {
	item := CreekAndContinent{
		Name:    "Fashion Show Tickets",
		Quality: 50,
		SellIn:  10,
	}
	item.tick()
	assertEquals(t, 50, item.Quality)
	assertEquals(t, 9, item.SellIn)
}

func TestUpdatesFashionShowTicketsItemsVeryCloseToTheSellDate(t *testing.T) {
	item := CreekAndContinent{
		Name:    "Fashion Show Tickets",
		Quality: 10,
		SellIn:  5,
	}
	item.tick()
	assertEquals(t, 13, item.Quality)
	assertEquals(t, 4, item.SellIn)
}

func TestUpdatesFashionShowTicketsItemsVeryCloseToTheSellDateAtMaxQuality(t *testing.T) {
	item := CreekAndContinent{
		Name:    "Fashion Show Tickets",
		Quality: 50,
		SellIn:  5,
	}
	item.tick()
	assertEquals(t, 50, item.Quality)
	assertEquals(t, 4, item.SellIn)
}

func TestUpdatesFashionShowTicketsItemsWithOneDayLeftToSell(t *testing.T) {
	item := CreekAndContinent{
		Name:    "Fashion Show Tickets",
		Quality: 10,
		SellIn:  1,
	}
	item.tick()
	assertEquals(t, 13, item.Quality)
	assertEquals(t, 0, item.SellIn)
}

func TestUpdatesFashionShowTicketsItemsWithOneDayLeftToSellAtMaxQuality(t *testing.T) {
	item := CreekAndContinent{
		Name:    "Fashion Show Tickets",
		Quality: 50,
		SellIn:  1,
	}
	item.tick()
	assertEquals(t, 50, item.Quality)
	assertEquals(t, 0, item.SellIn)
}

func TestUpdatesFashionShowTicketsItemsOnSellDate(t *testing.T) {
	item := CreekAndContinent{
		Name:    "Fashion Show Tickets",
		Quality: 10,
		SellIn:  0,
	}
	item.tick()
	assertEquals(t, 0, item.Quality)
	assertEquals(t, -1, item.SellIn)
}

func TestUpdatesFashionShowTicketsItemsAfterTheSellDate(t *testing.T) {
	item := CreekAndContinent{
		Name:    "Fashion Show Tickets",
		Quality: 10,
		SellIn:  -1,
	}
	item.tick()
	assertEquals(t, 0, item.Quality)
	assertEquals(t, -2, item.SellIn)
}

func _TestUpdatesConjuredItemsBeforeTheSellDate(t *testing.T) {
	item := CreekAndContinent{
		Name:    "Homeware Hand Towel",
		Quality: 10,
		SellIn:  10,
	}
	item.tick()
	assertEquals(t, 8, item.Quality)
	assertEquals(t, 9, item.SellIn)
}

func _TestUpdatesHomewareHandTowelItemsAtZeroQuality(t *testing.T) {
	item := CreekAndContinent{
		Name:    "Homeware Hand Towel",
		Quality: 0,
		SellIn:  10,
	}
	item.tick()
	assertEquals(t, 0, item.Quality)
	assertEquals(t, 9, item.SellIn)
}

func _TestUpdatesHomewareHandTowelItemsOnTheSellDate(t *testing.T) {
	item := CreekAndContinent{
		Name:    "Homeware Hand Towel",
		Quality: 10,
		SellIn:  0,
	}
	item.tick()
	assertEquals(t, 6, item.Quality)
	assertEquals(t, -1, item.SellIn)
}

func _TestUpdatesHomewareHandTowelItemsBeforeTheSellDateAtZeroQuality(t *testing.T) {
	item := CreekAndContinent{
		Name:    "Homeware Hand Towel",
		Quality: 0,
		SellIn:  0,
	}
	item.tick()
	assertEquals(t, 0, item.Quality)
	assertEquals(t, -1, item.SellIn)
}

func _TestUpdatesHomewareHandTowelItemsAfterTheSellDate(t *testing.T) {
	item := CreekAndContinent{
		Name:    "Homeware Hand Towel",
		Quality: 10,
		SellIn:  -10,
	}
	item.tick()
	assertEquals(t, 6, item.Quality)
	assertEquals(t, -11, item.SellIn)
}

func _TestUpdatesHomewareHandTowelItemsAfterTheSellDateAtZeroQuality(t *testing.T) {
	item := CreekAndContinent{
		Name:    "Homeware Hand Towel",
		Quality: 0,
		SellIn:  10,
	}
	item.tick()
	assertEquals(t, 0, item.Quality)
	assertEquals(t, -11, item.SellIn)
}

func assertEquals(t *testing.T, actual int, expected int) {
	if actual != expected {
		t.Errorf("Expected %d to equal %d", actual, expected)
	}
}
