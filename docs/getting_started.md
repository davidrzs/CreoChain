---
layout: page
title: Getting Started
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


## Auth Tokens and Database Administration

Database support for mySQL and SQLite is built into CreoChain. mySQL is expreimental only and should be used with caution, SQLite should work just fine.
In order to use SQLite create an emtpy file in the directory where the CreoChain executable will be executed and give it a name. You then have to tell CreoChain in the `config.yml` that you want to use that file as your database file.

In order to perform administrative work on your CreoChain server you need a masterpassword. Please choose a very secure password (e.g. at [https://passwordsgenerator.net/](https://passwordsgenerator.net/)) and place it in the `config.yml` file under `globalauthcode`.


# Programming the Example

We can now come back to our example. First we will start our CreoChain server by executing the executable. If you don't get any errors you are good to go. (You might be getting a `(near "CONSTRAINT": syntax error)` error, you can ignore it, it will be fixed in the future).

We will now take a look at the different operations CreoChain supports.

As a quick side note: CreoChain uses a versioned API, currently we are on v1, so all URLs will be prefixed with `/v1/`.

#### Creating a new Blochckain

If you want to create a new blockchain send a post request to `http://yourdomain:8080/v1/chain/` with the following JSON payload:

```json
{
	"name": "",
	"globalauthcode": "",
	"chainAccessToken": ""
}
```

- **name** is the name of the blockchain you are creating
- **globalauthcode** is the global authorization code you have specified in the `config.yaml` above.
- **chainAccessToken** is the authorization code which allows you to manipulate this blockchain (e.g. add blocks)

#### Adding a Block to a Blochckain

To add a block to a blockchain send a post request to `http://yourdomain:8080/v1/chain/chainname/`, where `chainname` is the name of the blockchain you have created in the previous step. Send it with the following JSON payload:

```json
{
	"data": "",
	"authcode": ""
}
```

- **data** the data you are storing in the blockchain. This can be anything, from binary, to JSON, literally anything. Just make sure you supply it as a string since it is stored as a string in the database.
- **authcode** is the chains chainAccessToken you have defined above, it is used to validate whether you are allowed to add blocks or not.


#### Getting a Copy of the Whole Blockchain

To see the whole blockchain send a get request to `http://yourdomain:8080/v1/chain/chainname/`, where `chainname` is the name of the blockchain. No JSON payload is required. 

This request will return a long JSON string with all blocks and a lot of other useful data on your blockchain.


#### Getting a Sinlge Block of Blockchain (experimental)

To see a single block of the blockchain send a get request to `http://yourdomain:8080/v1/chain/chainname/block/blockid`, where `chainname` is the name of the blockchain and `blockid` is the id-number of the block in the blockchain. No JSON payload is required. 

This request will return a JSON string with your block.

# Checking the correctness of the Blockchain

(More info will be added later)

Currently, in order to check correctness of your blockchain you can recalculate all the hashes with the following piece of Golang code. A simpler way of rechecking all the hashes will be supplied in the future.

``` go
func GetHash(b *Block) string {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers := bytes.Join([][]byte{[]byte(b.PrevBlockHash), []byte(b.Data), timestamp}, []byte{})
	hash := sha256.Sum256(headers)
	return b64.StdEncoding.EncodeToString(hash[:])
}
```
