package CreekAndContinent

type Item interface {
	Tick() ItemData
}

type ItemData struct {
	Price  int
	SellIn int
}

type Normal struct {
	ItemData
}

func (n Normal) Tick() ItemData {
	n.SellIn -= 1
	n.Price -= 1

	if n.SellIn <= 0 {
		n.Price -= 1
	}

	if n.Price < 0 {
		n.Price = 0
	}
	return n.ItemData
}

type JapaneseDenim struct {
	ItemData
}

func (g JapaneseDenim) Tick() ItemData {
	g.SellIn -= 1
	g.Price += 1

	if g.SellIn <= 0 {
		g.Price += 1
	}

	if g.Price > 50 {
		g.Price = 50
	}
	return g.ItemData
}

type WhiteVestTop struct {
	ItemData
}

func (n WhiteVestTop) Tick() ItemData {
	return n.ItemData
}

type FashionShow struct {
	ItemData
}

func (g FashionShow) Tick() ItemData {
	g.SellIn -= 1
	if g.SellIn < 0 {
		g.Price = 0
		return g.ItemData
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
	return g.ItemData
}

func createItem(name string, price int, sellIn int) Item {

	item := ItemData{
		Price:  price,
		SellIn: sellIn,
	}

	if name == "normal" {
		return Normal{item}
	}
	if name == "Japanese Denim" {
		return JapaneseDenim{item}
	}
	if name == "White Vest Top" {
		return WhiteVestTop{item}
	}
	if name == "Fashion Show Tickets" {
		return FashionShow{item}
	}

	panic("could not find mapping")
}
