/*
=============================================================================
  EXERCISE 4: GO (GOLANG) BASICS
=============================================================================
  Go is a simple but powerful language. This file teaches you:
  1. Variables & types
  2. Functions with multiple return values
  3. Structs & methods (Go's version of classes)
  4. Slices & Maps (arrays & dictionaries)
  5. Error handling (Go's unique approach)
  6. Goroutines & Mutexes (concurrency!)

  HOW TO USE:
  1. Read each section — examples show you the pattern
  2. Fill in the TODO sections
  3. Run: go test -v
  4. For concurrency: go test -race -v
=============================================================================
*/

package main

import (
	"fmt"
	"math"
	"strings"
	"sync"
)

// Avoid unused import errors
var _ = fmt.Sprintf
var _ = math.Pi
var _ = strings.ToUpper

// ============================================================================
// PART 1: VARIABLES & TYPES
// ============================================================================
//
// Go is statically typed — every variable has a type.
//
// Two ways to declare variables:
//   var name string = "Alice"     ← explicit type
//   name := "Alice"               ← short form (Go figures out the type)
//
// Types:
//   string    → "hello"
//   int       → 42
//   float64   → 3.14
//   bool      → true, false
//   []int     → slice (dynamic array) of ints
//   map[string]int → dictionary with string keys and int values
//
// IMPORTANT: Go has NO classes, NO exceptions, NO generics (before 1.18)
//             Instead: structs, multiple return values, and interfaces
//
// EXAMPLE:
func exampleVariables() (string, int, float64) {
	name := "Alice"
	age := 25
	gpa := 8.5
	return name, age, gpa
}

// TODO 1: Write a function that takes a Celsius temperature and returns Fahrenheit
// Formula: F = C × 9/5 + 32
// EXAMPLE: celsiusToFahrenheit(0) → 32.0
// EXAMPLE: celsiusToFahrenheit(100) → 212.0
func celsiusToFahrenheit(celsius float64) float64 {
	// Write your code below this line
	return 0
}

// ============================================================================
// PART 2: FUNCTIONS WITH MULTIPLE RETURN VALUES
// ============================================================================
//
// Go functions can return MULTIPLE values. This is how Go handles errors:
//   func divide(a, b float64) (float64, error) {
//       if b == 0 {
//           return 0, fmt.Errorf("cannot divide by zero")
//       }
//       return a / b, nil    // nil = no error
//   }
//
// CALLING:
//   result, err := divide(10, 3)
//   if err != nil {
//       fmt.Println("Error!", err)
//       return
//   }
//   fmt.Println("Result:", result)
//
// EXAMPLE:
func exampleSqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, fmt.Errorf("cannot take square root of negative number: %f", x)
	}
	return math.Sqrt(x), nil
}

// TODO 2: Write a function that divides two integers
// Return the result and remainder as two values, plus an error
// If b is 0, return (0, 0, error) with error message "division by zero"
// Otherwise return (a/b, a%b, nil)
// EXAMPLE: divmod(10, 3) → (3, 1, nil)
// EXAMPLE: divmod(10, 0) → (0, 0, error)
func divmod(a, b int) (int, int, error) {
	// Write your code below this line
	return 0, 0, nil
}

// ============================================================================
// PART 3: STRUCTS & METHODS
// ============================================================================
//
// Go doesn't have classes. Instead, it uses STRUCTS (data) + METHODS (behavior).
//
// STRUCT = a collection of fields (like a class without methods):
//   type Person struct {
//       Name string     ← Capital letter = public (exported)
//       age  int        ← lowercase = private (unexported)
//   }
//
// METHOD = a function attached to a struct:
//   func (p *Person) Greet() string {     ← (p *Person) = receiver
//       return "Hello, " + p.Name
//   }
//
// CREATING:
//   alice := Person{Name: "Alice", age: 25}
//   alice := &Person{Name: "Alice", age: 25}  ← returns a pointer
//
// EXAMPLE:
type ExampleRectangle struct {
	Width  float64
	Height float64
}

func (r *ExampleRectangle) Area() float64 {
	return r.Width * r.Height
}

