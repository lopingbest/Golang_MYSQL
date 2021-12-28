package Golang_MYSQL

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"
)

func TestExecSql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "INSERT INTO customer(id, name) VALUES('joko','Joko')"
	//exec tidak akan mengembalikan result
	_, err := db.ExecContext(ctx, script)
	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert new customer")
}

func TestQuerySql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "SELECT id, name FROM customer"
	rows, err := db.QueryContext(ctx, script)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	//selama masih ada data di rows, maka akan terus diambil
	for rows.Next() {
		var id, name string
		//didalam kurung dimasukan parameterdari apa yang akan  kita ambil
		//Pointer dipakai karena kita akan ngeset data dari parameter. Kalo enggak pointer, maka data tidak akan dipakai
		err = rows.Scan(&id, &name)
		//kalo data sudah tidak ada makan akan muncul panic
		if err != nil {
			panic(err)
		}
		//output
		fmt.Println("Id:", id)
		fmt.Println("Name:", name)
	}
}

func TestQuerySqlComplex(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "SELECT id, name, email, balance, rating, birth_date, married, created_at FROM customer"
	rows, err := db.QueryContext(ctx, script)
	//kalo data sudah tidak ada makan akan muncul panic
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	//selama masih ada data di rows, maka akan terus diambil
	for rows.Next() {
		var id, name string
		var email sql.NullString
		var balance int32
		var rating float64
		//data/timestamp tetep menggunakan time.Time kalau database cuma sampai tanggal, maka nanti jam menit detiknya akan kosong semua
		var birthDate sql.NullTime
		var createdAt time.Time
		var married bool

		//didalam kurung dimasukan parameterdari apa yang akan  kita ambil
		//Pointer dipakai karena kita akan ngeset data dari parameter. Kalo enggak pointer, maka data tidak akan dipakai
		err = rows.Scan(&id, &name, &email, &balance, &rating, &birthDate, &married, &createdAt)
		if err != nil {
			panic(err)
		}
		fmt.Println("================")
		fmt.Println("Id:", id)
		fmt.Println("Name:", name)
		if email.Valid {
			fmt.Println("Email:", email.String)
		}
		fmt.Println("Balance:", balance)
		fmt.Println("Rating:", rating)
		if birthDate.Valid {
			fmt.Println("Birth Date:", birthDate.Time)
		}
		fmt.Println("Married:", married)
		fmt.Println("Created At:", createdAt)
	}
}
