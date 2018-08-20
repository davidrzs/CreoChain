---
layout: default
title: CreoChain
subtitle: The Easily Programmable Blockchain Server
---
<h1 class="hasSubtitle">CreoChain</h1>
<h2 class="subtitle">An experimental non-distributed blockchain server with guaranteed data integrity.
</h2>
<br/>


The CreoChain server is a simple to use blockchain server which can savely store your data with built in security measures to ensure data integrity.

Think of it like an advanced noSQL database built on blockchain technology.
You can deploy it on any server on any operating system and any architecture, being built on top of Golang gives you complete flexibility.

It provides you with a JSON-API which you can use to:

* define new databases
* store data in a database
* get checksums to check data integrity (more on this later)


Since it is built on blockchain technology, once a record has been created it can't be deleted or changed.
Every time a new object is added to the database, the server will return a 256-bit long hash. Make sure to keep this hash somewhere, it is the key to making sure the blockchain hasn't been manipulated.

If you want to check the integrity of the data you can either let the sever check itself or if you want to make sure that the server code itself hasn't been manipulated, request all the data stored in the chain (using the simple API) and recalculate all the hashes by yourself. If the hashes don't match, your data has been manipulated.


## Notes on the state of this project

This software is not production ready, it's mainly a project to investigate whether the concept of storing data in a non distributed blockchain is viable. If it is, a more thorough and supported version will be released, if not the project will stay in non production ready.


## Copying, Forking, Copy-Pasting &amp; Enhancing

All the code is licensed under the very liberal MIT license. Feel free to do anything with it as long as it complies with the license. If you discover bugs or have patches you can use GitHub to improve this project.
