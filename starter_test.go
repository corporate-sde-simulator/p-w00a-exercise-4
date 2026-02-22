/*
  Tests for Exercise 4: Go Basics
  Run with: go test -v
  With race detector: go test -race -v
*/

package main

import (
	"math"
	"sync"
	"testing"
)

// ─── Part 1: Variables ────────────────────────────────────────────
func TestCelsiusToFahrenheit(t *testing.T) {
	tests := []struct {
		celsius    float64
		fahrenheit float64
	}{
		{0, 32},
		{100, 212},
		{-40, -40},
		{37, 98.6},
	}
	for _, tt := range tests {
		result := celsiusToFahrenheit(tt.celsius)
		if math.Abs(result-tt.fahrenheit) > 0.01 {
			t.Errorf("celsiusToFahrenheit(%v) = %v, want %v", tt.celsius, result, tt.fahrenheit)
		}
	}
}

// ─── Part 2: Multiple Return Values ──────────────────────────────
func TestDivmod(t *testing.T) {
	q, r, err := divmod(10, 3)
	if err != nil {
		t.Errorf("divmod(10, 3) returned unexpected error: %v", err)
	}
	if q != 3 || r != 1 {
		t.Errorf("divmod(10, 3) = (%d, %d), want (3, 1)", q, r)
	}

	q, r, err = divmod(7, 2)
	if q != 3 || r != 1 {
		t.Errorf("divmod(7, 2) = (%d, %d), want (3, 1)", q, r)
	}

	_, _, err = divmod(10, 0)
	if err == nil {
		t.Error("divmod(10, 0) should return an error")
	}
}

// ─── Part 3: Structs & Methods ──────────────────────────────────
func TestBankAccount(t *testing.T) {
	acc := &BankAccount{Owner: "Alice", Balance: 100}

	// Deposit
	err := acc.Deposit(50)
	if err != nil {
		t.Errorf("Deposit(50) returned error: %v", err)
	}
	if acc.GetBalance() != 150 {
		t.Errorf("Balance after deposit should be 150, got %v", acc.GetBalance())
	}

	// Negative deposit
	err = acc.Deposit(-10)
	if err == nil {
		t.Error("Deposit(-10) should return an error")
	}

	// Withdraw
	err = acc.Withdraw(30)
	if err != nil {
		t.Errorf("Withdraw(30) returned error: %v", err)
	}
	if acc.GetBalance() != 120 {
		t.Errorf("Balance after withdrawal should be 120, got %v", acc.GetBalance())
	}

	// Overdraft
	err = acc.Withdraw(200)
	if err == nil {
		t.Error("Withdraw(200) with balance 120 should return error")
	}

	// Negative withdrawal
	err = acc.Withdraw(-5)
	if err == nil {
		t.Error("Withdraw(-5) should return an error")
	}
}

// ─── Part 4: Slices & Maps ──────────────────────────────────────
func TestWordCount(t *testing.T) {
	result := wordCount("hello world hello")
	if result["hello"] != 2 {
		t.Errorf("wordCount: expected hello=2, got hello=%d", result["hello"])
	}
	if result["world"] != 1 {
		t.Errorf("wordCount: expected world=1, got world=%d", result["world"])
	}
}

func TestTopN(t *testing.T) {
	result := topN([]int{5, 3, 8, 1, 9}, 3)
	if len(result) != 3 {
		t.Errorf("topN(5 items, 3) should return 3 items, got %d", len(result))
	}
	expected := []int{5, 3, 8}
	for i, v := range expected {
		if result[i] != v {
			t.Errorf("topN result[%d] = %d, want %d", i, result[i], v)
		}
	}

	// N bigger than slice
	result2 := topN([]int{1, 2}, 5)
	if len(result2) != 2 {
		t.Errorf("topN(2 items, 5) should return 2 items, got %d", len(result2))
	}
}

// ─── Part 5: Goroutines & Mutexes ───────────────────────────────
func TestScoreBoard(t *testing.T) {
	sb := NewScoreBoard()

	sb.AddScore("Alice", 10)
	sb.AddScore("Bob", 20)
	sb.AddScore("Alice", 5)

	if sb.GetScore("Alice") != 15 {
		t.Errorf("Alice score should be 15, got %d", sb.GetScore("Alice"))
	}
	if sb.GetScore("Bob") != 20 {
		t.Errorf("Bob score should be 20, got %d", sb.GetScore("Bob"))
	}
	if sb.GetScore("Charlie") != 0 {
		t.Errorf("Charlie score should be 0, got %d", sb.GetScore("Charlie"))
	}

	all := sb.GetAllScores()
	if len(all) != 2 {
		t.Errorf("Should have 2 players, got %d", len(all))
	}
}

func TestScoreBoardConcurrent(t *testing.T) {
	sb := NewScoreBoard()
	var wg sync.WaitGroup

	// 100 goroutines each adding 1 point to "Player"
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			sb.AddScore("Player", 1)
		}()
	}
	wg.Wait()

	if sb.GetScore("Player") != 100 {
		t.Errorf("After 100 concurrent adds, score should be 100, got %d", sb.GetScore("Player"))
	}
}
