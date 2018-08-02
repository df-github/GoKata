package set_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"dfiedler/kata/Set"
)

var _ = Describe("Set", func() {

	Describe("NewSet", func() {
		It("creates a new set", func() {
			set := set.NewSet()
			Expect(set).NotTo(BeNil())
		})
	})

	Describe("isEmpty", func() {
		It("returs true on empty set", func() {
			emptySet := set.NewSet()
			Expect(emptySet.IsEmpty()).To(BeTrue())
		})
	})

	Describe("Add", func() {
		It("adds one element to set", func() {
			one := set.NewSet()

			one.Add(1)

			Expect(one.IsEmpty()).To(BeFalse())
			Expect(one.Size()).To(Equal(1))
		})

		It("ensures identity of the added element", func() {
			set := set.NewSet()

			set.Add(1)
			set.Add(1)

			Expect(set.Size()).To(Equal(1))
		})

		It("adds elements of different types", func() {
			set := set.NewSet(1, "1", true)

			set.Add([]byte{})

			Expect(set.Size()).To(Equal(4))
		})
	})

	Describe("Delete", func() {
		It("does not panic on not found element", func() {
			emptySet := set.NewSet()

			Expect(func() { emptySet.Delete(1) }).ToNot(Panic())
		})

		It("removes element from a set", func() {
			one := set.NewSet("Hallo Welt", 1, "2", []byte{})

			one.Delete("2")

			Expect(one.Size()).To(Equal(3))
		})

	})

})
