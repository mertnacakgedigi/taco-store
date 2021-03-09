package memstore

type TacoStore struct {
	TacoList []Taco
}

func NewTacoStore() *TacoStore {
	return &TacoStore{TacoList: []Taco{
		{
			ID:          1,
			Title:       "Fried Avocado",
			Description: "Fried, creamy, crunchy, avocado-y goodness, shall we? These yummy tacos are super easy to make and they’re perfect for Meatless Monday, Taco Tuesday, or Get In Mah Belly Day.",
			ImageUrl:    "https://i.imgur.com/4v5UYxZ.jpg",
			Price:       299,
		},
		{
			ID:          2,
			Title:       "Blackberry Chicken",
			Description: "Shredded Hoisin-Blackberry Chicken Tacos with Crunchy Slaw. The chicken is cooked in the slow cooker until it’s so tender that it basically falls apart with no effort.",
			ImageUrl:    "https://i.imgur.com/46BwHbG.jpg",
			Price:       349,
		},
		{
			ID:          3,
			Title:       "Cilantro Lime Shrimp",
			Description: "Although the cilantro lime shrimp would be good in just about everything, they are particularly nice in these fresh summery tacos! ",
			ImageUrl:    "https://i.imgur.com/p7NGATP.jpg",
			Price:       399,
		},
		{
			ID:          4,
			Title:       "Mango Mahi-Mahi",
			Description: "Mahi-Mahi Fish Tacos with Mango and Avocado Salsa. If your not too familiar with mahi-mahi it’s a white fish that is a little firmer and perfect for fish tacos!",
			ImageUrl:    "https://i.imgur.com/bffmw0P.jpg",
			Price:       499,
		},
		{
			ID:          5,
			Title:       "Wild Caught Lobster",
			Description: "Fresh lobster with sweet mango, refreshing cucumber, and creamy avocado with a hint of lime, topped with bright cilantro—it's the ultimate summer taco!",
			ImageUrl:    "https://i.imgur.com/74xhgNW.jpg",
			Price:       599,
		},
		{
			ID:          6,
			Title:       "The Hungry Vegan ",
			Description: "These Meatless wonders will have you feeling great about your new years resolution. Mushroom and vegan beef crumbles with avocado and plenty of cilantro. Yummy!",
			ImageUrl:    "https://i.imgur.com/0Ddydf1.jpg",
			Price:       299,
		},
		{
			ID:          7,
			Title:       "Carne Asada",
			Description: "The undisputed king of the taco world. Classic with a twist... of lime! Flank steak marinated in our signature Mojo with homemade pico de gallo. Enough saID.",
			ImageUrl:    "https://i.imgur.com/DbE3iX8.jpg",
			Price:       349,
		},
	}}
}

type Taco struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	ImageUrl    string  `json:"image_url"`
	Price       float32 `json:"price"`
}

func (t *TacoStore) SelectTacoList() []Taco {
	return t.TacoList
}
