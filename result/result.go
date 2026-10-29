package result

func CalculateTotal(marks []float64) float64 {
var total float64


for _, mark := range marks {
	total += mark
}

return total


}

func CalculateAverage(marks []float64) float64 {
if len(marks) == 0 {
return 0
}


return CalculateTotal(marks) / float64(len(marks))


}

func CalculateGrade(average float64) string {
switch {
case average >= 90:
return "O"
case average >= 80:
return "A+"
case average >= 70:
return "A"
case average >= 60:
return "B+"
case average >= 50:
return "B"
default:
return "RA"
}
}
