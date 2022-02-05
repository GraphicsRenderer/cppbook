CPPBOOK
---

`cppbook` is a small tool I created for inlining `c++` source code into `markdown` file.

`cppbook` finds the `cpp` code blocks inside the markdown file and compile to an executable then run it automatically.

For example, a `c++` code block below

```c++, hello.cpp
#include <stdio.h>

int main() {
    printf("Hello World!");
}
```

By calling `cppbook.exe README.md`, it shows

```
g++ -Wall -O3 -o README.exe README.cpp
./README.exe
Hello World!
```

You may also define your own filename by adding the filename behine the language tag.

<pre>
```cpp, hello2.cpp
#include ...
...
```
</pre>
