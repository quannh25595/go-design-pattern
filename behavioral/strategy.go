package behavioral

type Operator interface {
	Apply(int, int) int
}

type Operation struct {
	Operator Operator
}

func (o *Operation) Operate(left, right int) int {
	return o.Operator.Apply(left, right)
}

type Addition struct{}

func (Addition) Apply(left, right int) int {
	return left + right
}

type Multiplication struct{}

func (Multiplication) Apply(left, right int) int {
	return left + right
}
