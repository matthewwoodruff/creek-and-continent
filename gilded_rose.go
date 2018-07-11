package GildedRose

type GildedRose struct {
	Name    string
	Quality int
	SellIn  int
}

func (g *GildedRose) tick() {
	if g.Name != "Aged Brie" && g.Name != "Backstage passes to a TAFKAL80ETC concert" {
		if g.Quality > 0 {
			if g.Name != "Sulfuras, Hand of Ragnaros" {
				g.Quality = g.Quality - 1
			}
		}
	} else {
		if g.Quality < 50 {
			g.Quality = g.Quality + 1
			if g.Name == "Backstage passes to a TAFKAL80ETC concert" {
				if g.SellIn < 11 {
					if g.Quality < 50 {
						g.Quality = g.Quality + 1
					}
				}
				if g.SellIn < 6 {
					if g.Quality < 50 {
						g.Quality = g.Quality + 1
					}
				}
			}
		}
	}

	if g.Name != "Sulfuras, Hand of Ragnaros" {
		g.SellIn = g.SellIn - 1
	}

	if g.SellIn < 0 {
		if g.Name != "Aged Brie" {
			if g.Name != "Backstage passes to a TAFKAL80ETC concert" {
				if g.Quality > 0 {
					if g.Name != "Sulfuras, Hand of Ragnaros" {
						g.Quality = g.Quality - 1
					}
				}
			} else {
				g.Quality = g.Quality - g.Quality
			}
		} else {
			if g.Quality < 50 {
				g.Quality = g.Quality + 1
			}
		}
	}
}

