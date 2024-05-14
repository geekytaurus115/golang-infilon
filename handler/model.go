package handler

type Person struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Phone struct {
	Id       int    `json:"id"`
	Number   string `json:"number"`
	PersonId int    `json:"person_id"`
}

type Address struct {
	Id      int    `json:"id"`
	City    string `json:"city"`
	State   string `json:"state"`
	Street1 string `json:"street1"`
	Street2 string `json:"street2"`
	ZipCode string `json:"zip_code"`
}

type AddressJoin struct {
	Id        int `json:"id"`
	PersonId  int `json:"person_id"`
	AddressId int `json:"address_id"`
}

type PersonRequestBody struct {
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
	City        string `json:"city"`
	State       string `json:"state"`
	Street1     string `json:"street1"`
	Street2     string `json:"street2"`
	ZipCode     string `json:"zip_code"`
}
