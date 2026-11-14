package main

import (
	"fmt"
	"log"
	"reflect"

	"github.com/cayleygraph/cayley"
	"github.com/cayleygraph/cayley/graph"
	_ "github.com/cayleygraph/cayley/graph/bolt"
	"github.com/cayleygraph/cayley/quad"
)

// ref. /cayley/graph/memstore/quadstore_test.go

// This is a simple test graph.
//
//    +---+                        +---+
//    | A |-------               ->| F |<--
//    +---+       \------>+---+-/  +---+   \--+---+
//                 ------>|#B#|      |        | E |
//    +---+-------/      >+---+      |        +---+
//    | C |             /            v
//    +---+           -/           +---+
//      ----    +---+/             |#G#|
//          \-->|#D#|------------->+---+
//              +---+
//
var simpleGraph = []quad.Quad{
	quad.Make("A", "follows", "B", ""),
	quad.Make("C", "follows", "B", ""),
	quad.Make("C", "follows", "D", ""),
	quad.Make("D", "follows", "B", ""),
	quad.Make("B", "follows", "F", ""),
	quad.Make("F", "follows", "G", ""),
	quad.Make("D", "follows", "G", ""),
	quad.Make("E", "follows", "F", ""),
	quad.Make("B", "status", "cool", "status_graph"),
	quad.Make("D", "status", "cool", "status_graph"),
	quad.Make("G", "status", "cool", "status_graph"),
}

func main() {

	// Initialize the database

	// Open and use the database
	store, err := cayley.NewMemoryGraph()
	if err != nil {
		log.Fatalln(err)
	}

	for _, each := range simpleGraph {
		store.AddQuad(each)
	}

	store.AddQuad(quad.Make("phrase of the day", "is of course", "Hello MemoryDB!", "demo graph"))

	fmt.Printf("%#v\n", store.Size())

	store.AddQuad(quad.Make("whoever", "follows", "B", ""))
	store.AddQuad(quad.Make("whoever_another", "follows", "B", ""))

	//quadValB := quad.String("B")
	//store.RemoveNode(store.ValueOf(quadValB))

	// Now we create the path, to get to our data
	//p := cayley.StartPath(store, quad.String("phrase of the day")).Out(quad.String("is of course"))
	//p := cayley.StartPath(store, quad.String("phrase of the day")).Out("is of course")
	//p := cayley.StartPath(store, quad.String("A")).Out("follows")
	//p := cayley.StartPath(store, quad.String("B")).In("follows")
	// p:= cayley.StartPath(store, quad.String("A")).Out("follows")
	//p := cayley.StartPath(store, quad.String("A"))
	p := cayley.StartPath(store, quad.String("B")).In("follows")

	it, _ := p.BuildIterator().Optimize()
	it, _ = store.OptimizeIterator(it)

	defer it.Close()

	//	println("by iterator ==================== ")
	//	for it.Next() {
	//		token := it.Result()                // get a ref to a node (backend-specific)
	//		value := store.NameOf(token)        // get the value in the node (RDF)
	//		nativeValue := quad.NativeOf(value) // convert value to normal Go type
	//
	//		println("----------------")
	//		fmt.Printf("%#v\n", token)
	//		fmt.Printf("quad value %#v\n", value) //panic: interface conversion: graph.Value is iterator.Int64Quad, not iterator.Int64Node
	//		fmt.Printf("native %#v\n", nativeValue)
	//	}
	//	if err := it.Err(); err != nil {
	//		log.Fatalln(err)
	//	}
	//
	println(" all quad  iterate ==================== ")

	{
		qit := store.QuadsAllIterator()
		defer qit.Close()
		for qit.Next() {
			token := qit.Result()

			//nativeValue := quad.NativeOf(value)
			println("----------------")

			//value := store.NameOf(token) // panic not node

			//asQuad := store.Quad(token)
			fmt.Printf("%#v\n", reflect.TypeOf(token).String())
			fmt.Printf("%#v\n", store.Quad(token))

			fmt.Printf("%#v\n", token)

			fmt.Printf("direction of subject %#v, %#v\n", store.QuadDirection(token, quad.Subject), store.NameOf(store.QuadDirection(token, quad.Subject)))

			fmt.Printf("direction of predicate %#v, %#v\n", store.QuadDirection(token, quad.Predicate), store.NameOf(store.QuadDirection(token, quad.Predicate)))

			fmt.Printf("direction of object %#v, %#v\n", store.QuadDirection(token, quad.Object), store.NameOf(store.QuadDirection(token, quad.Object)))

			//fmt.Printf("name of %#v\n", store.NameOf(token)) // panic
			//fmt.Printf("%#v\n", nativeValue)
		}

		if err := qit.Err(); err != nil {
			log.Fatalln(err)
		}
	}

	println(" all node iterator ==================== ")

	{
		nit := store.NodesAllIterator()
		defer nit.Close()
		for nit.Next() {
			token := nit.Result()

			//value := store.NameOf(token)
			//nativeValue := quad.NativeOf(value)

			println("----------------")

			fmt.Printf("%#v\n", store.NameOf(token))

			fmt.Printf("%#v\n", reflect.TypeOf(token).String())

			fmt.Printf("%#v\n", token)

			//fmt.Printf("name of %#v\n", quad.NameOf(token))
			//fmt.Printf("%#v\n", nativeValue)
		}

		if err := nit.Err(); err != nil {
			log.Fatalln(err)
		}
	}

	///store.ApplyDeltas([]Delta, IgnoreOpts) error
	//	store.NodesAllIterator() Iterator

	println(" iterate ==================== ")
	err = p.Iterate(nil).Each(func(value graph.Value) {

		quadValue := store.NameOf(value)
		nativeValue := quad.NativeOf(quadValue)

		///	store.NameOf(Value) quad.Value
		///
		///	//store.QuadIterator(quad.Direction, Value) Iterator
		///	store.ValueOf(quad.Value) Value
		///	store.QuadDirection(id Value, d quad.Direction) Value

		println("----------------")
		//fmt.Printf("%#v\n", store.Quad(value))

		fmt.Printf("%#v\n", reflect.TypeOf(value).String())
		//fmt.Printf("key %#v\n", value.Key())
		fmt.Printf("name of %#v\n", store.NameOf(value))
		fmt.Printf("value of %#v\n", store.ValueOf(quadValue))

		//fmt.Printf("quad %#v\n", store.Quad(value)) //panic: interface conversion: graph.Value is iterator.Int64Node, not

		fmt.Printf("%#v\n", value)
		fmt.Printf("%#v\n", quadValue)
		fmt.Printf("%#v\n", nativeValue)
	})
	if err != nil {
		log.Fatalln(err)
	}

}
