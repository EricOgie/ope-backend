package requestdto

import "github.com/EricOgie/ope-be/domain/models"

type VerifyRequest struct {
	Id         string `json:"id"`
	FirstName  string `json:"firstname"`
	Lastname   string `json:"lastname"`
	Email      string `json:"email"`
	Created_at string `json:"created_at"`
	When       string `json:"when"`
}

func (v VerifyRequest) GetUserFromVerify() models.User {
	return models.User{
		Id:        v.Id,
		FirstName: v.FirstName,
		LastName:  v.Lastname,
		Email:     v.Email,
		CreatedAt: v.Created_at,
	}
}
