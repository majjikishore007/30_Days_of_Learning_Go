package day6

import "testing"

func asserStrings(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("%q want and %q got", want, got)
	}

}
func assertError(t *testing.T, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}
func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"} // dictionary := map[string]string{"test": "this is just a test"}

	t.Run("know word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"
		asserStrings(t, got, want)
	})
	t.Run("unknow word", func(t *testing.T) {

		_, got := dictionary.Search("hey")
		assertError(t, got, ErrorNotFound)
	})
}

func TestAdd(t *testing.T) {
	dictionary := Dictionary{}
	dictionary.Add("test", "This is just a test")
	want := "This is just a test"
	got, err := dictionary.Search("test")
	if err != nil {
		t.Fatal("Should find added word", err)
	}
	if got != want {
		t.Errorf("got %q and want %q", got, want)
	}
}
func TestDelete(t *testing.T) {
	word := "test"
	dictionary := Dictionary{word: "test definition"}

	dictionary.Delete(word)

	_, err := dictionary.Search(word)
	if err != ErrorNotFound {
		t.Errorf("Expected %q to be deleted", word)
	}
}
