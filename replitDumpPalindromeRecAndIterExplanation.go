// I do remember this; it was a big explanation I wrote out explaining recursive and iterative approaches for someone
// all that time and care, wasted lmao
// I don't know why I added fib (and the implementation is broken)
// I don't know why I added permute either, but I do think it's cool
package main

import "fmt"

// there are languages (like JS) with built in syntax that make this particular algo a liiiittle prettier
// but go is the most readable language I know

func main() {
	// fmt.Println("rec, true")
	// fmt.Println(isPalindromeRec(""))
	// fmt.Println(isPalindromeRec("a"))
	// fmt.Println(isPalindromeRec("aa"))
	// fmt.Println(isPalindromeRec("abba"))
	// fmt.Println(isPalindromeRec("abcba"))

	// fmt.Println("recOpt, true")
	// fmt.Println(isPalindromeRecOpt("", 0))
	// fmt.Println(isPalindromeRecOpt("a", 0))
	// fmt.Println(isPalindromeRecOpt("aa", 0))
	// fmt.Println(isPalindromeRecOpt("abba", 0))
	// fmt.Println(isPalindromeRecOpt("abcba", 0))

	// fmt.Println("iter, true")
	// fmt.Println(isPalindromeIter(""))
	// fmt.Println(isPalindromeIter("a"))
	// fmt.Println(isPalindromeIter("aa"))
	// fmt.Println(isPalindromeIter("abba"))
	// fmt.Println(isPalindromeIter("abcba"))

	// // false
	// fmt.Println("rec, false")
	// fmt.Println(isPalindromeRec("ab"))
	// fmt.Println(isPalindromeRec("abca"))
	// fmt.Println(isPalindromeRec("abcda"))

	// fmt.Println("recOpt, false")
	// fmt.Println(isPalindromeRecOpt("ab", 0))
	// fmt.Println(isPalindromeRecOpt("abca", 0))
	// fmt.Println(isPalindromeRecOpt("abcda", 0))

	// fmt.Println("iter, false")
	// fmt.Println(isPalindromeIter("ab"))
	// fmt.Println(isPalindromeIter("abca"))
	// fmt.Println(isPalindromeIter("abcda"))
	permute([]int{0, 1}, 1)
	permute([]int{0, 1}, 2)
	permute([]int{0, 1}, 3)
	permute([]int{0, 1}, 4)
}

// good; so neat! so simple!
// bad; extra memory for each new string allocated ("abba", "bb", and "") and in languages without tail call optimization https://stackoverflow.com/questions/310974/what-is-tail-call-optimization , you use extra memory for the callstack
func isPalindromeRec(v string) bool {
	if len(v) == 0 || len(v) == 1 {
		return true
	}

	if v[0] != v[len(v)-1] {
		return false
	}

	return isPalindromeRec(v[1 : len(v)-1])
}

// good; don't have to allocate memory new strings anymore
// bad; wow that's ugly and confusing. Weird call signature. Bookkeeping is even worse than in the iterative version
func isPalindromeRecOpt(v string, i int) bool {
	remainingLen := len(v) - (i * 2)
	if remainingLen == 0 || remainingLen == 1 {
		return true
	}

	if v[i] != v[len(v)-1-i] {
		return false
	}

	return isPalindromeRecOpt(v, i+1)
}

// good; no extra memory from callstack, no extra memory for allocations
// bad; weird ugly bookkeeping with the index
func isPalindromeIter(v string) bool {
	i := 0
	for i < len(v)/2 {
		if v[i] != v[len(v)-1-i] {
			return false
		}
		i++
	}

	return true
}

func fib(n int) {
	if n <= 0 {
		return
	}

	fib(n - 1)
	fib(n - 2)
}

func permute(alphabet []int, n int) [][]int {
	fmt.Println(">>>", alphabet, n)
	res := make([][]int, 0)
	temp := make([]int, n)

	var recurse func(stepsDeep int)
	recurse = func(stepsDeep int) {
		if stepsDeep == n {
			fmt.Println(len(res), temp)
			perm := make([]int, n)
			copy(perm, temp)
			res = append(res, perm)
			return
		}

		for _, char := range alphabet {
			temp[stepsDeep] = char
			recurse(stepsDeep + 1)
		}
	}

	recurse(0)

	return res
}
