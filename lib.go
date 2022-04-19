package feistel

import (
    "math"
    "hash/fnv"
)

// collision-free pseudorandom sequence using a feistel cipher in cycle walking mode
// useful for avoiding hotspots in free-lists, or otherwise obfuscating a unique index.
// this is NOT encryption. it just makes a sequence look random to the human eye
//
// for example this generates a sequence from 0 to 1000 in deterministic but random-looking order
//
//    var max uint32 = 1000;
//    for i:=uint32(0);i<max;i++ {
//        fmt.Println(feistel(i, max, "feistel is cool"));
//    }
//
// outputs all numbers between 0 and 1000 like [2, 821, 76, etc..
// when using a different seed, it creates a different sequence
// remember this is not actual cryptography and a determined attacker will figure out the seed
//
func Map(index uint32, bound uint32, seed string) uint32 {

    if index >= bound {
        panic("feistel with index>=bound would never terminate");
    }

    var bias float64 = 1.0
    if seed != "" {
        var h = fnv.New32()
        h.Write([]byte(seed))
        bias = float64(h.Sum32()) / float64(^uint32(0))
    }

    var blocksize = math.Ceil(math.Log2(float64(bound) + 1)/8.0) * 8.0

    for i := float64(0); i < 2*math.Pow(2, float64(blocksize)); i++ {
        var r = feistel(uint8(blocksize), index, bias)

        if r < bound {
            return r;
        }
        index = r;
    }

    panic("BUG: bounded feistel wasnt bounded after all");
}

func feistel(blocksize uint8, index uint32, seed float64) uint32 {

    var shift = blocksize/2;
    var mask  = math.Pow(2, float64(blocksize)/2) - 1;

    var val = int32(index);

    var l1 = (val >> shift) & int32(mask);
    var r1 = val & int32(mask);
    var l2, r2  int32;

    for i := 0; i < 3; i++ {
        var xp = (seed + float64( (1366 * r1 + 150889) % 714025) / 714025.0) / 2.0
        l2 = r1;
        r2 = l1 ^ int32(math.Round(xp * mask));
        l1 = l2;
        r1 = r2;
    }

    var r = uint32((l1 << shift) + r1);
    return r;
}

