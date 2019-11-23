package newsfeed

import "testing"

// Execute go test -cover ./... for run testing
func TestAdd(t *testing.T) {
	feed := New()
	feed.Add(Item{})

	if len(feed.Items) != 1 {
		t.Errorf("Item was not added")
	}
}

func TestGetAll(t *testing.T) {
	feed := New()
	feed.Add(Item{})

	results := feed.GetAll()
	if len(results) != 1 {
		t.Errorf("Item was not added")
	}
}
