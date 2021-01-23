// Additional math funcions
package math

// Calculate a Fibonacci number
func Fibonacci(n int) int {
        if n < 2 {
                return n
        }
        return Fibonacci(n-1) + Fibonacci(n-2)
}
