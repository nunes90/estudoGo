// Ex 1.16 -constants
/*
Go's `map` collection type, which will act as the cache. There is a global limit on the number of items that can be in the cache. We'll use one `map` to help keep track of the number of items in the cache. We have two types of data we need to cache: books and CDs.

Both use the ID, so we need a way to separate the two types of items in the shared cache. We need a way to set and get items from the cache.
*/
package main

import "fmt"

// global limit size
const GlobalLimit = 100

const MaxCacheSize int = 10 * GlobalLimit

// Cache prefix
const (
	CacheKeyBook = "book_"
	CacheKeyCD   = "cd_"
)

// Cache
var cache map[string]string

// get items
func cacheGet(key string) string {
	return cache[key]
}

// set items
func cacheSet(key string, val string) {
	if len(cache)+1 >= MaxCacheSize {
		return
	}
	cache[key] = val
}

func GetBook(isbn string) string {
	return cacheGet(CacheKeyBook + isbn)
}

func SetBook(isbn string, name string) {
	cacheSet(CacheKeyBook+isbn, name)
}

// CD data from cache
func GetCD(sku string) string {
	return cacheGet(CacheKeyCD + sku)
}

// add CD data to cache
func SetCD(sku string, title string) {
	cacheSet(CacheKeyCD+sku, title)
}

func main() {
	cache = make(map[string]string)
	SetBook("1234-5678", "Get Ready To Go")
	SetCD("1234-5678", "Get Ready To Go Audio Book")
	fmt.Println("Book :", GetBook("1234-5678"))
	fmt.Println("CD   :", GetCD("1234-5678"))
}
