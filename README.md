
# Envoy

A cosmos module that will be able to instantiate "named locks" so that an envoy participant of a decentralized cluster can obtain a "named lock" in order to perform a predefined action in a "singleton like" manner.

# References

Inspired by cosmos checkers tutorial which describes the process of wiring up modules https://tutorials.cosmos.network/hands-on-exercise/0-native/2-build-module.html

## chain minimal test setup

```shell
minid keys list --keyring-backend test
```
```shell
- address: mini16ajnus3hhpcsfqem55m5awf3mfwfvhpp36rc7d
  name: alice
  pubkey: '{"@type":"/cosmos.crypto.secp256k1.PubKey","key":"A0gUNtXpBqggTdnVICr04GHqIQOa3ZEpjAhn50889AQX"}'
  type: local
- address: mini1hv85y6h5rkqxgshcyzpn2zralmmcgnqwsjn3qg
  name: bob
  pubkey: '{"@type":"/cosmos.crypto.secp256k1.PubKey","key":"ArXLlxUs2gEw8+clqPp6YoVNmy36PrJ7aYbV+W8GrcnQ"}'
  type: local
```

## create a lock

```shell
minid tx envoy create lock1 mini16ajnus3hhpcsfqem55m5awf3mfwfvhpp36rc7d 666 12 --from alice --yes
```

## read a lock

```shell
minid query envoy get-lock lock1
```
```
lock:
  at_block: 666
  envoy: mini16ajnus3hhpcsfqem55m5awf3mfwfvhpp36rc7d
  name: lock1
  num_blocks: 12
```
