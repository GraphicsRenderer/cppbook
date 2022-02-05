Example
---

This is an example markdown file for inlining `c++` source code.

Let's create a function that prints `Hello World!`.

```cpp
#include <stdio.h>

void hello() {
    printf("Hello World!\n");
}
```

Now we can have the `main` function,

```cpp
int main() {
    hello();
    return 0;
}
```

Here we can try another `c++` code with filename attached behined.

```cpp, test.cpp
#include <stdio.h>

int main() {
    printf("Hello Test!\n");
}
```