func (r *ExampleRectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// TODO 3: Create a BankAccount struct and its methods
//
// Struct fields:
//   Owner   string
//   Balance float64
//
// Methods:
//   Deposit(amount float64) error
//     - If amount <= 0, return error "deposit amount must be positive"
//     - Otherwise add to Balance, return nil
//
//   Withdraw(amount float64) error
//     - If amount <= 0, return error "withdrawal amount must be positive"
//     - If amount > Balance, return error "insufficient funds"
//     - Otherwise subtract from Balance, return nil
//
//   GetBalance() float64
//     - Return the current Balance
//
// EXAMPLE:
//   acc := &BankAccount{Owner: "Alice", Balance: 100}
//   acc.Deposit(50)     → nil (balance is now 150)
//   acc.Withdraw(30)    → nil (balance is now 120)
//   acc.GetBalance()    → 120.0

type BankAccount struct {
	Owner   string
	Balance float64
}

func (b *BankAccount) Deposit(amount float64) error {
	// Write your code below this line
	return nil
}

func (b *BankAccount) Withdraw(amount float64) error {
	// Write your code below this line
	return nil
}

func (b *BankAccount) GetBalance() float64 {
	// Write your code below this line
	return 0
}

// ============================================================================
// PART 4: SLICES & MAPS
// ============================================================================
//
// SLICES (dynamic arrays):
//   nums := []int{1, 2, 3}
//   nums = append(nums, 4)        // Add: [1, 2, 3, 4]
//   len(nums)                     // Length: 4
//   nums[0]                       // Access: 1
//
// Looping:
//   for i, val := range nums {    // i = index, val = value
//       fmt.Println(i, val)
//   }
//   for _, val := range nums {    // _ = ignore index
//       fmt.Println(val)
//   }
//
// MAPS (dictionaries):
//   ages := map[string]int{
//       "Alice": 25,
//       "Bob":   30,
//   }
//   ages["Charlie"] = 35          // Add
//   val, ok := ages["Alice"]      // ok = true if key exists
//   delete(ages, "Bob")           // Remove
//   len(ages)                     // Count

// TODO 4: Write a function that counts how many times each word appears
// EXAMPLE: wordCount("hello world hello") → map[string]int{"hello": 2, "world": 1}
// HINT: Use strings.Fields(text) to split by whitespace
func wordCount(text string) map[string]int {
	// Write your code below this line
	return nil
}

// TODO 5: Write a function that returns the top N items from a slice
// If N is greater than the slice length, return the entire slice
// EXAMPLE: topN([]int{5, 3, 8, 1, 9}, 3) → [5, 3, 8]
// EXAMPLE: topN([]int{1, 2}, 5) → [1, 2]
func topN(items []int, n int) []int {
	// Write your code below this line
	return nil
}

// ============================================================================
// PART 5: GOROUTINES & MUTEXES (Concurrency!)
// ============================================================================
//
// This is what makes Go special! Go makes concurrency easy.
//
// GOROUTINE = a lightweight thread. Start one with the 'go' keyword:
//   go myFunction()     ← runs myFunction in the background
//
// PROBLEM: If two goroutines access the same variable, you get a RACE CONDITION
// (unpredictable results, data corruption)
//
// SOLUTION: Use a MUTEX (mutual exclusion lock):
//   var mu sync.Mutex
//
//   mu.Lock()           ← only one goroutine can pass this at a time
//   sharedVariable++    ← safe to modify!
//   mu.Unlock()         ← let others through
//
// EXAMPLE:
type ExampleSafeCounter struct {
	mu    sync.Mutex
	count int
}

func (c *ExampleSafeCounter) Increment() {
	c.mu.Lock()         // Lock before modifying
	c.count++
	c.mu.Unlock()       // Unlock after modifying
}

func (c *ExampleSafeCounter) GetCount() int {
	c.mu.Lock()
	defer c.mu.Unlock() // 'defer' = runs when function returns
	return c.count
}

// TODO 6: Create a thread-safe ScoreBoard struct
//
// Fields:
//   mu     sync.Mutex
//   scores map[string]int
//
// Methods:
//   AddScore(player string, points int)
//     - Safely add points to the player's score
//     - If player doesn't exist, create them with the given points
//
//   GetScore(player string) int
//     - Safely return the player's score (0 if not found)
//
//   GetAllScores() map[string]int
//     - Safely return a COPY of the scores map
//     - (Important: return a copy, not the original, for thread safety)
//
// HINT: Always Lock() before accessing scores, Unlock() after

type ScoreBoard struct {
	mu     sync.Mutex
	scores map[string]int
}

func NewScoreBoard() *ScoreBoard {
	return &ScoreBoard{scores: make(map[string]int)}
}

func (s *ScoreBoard) AddScore(player string, points int) {
	// Write your code below this line
}

func (s *ScoreBoard) GetScore(player string) int {
	// Write your code below this line
	return 0
}

func (s *ScoreBoard) GetAllScores() map[string]int {
	// Write your code below this line
	return nil
}
