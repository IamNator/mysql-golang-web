package controllers

import (
	"database/sql"
	"fmt"
	//"golang.org/x/net/html"
	"github.com/IamNator/mysql-golang-web/models"
	"net/http"
	"strconv"
)

func Index(w http.ResponseWriter, req *http.Request) {

	db, err := sql.Open("mysql", "root:299792458m/s@tcp(127.0.0.1:3306)/test")
	check(err)

	err = db.Ping()
	check(err)

	rows, err := db.Query(`SELECT fname, lname, phone_number, id FROM phonenumber`)
	check(err)

	var s string
	var user models.User
	s = `<!DOCTYPE html>
	<html lang="en">
		<head>

			<title>
				Phone Record
			</title>
			<meta charset="utf-8">
			<meta name="viewport" content="width=device-width, initial-scale=0.7">

		<link type="text/css"  href="/css/bootstrap.min.css"
		rel="stylesheet">
			
		</head>
		<body>
			<p>
				<a href="/insert">
					Click here to Populate 
				</a>
			</p>
			<p>
				Retrieved Data :
			</p>
			<table class="table table-bordered" style="width:65%" >
				<thead>
					<tr>
						<th scope="col">fname</th>
						<th scope="col">lname</th>
						<th scope="col">phone_number</th>
						<th scope="col">id</th>
					</tr>
				</thead>
				<tbody>`

	for rows.Next() {
		err = rows.Scan(&user.Fname, &user.Lname, &user.Phone_number, &user.ID)
		check(err)
		idStr := strconv.Itoa(user.ID)

		s += `  <tr>
				   <td>` + user.Fname + `</td>
				   <td>` + user.Lname + `</td>
				   <td>` + user.Phone_number + `</td>
				   <td>` + idStr + `</td>
		        </tr>`
	}

	s += `
				</tbody>
			</table>
			<form metod="POST" action="del" style="width:35%">
			<div class="col-sm-15" style="width:12%">
				<input name="Del_id" type="text" value="" class="form-control" placeholder="id number to Del."/>
				<input type="submit" value="submit" class="col-sm-15" />
			</div>
			</form>
		</body>
	<html>`
	fmt.Fprintln(w, s)
	db.Close()
}

func Delete(w http.ResponseWriter, req *http.Request) {
	if err := req.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	del_id := req.FormValue("Del_id")

	db, err := sql.Open("mysql", "root:299792458m/s@tcp(127.0.0.1:3306)/test") //##################
	check(err)

	stmt, err := db.Prepare(`DELETE FROM phonenumber WHERE id = ? ;`)

	_, err = stmt.Exec(del_id)
	check(err)
	db.Close() //#######################

	fmt.Fprintf(w, `<!DOCTYPE html>
	<html lang="en">
		<head>

			<title>
				 Successful
			</title>
			<meta charset="utf-8">
			<meta name="viewport" content="width=device-width, initial-scale=0.7">

		<link type="text/css"  href="/css/bootstrap.min.css"
		rel="stylesheet">
			
		</head>
		<body>
			<p>
				<a href="/index">
					Click here to return to table 
				</a>
			</p>
			<p>
					Delete successful
			</p>
		</body>
	</html>`)

}

func FormHandler(w http.ResponseWriter, req *http.Request) {

	if err := req.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	var user models.User
	user.Fname = req.FormValue("fname")
	user.Lname = req.FormValue("lname")
	user.Phone_number = req.FormValue("phone_number")
	user.ID, _ = strconv.Atoi(req.FormValue("id"))

	if req.FormValue("fname") != "" && req.FormValue("lname") != "" && req.FormValue("phone_number") != "" && req.FormValue("id") != "" {

		db, err := sql.Open("mysql", "root:299792458m/s@tcp(127.0.0.1:3306)/test") //##################
		check(err)

		stmt, err := db.Prepare(`INSERT INTO phonenumber (fname,lname,phone_number,id)
	VALUES (?,?,?,?)`)

		_, err = stmt.Exec(user.Fname, user.Lname, user.Phone_number, user.ID)
		check(err)
		db.Close() //#######################

		if err != nil {
			fmt.Fprintln(w, err)
		} else {
			fmt.Println("\nData Successfully Added")
		}

		fmt.Fprintf(w, `<!DOCTYPE html>
	<html lang="en">
		<head>

			<title>
				 Successful
			</title>
			<meta charset="utf-8">
			<meta name="viewport" content="width=device-width, initial-scale=0.7">

		<link type="text/css"  href="/css/bootstrap.min.css"
		rel="stylesheet">
			
		</head>
		<body>
			<p>
				<a href="index">
					Click here to return to table 
				</a>
			</p>
			<p>
					successful added
			</p>
		</body>
	</html>`)
	} else {
		fmt.Fprintf(w, `<!DOCTYPE html>
	<html lang="en">
		<head>

			<title>
				 Successful
			</title>
			<meta charset="utf-8">
			<meta name="viewport" content="width=device-width, initial-scale=0.7">

		<link type="text/css"  href="/css/bootstrap.min.css"
		rel="stylesheet">
			
		</head>
		<body>
			<p>
				<a href="/index">
					Click here to return to table 
				</a>
			</p>
			<p>
					Please fill in all the fields
			</p>
		</body>
	</html>`)
	}

}

func Insert(w http.ResponseWriter, req *http.Request) {

	r := `<!DOCTYPE html>
			<html>
				  <head>
				  	<meta charset="UTF-8" />
					<title>
						Fill data
					</title>

					<link type="text/css"  href="/css/bootstrap.min.css"
					rel="stylesheet">
					
				</head>
				<body>
					<p>
						<a href="/index"> 
							Click here to view filed table
						</a>
					</p>
						<form method="POST" action="/form" style="width:35%">
							<div class="form-row">
								<label class="col-sm-4 col-form-label"> First Name</label>
								<div class="col-sm-15">	
									<input name="fname" type="text" value="" class="form-control" placeholder="First Name"/>
								</div>
							</div>	
							<div class="form-row">
								<label class="col-sm-4 col-form-label"> Last Name</label>
								<div class="col-sm-15">			
									<input name="lname" type="text" value="" class="form-control" placeholder="Last Name" />
								</div>
							</div>
							<div class="form-row">
								<label class="col-sm-4 col-form-label"> Phone Number</label>
								<div class="col-sm-15">
									<input name="phone_number" type="text" value="" class="form-control" placeholder="Phone Number"/>
								</div>
							</div>	
							<div class="form-row">	
								<label class="col-sm-4 col-form-label">id</label>
							    <div class="col-sm-15">
									<input name="id" type="text" value="" class="form-control" placeholder="id number"/>
								</div>
							</div>
							<input type="submit" value="submit" class="col-sm-15" />
						</form>
				</body>
			</html>`

	fmt.Fprintln(w, r)
}

func check(err error) {
	if err != nil {
		fmt.Printf("%+v\n", err)
	}
}
