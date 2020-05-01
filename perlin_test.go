package perlin

import (
	"math"
	"math/rand"
	"sync"
	"testing"
)

const floatTolerance = 0.000001

func TestPerlin(t *testing.T) {
	type args struct {
		x float64
		y float64
		z float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "whole numbers",
			args: args{
				x: 0.0,
				y: 1.0,
				z: 2.0,
			},
			want: 0,
		}, {
			name: "1d",
			args: args{
				x: 0.0,
				y: 0.0,
				z: 4.75,
			},
			want: -0.22412109375,
		}, {
			name: "2d",
			args: args{
				x: 0.0,
				y: -8.42,
				z: 4.75,
			},
			want: -2.3982550411154984,
		}, {
			name: "3d",
			args: args{
				x: 0.1,
				y: -0.1,
				z: .7,
			},
			want: 0.301003,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Perlin(tt.args.x, tt.args.y, tt.args.z)
			if diff := math.Abs(got - tt.want); diff > floatTolerance {
				t.Errorf("Perlin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_grad(t *testing.T) {
	type args struct {
		hash int
		x    float64
		y    float64
		z    float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "hash: 0b0000",
			args: args{
				hash: 0b0,
				x:    100,
				y:    -1,
				z:    1000,
			},
			want: 99,
		}, {
			name: "hash: 0b0001",
			args: args{
				hash: 0b1,
				x:    100,
				y:    -1,
				z:    1000,
			},
			want: -101,
		}, {
			name: "hash: 0b0010",
			args: args{
				hash: 0b10,
				x:    100,
				y:    -1,
				z:    1000,
			},
			want: 101,
		}, {
			name: "hash: 0b0011",
			args: args{
				hash: 0b11,
				x:    100,
				y:    -1,
				z:    1000,
			},
			want: -99,
		}, {
			name: "hash: 0b0100",
			args: args{
				hash: 0b100,
				x:    100,
				y:    -1,
				z:    1000,
			},
			want: 1100,
		}, {
			name: "hash: 0b0101",
			args: args{
				hash: 0b101,
				x:    100,
				y:    -1,
				z:    1000,
			},
			want: 900,
		}, {
			name: "hash: 0b0110",
			args: args{
				hash: 0b110,
				x:    100,
				y:    -1,
				z:    1000,
			},
			want: -900,
		}, {
			name: "hash: 0b0111",
			args: args{
				hash: 0b111,
				x:    100,
				y:    -1,
				z:    1000,
			},
			want: -1100,
		}, {
			name: "hash: 0b1000",
			args: args{
				hash: 0b1000,
				x:    100,
				y:    -1,
				z:    1000,
			},
			want: 999,
		}, {
			name: "hash: 0b1001",
			args: args{
				hash: 0b1001,
				x:    100,
				y:    -1,
				z:    1000,
			},
			want: 1001,
		}, {
			name: "hash: 0b1010",
			args: args{
				hash: 0b1010,
				x:    100,
				y:    -1,
				z:    1000,
			},
			want: -1001,
		}, {
			name: "hash: 0b1011",
			args: args{
				hash: 0b1011,
				x:    100,
				y:    -1,
				z:    1000,
			},
			want: -999,
		}, {
			name: "hash: 0b1100",
			args: args{
				hash: 0b1100,
				x:    100,
				y:    -1,
				z:    1000,
			},
			want: 99,
		}, {
			name: "hash: 0b1101",
			args: args{
				hash: 0b1101,
				x:    100,
				y:    -1,
				z:    1000,
			},
			want: 1001,
		}, {
			name: "hash: 0b1110",
			args: args{
				hash: 0b1110,
				x:    100,
				y:    -1,
				z:    1000,
			},
			want: -101,
		}, {
			name: "hash: 0b1111",
			args: args{
				hash: 0b1111,
				x:    100,
				y:    -1,
				z:    1000,
			},
			want: -999,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := grad(tt.args.hash, tt.args.x, tt.args.y, tt.args.z)
			if diff := math.Abs(got - tt.want); diff > floatTolerance {
				t.Errorf("grad() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fade(t *testing.T) {
	type args struct {
		t float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "Zero",
			args: args{t: 0},
			want: 0,
		}, {
			name: "One",
			args: args{t: 1},
			want: 1,
		}, {
			name: "rnd_1",
			args: args{t: 1.5},
			want: +3.375000e+000,
		}, {
			name: "rnd_2",
			args: args{t: .75},
			want: +8.964844e-001,
		}, {
			name: "rnd_3",
			args: args{t: 1.05},
			want: +1.001346e+000,
		}, {
			name: "rnd_4",
			args: args{t: -1.5},
			want: -1.552500e+002,
		}, {
			name: "rnd_5",
			args: args{t: 0.4},
			want: 3.174400e-001,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := fade(tt.args.t)
			if diff := math.Abs(got - tt.want); diff > floatTolerance {
				t.Errorf("fade() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_lerp(t *testing.T) {
	type args struct {
		v0    float64
		v1    float64
		alpha float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "if alpha 1 then result = v1",
			args: args{
				v0:    rand.Float64(),
				v1:    .975,
				alpha: 1,
			},
			want: .975,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lerp(tt.args.v0, tt.args.v1, tt.args.alpha)
			if diff := math.Abs(got - tt.want); diff > floatTolerance {
				t.Errorf("lerp() = %v, want %v", got, tt.want)
			}
		})
	}

	// alternate algorithm, see if same result
	var v0, v1, alpha, got float64

	for i := 0; i < 25; i++ {
		// random values between -5.0 and 5.0
		v0 = (rand.Float64() - 0.5) * 10
		v1 = (rand.Float64() - 0.5) * 10
		alpha = (rand.Float64() - 0.5) * 10

		check := v0 + alpha*(v1-v0)
		got = lerp(v0, v1, alpha)

		if diff := math.Abs(got - check); diff > floatTolerance {
			t.Errorf("lerp() = %f; want %f", got, check)
		}
	}
}

func Benchmark(b *testing.B) {
	if testing.Short() {
		b.Skip("Skipping benchmark in short mode.")
	}

	benchmarks := []struct {
		name   string
		width  int
		height int
		depth  int
	}{
		{
			name:   "line of 64",
			width:  1,
			height: 1,
			depth:  64,
		}, {
			name:   "line of 128",
			width:  1,
			height: 1,
			depth:  128,
		}, {
			name:   "line of 256",
			width:  1,
			height: 1,
			depth:  256,
		}, {
			name:   "square 64",
			width:  1,
			height: 64,
			depth:  64,
		}, {
			name:   "square 128",
			width:  1,
			height: 128,
			depth:  128,
		}, {
			name:   "square 256",
			width:  1,
			height: 256,
			depth:  256,
		}, {
			name:   "cube 64",
			width:  64,
			height: 64,
			depth:  64,
		}, {
			name:   "cube 128",
			width:  128,
			height: 128,
			depth:  128,
		}, {
			name:   "cube 256",
			width:  256,
			height: 256,
			depth:  256,
		},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				wg := sync.WaitGroup{}
				noise := make([][][]float64, bm.width)

				for x := 0; x < len(noise); x++ {
					if noise[x] == nil {
						noise[x] = make([][]float64, bm.height)
					}

					for y := 0; y < len(noise[x]); y++ {
						if noise[x][y] == nil {
							noise[x][y] = make([]float64, bm.depth)
						}

						for z := 0; z < len(noise[x][y]); z++ {
							wg.Add(1)
							go func(x, y, z int, persistence, frequencyMultiplier float64, octaves int, noise *[][][]float64) {
								v, err := OctavePerlin(float64(x)/float64(len(*noise)), float64(y)/float64(len((*noise)[x])), float64(z)/float64(len((*noise)[x][y])), persistence, frequencyMultiplier, octaves)
								if err != nil {
									panic(err)
								}

								(*noise)[x][y][z] = v

								wg.Done()
							}(x, y, z, 1.25, 2.0, 6, &noise)
						}
					}
				}
				wg.Wait()
			}
		})
	}
}

func Test_resolveHashes(t *testing.T) {
	type args struct {
		xi int
		yi int
		zi int
	}
	tests := []struct {
		name    string
		args    args
		wantAaa int
		wantAba int
		wantAab int
		wantAbb int
		wantBaa int
		wantBba int
		wantBab int
		wantBbb int
	}{
		{
			name: "maxVals",
			args: args{
				xi: 0b11111111,
				yi: 0b11111111,
				zi: 0b11111111,
			},
			wantAaa: 215,
			wantAba: 103,
			wantAab: 61,
			wantAbb: 30,
			wantBaa: 20,
			wantBba: 140,
			wantBab: 125,
			wantBbb: 36,
		},
		{
			name: "minVals",
			args: args{
				xi: 0b0,
				yi: 0b0,
				zi: 0b0,
			},
			wantAaa: 36,
			wantAba: 108,
			wantAab: 103,
			wantAbb: 110,
			wantBaa: 86,
			wantBba: 128,
			wantBab: 164,
			wantBbb: 195,
		},
		{
			name: "randomVals",
			args: args{
				xi: 0b110100,
				yi: 0b1011101,
				zi: 0b101101,
			},
			wantAaa: 12,
			wantAba: 69,
			wantAab: 191,
			wantAbb: 142,
			wantBaa: 200,
			wantBba: 158,
			wantBab: 196,
			wantBbb: 231,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotAaa, gotAba, gotAab, gotAbb, gotBaa, gotBba, gotBab, gotBbb := resolveHashes(tt.args.xi, tt.args.yi, tt.args.zi)
			if gotAaa != tt.wantAaa {
				t.Errorf("resolveHashes() gotAaa = %v, want %v", gotAaa, tt.wantAaa)
			}
			if gotAba != tt.wantAba {
				t.Errorf("resolveHashes() gotAba = %v, want %v", gotAba, tt.wantAba)
			}
			if gotAab != tt.wantAab {
				t.Errorf("resolveHashes() gotAab = %v, want %v", gotAab, tt.wantAab)
			}
			if gotAbb != tt.wantAbb {
				t.Errorf("resolveHashes() gotAbb = %v, want %v", gotAbb, tt.wantAbb)
			}
			if gotBaa != tt.wantBaa {
				t.Errorf("resolveHashes() gotBaa = %v, want %v", gotBaa, tt.wantBaa)
			}
			if gotBba != tt.wantBba {
				t.Errorf("resolveHashes() gotBba = %v, want %v", gotBba, tt.wantBba)
			}
			if gotBab != tt.wantBab {
				t.Errorf("resolveHashes() gotBab = %v, want %v", gotBab, tt.wantBab)
			}
			if gotBbb != tt.wantBbb {
				t.Errorf("resolveHashes() gotBbb = %v, want %v", gotBbb, tt.wantBbb)
			}
		})
	}
}
