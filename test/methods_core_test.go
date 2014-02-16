package testrdio

import (
	"testing"
    "fmt"
)

func TestGet(t *testing.T) {
	c := createClient(t)

	objects, err := c.Get([]string{})
	if err != nil {
		t.Fatal(err)
	}

	if len(objects) != 0 {
		t.Errorf("Expected 0 objects, got %d", len(objects))
	}

	objects, err = c.Get([]string{"r49021"})
	if err != nil {
		t.Fatal(err)
	}

	if len(objects) != 1 {
		t.Errorf("Expected 1 object, got %d", len(objects))
	}

	objects, err = c.Get([]string{"r49021", "l755"})
	if err != nil {
		t.Fatal(err)
	}

	if len(objects) != 2 {
		t.Errorf("Expected 2 objects, got %d", len(objects))
	}
}

func TestGetObjectFromShortCode(t *testing.T) {
	c := createClient(t)

	object, err := c.GetObjectFromShortCode("")
	if err == nil {
		t.Error("Expected API error, got none")
	}

	if object != nil {
		t.Errorf("Expected no object, got %s", object)
	}

	object, err = c.GetObjectFromShortCode("QFO0PnCM-A")
	if err != nil {
		t.Fatal(err)
	}

	if object == nil {
		t.Errorf("Expected object, got %s", object)
	}
}

func TestGetObjectFromUrl(t *testing.T) {
		fmt.Print("running!")
	c := createClient(t)

	object, err := c.GetObjectFromUrl("")
	if err == nil {
		t.Error("Expected API error, got none")
	}

	if object != nil {
		t.Errorf("Expected no object, got %s", object)
	}

	object, err = c.GetObjectFromUrl("/artist/Janelle_Mon%C3%A1e/album/The_Electric_Lady/")
	if err != nil {
		t.Fatal(err)
	}

		fmt.Print(object)

	if object == nil {
		t.Errorf("Expected object, got %s", object)
	}
}
