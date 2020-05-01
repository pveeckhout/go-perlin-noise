# go-perlin-noise
Basic implementation of a noise generator in golang based on [Ken Perlin's Improved Perlin Noise](http://mrl.nyu.edu/~perlin/noise/),  published in 2002.
This does not handle randomness, the generated value for each input is deterministic.

Randomness needs to be introduced in the input, the method of which is left to the user.

# Install
```bash
go get github.com/pveeckhout/go-perlin-noise
```

# usage
## Perlin
generates a perlin noise value for the x, y and z coordiantes passed.

## OctavePerlin
While perlin noise generates natural randomness to a degree, it cannot does not cover teh full range of variability 
expected in nature. For example: terrain has large feature such as mountains, but also smaller variations such as hills
and depressions, rocks, ...

To cover this you take multiple noise functions with varying frequencies and amplitudes, and add them together.
  

