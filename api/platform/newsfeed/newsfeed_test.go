package newsfeed

import "testing"

func TestRepo_Add(t *testing.T) {
	feed := New()
	feed.Add(Item{})

	if len(feed.Items) != 1 {
		t.Error("Item was not added!")
	}
}

func TestRepo_GetAll(t *testing.T) {
	feed := New()
	feed.Add(Item{})

	results := feed.GetAll()

	if len(results) != 1 {
		t.Error("Items were not obtained!")
	}
}
