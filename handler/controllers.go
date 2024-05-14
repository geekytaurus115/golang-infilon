package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetPersonById(c *gin.Context) {
	personId := c.Param("person_id")

	person, err := GetPersonByIdQuery(personId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch person information"})
		return
	}

	phone, err := GetPhoneByPersonIdQuery(personId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch phone information"})
		return
	}

	address, err := GetAddressByPersonIdQuery(personId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch address information"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"name":         person.Name,
		"phone_number": phone.Number,
		"city":         address.City,
		"state":        address.State,
		"street1":      address.Street1,
		"street2":      address.Street2,
		"zip_code":     address.ZipCode,
	})
}

func CreatePerson(c *gin.Context) {
	var requestBody PersonRequestBody

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		log.Println("Invalid request body")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	personId, err := CreateNewPersonQuery(requestBody.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create person"})
		return
	}

	err = CreatePhoneQuery(personId, requestBody.PhoneNumber)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create phone"})
		return
	}

	addressId, err := CreateAddressQuery(requestBody.City, requestBody.State, requestBody.Street1, requestBody.Street2, requestBody.ZipCode)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create address"})
		return
	}

	_, err = db.Exec("INSERT INTO address_join (person_id, address_id) VALUES (?, ?)", personId, addressId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add mapping to address_join table"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":   "Person created successfully",
		"person_id": personId,
	})
}
