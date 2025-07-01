package main

import "testing"

func TestSearch(t *testing.T) {

	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("Testing Present Word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"

		assertStrings(t, got, want)
	})

	t.Run("Testing Not Present Word", func(t *testing.T) {
		_, got := dictionary.Search("found")
		if got == nil {
			t.Fatal("expected to get error")
		}

		assertError(t, got, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {

	t.Run("new word", func(t *testing.T) {
		dictonary := Dictionary{}
		word := "test"
		definition := "this is just a test"

		dictonary.Add(word, definition)

		assertDefinition(t, dictonary, word, definition)
	})

	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definintion := "this is just a test"
		dictionary := Dictionary{word: definintion}
		err := dictionary.Add(word, "new test")

		assertError(t, err, ErrWordExists)
		assertDefinition(t, dictionary, word, definintion)

	})
}

func TestUpdate(t *testing.T) {

	t.Run("update a word that exists", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}
		newDefinition := "new definition"

		err := dictionary.Update(word, newDefinition)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, newDefinition)
	})

	t.Run("Already Existing Definition", func(t *testing.T) {
		word := "test"
		def := "this is just a test"
		dictionary := Dictionary{}

		err := dictionary.Update(word, def)

		assertError(t, err, ErrWordDoesNotExist)
	})
}

func TestDelete(t *testing.T) {

	t.Run("remove word that exsits", func(t *testing.T) {
		word := "test"
		dictionary := Dictionary{word: "test definition"}

		dictionary.Delete(word)
		_, err := dictionary.Search(word)
		assertError(t, err, ErrNotFound)
	})

	t.Run("when a word does not exist", func(t *testing.T) {
		word := "test"
		dictionary := Dictionary{}

		err := dictionary.Delete(word)

		assertError(t, err, ErrWordDoesNotExist)
	})
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertDefinition(t testing.TB, dictionary Dictionary, word string, definition string) {
	t.Helper()

	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("should find added word ", err)
	}
	assertStrings(t, got, definition)
}
