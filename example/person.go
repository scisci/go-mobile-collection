package society

// @collection-wrapper
type Person struct {
  Name string
  YearOfBirth int
  CanRideABike bool
}

func (*Person) Equal(rhs *Person) bool{
	return true
}