# 🔐 Random Password Generator (Go)

A secure, policy-enforced password generator written in Go, designed with cryptographic correctness, clean architecture, and benchmark-driven decisions.

This project demonstrates how to build a production-ready CLI tool in Go while balancing security, performance, and simplicity.

---

## ✨ Features

- Cryptographically secure password generation using `crypto/rand`
- Enforces strong password policy:
  - At least 1 lowercase letter
  - At least 1 uppercase letter
  - At least 1 digit
  - At least 1 special character
- Configurable password length and count via CLI flags
- Clean separation between CLI and core logic
- Unit tests and benchmarks included
- Designed with explicit performance and security trade-offs

---

## 📁 Project Structure

```
RandomPasswordGenerator/
├── main.go
└── password/
    ├── password_generator.go
    ├── password_generator_test.go
    └── password_generator_benchmarking_test.go
```

| File | Purpose |
|----|----|
| main.go | CLI entry point |
| password_generator.go | Core password generation logic |
| password_generator_test.go | Unit tests |
| password_generator_benchmarking_test.go | Benchmarks |

---

## 🚀 Usage

### Build

```bash
go build -o passwordgen
```

### Run

```bash
./passwordgen -length=16 -count=3
```

### CLI Flags

| Flag | Description | Default |
|----|----|----|
| -length | Length of each password (minimum 8) | 16 |
| -count | Number of passwords to generate | 1 |

### Example Output

```
G9@kP!zA8eR2L#Dq
wF4%QZ@9X&kM8A!
```

---

## 🧠 Design Rationale

### 1. Why `crypto/rand` instead of `math/rand`?

Passwords are security-critical.

`math/rand` is deterministic and predictable, even when seeded.

This project uses:

```go
crypto/rand
```

to ensure:

- Unpredictable output
- OS-backed entropy
- Resistance to replay and prediction attacks

---

### 2. Policy-First Generation

Instead of generating random strings and validating later, the generator:

- Explicitly places one character from each required category
- Fills the remaining characters randomly
- Shuffles the final byte slice

This guarantees:

- Strong passwords
- No retry loops
- Predictable performance

---

### 3. Secure Shuffling Strategy

Characters are selected using `crypto/rand`.
The final password is shuffled to avoid predictable character positions.

This design avoids:

- Deterministic layouts
- Policy leakage (e.g., always first char is lowercase)

---

### 4. Why No Goroutines?

Benchmarks showed that:

- Password generation is entropy-bound, not CPU-bound
- Goroutines add overhead without measurable benefit

Sequential execution is simpler and faster for a CLI tool.

Concurrency was intentionally avoided to keep:

- Code simpler
- Performance predictable
- Debugging easy

---

## 🧪 Testing

Run unit tests:

```bash
go test ./...
```

Tests cover:

- Minimum length validation
- Correct password length
- Error handling
- Randomness sanity checks

---

## 📊 Benchmarking

### Run Benchmarks

```bash
go test ./password -bench=. -benchmem -run=^$ -count=1
```

### Sample Benchmark Results

**Environment**
- OS: Windows
- CPU: Intel i5-1135G7

#### Generate 16-character password

```
BenchmarkGenerate16-8
168,338 iterations
7,227 ns/op
1,520 B/op
95 allocs/op
```

#### Generate 64-character password

```
BenchmarkGenerate64-8
35,308 iterations
30,819 ns/op
6,224 B/op
383 allocs/op
```

---

## 📈 Benchmark Analysis

- Performance scales linearly with password length
- Allocation count increases predictably with cryptographic operations
- Trade-off favors security over micro-optimizations
- Suitable for CLI tools, batch jobs, and backend services

---

## 🏆 Strengths

- Cryptographically sound
- Predictable and policy-enforced output
- Clean Go idioms and structure
- Benchmark-validated behavior
- Easy to extend or embed as a library

---

## 📜 License

MIT License
