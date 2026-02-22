# Exercise 4: Go (Golang) Basics 🐹

> **Goal:** Learn Go basics you'll need for ~20% of Product Track tasks.
> Go is used for concurrent systems, so this exercise covers goroutines and mutexes too.

---

## How This Exercise Works

1. Open `starter.go`
2. Read each section and fill in the TODO parts
3. Run the tests:
   ```bash
   go test -v
   ```
4. For the concurrency part, also run with the race detector:
   ```bash
   go test -race -v
   ```

---

## Why Go?

Go is Google's language, widely used for backend services, DevOps tools, and cloud infrastructure. It's known for being simple and fast with built-in concurrency.

## Key Differences from Python/JS

| Python/JS | Go |
|-----------|-----|
| Dynamic types | Static types (you must declare types) |
| `class` keyword | `struct` + methods |
| `try/except` | Multiple return values `(result, error)` |
| Threads are hard | Goroutines are easy! |
| `self` or `this` | Receiver `(c *Calculator)` |
| Garbage collected | Garbage collected ✅ |
| `None` / `null` | `nil` |
