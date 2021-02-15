package another_package

import "fmt"

var scopedVariable string = "This variable is visible only another_package scope"
var ExportedVariable string = "This variable is public"

func PrintPublic() {
	fmt.Println("This function is public/exported")
	fmt.Println(scopedVariable)
}
