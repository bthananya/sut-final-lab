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

	t.Run("EmployeeID invalid format" , func(t *testing.T) {
		emp := Employee
			Name: "Thananya",
			Email: "thny@gmail.com",
			EmployeeID: "B630", //ผิด
	})
	ok, err := govalidator.ValidateStruct(emp)

	g.Expect(ok).NotTo(BeTrue())

	g.Expect(err).ToNot(BeNil())

	g.Expect(err.Error()).To(Equal("รูปแบบรหัสพนักงานไม่ถูกต้อง"))
}
