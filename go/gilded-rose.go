package main

type Item struct {
	name            string
	sellIn, quality int
}

const BRIE = "Aged Brie"
const SULFURAS = "Sulfuras, Hand of Ragnaros"
const BACKSTAGE = "Backstage passes to a TAFKAL80ETC concert"

var maxQuality = 50
var minQuality = 0

func (i *Item) newItemSellIn() int {
	if i.name == SULFURAS {
		return i.sellIn
	} else {
		return i.sellIn - 1
	}
}

func (i *Item) newItemQuality() int {
	switch name := i.name; name {
	case SULFURAS:
		return i.quality
	case BRIE:
		if i.sellIn < 0 {
			return i.quality + 2
		}
		return i.quality + 1
	case BACKSTAGE:
		if i.sellIn > 10 {
			return i.quality + 1
		} else if i.sellIn > 5 {
			return i.quality + 2
		} else if i.sellIn > 0 {
			return i.quality + 3
		} else {
			return 0
		}
	default:
		if i.sellIn < 0 {
			return i.quality - 2
		}
		return i.quality - 1
	}
}

func UpdateQuality(items []*Item) {
	for i := 0; i < len(items); i++ {

		items[i].sellIn = items[i].newItemSellIn()
		items[i].quality = items[i].newItemQuality()

		if items[i].quality < minQuality {
			items[i].quality = minQuality
		}
		if items[i].quality > maxQuality {
			items[i].quality = maxQuality
		}
	}
}
