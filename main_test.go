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

	// testing
	g.Expect(actual).To(Equal(expected))
}
