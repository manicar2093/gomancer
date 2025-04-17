package env_test

import (
	"github.com/manicar2093/gomancer/bootstrap/core/env"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Getenv", func() {
	It("should return data if exists", func() {
		envVar := "EXISTS"
		envVarValue := "true"
		GinkgoT().Setenv(envVar, envVarValue)

		got := env.GetEnvOrPanic(envVar)

		Expect(got).To(Equal(envVarValue))
	})

	It("should panics if data not exists", func() {
		envVar := "NOTEXIST"

		Expect(func() { env.GetEnvOrPanic(envVar) }).Should(Panic())

	})

})
