package models

import "time"

// Payment Payment made for the course enrollment
type Payment struct {
	ID                int         `gorm:"primary_key;column:id;type:int(11);not null" json:"-"`
	EnrollmentID      int         `gorm:"index:fk_payment_enrollment_id_idx;column:enrollment_id;type:int(11);not null" json:"enrollment_id"`
	Enrollment        *Enrollment `gorm:"association_foreignkey:enrollment_id;foreignkey:id" json:"enrollment_list"` // Stores user-course information
	TransactionNumber string      `gorm:"column:transaction_number;type:varchar(16);not null" json:"transaction_number"`
	Amount            int16       `gorm:"column:amount;type:smallint(6);not null" json:"amount"`
	Mode              string      `gorm:"column:mode;type:enum('credit card','debit card','upi','wallet');not null" json:"mode"`
	State             string      `gorm:"column:state;type:enum('new','success','failure');not null" json:"state"`
	CreatedAt         time.Time   `gorm:"column:created_at;type:datetime;not null" json:"created_at"`
	UpdatedAt         time.Time   `gorm:"column:updated_at;type:datetime" json:"updated_at"`
}
