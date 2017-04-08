package main

import(
    "crypto/rand"
    "crypto/md5"
    "crypto/sha256"
    "runtime"
	"time"
	"fmt"
)
import sha256simd "github.com/minio/sha256-simd"

const TestBufs int = 32
const L3CacheLen int = 32*1024*1024
const TestLen int = L3CacheLen / TestBufs
const TestLoops int = 100
const TestMem int = L3CacheLen * TestLoops

func ReportSpeed(title string, memsize int, d time.Duration) {
    sp := float64(memsize) / 1024 / 1024 / d.Seconds()
    sz := memsize / 1024 / 1024
    fmt.Printf("%s : size = %d MB, runtime = %f, speed = %f MB/s \n", title, sz, d.Seconds(), sp)
}

func NoneTesting() {
    time.Sleep(6 * time.Second)
} 

func BuildInMd5Testing(a [][]byte) {
    for i := 0; i < TestLoops; i++ {
        for j := range a {
	        md5.Sum(a[j])		
        }
    }
}

func BuildInSha256Testing(a [][]byte) {
    for i := 0; i < TestLoops; i++ {
        for j := range a {
	        sha256.Sum256(a[j])		
        }
    }
}

//https://github.com/minio/sha256-simd
func SimdSha256Testing(a [][]byte) {
    for i := 0; i < TestLoops; i++ {
        for j := range a {
	        sha256simd.Sum256(a[j])		
        }
    }
}

//https://blog.golang.org/pipelines
func PipelineSha256Testing(a [][]byte) {
}

func main() {
    runtime.GOMAXPROCS(4)
    a := make([][]byte, TestBufs)
    for i := range a {
        a[i] = make([]byte, TestLen)
		_, err := rand.Read(a[i])
		if err != nil {
			fmt.Println("error:", err)
			return
		}
    }

    t0 := time.Now()
    NoneTesting()
    t1 := time.Now()
    ReportSpeed("NoneTesting", TestMem, t1.Sub(t0))

    t0 = time.Now()
    BuildInMd5Testing(a)
    t1 = time.Now()
    ReportSpeed("BuildInMd5Testing", TestMem, t1.Sub(t0))

    t0 = time.Now()
    BuildInSha256Testing(a)
    t1 = time.Now()
    ReportSpeed("BuildInSha256Testing", TestMem, t1.Sub(t0))

    t0 = time.Now()
    SimdSha256Testing(a)
    t1 = time.Now()
    ReportSpeed("SimdSha256Testing", TestMem, t1.Sub(t0))
}
