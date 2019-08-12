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

