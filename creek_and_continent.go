package CreekAndContinent

type CreekAndContinent struct {
	Name   string
	Price  int
	SellIn int
}

func (g *CreekAndContinent) normalTick() {
	g.SellIn -= 1
	g.Price -= 1

	if g.SellIn <= 0 {
		g.Price -= 1
	}

	if g.Price < 0 {
		g.Price = 0
	}
}

func (g *CreekAndContinent) japaneseDenimTick() {
	g.SellIn -= 1
	g.Price += 1

	if g.SellIn <= 0 {
		g.Price += 1
	}

	if g.Price > 50 {
		g.Price = 50
	}
}

func (g *CreekAndContinent) whiteVestTopTick() {
	return
}

func (g *CreekAndContinent) fashionShowTick() {
	g.SellIn -= 1
	if g.SellIn < 0 {
		g.Price = 0
		return
	}
	g.Price += 1
	if g.SellIn < 10 {
		g.Price += 1
	}
	if g.SellIn < 5 {
		g.Price += 1
	}
	if g.Price > 50 {
		g.Price = 50
	}
}

func (g *CreekAndContinent) tick() {
	if g.Name == "normal" {
		g.normalTick()
		return
	}
	if g.Name == "Japanese Denim" {
		g.japaneseDenimTick()
		return
	}
	if g.Name == "White Vest Top" {
		g.whiteVestTopTick()
		return
	}
	if g.Name == "Fashion Show Tickets" {
		g.fashionShowTick()
		return
	}

	if g.Name != "Japanese Denim" && g.Name != "Fashion Show Tickets" {
		if g.Price > 0 {
			if g.Name != "White Vest Top" {
				g.Price = g.Price - 1
			}
		}
	} else {
		if g.Price < 50 {
			g.Price = g.Price + 1
			if g.Name == "Fashion Show Tickets" {
				if g.SellIn < 11 {
					if g.Price < 50 {
						g.Price = g.Price + 1
					}
				}
				if g.SellIn < 6 {
					if g.Price < 50 {
						g.Price = g.Price + 1
					}
				}
			}
		}
	}

	if g.Name != "White Vest Top" {
		g.SellIn = g.SellIn - 1
	}

	if g.SellIn < 0 {
		if g.Name != "Japanese Denim" {
			if g.Name != "Fashion Show Tickets" {
				if g.Price > 0 {
					if g.Name != "White Vest Top" {
						g.Price = g.Price - 1
					}
				}
			} else {
				g.Price = g.Price - g.Price
			}
		} else {
			if g.Price < 50 {
				g.Price = g.Price + 1
			}
		}
	}
}
