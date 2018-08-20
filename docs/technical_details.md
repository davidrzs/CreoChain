---
layout: page
title: Technical Details
---

<p class="message">
The following bullet points explain which technologies are used in CreoChain.
</p>

-	CreoChain is written in [Go](https://golang.org/), this guarantees you memory safety and high performance.

-	As a database you can choose between SQLite and PostgreSQL, mySQL and msSQL will be supported in the future.

-	As a server we are using [Gin](https://github.com/gin-gonic/gin).

-	For testing we use Ruby scripts.

-	To model our blockchain we are using [GORM](https://github.com/jinzhu/gorm).

-	Last but not least, all variables in the server are completely thread save using mutexes.
