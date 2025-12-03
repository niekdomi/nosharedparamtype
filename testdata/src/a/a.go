// Package a contains test cases.
package a

func funcWithSharedParamTypes(x, y int) int { // want "function funcWithSharedParamTypes has type sharing parameters"
	return x + y
}

func funcWithoutSharedParamTypes(x int, y int) int {
	return x + y
}

func funcWithThreeSharedParams(a, b, c string) string { // want "function funcWithThreeSharedParams has type sharing parameters"
	return a + b + c
}

func funcWithMixedParams(a string, b, c int) string { // want "function funcWithMixedParams has type sharing parameters"
	return a
}

func funcWithSharedPointers(x, y *int) int { // want "function funcWithSharedPointers has type sharing parameters"
	return *x + *y
}

func funcWithSharedSlices(x, y []string) int { // want "function funcWithSharedSlices has type sharing parameters"
	return len(x) + len(y)
}

func funcWithSingleParam(x int) int {
	return x
}

func funcWithNoParams() int {
	return 42
}

type MyType struct{}

func (m MyType) methodWithSharedParams(x, y int) int { // want "function methodWithSharedParams has type sharing parameters"
	return x + y
}

func (m MyType) methodWithSeparateParams(x int, y int) int {
	return x + y
}
