package miller_test

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/ninjadojo/miller"
)

func TestMiller(t *testing.T) {
	spew.Config.Indent = "\t"

	columns := miller.NewColumns([]string{"."})

	columns.Descend("test")
	lastCategory := columns.Categories[len(columns.Categories)-1]
	for _, item := range lastCategory.Items {
		if item.Name != "child" {
			t.Errorf("Expected child, got %s", item.Name)
			return
		}
	}

	columns.Descend("child")
	lastCategory = columns.Categories[len(columns.Categories)-1]
	for _, item := range lastCategory.Items {
		if item.Name != "anotherchild" {
			t.Errorf("Expected anotherchild, got %s", item.Name)
			return
		}
	}

	columns.Descend("anotherchild")
	lastCategory = columns.Categories[len(columns.Categories)-1]
	if len(lastCategory.Items) != 0 {
		t.Errorf("Expected no folders, got %d folders", len(lastCategory.Items))
		return
	}

	columns.Ascend()
	lastCategory = columns.Categories[len(columns.Categories)-1]
	for _, item := range lastCategory.Items {
		if item.Name != "anotherchild" {
			t.Errorf("Expected anotherchild, got %s", item.Name)
			return
		}
	}

	columns.Ascend()
	lastCategory = columns.Categories[len(columns.Categories)-1]
	for _, item := range lastCategory.Items {
		if item.Name != "child" {
			t.Errorf("Expected child, got %s", item.Name)
			return
		}
	}

	columns.Ascend()
	lastCategory = columns.Categories[len(columns.Categories)-1]
	for _, item := range lastCategory.Items {
		if item.Name != "test" && item.Name != ".git" {
			t.Errorf("Expected child, got %s", item.Name)
			return
		}
	}

	columns.Ascend()
	lastCategory = columns.Categories[len(columns.Categories)-1]
	for _, item := range lastCategory.Items {
		if item.Name != "test" && item.Name != ".git" {
			t.Errorf("Expected child, got %s", item.Name)
			return
		}
	}

}
