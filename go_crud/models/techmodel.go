package models

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/elysiamori/go-crud/config"
	"github.com/elysiamori/go-crud/entities"
)

// Koneksi Database
type TechModel struct {
	conn *sql.DB
}

func NewTechModel() *TechModel {
	conn, err := config.DBConnection()
	if err != nil {
		panic(err)
	}

	return &TechModel{
		conn: conn,
	}
}

// Function untuk menampilkan data dari database ke dalam form
func (p *TechModel) FindAll() ([]entities.Tech, error) {
	rows, err := p.conn.Query("select * from tech")
	if err != nil {
		return []entities.Tech{}, err
	}
	defer rows.Close()

	var dataTech []entities.Tech
	for rows.Next() {
		var tech entities.Tech
		rows.Scan(&tech.Id,
			&tech.Name,
			&tech.Job,
			&tech.Programming,
			&tech.Date)

		date, _ := time.Parse("2006-01-02", tech.Date)
		tech.Date = date.Format("02-01-2006")

		dataTech = append(dataTech, tech)
	}

	return dataTech, nil
}

// Function untuk membuat data kedalam database
func (p *TechModel) Create(tech entities.Tech) bool {
	result, err := p.conn.Exec("insert into tech (name, job, programming, date) values (?,?,?,?)",
		tech.Name, tech.Job, tech.Programming, tech.Date)

	if err != nil {
		fmt.Println(err)
		return false
	}

	lastInsertId, _ := result.LastInsertId()

	return lastInsertId > 0
}

// Function untuk mengklik bagian edit untuk setiap data
func (p *TechModel) Find(id int64, tech *entities.Tech) error {

	return p.conn.QueryRow("select * from tech where id = ?", id).Scan(&tech.Id,
		&tech.Name,
		&tech.Job,
		&tech.Programming,
		&tech.Date)
}

// Function Update data
func (p *TechModel) Update(tech entities.Tech) error {

	_, err := p.conn.Exec(
		"update tech set name = ?, job = ?, programming = ?, date = ? where id = ?",
		tech.Name, tech.Job, tech.Programming, tech.Date, tech.Id)

	if err != nil {
		return err
	}

	return nil
}

// Function Delete Data
func (p *TechModel) Delete(id int64) {
	p.conn.Exec("delete from tech where id = ?", id)
}
