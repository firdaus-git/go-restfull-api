package model
 
import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

type Model struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time 
	UpdatedAt time.Time 
	DeletedAt *time.Time 
  }

type Employee struct {
	gorm.Model
	Name    string `gorm:"unique" json:"name"`
	Age     int    `json:"age"`
	Email   string `json:"email"`
	Phone  string `json:"phone"`
	Address string `json:"address"`
	City    string `json:"city"`
	Status  bool `json:"status"`
}
 
func (e *Employee) Disable() {
	e.Status = false
}
 
func (p *Employee) Enable() {
	p.Status = true
}
 
func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&Employee{})
	return db
}