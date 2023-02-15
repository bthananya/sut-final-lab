package entity

import (
	"testing"

	"github.com/asaskevich/govalidator"
	"github.com/onsi.gomega"
	"gorm.io/gorm"
)

type Empolyee struct {
	gorm.Model
	Name       string `valid:"required"`
	Email      string `valid:"email"`
	EmployeeID string `valid:"matches(^\\[JMS][0-9]{8}$) ~ รูปแบบรหัสพนักงานไม่ถูกต้อง"`
}

func TestEmpolyeeValidate(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	t.Run("correct format" , func(t *testing.T) {
		emp := Employee
			Name: "Thananya",
			Email: "thny@gmail.com",
			EmployeeID: "J01234567",
	})
	ok, err := govalidator.ValidateStruct(emp)

	g.Expect(ok).To(BeTrue())

	g.Expect(err).To(BeNil())

	
	
}
