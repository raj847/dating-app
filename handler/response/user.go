package response

import (
	"dating-app/entity"
	"time"
)

type User struct {
	UserID    uint      `json:"user_id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Name      string    `json:"name"`
	Gender    string    `json:"gender"`
	DOB       time.Time `json:"date_of_birth"`
	Nickname  string    `json:"nickname"`
	Domicile  string    `json:"domicile"`
	Photo     string    `json:"photo"`
	Job       string    `json:"job"`
	Interest  string    `json:"interest"`
	CreatedAt string    `json:"created_at"`
	UpdatedAt string    `json:"updated_at"`
	DeletedAt string    `json:"deleted_at"`
}

func BuildUser(member entity.User) User {
	return User{
		UserID:    member.ID,
		Username:  member.Username,
		Email:     member.Email,
		Name:      member.Name,
		Gender:    member.Gender,
		DOB:       member.DOB,
		Nickname:  member.Nickname,
		Domicile:  member.Domicile,
		Photo:     member.Photo,
		Job:       member.Job,
		Interest:  member.Interest,
		CreatedAt: member.CreatedAt.String(),
		UpdatedAt: member.UpdatedAt.String(),
		DeletedAt: member.DeletedAt.Time.String(),
	}
}
