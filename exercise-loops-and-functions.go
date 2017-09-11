package main

import (
  "fmt"
  "math"
)

func Sqrt(x float64) float64 {
  res, res_tmp := 1.0, 1.1

  for delta := res - res_tmp; math.Abs(delta) > 0.001; delta = res - res_tmp {
    res_tmp = res
    res = res - ((res*res - x) / (2 * res))
  }

  return res
}

func main() {
  fmt.Println(Sqrt(2))
}
