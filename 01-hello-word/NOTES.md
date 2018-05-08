Separate "domain" code from the outside world (side-effetcs). In hello.go was created the function Hello and the result is used
to Println. With this the test would be more easier

Rules for writing tests:

- It needs to be in a file with name xxx_test.go. By convention it's the same name of file that will be tested e.g. hello.go -> hello_test.go
- The test function must start with the word Test. e.g. Hello() -> TestHello()
- The test function takes one argument only t *testing.T. e.g. TestHello(t *testing.T)