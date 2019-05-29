
# shamirssgo

A Go Implementation of Shamir's Secret Sharing.

### How to get it?

```go
$ go get -u "github.com/hesahesa/shamirssgo"
```

### What is this?

This is an implementation of Shamir's Secret Sharing in the Go programming language.

### How to use?

Supposed that you have a secret that is encoded as an integer (in this case: 58034)
```go
secret  := big.NewInt(58034)
```
pick a modulus that are a prime number (or prime enough) that are greater than our secret
```go
modulus  := big.NewInt(1000003) // this is a prime number
```
Let's say that you want to share this secret with 5 of your most trusted friends but they can only recover this secret if and only if at least 3 of them work together.
```go
threshold     :=  3 // min number of friends required
shamirSecret  := shamirssgo.New(secret, threshold, modulus)

shares        :=  make([]*big.Int, 6) // 6 is 5 (#friends) + 1
// notices that we don't need to get shares[0], this is intended
shares[1], _  = shamirSecret.Shares(1)
shares[2], _  = shamirSecret.Shares(2)
shares[3], _  = shamirSecret.Shares(3)
shares[4], _  = shamirSecret.Shares(4)
shares[5], _  = shamirSecret.Shares(5)

log.Println(shares)
// [<nil> 490906 205250 201069 478363 37129] (you will get different numbers of course)
```
each of your 5 friends then get their respective share (note that they don't know the secret number).
Supposed that 3 of your friend gather and want to recover the secret
```go
selectedShares   :=  make(map[int]*big.Int, 3)
// friend number 1, 4, and 5
selectedShares[1] = shares[1]
selectedShares[4] = shares[4]
selectedShares[5] = shares[5]
```
They can then recover your secret number
```go
computedSecret, _  := shamirssgo.ReconstructSecret(selectedShares, modulus)
// computedSecret = 58034
```
### Other programming language

[https://github.com/hesahesa/shamir-secretshare](https://github.com/hesahesa/shamir-secretshare) (in Java)

### References

- Shamir, Adi. "How to share a secret." Communications of the ACM 22.11 (1979): 612-613.Shamir, Adi. "How to share a secret." Communications of the ACM 22.11 (1979): 612-613.
- Cramer, Ronald, and Ivan Bjerre Damgård. "Secure multiparty computation." (2015).
- Hoepman, Jaap-Henk. "Privacy Friendly Aggregation of Smart Meter Readings, Even When Meters Crash." Proceedings of the 2nd Workshop on Cyber-Physical Security and Resilience in Smart Grids. ACM, 2017.

###### Made with <3 by [@hesahesa]

[@hesahesa]: <http://prahesa.id>
