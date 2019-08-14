### Patchctl Commands

To create a repository of data called myrepo, you can run this code:

```sh
$ pachctl create-repo myrepo
```

You can then confirm that the repository exists with list-repo:


```sh
$ pachctl list-repo

  NAME CREATED SIZE 
  myrepo 2 seconds ago 0 B
```

### Putting Data into data repositories

Let's say that we have a simple text file:

```sh
$ cat blah.txt 

This is an example file.
```

If this file is part of the data we are utilizing in our ML workflow, we should version it. To version this file in our repository, myrepo, we just need to commit it into that repository:

```sh
$ pachctl put-file myrepo master -c -f blah.txt
```

As a sanity check, we can confirm that our file was versioned in the repository:

```sh
$ pachctl list-repo

NAME CREATED SIZE 
myrepo 10 minutes ago 25 B 
```

```sh
$ pachctl list-file myrepo master

NAME TYPE SIZE 
blah.txt file 25 B
```
However, if we manually want to pull certain sets of versioned data out of Pachyderm, analyze them interactively, then we can use the pachctl CLI to get data:


```sh
$ pachctl get-file myrepo master blah.txt
This is an example file.
```

### License
```sh
MIT License

Copyright (c) 2019 Manuh.__

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
```

