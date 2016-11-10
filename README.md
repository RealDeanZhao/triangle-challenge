# triangle-challenge
I am learning golang currently, so I choose golang to implement this application.

## Requirement
Write a program that will determine the type of a triangle. It should take the lengths of the triangle's three sides as input, and return whether the triangle is equilateral, isosceles or scalene.

## Business Analysis
A triangle must meet the following conditions:

1. Any side should be a positive number.
2. The sum of any two sides should be greater than the third.

## Technical Design

### Programming language:
* go

### Environment:
* Ubuntu

### Unit Test:
* The go test tool

### Project structure:
It is a simple application but it also can be modulized. The main logic can be splitted into serveral small pieces.
So it will be much easier to maintain, read and test.

**main.go**: Application entry.

**app/app.go** : Functions that exported to determine if it is a triangle.
* DetermineTriangleType: Determine the type of the triangle.
* IsTriangle: Check if it is a triangle or not first.
* AllSidesArePositive: Check if all the sides are positive numbers.
* AnyTwoSidesAreGreaterThanTheThrid: Check if the sum of any two sides are greater than the third side or not. **There might be a float64 overflow.**
* AllSidesAreEqual: Check all the sides are equal or not.
* TwoSidesAreEqual: Check any two sides are equal or not.

**app/app_test.go**: Unit Test, the go test tool seems doesn't have something like the assertion in other programming language.
```
➜  app git:(master) ✗ go test
PASS
ok  	_/home/dean/CodeRepos/triangle-challenge/app	0.003s
```

### Run and Test:
Run the application:
```
➜  triangle-challenge git:(master) ✗ go run main.go
0
0
0
1
2
3
```
Test the application:
```
➜  triangle-challenge git:(master) ✗ cd app
➜  app git:(master) ✗ go test 
PASS
ok  	_/home/dean/CodeRepos/triangle-challenge/app	0.002s
```


