# imageio
Input / output for image files in Golang

## RGB Reader

Read a RGB file from a io.reader.

RGB files are 24 bits true color (8 bits per channels).

```go
img, _ := rgb24.Decode(f, &rgb24.Options{Width: width, Height: height})
```


## RGB Writer

Write a RGB file from a io.writer and go image.

The writer may lose some informations due to the conversion between uint32 into byte.

```go
img,_ := jpeg.Decode(infile)
rgb24.Encode(outfile, img)
```