package main

// middleware are used to see cross cutting concern
// validation,authentication
// middleware takes a handlerfunc as a input and returns an output as a handlerfunc
// and handlerfunc takes a context as a input and returns an error
// serveMessage in tronics will go
// ... in middleware means u can pass any number of middleware in this
// we can have a middleware that is to be used by everyone
// e.use()
// u can use e.pre(middleware.RemoveTrailingSlash()) as a part of echo function
// type middleware.  u will get the list
func main() {

}
