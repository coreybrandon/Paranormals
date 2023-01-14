package models

type Product struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name"`
	Life int    `json:"life,omitempty"`
}

// type Tendency struct {
// 	Good    string `json:"good,omitempty"`
// 	Neutral string `json:"neutral,omitempty"`
// 	Evil    string `json:"evil,omitempty"`
// }
