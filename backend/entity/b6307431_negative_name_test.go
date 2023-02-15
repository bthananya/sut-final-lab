package entity

import (
	"testing"

	"github.com/asaskevich/govalidator"
	"github.com/onsi.gomega"
	"gorm.io/gorm"
)

type Empolyee struct {
	gorm.Model
	Name       string `valid:"required~Name cannot be blank"`
	Email      string `valid:"email"`
	EmployeeID string `valid:"matches(^\\[JMS][0-9]{8}$) ~ รูปแบบรหัสพนักงานไม่ถูกต้อง"`
}

func TestEmpolyeeValidate(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	g.Expect(err).To(BeNil())


	t.Run("Name cannot be blank" , func(t *testing.T) {
		emp := Employee
			Name: "", //ผิด
			Email: "thny@gmail.com",
			EmployeeID: "J01234567",
	})
	ok, err := govalidator.ValidateStruct(emp)

	g.Expect(ok).NotTo(BeTrue())

	g.Expect(err).ToNot(BeNil())

	g.Expect(err.Error()).To(Equal("Name cannot be blank"))

		
}