package main

import (
	"fmt"
	"sort"
)

type people []string

func (p people) Len() int           { return len(p) }
func (p people) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p people) Less(i, j int) bool { return p[i] < p[j] }

func main(){
	//There are some functions which you need to implement an interface to use them, for example, sort
	//sort takes an interface which implements Less, Len and Swap
	studyGroup := people{"Zeno", "John", "Al", "Jenny"}

	fmt.Println(studyGroup)
	sort.Sort(studyGroup)
	fmt.Println(studyGroup)

	//sort have other functions for primitive types like int

	n := []int{7, 4, 8, 2, 9, 19, 12, 32, 3}
	fmt.Println(n)
	sort.Sort(sort.IntSlice(n))
	fmt.Println(n)

	//same as
	//sort.Ints(n)

}

////Also you can have different sorting methods:
//type person struct {
//	Name string
//	Age  int
//}
//// ByAge implements sort.Interface for []person based on
//// the Age field.
//type ByAge []person
//
//func (a ByAge) Len() int           { return len(a) }
//func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
//func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }
//
//sort.Sort(ByAge(people))

//When you declare an empty interface every type is part of that interface
//type vehicles interface{}

//If you want to accept any parameter you just have to receive an empty interface "a interface{}"
//Same with slices []interface{}{...}

//ASSERTIONS (check if a var belongs to a type)

//value, ok := <var>.(<type>)

//WAIT GROUPS
//with var wg sync.WaitGroup
//we can wait for a number of goroutines to finish
//we add the number with wg.Add(number)
//and once called we wait with wg.Wait()
//in the func called with go, once finished we need to call wg.Done()

//CPUs
//With this init
//func init() {
//	runtime.GOMAXPROCS(runtime.NumCPU())
//}
//We can set the number of CPUs we want to use in our parallelism
//On parallelism there can be race condition
//We can use mutex.Lock() and mutex.Unlock() to avoid it in the part of code we want to protect

//We can also use the library atomic (atomic.AddInt64(&counter, 1)) which manages variables in parallel environments
//to access atomic.LoadInt64(&counter)

//When using channels we can use a channel to be used as a semaphore and when the channel is true we can close the channel with the data

//log.Fatalln(err) is equivalent to println and exit
//We can set a file to log in, with log.SetOutput(nf)

//We can create custom errors with errors.New("norgate math: square root of negative number")