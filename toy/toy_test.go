package toy

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Toy", func() {
	var (
		testBox *box
		height, width, length int64
	)

	Context("box", func() {
		Context("NewBox", func() {
			Context("without no sizes specified", func() {
				BeforeEach(func() {
					testBox = NewBox()
				})

				It("should create a zero size box", func() {
					var z int64
					Expect(testBox.height).To(Equal(z))
					Expect(testBox.width).To(Equal(z))
					Expect(testBox.length).To(Equal(z))
				})
			})

			Context("with sizes specified", func() {
				BeforeEach(func() {
					height = 10
					width = 10
					length = 10
					testBox = NewBox(height, width, length)
				})

				It("should create a box with the expected dimensions", func() {
					Expect(testBox.height).To(Equal(height))
					Expect(testBox.width).To(Equal(width))
					Expect(testBox.length).To(Equal(length))
				})

				It("should not create a full box", func() {
					Expect(testBox.isFull).To(BeFalse())
				})
			})
		})
	})

	Context(".Fill", func() {
		var err error

		BeforeEach(func() {
			height = 10
			width = 10
			length = 10
			testBox = &box{height: height, width: width, length: length}
		})

		Context("when not full", func() {
			BeforeEach(func() {
				testBox.isFull = false
				err = testBox.Fill()
			})

			It("should fill the box", func() {
				Expect(err).ToNot(HaveOccurred())
				Expect(testBox.isFull).To(BeTrue())
			})
		})

		Context("when full", func() {
			BeforeEach(func() {
				testBox.isFull = true
				err = testBox.Fill()
			})

			It("should error", func() {
				Expect(err).To(MatchError(ErrFull))
			})
		})
	})

	Context(".Empty", func() {
		var err error

		BeforeEach(func() {
			height = 10
			width = 10
			length = 10
			testBox = &box{height: height, width: width, length: length}
		})

		Context("when note empty", func() {
			BeforeEach(func() {
				testBox.isFull = true
				err = testBox.Empty()
			})

			It("should empty the box", func() {
				Expect(err).ToNot(HaveOccurred())
				Expect(testBox.isFull).To(BeFalse())
			})
		})

		Context("when empty", func() {
			BeforeEach(func() {
				testBox.isFull = false
				err = testBox.Empty()
			})

			It("should error", func() {
				Expect(err).To(MatchError(ErrEmpty))
			})
		})
	})

	Context(".Size", func() {
		BeforeEach(func() {
			height = 24
			width = 12
			length = 42
			testBox = &box{height: height, width: width, length: length}
		})

		It("should return the expected dimensions", func() {
			Expect(testBox.height).To(Equal(height))
			Expect(testBox.width).To(Equal(width))
			Expect(testBox.length).To(Equal(length))
		})
	})
})
