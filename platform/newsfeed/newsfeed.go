package newsfeed

type Getter interface {
	GetAll() []Item
}

type Adder interface {
	Add(item Item)
}

// Struct Item with json binding
type Item struct {
	Title string `json:"title"`
	Post  string `json:"post"`
}

type Repo struct {
	// The type is a slice of Item struct
	Items []Item
}

func New() *Repo {
	return &Repo{
		Items: []Item{},
	}
}

func (r *Repo) Add(item Item) {
	r.Items = append(r.Items, item)
}

func (r *Repo) GetAll() []Item {
	return r.Items
}
