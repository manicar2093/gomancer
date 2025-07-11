package validator_test

import (
	"github.com/manicar2093/gomancer/bootstrap/core/apperrors"
	"github.com/manicar2093/gomancer/bootstrap/core/validator"
	"net/http"

	"github.com/go-faker/faker/v4"
	"github.com/google/uuid"
	"github.com/manicar2093/goption"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Gookitvalidator", func() {

	Describe("StructValidator", func() {

		It("returns a list of errors if any exists", func() {
			expectedDataToValidate := struct {
				Name        string                   `validate:"required|min_len:7" json:"name,omitempty"`
				LastName    goption.Optional[string] `validate:"required" json:"last_name,omitempty"`
				LastNamePtr string                   `validate:"required" json:"last_name_ptr,omitempty"`
				Pass        string                   `json:"pass" validate:"eq_field:NewPass"`
				NewPass     string                   `json:"new_pass"`
			}{
				Pass:    "hello",
				NewPass: "hello",
			}

			got := validator.ValidateAndTransform(validator.NewStructValidatorConfigured(&expectedDataToValidate, "en"))

			Expect(got).ToNot(BeNil())
			Expect(got.(*validator.ValidationError).Errors).To(HaveLen(3))
			Expect(got.(apperrors.HandleableError).StatusCode()).To(Equal(http.StatusBadRequest))
		})

		When("there is any error", func() {
			It("returns nil", func() {
				expectedDataToValidate := struct {
					Name string `validate:"required|min_len:7" json:"name,omitempty"`
				}{
					Name: faker.Name(),
				}

				got := validator.ValidateAndTransform(validator.NewStructValidatorConfigured(&expectedDataToValidate, "en"))

				Expect(got).To(BeNil())
			})
		})

		It("validates google uuid.UUID is required", func() {
			var (
				expectedDataToValidate = struct {
					AnId uuid.UUID `validate:"required_uuid" json:"an_id"`
				}{}
			)
			got := validator.ValidateAndTransform(validator.NewStructValidatorConfigured(&expectedDataToValidate, "en"))

			err := got.(*validator.ValidationError)
			errMap := err.Errors
			Expect(errMap["an_id"]["required_uuid"]).To(ContainSubstring("needs to be on request"))
		})
	})

})
