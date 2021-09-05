package ncm

import (
	"fmt"
	"testing"
)

func Test_randomMusics(t *testing.T) {
	musics := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	num := 5
	randomMusics(musics, num)
	fmt.Println(musics[:num])
}
