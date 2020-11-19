package views

import (
	"fmt"
	//"golang.org/x/net/html"
	"net/http"
)

func Index(w http.ResponseWriter, req *http.Request) {

	var s string
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

	// for rows.Next() {
	// 	err = rows.Scan(&user.Fname, &user.Lname, &user.Phone_number, &user.ID)
	// 	check(err)
	// 	idStr := strconv.Itoa(user.ID)

	s += `  <tr>
					   <td>` + `user.Fname` + `</td>
					   <td>` + `user.Lname` + `</td>
					   <td>` + `user.Phone_number` + `</td>
					   <td>` + `idStr` + `</td>
			           </tr>
	
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
