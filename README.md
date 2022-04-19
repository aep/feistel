collision-free pseudorandom sequence using a feistel cipher in cycle walking mode.

useful for avoiding hotspots in free-lists, or otherwise obfuscating a unique index.
This is NOT encryption. it just makes a sequence look random to the human eye

for example this generates a sequence from 0 to 1000 in deterministic but random-looking order

```golang
package main

import (
    "github.com/aep/feistel"
    "fmt"
)

func main() {
    var max uint32 = 1000;
    for i:=uint32(0);i<max;i++ {
        fmt.Println(feistel.Map(i, max, "feistel is cool"));
    }
}
```

outputs all numbers between 0 and 1000 like [2, 821, 76, etc..

When using a different seed, it creates a different sequence.
Remember this is not actual cryptography and a determined attacker will figure out the seed

mostly inspired by and maybe compatible with https://wiki.postgresql.org/wiki/Pseudo_encrypt_constrained_to_an_arbitrary_range
