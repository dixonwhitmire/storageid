# storageid

Generates a valid "ID" for Azure Storage Accounts. A valid Azure Storage Account identifier:

* Is a maximum of 24 characters.
* Consists of lowercase alpha or numeric characters.

I use this project a bit in my day to day job, and am evolving it as I learn more about Rust using the 
[Interactive Book](https://rust-book.cs.brown.edu/).

## Quickstart

```shell
storageid % cargo run
   Compiling storageid v0.1.0 (/Users/dwhitmire/code/storageid)
    Finished dev [unoptimized + debuginfo] target(s) in 0.16s
     Running `target/debug/storageid`
Please enter a storage account prefix or hit enter for None
int
New ID = intWWInIkeTQbBrlab179NC3
```