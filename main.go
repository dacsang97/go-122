package main

import (
	"fmt"
	"math/rand/v2"
	"net/http"
	"slices"
	"sync"
	"time"
)

func loop_closure() {
	wg := sync.WaitGroup{}
	wg.Add(5)

	nums := []int{1, 2, 3, 4, 5}
	for _, num := range nums {
		go func() {
			defer wg.Done()
			fmt.Println(num)
		}()
	}

	wg.Wait()
}

func range_number() {
	for i := range 10 {
		fmt.Println(i)
	}
}

// func Split(s string) func(func(int, string) bool) {
// 	parts := strings.Split(s, " ")
// 	return func(f func(int, string) bool) {
// 		for i, part := range parts {
// 			if !f(i, part) {
// 				break
// 			}
// 		}
// 	}
// }

// // run with GOEXPERIMENT=rangefunc go run main.go
// func range_func() {
// 	str := "Happy new year 2024"
// 	for i, x := range Split(str) {
// 		fmt.Println(i, x)
// 	}
// }

func math_rand_v2() {
	for _ = range 10 {
		duration := rand.N(time.Minute * 5)
		fmt.Println(duration)
		time.Sleep(500 * time.Millisecond)
	}
}

func slices_concat() {
	s1 := []int{1, 2, 3}
	s2 := []int{4, 5, 6}
	s3 := []int{7, 8, 9}
	fmt.Println(slices.Concat(s1, s2, s3))
}

// new mux
func new_mux() {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /users", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("create user"))
	})

	mux.HandleFunc("/users/{id}", func(w http.ResponseWriter, r *http.Request) {
		userID := r.PathValue("id")

		fmt.Printf("get user %s\n", userID)
	})

	http.ListenAndServe(":8080", mux)
}

func main() {
	// loop_closure()
	// range_number()
	// range_func() // run with GOEXPERIMENT=rangefunc go run main.go
	// math_rand_v2()
	// slices_concat()
	new_mux()
}
