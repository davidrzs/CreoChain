---
layout: page
title: Getting Sarted
subtitle: Building a blockchain based land registry.
---

<p class="message">
  Thank you for your interest. This project is still at a very early stage. It is intended for personal use and should NEVER be used where security is your primary concern.
</p>


## Introduction

Let's assume we live in a kleptocracy where you cannot be sure whether the land you have bought 20 years ago is still registered on your name or if some corrupt politician has changed the official record and claimed your land.

Nowadays, we just trust our governments to not tamper with our record; this is fine for now but it raises the questions if there isn't a better solution to our problem. What if there exists a method writing something down, that cannot be changed in the future? Using CreoChain we can do just that.

But enough hypothetical writing, let's get started:

## Installing CreoChain on a server

Go to the downloads section of this website (can be found in the sidebar on the left). Download the version that fits your computer architecture and operating system.

Place the executable in a folder, on Linux this could be `/var/www/creochain`. For MacOS and Windows just ensure that you place the executable in a folder named `/creochain`

We now need to add a configuration file, where the server will be able to read all the settings it needs.
Within `/creochain` create a new file called `config.yml`. This will be a `yaml` file.

Open the `config.yml` file in a text editor of your choice (Notepad++, Notepad, Atom, Sublime etc.) and start to add your configuration in `yaml` format. Keep in mind that `yaml` doesn't support tabs but only supports spaces, if you do use tabs you risk that the server will not be able to parse your configuration.

An example `config.yml` file can be seen below:

```yaml
server:
  globalauthcode: "auth"
  usessl: false
database:
  adapter: sqlite
  path: "./main.db"
```
This configuration file will start the server with the global authcode "auth" (more on this later), will not use SSL and will have an SQLite backend where the database is located in the same directory at `main.db`

## Compiling CreoChain on your own.

It is reccomended that you use a linux server for this.

First you need to install Go's dependency management tool called `dep`.

You can install it on your machine by executing the following command:
```bash
https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
```

After it has succeeded, clone the CreoChain repository onto your computer.
You can use the following Git URL: [https://github.com/davidrzs/CreoChain.git](https://github.com/davidrzs/CreoChain.git)

Now `cd` into the `CreoChain` directory that you have just downloaded and run `dep ensure` in order to get all the necessary dependencies onto your computer. If something goes wrong with the SQLite dependency its probably because
you do not have a 64-bit GCC compiler.

To build the project just type:

```bash
go build main/main.go
```

Et voilà, you now have a working CreoChain executable.
