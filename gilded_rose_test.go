package GildedRose

import (
	"testing"
)

func TestUpdatesNormalItemsBeforeSellDate(t *testing.T) {
	item := GildedRose{
		Name:    "normal",
		Quality: 10,
		SellIn:  5,
	}
	item.tick()
	assertEquals(t, 9, item.Quality)
	assertEquals(t, 4, item.SellIn)
}

func TestUpdatesNormalItemsOnTheSellDate(t *testing.T) {
	item := GildedRose{
		Name:    "normal",
		Quality: 10,
		SellIn:  0,
	}
	item.tick()
	assertEquals(t, 8, item.Quality)
	assertEquals(t, -1, item.SellIn)
}

func TestUpdatesNormalItemsAfterTheSellDate(t *testing.T) {
	item := GildedRose{
		Name:    "normal",
		Quality: 10,
		SellIn:  -5,
	}
	item.tick()
	assertEquals(t, 8, item.Quality)
	assertEquals(t, -6, item.SellIn)
}

func TestUpdatesNormalItemsWithQualityOf0(t *testing.T) {
	item := GildedRose{
		Name:    "normal",
		Quality: 0,
		SellIn:  5,
	}
	item.tick()
	assertEquals(t, 0, item.Quality)
	assertEquals(t, 4, item.SellIn)
}

func TestUpdatesBrieItemsBeforeTheSellDate(t *testing.T) {
	item := GildedRose{
		Name:    "Aged Brie",
		Quality: 10,
		SellIn:  5,
	}
	item.tick()
	assertEquals(t, 11, item.Quality)
	assertEquals(t, 4, item.SellIn)
}

func TestUpdatesBrieItemsBeforeTheSellDateWithMaxQuality(t *testing.T) {
	item := GildedRose{
		Name:    "Aged Brie",
		Quality: 50,
		SellIn:  5,
	}
	item.tick()
	assertEquals(t, 50, item.Quality)
	assertEquals(t, 4, item.SellIn)
}

func TestUpdatesBrieItemsOnTheSellDate(t *testing.T) {
	item := GildedRose{
		Name:    "Aged Brie",
		Quality: 10,
		SellIn:  0,
	}
	item.tick()
	assertEquals(t, 12, item.Quality)
	assertEquals(t, -1, item.SellIn)
}

func TestUpdatesBrieItemsOnTheSellDateNearMaxQuality(t *testing.T) {
	item := GildedRose{
		Name:    "Aged Brie",
		Quality: 49,
		SellIn:  0,
	}
	item.tick()
	assertEquals(t, 50, item.Quality)
	assertEquals(t, -1, item.SellIn)
}

func TestUpdatesBrieItemsOnTheSellDateWithMaxQuality(t *testing.T) {
	item := GildedRose{
		Name:    "Aged Brie",
		Quality: 50,
		SellIn:  0,
	}
	item.tick()
	assertEquals(t, 50, item.Quality)
	assertEquals(t, -1, item.SellIn)
}

func TestUpdatesBrieItemsAfterTheSellDate(t *testing.T) {
	item := GildedRose{
		Name:    "Aged Brie",
		Quality: 10,
		SellIn:  -10,
	}
	item.tick()
	assertEquals(t, 12, item.Quality)
	assertEquals(t, -11, item.SellIn)
}

func TestUpdatesBrieItemsAfterTheSellDateWithMaxQuality(t *testing.T) {
	item := GildedRose{
		Name:    "Aged Brie",
		Quality: 50,
		SellIn:  -10,
	}
	item.tick()
	assertEquals(t, 50, item.Quality)
	assertEquals(t, -11, item.SellIn)
}

func TestUpdatesSulfurasItemsBeforeTheSellDate(t *testing.T) {
	item := GildedRose{
		Name:    "Sulfuras, Hand of Ragnaros",
		Quality: 10,
		SellIn:  5,
	}
	item.tick()
	assertEquals(t, 10, item.Quality)
	assertEquals(t, 5, item.SellIn)
}

func TestUpdatesSulfurasItemsOnTheSellDate(t *testing.T) {
	item := GildedRose{
		Name:    "Sulfuras, Hand of Ragnaros",
		Quality: 10,
		SellIn:  5,
	}
	item.tick()
	assertEquals(t, 10, item.Quality)
	assertEquals(t, 5, item.SellIn)
}

func TestUpdatesSulfurasItemsAfterTheSellDate(t *testing.T) {
	item := GildedRose{
		Name:    "Sulfuras, Hand of Ragnaros",
		Quality: 10,
		SellIn:  -1,
	}
	item.tick()
	assertEquals(t, 10, item.Quality)
	assertEquals(t, -1, item.SellIn)
}

func TestUpdatesBackstagePassItemsBeforeTheSellDate(t *testing.T) {
	item := GildedRose{
		Name:    "Backstage passes to a TAFKAL80ETC concert",
		Quality: 10,
		SellIn:  11,
	}
	item.tick()
	assertEquals(t, 11, item.Quality)
	assertEquals(t, 10, item.SellIn)
}

