package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("UpdateQuality", func() {
	var (
		items   []*Item
		quality int
		sellIn  int
	)

	Context("multiple items", func() {
		JustBeforeEach(func() {
			items = []*Item{
				&Item{"item1", 10, 1},
				&Item{"item2", 20, 2},
			}
			UpdateQuality(items)
		})

		It("updates the quality and sellIn date of each item", func() {
			Expect(items[0].quality).To(Equal(0))
			Expect(items[1].quality).To(Equal(1))
			Expect(items[0].sellIn).To(Equal(9))
			Expect(items[1].sellIn).To(Equal(19))
		})

		It("does not change the name", func() {
			Expect(items[0].name).To(Equal("item1"))
			Expect(items[1].name).To(Equal("item2"))
		})
	})

	Context("regular items", func() {
		BeforeEach(func() {
			quality = 20
			sellIn = 10
		})
		JustBeforeEach(func() {
			items = []*Item{
				&Item{"foo", sellIn, quality},
			}
			UpdateQuality(items)
		})

		Context("before sellIn date", func() {
			BeforeEach(func() {
				sellIn = 10
			})
			It("decreases 1 in quality and sellIn date", func() {
				Expect(items[0].quality).To(Equal(19))
				Expect(items[0].sellIn).To(Equal(9))
			})
		})

		Context("after sellIn date", func() {
			BeforeEach(func() {
				sellIn = 0
			})

			It("decreases 1 in sellIn date, decreases 2 in quality", func() {
				Expect(items[0].quality).To(Equal(quality - 2))
				Expect(items[0].sellIn).To(Equal(sellIn - 1))
			})
		})

		Context("quality is zero", func() {
			BeforeEach(func() {
				quality = 0
			})

			It("stays at zero quality", func() {
				Expect(items[0].quality).To(Equal(0))
			})
		})
	})
	Context("Aged Brie items", func() {
		BeforeEach(func() {
			quality = 20
			sellIn = 10
		})
		JustBeforeEach(func() {
			items = []*Item{
				&Item{"Aged Brie", sellIn, quality},
			}
			UpdateQuality(items)
		})

		Context("before sellIn date", func() {
			BeforeEach(func() {
				sellIn = 10
			})
			It("increases 1 in quality and decreases 1 in sellIn date", func() {
				Expect(items[0].quality).To(Equal(quality + 1))
				Expect(items[0].sellIn).To(Equal(sellIn - 1))
			})
		})

		Context("after sellIn date", func() {
			BeforeEach(func() {
				sellIn = 0
			})

			It("increases 2 in quality and decreases 1 in sellIn date", func() {
				Expect(items[0].quality).To(Equal(22))
				Expect(items[0].sellIn).To(Equal(sellIn - 1))
			})
		})

		Context("quality reaches max", func() {
			BeforeEach(func() {
				quality = 50
			})

			It("stays at max quality", func() {
				Expect(items[0].quality).To(Equal(quality))
			})
		})
	})
	Context("Sulfuras items", func() {
		BeforeEach(func() {
			quality = 20
			sellIn = 10
		})
		JustBeforeEach(func() {
			items = []*Item{
				&Item{"Sulfuras, Hand of Ragnaros", sellIn, quality},
			}
			UpdateQuality(items)
		})

		Context("before sellIn date", func() {
			BeforeEach(func() {
				sellIn = 10
			})
			It("stays the same in quality and sellIn date", func() {
				Expect(items[0].quality).To(Equal(quality))
				Expect(items[0].sellIn).To(Equal(sellIn))
			})
		})

		Context("after sellIn date", func() {
			BeforeEach(func() {
				sellIn = 0
			})

			It("stays the same in quality and sellIn date", func() {
				Expect(items[0].quality).To(Equal(quality))
				Expect(items[0].sellIn).To(Equal(sellIn))
			})
		})
	})
	Context("backstage ticket items", func() {
		BeforeEach(func() {
			quality = 20
			sellIn = 10
		})
		JustBeforeEach(func() {
			items = []*Item{
				&Item{"Backstage passes to a TAFKAL80ETC concert", sellIn, quality},
			}
			UpdateQuality(items)
		})

		Context("more than 10 days before sellIn date", func() {
			BeforeEach(func() {
				sellIn = 30
			})
			It("increases by 1 in quality and decreases sellIn date", func() {
				Expect(items[0].quality).To(Equal(quality + 1))
				Expect(items[0].sellIn).To(Equal(sellIn - 1))
			})
		})
		Context("10 days or less before sellIn date", func() {
			BeforeEach(func() {
				sellIn = 10
			})
			It("increasess 2 in quality and decreases sellIn date", func() {
				Expect(items[0].quality).To(Equal(quality + 2))
				Expect(items[0].sellIn).To(Equal(sellIn - 1))
			})
		})
		Context("5 days or less before sellIn date", func() {
			BeforeEach(func() {
				sellIn = 5
			})
			It("increasess 3 in quality and decreases sellIn date", func() {
				Expect(items[0].quality).To(Equal(quality + 3))
				Expect(items[0].sellIn).To(Equal(sellIn - 1))
			})
		})

		Context("at max quality", func() {
			BeforeEach(func() {
				sellIn = 5
				quality = 50
			})

			It("quality stays at the max quality", func() {
				Expect(items[0].quality).To(Equal(50))
				Expect(items[0].sellIn).To(Equal(sellIn - 1))
			})
		})
		Context("after sellIn date", func() {
			BeforeEach(func() {
				sellIn = 0
			})

			It("its quality drops to zero", func() {
				Expect(items[0].quality).To(Equal(0))
				Expect(items[0].sellIn).To(Equal(sellIn - 1))
			})
		})
	})
})
