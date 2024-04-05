package main

import (
	"bytes"
	"context"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"sync"
	"time"

	"github.com/google/uuid"
	"golang.org/x/sync/errgroup"
)

func Main_chunk() {

	fmt.Println("------ testing---------")

	path := "/Users/nisingh/Downloads/test40GB"
	// path := "/Users/nisingh/Downloads/03.mpg"
	// createFileSize(path, file_Size)

	// Create a reader for the file to upload
	fileReader, err := getFileReader(path)
	if err != nil {
		fmt.Println("Error getting file reader:", err)
		return
	}

	// Get the total file size
	fileSize := getFileSize(path)

	totalchunks := fileSize / DefaultChunkUploadSize
	remaining := fileSize % DefaultChunkUploadSize

	fmt.Println("Total chunks----> ", totalchunks, "remaining:", remaining)

	err = chunkedUpload(context.Background(), path, fileReader, fileSize)
	if err != nil {
		fmt.Println("Error getting file size:", err)
		return
	}

	fmt.Println("Total chunks----> ", totalchunks, "remaining:", remaining)
	fmt.Println("------ testing done---------")
}

func createFileSize(path_name string, fileSize int64) {

	// Create the file
	file, err := os.Create(path_name)
	if err != nil {
		panic(err)
	}
	// Set the file size
	if err := file.Truncate(fileSize); err != nil {
		panic(err)
	}
	// Close the file
	if err := file.Close(); err != nil {
		panic(err)
	}
}

/// getFileReader returns a reader for the specified file path
func getFileReader(filePath string) (io.ReadSeeker, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	return file, nil
}

// getFileSize returns the size of the file represented by the reader
func getFileSize(path string) int64 {
	fInfo, err := os.Stat(path)
	if err != nil {
		log.Fatal(err)
	}

	return fInfo.Size()

}

// package main

// import (
// 	"bytes"
// 	"context"
// 	"fmt"
// 	"io"
// 	"time"

// 	"github.com/google/uuid"
// 	"golang.org/x/sync/errgroup"
// )

const (
	DefaultChunkUploadSize         = 104857600 //100MB size of chunk to upload
	chunkRetry                     = 10        // number of retries with pacer
	defaultChunkUploadThreadsCount = 4         // default threads for upload chunk
	minSleep                       = 2 * time.Second
)

// upload first chunk of large files
func uploadFirstChunk(ctx context.Context, path string, chunkInfo *ChunkUploadInfo) (UploadInfo, error) {

	chunkData, rem, chunkNum, err := chunkInfo.GetChunk()
	if err != nil {
		return UploadInfo{}, err
	}
	uploadInfo := UploadInfo{
		Path:      path,
		Data:      bytes.NewReader(chunkData),
		Csum:      SHA512Digest(chunkData),
		ChunkNum:  chunkNum,
		ChunkSize: DefaultChunkUploadSize,
		UploadID:  "",
	}

	uploadInfo.UploadID = uuid.New().String()

	fmt.Printf("\n uploadFirstChunk :: chunkNum-->: %d remaining bytes: %d, \n Upload info---->: %+v, \n\n", chunkNum, rem, uploadInfo)

	return uploadInfo, err
}

func uploadMiddleChunk(ctx context.Context, path string, chunkInfo *ChunkUploadInfo, sema chan struct{}, UploadID string) error {

	defer func() {
		<-sema
	}()
	chunkData, rem, chunkNum, err := chunkInfo.GetChunk()
	if err != nil {
		return err
	}

	uploadInfo := UploadInfo{
		Path:      path,
		Data:      bytes.NewReader(chunkData),
		Csum:      SHA512Digest(chunkData),
		ChunkNum:  chunkNum,
		ChunkSize: DefaultChunkUploadSize,
		UploadID:  UploadID,
	}

	fmt.Printf("\n uploadMiddleChunk :: chunkNum-->: %d, remaining bytes: %d, \n Upload info---->: %+v, \n\n", chunkNum, rem, uploadInfo)

	time.Sleep(time.Millisecond * 100)

	chunkInfo.SetChunkCheckSum(uploadInfo.ChunkNum, uploadInfo.Csum)
	return err
}

