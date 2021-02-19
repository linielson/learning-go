package car

//type car struct { //unexported
type Car struct {
	Name string
	color string //unexported
}

//func (c Car) start() string { //unexported
func (c Car) Start() string {
	return c.Name + " has been started"
}