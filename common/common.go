package common

type Users struct {
	ID           int    `json:"id,omitempty"`
	NamaDepan    string `json:"nama_depan,omitempty"`
	NamaBelakang string `json:"nama_belakang,omitempty"`
	Email        string `json:"email, omitempty"`
	Username     string `json:"username, omitempty"`
	Password     string `json:"password, omitempty"`
}

type Customers struct {
	CustomerID   string `json:"CustomerID, omitempty"`
	CompanyName  string `json:"CompanyName, omitempty"`
	ContactName  string `json:"ContactName, omitempty"`
	ContactTitle string `json:"ContactTitle, omitempty"`
	Address      string `json:"Address, omitempty"`
	City         string `json:"City, omitempty"`
	Country      string `json:"Country, omitempty"`
	Phone        string `json:"Phone, omitempty"`
	PostalCode   string `json:"PostalCode, omitempty"`
}

type Employees struct {
	LastName  string `json:"lastName"`
	FirstName string `json:"firstName"`
	Title     string `json:"title"`
	Address   string `json:"address"`
}