func TestUpdatesBackstagePassItemsCloseToTheSellDate(t *testing.T) {
	item := GildedRose{
		Name:    "Backstage passes to a TAFKAL80ETC concert",
		Quality: 10,
		SellIn:  10,
	}
	item.tick()
	assertEquals(t, 12, item.Quality)
	assertEquals(t, 9, item.SellIn)
}

func TestUpdatesBackstagePassItemsBeforeTheSellDateAtMaxQuality(t *testing.T) {
	item := GildedRose{
		Name:    "Backstage passes to a TAFKAL80ETC concert",
		Quality: 50,
		SellIn:  10,
	}
	item.tick()
	assertEquals(t, 50, item.Quality)
	assertEquals(t, 9, item.SellIn)
}

func TestUpdatesBackstagePassItemsVeryCloseToTheSellDate(t *testing.T) {
	item := GildedRose{
		Name:    "Backstage passes to a TAFKAL80ETC concert",
		Quality: 10,
		SellIn:  5,
	}
	item.tick()
	assertEquals(t, 13, item.Quality)
	assertEquals(t, 4, item.SellIn)
}

func TestUpdatesBackstagePassItemsVeryCloseToTheSellDateAtMaxQuality(t *testing.T) {
	item := GildedRose{
		Name:    "Backstage passes to a TAFKAL80ETC concert",
		Quality: 50,
		SellIn:  5,
	}
	item.tick()
	assertEquals(t, 50, item.Quality)
	assertEquals(t, 4, item.SellIn)
}

func TestUpdatesBackstagePassItemsWithOneDayLeftToSell(t *testing.T) {
	item := GildedRose{
		Name:    "Backstage passes to a TAFKAL80ETC concert",
		Quality: 10,
		SellIn:  1,
	}
	item.tick()
	assertEquals(t, 13, item.Quality)
	assertEquals(t, 0, item.SellIn)
}

func TestUpdatesBackstagePassItemsWithOneDayLeftToSellAtMaxQuality(t *testing.T) {
	item := GildedRose{
		Name:    "Backstage passes to a TAFKAL80ETC concert",
		Quality: 50,
		SellIn:  1,
	}
	item.tick()
	assertEquals(t, 50, item.Quality)
	assertEquals(t, 0, item.SellIn)
}

func TestUpdatesBackstagePassItemsOnSellDate(t *testing.T) {
	item := GildedRose{
		Name:    "Backstage passes to a TAFKAL80ETC concert",
		Quality: 10,
		SellIn:  0,
	}
	item.tick()
	assertEquals(t, 0, item.Quality)
	assertEquals(t, -1, item.SellIn)
}

func TestUpdatesBackstagePassItemsAfterTheSellDate(t *testing.T) {
	item := GildedRose{
		Name:    "Backstage passes to a TAFKAL80ETC concert",
		Quality: 10,
		SellIn:  -1,
	}
	item.tick()
	assertEquals(t, 0, item.Quality)
	assertEquals(t, -2, item.SellIn)
}

func _TestUpdatesConjuredItemsBeforeTheSellDate(t *testing.T) {
	item := GildedRose{
		Name:    "Conjured Mana Cake",
		Quality: 10,
		SellIn:  10,
	}
	item.tick()
	assertEquals(t, 8, item.Quality)
	assertEquals(t, 9, item.SellIn)
}

func _TestUpdatesConjuredItemsAtZeroQuality(t *testing.T) {
	item := GildedRose{
		Name:    "Conjured Mana Cake",
		Quality: 0,
		SellIn:  10,
	}
	item.tick()
	assertEquals(t, 0, item.Quality)
	assertEquals(t, 9, item.SellIn)
}

func _TestUpdatesConjuredItemsOnTheSellDate(t *testing.T) {
	item := GildedRose{
		Name:    "Conjured Mana Cake",
		Quality: 10,
		SellIn:  0,
	}
	item.tick()
	assertEquals(t, 6, item.Quality)
	assertEquals(t, -1, item.SellIn)
}

func _TestUpdatesConjuredItemsBeforeTheSellDateAtZeroQuality(t *testing.T) {
	item := GildedRose{
		Name:    "Conjured Mana Cake",
		Quality: 0,
		SellIn:  0,
	}
	item.tick()
	assertEquals(t, 0, item.Quality)
	assertEquals(t, -1, item.SellIn)
}

func _TestUpdatesConjuredItemsAfterTheSellDate(t *testing.T) {
	item := GildedRose{
		Name:    "Conjured Mana Cake",
		Quality: 10,
		SellIn:  -10,
	}
	item.tick()
	assertEquals(t, 6, item.Quality)
	assertEquals(t, -11, item.SellIn)
}

func _TestUpdatesConjuredItemsAfterTheSellDateAtZeroQuality(t *testing.T) {
	item := GildedRose{
		Name:    "Conjured Mana Cake",
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
