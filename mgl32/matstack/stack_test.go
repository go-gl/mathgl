package matstack

import (
	"fmt"
	"github.com/go-gl/mathgl/mgl32"
	"testing"
)

func TestStackNew(t *testing.T) {
	stack := NewMatStack()

	if !(*stack)[0].ApproxEqual(mgl32.Ident4()) {
		t.Errorf("Cannot construct stack correctly")
	}
}

func TestStackPushPopPeek(t *testing.T) {
	stack := NewMatStack()

	if !stack.Peek().ApproxEqual(mgl32.Ident4()) {
		t.Errorf("Peek not working")
	}

	stack.Push(mgl32.HomogRotate3DY(mgl32.DegToRad(90)))

	if !stack.Peek().ApproxEqual(mgl32.HomogRotate3DY(mgl32.DegToRad(90))) {
		t.Errorf("Peek not working")
	}

	if stack.Len() != 2 {
		t.Errorf("Peek alters stack length")
	}

	pop, err := stack.Pop()
	if err != nil || !pop.ApproxEqual(mgl32.HomogRotate3DY(mgl32.DegToRad(90))) {
		t.Errorf("Pop is unsuccessful")
	}

	if stack.Len() != 1 {
		t.Errorf("Pop does not actually shorten stack")
	}

	_, err = stack.Pop()

	if err == nil {
		t.Errorf("Popping stack with 1 element does not return error as expected")
	}
}

func TestStackMultiPush(t *testing.T) {
	stack := NewMatStack()

	scale := mgl32.Scale3D(2, 2, 2)
	rot := mgl32.HomogRotate3DY(mgl32.DegToRad(90))
	trans := mgl32.Translate3D(4, 5, 6)

	stack.Push(trans)
	stack.Push(rot)

	if !stack.Peek().ApproxEqualThreshold(trans.Mul4(rot), 1e-4) {
		t.Errorf("Stack does not multiply first two pushes correctly")
	}

	stack.Push(scale)

	if !stack.Peek().ApproxEqualThreshold(trans.Mul4(rot).Mul4(scale), 1e-4) {
		t.Errorf("Stack does not multiple third push correctly")
	}

	stack.Unwind(2)
	stack.Push(scale)

	if !stack.Peek().ApproxEqualThreshold(trans.Mul4(scale), 1e-4) {
		t.Errorf("Unwinding and multiplying does not work correctly")
	}
}

func TestReseed(t *testing.T) {
	stack := NewMatStack()

	scale := mgl32.Scale3D(2, 2, 2)
	rot := mgl32.HomogRotate3DY(mgl32.DegToRad(90))
	trans := mgl32.Translate3D(4, 5, 6)

	stack.Push(trans)
	stack.Push(rot)
	stack.Push(scale)

	trans2 := mgl32.Translate3D(1, 2, 3)
	err := stack.Reseed(1, trans2)

	if err != nil {
		t.Fatalf("Rebase returned error when it should not %v", err)
	}

	if !stack.Peek().ApproxEqualThreshold(trans2.Mul4(rot).Mul4(scale), 1e-4) {
		t.Fatalf("Rebase does not remultiply correctly. Got\n %v expected\n %v. (Previous state:\n %v)", stack.Peek(), trans2.Mul4(rot).Mul4(scale), trans.Mul4(rot).Mul4(scale))
	}
}

func ExampleReseed() {
	stack := NewMatStack()

	scale := mgl32.Scale3D(2, 2, 2)
	rot := mgl32.HomogRotate3DY(mgl32.DegToRad(90))
	trans := mgl32.Translate3D(4, 5, 6)

	stack.Push(trans)
	stack.Push(rot)
	stack.Push(scale)

	fmt.Println("Initial state:\n", stack.Peek())

	trans2 := mgl32.Translate3D(1, 2, 3)

	err := stack.Reseed(1, trans2)
	if err == nil {
		panic("Rebase failed")
	}

	fmt.Println("After rebase:\n", stack.Peek())
	fmt.Println("Should be:\n", trans2.Mul4(rot).Mul4(scale))
}