// uploadLastChunk of large files
func uploadLastChunk(ctx context.Context, path string, chunkInfo *ChunkUploadInfo, UploadID string) error {
	chunkData, chunkNum := chunkInfo.GetLastChunk()
	csum := SHA512Digest(chunkData)
	uploadInfo := UploadInfo{
		Path:      path,
		Data:      bytes.NewReader(chunkData),
		Csum:      SHA512Digest(chunkData),
		ChunkNum:  chunkNum,
		ChunkSize: DefaultChunkUploadSize,
		UploadID:  UploadID,
	}

	chunkInfo.SetChunkCheckSum(chunkNum, csum)
	finalCsum := chunkInfo.GetResultCsum()
	uploadInfo.ChunkNum = chunkNum
	uploadInfo.Data = bytes.NewReader(chunkData)
	uploadInfo.Csum = SHA512Digest([]byte(finalCsum))

	fmt.Printf("\n uploadLastChunk :: chunkNum-->: %d, remaining bytes: %d, \n Upload info---->: %+v,  \n\n", chunkNum, chunkInfo.remainingBytes, uploadInfo)

	return nil
}

// ChunkedUpload: Upload large files in chunks
func chunkedUpload(ctx context.Context, path string, in io.Reader, size int64) error {
	if size <= DefaultChunkUploadSize {
		return fmt.Errorf("chunk upload file size should be greater than %d B", DefaultChunkUploadSize)
	}
	chunkInfo := ChunkUploadInfo{}

	chunkInfo.Init(in, size, DefaultChunkUploadSize)
	var err error
	uploadInfo, err := uploadFirstChunk(ctx, path, &chunkInfo)
	if err != nil {
		return err
	}
	chunkInfo.SetChunkCheckSum(uploadInfo.ChunkNum, uploadInfo.Csum)

	fmt.Printf("\n starting :: Upload info : %+v, \n ChunkInfo: %+v", uploadInfo, chunkInfo)

	var parallelUploadThread int = defaultChunkUploadThreadsCount
	// checking no of chunks if less then defaultChunkUploadThreadsCount so we can resuce threads for chunkupload
	totalChunks := int(size / DefaultChunkUploadSize)
	if totalChunks < 3 {
		parallelUploadThread = 1
	} else if totalChunks <= parallelUploadThread {
		parallelUploadThread = totalChunks - 1
	}
	limiter := make(chan struct{}, parallelUploadThread)
	g, gCtx := errgroup.WithContext(ctx)
	for {

		switch {
		case gCtx.Err() == context.Canceled:
			return gCtx.Err()
		default:
			limiter <- struct{}{}
		}
		if chunkInfo.GetRemainingBytes() == 0 {
			break
		}

		g.Go(func() (err error) {
			return uploadMiddleChunk(gCtx, path, &chunkInfo, limiter, uploadInfo.UploadID)
		})

	}
	err = g.Wait()
	if err != nil {
		return err
	}

	if err = uploadLastChunk(ctx, path, &chunkInfo, uploadInfo.UploadID); err != nil {
		return err
	}
	fmt.Println(nil, "Chunk upload done for path: %v", path)

	return nil
}

type UploadInfo struct {
	Path      string
	Data      io.Reader
	Csum      string
	ChunkNum  int
	ChunkSize int64
	UploadID  string
}

// String prints the uplaod info params value
func (u *UploadInfo) String() string {
	return fmt.Sprintf("UploadInfo:: chunkNum: %d, ChunkSize: %d, UploadId: %s", u.ChunkNum, u.ChunkSize, u.UploadID)
}

