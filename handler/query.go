package handler

import (
	"database/sql"
	"fmt"
	"log"
)

func GetPersonByIdQuery(personID string) (*Person, error) {
	var person Person

	var age sql.NullInt64
	err := db.QueryRow("SELECT id, name, age FROM person WHERE id = ?", personID).Scan(&person.Id,
		&person.Name, &age)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("Person not found: %v", err)
		}
		return nil, err
	}

	if age.Valid {
		person.Age = int(age.Int64)
	} else {
		log.Println("Age is null")
	}

	return &person, nil

}

func GetPhoneByPersonIdQuery(personID string) (*Phone, error) {
	var phone Phone
	err := db.QueryRow("SELECT id, number, person_id FROM phone WHERE person_id = ?", personID).Scan(&phone.Id,
		&phone.Number, &phone.PersonId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("Phone not found: %v", err)
		}
		return nil, err
	}
	return &phone, nil
}

func GetAddressByPersonIdQuery(personId string) (*Address, error) {
	var address Address

	err := db.QueryRow(`SELECT a.id, a.city, a.state, a.street1, a.street2, a.zip_code 
		FROM address a JOIN address_join aj ON a.id = aj.address_id 
			WHERE aj.person_id = ?`, personId).Scan(&address.Id,
		&address.City, &address.State, &address.Street1, &address.Street2, &address.ZipCode)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("Address not found: %v", err)
		}
		return nil, err
	}
	return &address, nil
}

func CreateNewPersonQuery(name string) (int, error) {
	res, err := db.Exec("INSERT INTO person(name) VALUES(?)", name)
	if err != nil {
		return 0, err
	}

	personID, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(personID), nil
}

func CreatePhoneQuery(personId int, number string) error {
	_, err := db.Exec("INSERT INTO phone(person_id, number) VALUES(?, ?)", personId, number)
	if err != nil {
		return err
	}
	return nil
}

func CreateAddressQuery(city, state, street1, street2, zipCode string) (int, error) {
	res, err := db.Exec("INSERT INTO address(city, state, street1, street2, zip_code) VALUES(?, ?, ?, ?, ?)",
		city, state, street1, street2, zipCode)
	if err != nil {
		return 0, err
	}

	addressId, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(addressId), nil
}
