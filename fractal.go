// I copied this from the Julia set page on rosetta code and added stuff

package main	

import (
	"image"
	"image/color"
	"image/png"
//	"math/rand"
	"math"
        "log"
	"os"
	"sync"
)
 
func main() {
	const ( //width, height = 1920, 1080
                width, height = 15360, 8640
		max_iter      = 255		
//              c = cx + cy. c is a complex number!
//                cx, cy = 0.96, 0.04
//                cx, cy = -0.123, 0.745
//    		  cx, cy = 0, 0
                cx, cy = (math.Phi - 2), (math.Phi - 1)
//                cx, cy = -0.755, 0.06 
//                cx, cy = -0.835, -0.2321
//                cx, cy = -0.7, 0.27015
//                cx, cy = -0.755, 0.01
		file          = "julia.png"
                stretch       = -1.65 )
		// make it fit on the entire screen
		// 1.5 is too big and looks ugly
		//
		// if the value is negative, the fractal is flipped
		
	img := image.NewNRGBA(image.Rect(0, 0, width, height))
 
	var wg sync.WaitGroup
	wg.Add(width)
	for x := 0; x < width; x++ {
		thisx := float64(x)
		
		go func() {
			var tmp, zx, zy float64
			var i uint8
			for y := 0; y < height; y++ {
                                
                                // how stretched the             
                                // fractal will be       division by 2^1
				zx = stretch * (thisx - float64(width >> 1)) / (0.5 * width)
				zy = (float64(y)  - float64(height >> 1)) / (0.5 * height)
				i = max_iter
				           				
				for (zx*zx + zy*zy) - float64(i) < 4.0 && i > 0 {
        				// for z^2 + tan c
					tmp = zx*zx - zy*zy + math.Tan(cx)
					zy  = 2*zx * zy + math.Tan(cy)
					zx  = tmp
					i--
					
        				// for z^2 + c
        				/*
					tmp = zx*zx - zy*zy + cx	
					zy  = 2*zx * zy + cy
					zx  = tmp
					i--
					*/
					
                			// for z^n + c
                			// slow
                                /*
                                        var n int = 2
        			        tmp = math.Pow((zx*zx + zy*zy), float64(n>>1)) * math.Cos(float64(n) * math.Atan2(zy, zx)) + cx
        			        zy  = math.Pow((zx*zx + zy*zy), float64(n>>1)) * math.Sin(float64(n) * math.Atan2(zy, zx)) + cy
					zx  = tmp
					i--
                                */
				}
				
//				var alpha uint8 = uint8(rand.Uint32())
				
				img.Set(int(thisx), y, color.RGBA{ i * uint8(int(thisx) & (2*y)), i*(i-uint8(x)), i*i, 255 })
			}
			wg.Done()
		}()
	}
	wg.Wait()
	img_file, err := os.Create(file)
	if err != nil {
		log.Fatal(err)
	}
	defer img_file.Close()
	if err := png.Encode(img_file, img); err != nil {
		img_file.Close()
		log.Fatal(err)
	}
}