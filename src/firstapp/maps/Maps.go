package maps

//Weeks is public
var Weeks map[int]string

//PopulateMaps is public
func PopulateMaps() {
	Weeks = map[int]string{
		1: "Monday",
		2: "Tuesday",
		3: "Wednesday",
		4: "Thursday",
		5: "Friday",
		6: "Saturday",
		7: "Sunday",
	}

}
