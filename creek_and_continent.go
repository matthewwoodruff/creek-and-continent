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

func (g *CreekAndContinent) tick() {
	if g.Name == "normal" {
		g.normalTick()
		return
	}

	if g.Name != "Japanese Demin" && g.Name != "Fashion Show Tickets" {
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
		if g.Name != "Japanese Demin" {
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
