package entity

// MenuCategory represents a section of the menu, like "MENU" or "OTHER".
// It contains a title and a list of MenuItems.
type MenuCategory struct {
	Title string     `json:"title"`
	Items []MenuItem `json:"items"`
}

// MenuItem represents a single link item within a menu category.
type MenuItem struct {
	Icon  string `json:"icon"`
	Label string `json:"label"`
	Href  string `json:"href"`
}
