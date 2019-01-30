package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type EmployeeTable_20190131_004314 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &EmployeeTable_20190131_004314{}
	m.Created = "20190131_004314"

	migration.Register("EmployeeTable_20190131_004314", m)
}

// Run the migrations
func (m *EmployeeTable_20190131_004314) Up() {
    m.SQL("CREATE TABLE employee(id INT AUTO_INCREMENT PRIMARY KEY, name VARCHAR(100), email VARCHAR(100), city VARCHAR(100))")
}
 
// Reverse the migrations
func (m *EmployeeTable_20190131_004314) Down() {
    m.SQL("DROP TABLE employee")
}
