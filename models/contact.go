package models

type Contact struct {
	FirstName string `form:"first_name" json:"first_name" binding:"required,min=3,max=50"`
	LastName  string `form:"last_name" json:"last_name" binding:"required,min=3,max=50"`
	Mobile    string `form:"phone" json:"phone" binding:"required,min=10,max=10"`
}

func (c *Contact) SaveContact() (*Contact, error) {
	err := DB.Create(&c).Error
	if err != nil {
		return &Contact{}, err
	}
	return c, nil
}
