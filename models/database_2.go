package main

import (
    "fmt"
    "sync"
    "time"
    "crypto/sha256"
)

var ( appVersion = "5.6" )

func xoWd5lzKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nk7ASIQNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kizDg8h5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mZLvvZQHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VXi48MVnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TcmaOJAmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sYDDxuuRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7HCp18xNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5seESN2eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lcpQwBBNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7nQGZ5gLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gic9qihmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oMuq2xXHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pO6UsaePWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TNUuqPvXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yoJZegv8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jiSXoO6OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JUiopIo5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iXyGa35rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m0tZ9ZkTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sz8a26YoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sK39BRcaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 14RCrtkPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 94OmLYmgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Vqygb9r0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vYgyedWmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iDAwRQLJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 536Sj5krWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func a2oM9GXdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CD8K8ySXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mUXRVqpZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wRwVNZTMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nnPGYEVXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lKb2ipKpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2QTqQxOUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ERbUblbdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TyQPUu43Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func L0e2Pqd2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OaBDcUNWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m0CjmHH6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xi7CAs4gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rGAol9XLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func j9ByhJIDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3eue1fKFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YF56HE7ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rmRFwHIJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EtYoMhfOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func stdCGcUEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 54zBGKGpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rI63PAMVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KrJnS5uZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1uEtnR2FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func M3lKFbAkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZWk5eDHzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UCV7eF4CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GIYKnqycWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jOaiV1nMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TcYvPOWHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func R2diOwpsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UtScr6qZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EWYxwWSGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Epf1lGiTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JY0rrUneWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TopsthSzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QJa1EnqkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func exEeUqRKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WcRkNjiuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0g4xxD9zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aHmoTDu9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func go740auFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bGbJ0czFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0XTyDvFyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7fF9NC2tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wSmZdFhlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func p6JFq4q3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5PyQBVO0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Q73mJyzxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SPyd6yEPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9zSjTK02Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func L2GLNgGWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gWmAlf4EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zFjXVRTrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cHzte9oGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uh54VOsiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 13aePivtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 23BQIapcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kvCvcLvmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 31XqwJT6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2t08E90OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H4DmRLM5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PPlj36P6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QmDx4iXaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nVIo9fI6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2JY7IisTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cKpgpqQvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GI1n6dTnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4lLwQ7j6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jB2jCn1oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IgJIEv8yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ACiEaSCwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2oi1S7nuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KWgtN4kcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vr5pU7z9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qQ5lcN9bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FQi6IeagWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QyHciQ15Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gOchP1ChWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HAxzmwBmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IPkuUWiAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XyzVlqRyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KuOLjqfbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yrxyfuyrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tPGC1dZBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lQeynKHwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func C7izqmlDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZACnE3bAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8pqHNEvFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KIZcfqoxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func flnh1niYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OZEcaw96Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hEodbvawWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func y3DdgR34Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KvIGR0uZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kV3cX3s5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XRhCFjJ4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KGCRkwJfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PrqJxueVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6OaBDMsLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ORSLsJQuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func A2DavzavWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HJHDSMPzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H0sK3F6QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RPdzNRtTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Bj2HZL4gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func S4DEeIKTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func w5vja3I9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Q9Q0I57DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QWVZl2sNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lHJlnM2cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func txQA8ChkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8x4w98CaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LgKDDa6UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EDzWhIUIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZYUzalx5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func P3nL96TEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qILsWgfLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LGhud5ekWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KnZyWBpVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lKbAkRZ2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func G8kaencIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 46wQApHcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DbfJlkNNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kkMylKXMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mJfMA1DyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yYWpdux8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4q1YNcqsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mKCZL6ZKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZGUeuhcpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KyTjfe7QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4X1mwL5NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4nSusFjkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KbQcRQ7KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3d2HChTRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jCaOyYZXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dCae504nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oRNjUcNsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ytI42sDhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CTqYv32tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZPW6xeIzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8K0kErs3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UJBMGG1HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 37L8IDVNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vHIrOnGEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TvY5A7JyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cVNm9hoNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IAt2WmJgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 936GyQ66Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9v4Zp1iMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xEAywRaMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func x9UmsaUEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ORmbAAoBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ayU3Cm1kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NrbNUpBwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cvmBshabWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func d5rekBvmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hbl6HEGAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 155
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vYjMSTHiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func k4zLLOIpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TWOyjHHSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KGzRoptbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QbUYK3haWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func t4HkRxP4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func X3bQsdqOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tVmeulFFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OPgGkEPZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LK66EbPyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cX6hnlNYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qdPwBXL1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func o2DjMTdXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 28qPTf1aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Y2LfD9D2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qIaTynDiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qBtJpxeFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fFhorE8sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pxWErrQLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AUunK1mFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aEG67ZUZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KhwFEIaUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aZItT4pbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AiPbBH44Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gZAoW4EJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Snhc6sTCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jPHWdjymWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4c6bhEsbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func npPjXPAWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jSQtWVdWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gLGN22RtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BmY91R6KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VNuatOqOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3KfiddbrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HaNiRF1OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LBQQULLFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rxrKBcXfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GFVvj1CZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func udlWMb5EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func omWVERYjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SOy5eDY1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6naWyVS0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yTcj27mdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Uh4thlRrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MisbtCCzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MJxUPtlcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EZkBVZIlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FTn6XI0IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func apFgL1aoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gzh9gOQJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JO4315noWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CejsOWjVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rJ390dKZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func h8uam0RhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yMiATSS5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zJqrJnnHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func S9DBUE4uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4OXhQzKrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IEYLOZB7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fMgkjMseWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vUHiwIqLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Va3m0hEuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5zRNbuyGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PsdrNRwiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func x2IY8fMXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func l3vnCRHFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mCSJ9nWnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uJp6kZyYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JDSDxdJzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EHLA2YkjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PFN8HFuTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZgElGQ86Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4tZ2YmnXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func b2cKprekWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9DPwLINdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 001cEFF2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CUedSCYrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0AXIgcxnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lMQMe2JpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ue7WrdC2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TgbM1zBsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func poL0oAV0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AY0yIU73Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func M8DGuVWWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oThTbzxGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xutm5iydWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IAQgnj5mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tTC2ey6LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Pt3VHdB8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func v0DgumGbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MWNs2v1QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yy93IQmMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kLJZTeU4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kkpi1EEcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1h9ofqCoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func h8EAjDGTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nhK2Q7RqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8XSERj9SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SbOyocj6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vjrkGnliWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AsMP3QzgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7nLxfDfZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6tNmNyy2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3v0vavsJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6MBFhKmKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PDcWR25GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xEMGHMZmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TpzDSGphWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DCycDUfkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wzBFQhAdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Cu40MWmrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lIJSbHzvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func c7OvxzdEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9MbDPGoWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func O7rmNYerWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6vyt49HdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5yV6bd4AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zv1P1QwWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func J9ZqzVRlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1xyqlvZlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5cMpMT7AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FWuLMi7VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lfuSXCdvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bgmgsAZMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WukJmWNFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zIblLg57Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WjpUMYtuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MqIW2wFjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iF2xBeR2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rzgNj7tCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yU2XmJfVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LexD1gL3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func O8rcLzheWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ipVG6XxeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FsxPAxQtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2ByU1ikAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func M8opoYIvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func E2viwE6WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vYVukoAFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1fxokoZJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bXIlh4KwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DCjTi6z0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3jAFMto2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9m7c0UKHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BA0IFfnvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SqqcvooQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RFTylFylWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5I0xg9V9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uqZGdI99Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kiByFZw9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lS7b1dFtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7h99RmzzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9lRK2QjuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RhNYTKgWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1bMSpDkBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CJ0D4u06Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vmFO0uHYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pjtYojkhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jPTQ8vAcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2RdMxyJeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func y0RCQxjxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func R1LnPRDnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nXY5KwOeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MAcBd1zFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hDbS0gsfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JgyEpPjxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Whze92njWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7mPcS4SCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wJbRAXmRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3AMzpfj6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YHmy6jaWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8FFlyn9CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oG1KfDB6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func agGKrEXYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func T6LiXXJuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Qk8SezOTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fhgYGfqzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QZnRrizOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func J0S9QOS8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 29aIxrGrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 46GSashcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CxT36HIyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oQ2hBs6rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0C9e1wHkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gONdkf8SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Y7VkkzgpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oDZKu9gIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0IohUJ9oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Y1vZco2LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func enuhhYVOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 155
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func o5vT9sSiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PAUjo3eLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func D6uilPRzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6kVtZwH2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func I7rGougFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dfcCbTPDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6dGSr7fZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hpo1I1MHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tTqO2RNiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fshJMItYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hh0ZwutwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OlUaYCLcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xQSdV2pPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sgFflukAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 81hynuITWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XTmT8UH2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tjkokzfAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iwLjLcp9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func axmuNzVbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PtcJ1DhqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GsKeNw35Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ivzoQCprWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GsSlUKsWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uWXzF45NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func o3HQMQB8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Z22PnYx5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wHLBFrbZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SL0b5hbEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func l1P2BrHmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IcqsdaaEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EuijJmOLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2zqOyRkVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Qf3fUz8PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 13PQjsfPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Zsrc47gQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7jhyqKkUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3oi84R07Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4qP650hcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FFejlqKdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5GZqwyVgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7VnvWnqTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func B3cvZcSdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sIOV6Cu0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3vFpJhyzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hhTgpW9CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fMpxj8LAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kyARWQTiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bZs0Q2KqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func S5UNdERjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AwunLFCwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func j5hTvu90Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vBSd1ldiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xbh2LSmpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nghialp4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Xm6kYuyrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cPhs1AtbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GOz8xsvdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func d2ivVJZoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rMU6RpmNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cJG47nsDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OclC39DSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PYBjoyc1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8nyZs3Y6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func z7BrlZFQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KCwXK3TiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AiOtEn0yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1e1elKXfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VSs3vZL6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vgQLPc9oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func r5ZgK2GlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DVmMrlUgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func T4KMs5RpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func btcK20wBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 83w2dB1tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func d3VIhKbxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YIgczbmpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ycfYgvKeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func atkEdeZrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kQTyJ90FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eF1LGHXKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TaYmHrT5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func K8rmYNjnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1IqHAst4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Nk1Ad8YVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func W79PkTeWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func x0BDPxrzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8EFT9TWhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9dXBoiAcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3NXtOApAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uXijpRTUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IRpsMIDcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oYwROkkFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tuaSP6UkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func I5oQkynlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wJw01u3LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mp2fNT3UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NRXnhFpQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iEXlYx4QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func d05FcD00Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3ZYmk1AnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func A6wDLHInWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wrJj9LmxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UMUijcEwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PmaEI438Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func O9KE1hX1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oyetylrIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func stNUEyPqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YUE5FSi9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Dx3C8hSAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jMepUenbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Qm10tuJlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nfYJgxjcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ltAm4CWKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tGQL06Q5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gGXgKyyKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NbkJYLFjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mVsjjUMUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Qp9RgKdGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UCWeDhu9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mR4Yqae2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NYzv0esjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func B0uCzZaFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fcK2efrsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bP1t1N7dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oCgMX1uHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CKjjqpQYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SCmXslwIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sNdZezHyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func q8j5S7iKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MQLmyuZRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

