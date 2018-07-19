package miller

import (
	"os"
	"path"
	"time"
)

// Columns contains the state of the miller columns
type Columns struct {
	CurrentFolder []string
	Categories    []*Category
}

// Category is a column
type Category struct {
	Path          []string        `json:"path"`
	CategoryName  string          `json:"categoryName"`
	IsLowestLevel bool            `json:"isLowestLevel"`
	Items         []*CategoryItem `json:"items"`
}

// CategoryItem are the items in a column
type CategoryItem struct {
	Name    string      `json:"name,omitempty"`
	Size    int64       `json:"size,omitempty"`
	Mode    os.FileMode `json:"mode,omitempty"`
	ModTime time.Time   `json:"modTime,omitempty"`
}

// NewColumns initialises the cascading list
func NewColumns(dir []string) (*Columns, error) {
	c := &Columns{
		Categories:    []*Category{},
		CurrentFolder: dir,
	}

	items, err := c.ListDir(dir)
	if err != nil {
		return nil, err
	}
	c.Categories = append(c.Categories, &Category{
		Items:         items,
		Path:          dir,
		IsLowestLevel: true,
	})
	return c, nil
}

// Descend moves down a level in the tree, appending the category and folder to the struct
func (c *Columns) Descend(dir string) error {
	args := append(c.CurrentFolder, dir)
	file, err := os.Open(path.Join(args...))
	if err != nil {
		return err
	}
	defer file.Close()

	c.CurrentFolder = append(c.CurrentFolder, dir)
	items, err := c.ListDir(c.CurrentFolder)
	if err != nil {
		return err
	}
	c.Categories = append(c.Categories, &Category{
		Items:         items,
		Path:          c.CurrentFolder,
		IsLowestLevel: false,
	})
	return nil
}

// Ascend moves up a level in the tree, removing the last category and folder in the struct
func (c *Columns) Ascend() {
	if len(c.Categories) == 1 {
		return
	}
	c.CurrentFolder = c.CurrentFolder[:len(c.CurrentFolder)-1]
	c.Categories = c.Categories[:len(c.Categories)-1]
}

// ListDir returns all the folders in directory
func (c *Columns) ListDir(directory []string) ([]*CategoryItem, error) {
	dir, err := os.Open(path.Join(directory...))
	if err != nil {
		return nil, err
	}
	defer dir.Close()

	entries, err := dir.Readdir(0)
	if err != nil {
		return nil, err
	}

	list := []*CategoryItem{}

	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}
		f := &CategoryItem{
			Name:    entry.Name(),
			Size:    entry.Size(),
			Mode:    entry.Mode(),
			ModTime: entry.ModTime(),
		}
		list = append(list, f)
	}
	return list, nil
}
