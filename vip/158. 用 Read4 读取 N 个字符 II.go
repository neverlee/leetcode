// 158. 用 Read4 读取 N 个字符 II

package main

import "fmt"

/**
 * The read4 API is already defined for you.
 *
 *     read4 := func(buf4 []byte) int
 *
 * // Below is an example of how the read4 API can be called.
 * file := File("abcdefghijk") // File is "abcdefghijk", initially file pointer (fp) points to 'a'
 * buf4 := make([]byte, 4) // Create buffer with enough space to store characters
 * read4(buf4) // read4 returns 4. Now buf = ['a','b','c','d'], fp points to 'e'
 * read4(buf4) // read4 returns 4. Now buf = ['e','f','g','h'], fp points to 'i'
 * read4(buf4) // read4 returns 3. Now buf = ['i','j','k',...], fp points to end of file
 */

var src = []byte("abc")

func read4(buf4 []byte) int {
	n := 0
	for ; n < 4 && n < len(src); n++ {
		buf4[n] = src[n]
	}
	src = src[n:]
	return n
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

var solution = func(read4 func([]byte) int) func([]byte, int) int {
	// implement read below.
	var remain []byte
	return func(buf []byte, n int) int {
		cur := copy(buf[:min(len(buf), n)], remain)
		remain = remain[cur:]
		if cur >= n {
			return cur
		}

		buf4 := make([]byte, 4)
		rn := 0
		i := 0
		modi := 0
		for ; i < n-cur; i++ {
			modi = i % 4
			if modi == 0 {
				rn = read4(buf4)
			}
			if modi >= rn {
				break
			}
			buf[cur+i] = buf4[modi]
		}

		if modi+1 < rn {
			remain = buf4[modi+1 : rn]
		} else {
			remain = nil
		}
		return cur + i
	}
}

func main() {
	read := solution(read4)
	// query := []int{4, 1}
	query := []int{1, 1, 1, 1}
	for i, n := range query {
		buf := make([]byte, 1000)
		rn := read(buf, n)
		fmt.Println(">>>>", i, n, rn, buf[:n])
	}
}
