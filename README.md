superhash
========

[![Build Status](https://travis-ci.org/esco/superhash-node.svg?branch=master)](https://travis-ci.org/esco/superhash-go) [![Coverage Status](https://coveralls.io/repos/esco/superhash/badge.png)](https://coveralls.io/r/esco/superhash-go) [![Code Climate](https://codeclimate.com/github/esco/superhash/badges/gpa.svg)](https://codeclimate.com/github/esco/superhash-go)

![superhash](https://i.imgur.com/mVcwsbC.png)

Thread safe Hash Map that supports using one or more keys of any type

## Installation

```
$ go get esco/superhash-go
```

## Example

```go
import (
    "github.com/esco/superhash"
)

hashmap := superhash.New()
k1, k2, k3, value := 1, true, "3", 4
hashmap.Set(k1, k2, k3, value)
hashmap.Get(k1, k2, k3) // 4
hashmap.Delete(k1, k2, k3) 
```

## LICENSE
[MIT][license-url]

[license-url]: LICENSE
