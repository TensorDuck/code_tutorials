# Java Demo

Collection of random Java implementations for learning more about Java.
Previously studied some Java during my undergrad days, so mostly relearning.

# Helpful Commands

```
javac *.java
```

Will compile all the Java files inside this file.

```
java HelloWorld
```

Will run the compiled class.

# Fun Things I Learned

Coming from a mostly Python background, some fun gotchas I found when relearning
Java.

## Java primitives are like immutables but not really

I think a good mental model is that the primitives are what can actually be
stored in memory on a computer. This is sort of like Python immutables `float`,
`int`, `bool`, `set`, minus `set`.

They also must have a value initialized! or else the compiler will freak out. So
in Python, it would be quite natural to do:
```
a = None
if a is None:
    print(1)
else:
    print(0)
```

But I couldn't find a way to check of a primitive is `Null` or not in Java.

## `this` in Java is `self` in Python ... but rarely used

It seems like a lot of stackoverflow topics talk about how `this` keyword lets
 you reference specifically a value inside the object or from the input to a
 method. But the default behavior is to just assume the object's variables are
 being refernced. Thus, it is not often used and I haven't been able to find a
 good standard.

```
public void setVal(int x) {
    this.x = x; // distinguishes the
}

public int getVal() {
    return x; // returns this.x automatically
}
```

# Silly feelings I developed

## I hate semicolons
Python is a much less verbose language to learn with.
