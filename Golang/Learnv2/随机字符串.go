package main

// var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// func main() {
// 	logrus.Info(randStr(10))
// }
//
// func randStr(n int) string {
// 	b := make([]rune, n)
// 	for i := range b {
// 		b[i] = letters[grand.Intn(len(letters))]
// 	}
// 	return string(b)
// }
//
// func TestApproach1(t *testing.T) {
// 	rand.Seed(time.Now().UnixNano())
// 	fmt.Println(randStr(10))
// }
//
// func BenchmarkApproach1(b *testing.B) {
// 	rand.Seed(time.Now().UnixNano())
// 	for i := 0; i < b.N; i++ {
// 		_ = randStr(10)
// 	}
// }
