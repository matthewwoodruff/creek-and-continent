package CreekAndContinent

type CreekAndContinent struct {
	Name   string
	Price  int
	SellIn int
}

func (g *CreekAndContinent) tick() {
	tick := createItem(g.Name, g.Price, g.SellIn).Tick()
	g.Price = tick.Price
	g.SellIn = tick.SellIn
}
