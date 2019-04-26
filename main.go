package main

func main() {
	a := App{}
	// You need to set your Username and Password here
	a.Initialize("development", "Magicmicro1!", "sa", "golang-test")

	a.Run(":8080")
}
