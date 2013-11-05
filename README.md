[![Build Status](https://travis-ci.org/grantmd/go-rdio.png?branch=master)](https://travis-ci.org/grantmd/go-rdio)

go-rdio
=======

An API client for Rdio in Go (golang)

**This is not complete. All methods have been implemented but don't support optional arguments, etc.** If you would 
like to help, even just to add a single method or improve some documentation, please make _clean_ and _small_ pull 
requests with test cases. You will get credit for helping. Thanks! 

Usage
-----

* First, get yourself an Rdio api key: http://www.rdio.com/developers/

* Install the library:

        go get github.com/mrjones/oauth

* Include it in your project:

        import "github.com/grantmd/go-rdio"

* Setup your Rdio client:

        c := &rdio.Client{
                ConsumerKey:    config.ConsumerKey,
                ConsumerSecret: config.ConsumerSecret,
        }

* Authenticate against the Rdio API as a user:

        auth, err := c.StartAuth()
        // Redirect the user to Rdio and get their verifier
        // ...
        auth, err = c.CompleteAuth(string(verifier))
        albums, err := c.GetNewReleases()
        fmt.Printf("%#v\n", albums)

Examples
--------

There is an example in the `example/` directory. `go build` and then run `./example` for usage information.
