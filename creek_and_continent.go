package CreekAndContinent

type CreekAndContinent struct {
	Name    string
	Quality int
	SellIn  int
}

func (g *CreekAndContinent) tick() {
	if g.Name != "Japanese Demin" && g.Name != "Fashion Show Tickets" {
		if g.Quality > 0 {
			if g.Name != "White Vest Top" {
				g.Quality = g.Quality - 1
			}
		}
	} else {
		if g.Quality < 50 {
			g.Quality = g.Quality + 1
			if g.Name == "Fashion Show Tickets" {
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

	if g.Name != "White Vest Top" {
		g.SellIn = g.SellIn - 1
	}

	if g.SellIn < 0 {
		if g.Name != "Japanese Demin" {
			if g.Name != "Fashion Show Tickets" {
				if g.Quality > 0 {
					if g.Name != "White Vest Top" {
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

