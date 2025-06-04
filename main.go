package main

import (
	"fmt"
)

// Define the problem structure
type Problem struct{}

// Objective function: x^2 + y^2
func (p *Problem) Objective(x []float64, new_x bool) float64 {
	return x[0]*x[0] + x[1]*x[1]
}

// Gradient of the objective function
func (p *Problem) Gradient(x []float64, new_x bool, grad_f []float64) {
	grad_f[0] = 2 * x[0] // df/dx = 2x
	grad_f[1] = 2 * x[1] // df/dy = 2y
}

// Constraints: x + y <= 1 (formulated as x + y - 1 <= 0)
func (p *Problem) Constraints(x []float64, new_x bool, cons []float64) {
	cons[0] = x[0] + x[1] - 1.0 // x + y - 1 <= 0
}

// Jacobian of the constraints (sparse format)
func (p *Problem) Jacobian(x []float64, new_x bool, iRow, jCol []int, values []float64) {
	if values == nil {
		// Structure of the Jacobian (constant for this problem)
		iRow[0] = 0 // constraint 0
		jCol[0] = 0 // variable x
		iRow[1] = 0 // constraint 0
		jCol[1] = 1 // variable y
	} else {
		// Values of the Jacobian
		values[0] = 1.0 // d(g0)/dx = 1
		values[1] = 1.0 // d(g0)/dy = 1
	}
}

// Hessian of the Lagrangian (sparse format)
func (p *Problem) Hessian(x []float64, new_x bool, obj_factor float64, lambda []float64, new_lambda bool, iRow, jCol []int, values []float64) {
	if values == nil {
		// Structure of the Hessian (constant for this problem)
		iRow[0] = 0 // diagonal element xx
		jCol[0] = 0
		iRow[1] = 1 // diagonal element yy
		jCol[1] = 1
	} else {
		// Values of the Hessian (2x2 identity matrix multiplied by 2)
		values[0] = 2.0 * obj_factor // d²f/dx² = 2
		values[1] = 2.0 * obj_factor // d²f/dy² = 2
	}
}

func main() {
	// Create the problem
	problem := &Problem{}

	// Number of variables and constraints
	n := 2 // x and y
	m := 1 // one constraint (x + y <= 1)

	// Lower and upper bounds for variables
	x_L := []float64{0.0, 0.0} // x >= 0, y >= 0
	x_U := []float64{Ipopt.Infinity(), Ipopt.Infinity()}

	// Lower and upper bounds for constraints
	g_L := []float64{-Ipopt.Infinity()} // x + y - 1 <= 0
	g_U := []float64{0.0}

	// Create the Ipopt problem
	ipopt := Ipopt.NewIpoptProblem(n, x_L, x_U, m, g_L, g_U, 8, 10, problem)
	defer ipopt.Free()

	// Set options
	ipopt.AddOption("tol", 1e-7)
	ipopt.AddOption("mu_strategy", "adaptive")
	ipopt.AddOption("print_level", 5)

	// Initial guess
	x0 := []float64{0.5, 0.5} // initial point

	// Solve the problem
	status, x := ipopt.Solve(x0)

	// Output results
	fmt.Printf("\nSolution status: %v\n", status)
	fmt.Printf("Optimal point: x = %.6f, y = %.6f\n", x[0], x[1])
	fmt.Printf("Maximum value: %.6f\n", x[0]*x[0]+x[1]*x[1])
}
