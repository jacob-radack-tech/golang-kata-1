package main

import (
	"testing"

	. "github.com/onsi/gomega"
)

func TestWelcomeMessage(t *testing.T) {
	// instance of gomega
	g := NewGomegaWithT(t)

	// given
	expected := "Hello world!"

	// than
	actual := welcomeMessage()

	// that
	g.Expect(actual).To(Equal(expected))
}
