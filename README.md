A simple guestbook web app demo based on [this tutorial](http://shadynasty.biz/blog/2012/07/30/quick-and-clean-in-go/).

I just eliminated the Mongo dependency and instead used [Bolt](https://github.com/boltdb/bolt) for persistence.

To get the code and run the server, try ...

    $ go get github.com/joyrexus/guestbook
    $ cd $GOPATH/src/github.com/joyrexus/guestbook
    $ go run *go

Then open `http://localhost:8080/book` and try signing it.
