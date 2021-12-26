# FizPop

A light-weight FP utils for Go

## Usage

All functions are group by the data type they are processing, for
example we can `slices.Map` a function to a slice

```go
	strLen := func (a1 string) int {
			return len(a1)
		}
	fmt.Println(
		Map(strLen)([]string{"cat", "catfish", "caterpillar"}),
	)
	// Output: [3 7 11]
```

Similarily, we can also `Map` the function to a `chan` with `chans.Map`

```go
	wordChan := make(chan string, 3)
	wordLenChan := Map(strLen)(wordChan)

	wordChan <- "cat"
	wordChan <- "catfish"
	wordChan <- "caterpillar"
	
	fmt.Println(
		<-wordLenChan,
		<-wordLenChan,
		<-wordLenChan,
	)
	// Output: 3 7 11
```