// ChunkUploadInfo is for infromation of chunk needs to upload
type ChunkUploadInfo struct {
	dataMutex      *sync.Mutex
	resultMutex    *sync.Mutex
	filesize       int64
	chunkNum       int
	data           io.Reader
	remainingBytes int64
	checkSumMap    map[int]string
	chunkSize      int64
	lastChunk      []byte
}

// Init is to initialize the ChunkUploadInfo struct with passed values
func (c *ChunkUploadInfo) Init(data io.Reader, size int64, chunkSize int64) {
	c.checkSumMap = make(map[int]string)
	c.data = data
	c.remainingBytes = size
	c.filesize = size
	c.chunkSize = chunkSize
	c.chunkNum = 0
	c.lastChunk = nil
	c.dataMutex = &sync.Mutex{}
	c.resultMutex = &sync.Mutex{}

}

// GetRemainingBytes returns the remaining bytes in chunk upload process
func (c *ChunkUploadInfo) GetRemainingBytes() int64 {
	c.dataMutex.Lock()
	defer c.dataMutex.Unlock()
	return c.remainingBytes
}

// GetRemainingBytes returns the remaining bytes in chunk upload process
func (c *ChunkUploadInfo) GetChunkNumber() int {
	c.dataMutex.Lock()
	defer c.dataMutex.Unlock()
	return c.chunkNum
}

// GetChunk is used to get the chunk data
// it will check for remaining bytes
// create chunk buffer data based on size, if any error it will return error otherwise update the remaining bytes.
// return chunk data to upload, remainingbytes, chunk number, and error if any
func (c *ChunkUploadInfo) GetChunk() ([]byte, int64, int, error) {
	c.dataMutex.Lock()
	defer c.dataMutex.Unlock()
	if c.remainingBytes == 0 {
		return nil, 0, 0, nil
	}
	c.chunkNum += 1
	buf := make([]byte, c.chunkSize)
	n, err := io.ReadFull(c.data, buf)
	switch err {
	case nil:
		break
	case io.ErrUnexpectedEOF:
		break
	default:
		return nil, 0, 0, err
	}
	c.remainingBytes -= int64(n)
	buf = buf[:n]
	if c.remainingBytes == 0 {
		c.lastChunk = buf
	}
	return buf, c.remainingBytes, c.chunkNum, nil
}

// GetLastChunk returns the last chunk and the chunk number
func (c *ChunkUploadInfo) GetLastChunk() ([]byte, int) {
	c.dataMutex.Lock()
	defer c.dataMutex.Unlock()
	return c.lastChunk, c.chunkNum
}

// SetChunkCheckSum sets the chunk checksum into the ChunkUploadInfo object checkSumMap
func (c *ChunkUploadInfo) SetChunkCheckSum(chunkNum int, csum string) {
	c.resultMutex.Lock()
	defer c.resultMutex.Unlock()
	c.checkSumMap[chunkNum] = csum
}

// GetResultCsum return final check sum of all chunks
func (c *ChunkUploadInfo) GetResultCsum() string {
	c.resultMutex.Lock()
	defer c.resultMutex.Unlock()
	res := ""
	chunks := make([]int, len(c.checkSumMap))
	i := 0
	for k := range c.checkSumMap {
		chunks[i] = k
		i++
	}
	sort.Ints(chunks)
	for _, chunk := range chunks {
		res += c.checkSumMap[chunk]
	}
	return res
}

// SHA512Digest generates the checksume with SUM512 and hex encoded.
func SHA512Digest(buf []byte) string {
	csum := sha512.Sum512(buf)
	return hex.EncodeToString(csum[:])
}

// String prints the Chunk Upload Info params.
func (c *ChunkUploadInfo) String() string {
	return fmt.Sprintf("ChunkUploadInfo:: chunkNum: %d, remainingbytes: %d, chunkSize: %d, lastChunk: %d", c.chunkNum, c.remainingBytes, c.chunkSize, len(c.lastChunk))
}
