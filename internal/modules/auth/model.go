package auth

import "time"

type UserRole string

const (
	RoleUser      UserRole = "USER"
	RoleAdmin     UserRole = "ADMIN"
	RoleMODARETOR UserRole = "MODARETOR"
)

type User struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Auth0ID   string    `gorm:"uniqueIndex;not null" json:"auth0_id"`
	Email     string    `gorm:"uniqueIndex;not null" json:"email"`
	Role      UserRole  `gorm:"type:varchar(20);default:'USER';not null" json:"role"`
	Name      string    `json:"name"`
	Picture   string    `json:"picture"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
