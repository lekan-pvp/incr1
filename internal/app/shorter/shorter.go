package shorter

import "fmt"

//const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
//
//func Shorting() string {
//	b := make([]byte, 5)
//	for i := range b {
//		b[i] = letters[rand.Intn(len(letters))]
//	}
//	return string(b)
//}

const endpoint = "http://localhost:8080/"

func Shorting(id int) string {
	return fmt.Sprintf("%s%d", endpoint, id)
}
