package main

import "testing"

func TestHello(t *testing.T) {
	checkcorrectmessage := func(t *testing.T, result, expected string) {
		t.Helper()
		if result != expected {
			t.Errorf("result '%s', expected '%s'", result, expected)
		}
	}

	t.Run("say hello to people", func(t *testing.T) {
		result := Hello("Chris", "")
		expected := "Hello, Chris"
		checkcorrectmessage(t, result, expected)
	})

	t.Run("in english", func(t *testing.T) {
		result := Hello("", "")
		expected := "Hello, World"
		checkcorrectmessage(t, result, expected)
	})

	t.Run("in espanhol", func(t *testing.T) {
		result := Hello("Elodie", "espanhol")
		expected := "Hola, Elodie"
		checkcorrectmessage(t, result, expected)
	})

	t.Run("in frances", func(t *testing.T) {
		result := Hello("Elodie", "frances")
		expected := "Bonjour, Elodie"
		checkcorrectmessage(t, result, expected)
	})
}
