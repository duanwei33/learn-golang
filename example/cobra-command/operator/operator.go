package operator

import (
	"fmt"
	"time"
)

// RunOperator simulates starting the operator
func RunOperator() {
	fmt.Println("Starting My Operator...")
	time.Sleep(2 * time.Second) // Simulate some work
	fmt.Println("Operator is running successfully!")
}

