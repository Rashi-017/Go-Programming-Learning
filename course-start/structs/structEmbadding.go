package main

type Info struct {
	email string
	phone int
}
type PersonDetails struct {
	firstname    string
	lastname     string
	PersonalInfo Info
}

func Filldetails(firstname string, lastname string, email string, phone int) PersonDetails {
	return PersonDetails{
		firstname: firstname,
		lastname:  lastname,
		PersonalInfo: Info{
			email: email,
			phone: phone,
		},
	}
}
func (p *Person) changeInfo(name string) {

}
