package main

func main() {
	a := App{}
	a.Initialize("postgres", "abc12345", "dealtap", "public")
	a.StartServer(":8080")
}