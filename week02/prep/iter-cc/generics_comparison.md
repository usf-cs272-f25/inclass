# Go vs Java Generics Comparison

## 1. Implementation Strategy

### Java: Type Erasure
- Generics exist only at compile time
- At runtime, all generic types become Object
- No runtime type information for type parameters
- Backward compatible with pre-generics Java code

### Go: Type Instantiation (Monomorphization)
- Compiler generates specialized code for each type used
- Generic type information preserved at runtime
- Similar to C++ templates
- Better performance, larger binary size

```java
// Java - Type erasure
List<String> list = new ArrayList<>();
// At runtime, this is just List of Object

// Cannot do this in Java:
if (list instanceof List<String>) {} // Compile error
```

```go
// Go - Type preserved
var list []string
// Runtime knows this is []string

// Can check types at runtime
var x any = []string{"a", "b"}
if _, ok := x.([]string); ok {} // Works
```

## 2. Syntax Differences

### Java: Angle Brackets <>
```java
public class Box<T> {
    private T value;
    public T get() { return value; }
    public void set(T value) { this.value = value; }
}

public <T> T first(List<T> list) {
    return list.get(0);
}
```

### Go: Square Brackets []
```go
type Box[T any] struct {
    value T
}

func (b Box[T]) Get() T { return b.value }
func (b *Box[T]) Set(v T) { b.value = v }

func First[T any](slice []T) T {
    return slice[0]
}
```

## 3. Type Constraints

### Java: Bounded Type Parameters
```java
// Upper bound with extends
public <T extends Number> double sum(List<T> numbers) {
    double total = 0;
    for (T n : numbers) {
        total += n.doubleValue();
    }
    return total;
}

// Multiple bounds
public <T extends Comparable<T> & Serializable> T max(T a, T b) {
    return a.compareTo(b) > 0 ? a : b;
}
```

### Go: Interface Constraints and Type Sets
```go
// Interface constraint
func Sum[T ~int | ~float64](numbers []T) T {
    var sum T
    for _, n := range numbers {
        sum += n
    }
    return sum
}

// comparable constraint
func Max[T comparable](a, b T) T {
    if a > b {
        return a
    }
    return b
}

// Custom constraint
type Number interface {
    ~int | ~int32 | ~int64 | ~float32 | ~float64
}
```

## 4. Wildcards vs No Wildcards

### Java: Has Wildcards
```java
// ? wildcard - unknown type
List<?> unknownList;

// Upper bounded wildcard
List<? extends Number> numbers;

// Lower bounded wildcard  
List<? super Integer> integers;

// Use-site variance
void printAll(List<? extends Object> list) {
    for (Object o : list) {
        System.out.println(o);
    }
}
```

### Go: No Wildcards
```go
// Go doesn't have wildcards
// Must use concrete types or type parameters

// Can't do this:
// var slice []? // No wildcard concept

// Instead, use interface{} or any
func PrintAll[T any](slice []T) {
    for _, v := range slice {
        fmt.Println(v)
    }
}
```

## 5. Variance

### Java: Use-site Variance (Wildcards)
```java
// Covariance with ? extends
List<? extends Animal> animals = new ArrayList<Dog>();

// Contravariance with ? super
List<? super Dog> dogs = new ArrayList<Animal>();

// Invariant by default
List<Dog> dogs = new ArrayList<Animal>(); // Error
```

### Go: Invariant Only
```go
// Go generics are invariant
type Animal struct{}
type Dog struct{ Animal }

// Can't do this:
var animals []Animal = []Dog{} // Error

// Must explicitly convert or use interfaces
```

## 6. Type Inference

### Java: Limited Type Inference
```java
// Diamond operator (Java 7+)
List<String> list = new ArrayList<>(); // <> infers String

// Method type inference
Collections.emptyList(); // Type inferred from context

// Sometimes requires explicit types
List<String> list = Collections.<String>emptyList();
```

### Go: More Aggressive Type Inference
```go
// Function calls infer types
result := First([]int{1, 2, 3}) // T inferred as int

// Struct literals (with limitations)
box := Box[int]{value: 42} // Must specify type

// Can often omit types entirely
Double(42)    // Infers int
Double(3.14)  // Infers float64
```

## 7. Primitives vs Objects

### Java: Boxing/Unboxing Issues
```java
// Can't use primitives directly
List<int> numbers; // Error
List<Integer> numbers; // Must use wrapper

// Performance overhead from boxing
ArrayList<Integer> list = new ArrayList<>();
list.add(42); // Boxing happens here
```

### Go: No Boxing Needed
```go
// Primitives work directly
var slice []int // Direct int, no boxing
slice = append(slice, 42) // No boxing

// Better performance for primitive types
type Stack[T any] []T
var s Stack[int] // Direct int storage
```

## 8. Reification

### Java: Type Erasure Problems
```java
// Can't create arrays of generic types
T[] array = new T[10]; // Error

// Can't check generic types at runtime
if (list instanceof List<String>) {} // Error

// Can't catch generic exceptions
try {} catch(T e) {} // Error where T is type parameter
```

### Go: Types Are Reified
```go
// Can create slices of generic types
func MakeSlice[T any](size int) []T {
    return make([]T, size) // Works fine
}

// Type information available at runtime
func TypeName[T any](v T) string {
    return fmt.Sprintf("%T", v)
}
```

## 9. Static Methods/Functions

### Java: Complex Rules
```java
public class Container<T> {
    // Static method can't use class type parameter
    public static <U> U staticMethod(U value) { // Own type param
        return value;
    }
    
    // Can't do this:
    // public static T invalid() {} // Error
}
```

### Go: Functions Are Top-Level
```go
// Go doesn't have static methods, just functions
func StaticLike[T any](value T) T {
    return value
}

// No confusion about type parameters
type Container[T any] struct{}

// Methods on generic types work naturally
func (c Container[T]) Method(v T) T {
    return v
}
```

## 10. Performance Implications

### Java
- Type erasure means one compiled version
- Boxing/unboxing overhead for primitives
- Virtual method calls through Object
- Smaller bytecode size

### Go
- Specialized code for each type (monomorphization)
- No boxing overhead
- Direct method calls
- Larger binary size but faster execution
- Better cache locality

## Summary Table

| Feature | Java | Go |
|---------|------|-----|
| Implementation | Type Erasure | Monomorphization |
| Syntax | `<T>` | `[T]` |
| Wildcards | Yes (`? extends`, `? super`) | No |
| Variance | Use-site (wildcards) | Invariant only |
| Type at runtime | No (erased) | Yes (reified) |
| Primitives | Must box | Direct use |
| Type inference | Limited | More aggressive |
| Performance | Boxing overhead | Faster, larger binary |
| Constraints | `extends`/`implements` | Type sets/interfaces |
| Methods with type params | Not allowed | Not allowed |

## When to Use Which?

### Java Generics Excel At:
- Large existing codebases (backward compatibility)
- Complex type hierarchies with variance
- Smaller compiled size requirements
- Framework/library design with wildcards

### Go Generics Excel At:
- Performance-critical code
- Simple, straightforward generic code
- Working with primitives
- When you need runtime type information
- Systems programming scenarios