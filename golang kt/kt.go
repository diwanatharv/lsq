/*
 concurrency== splitting them into individual parts they execute independently
 and can be snychronized by communicating properly
concurrency is achieved through go routines and channels
go routines are independent function and channels synchronized
go routines is a thread
channels are the data structure which helps the go routines to communicate and do signaling'
ch:=make(chan string ,1);
ch<-"kamesh"; inserting values
<-ch
whenever use go routines write a comment when u think u should terminate
close it also
close(ch); if we do not this it will run in background

//package designing
//network layer--how are your services exposed
//buisness layer
//data layer--database
//put the package name same as directory name


//cmd --main.go
//internal
//pkg

//api and workflowScreenshot_20230226_224924.png


//import 1st standard like fmt
//2nd go packages
//3rd 3rd party

// whitespace before comments like this one






















*/