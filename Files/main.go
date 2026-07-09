package main

// const (
// 	deafultBufSize = 4096
// )

// var (
// 	ErrInvalidUnreadByte = errors.New("bufio: invalid use of UnreadByte")
// 	ErrInvalidUnreadRune = errors.New("bufio: invalid use of UnreadRune")
// 	ErrBufferFull        = errors.New("bufio: buffer full")
// 	ErrNegativeCount     = errors.New(" negative count ")
// )

// Buffered input

// Reader implements buffering for an io.Reader object.

func main() {
	// f, err := os.Open("example.txt")
	// if err != nil {
	// 	// log the error
	// 	panic(err)
	// }

	// fileInfo, err := f.Stat()
	// if err != nil {
	// 	// log the error
	// 	panic(err)
	// }

	// fmt.Println("File name:", fileInfo.Name())
	// fmt.Println("File or Folder:", fileInfo.IsDir())
	// fmt.Println("File Size:", fileInfo.Size())
	// fmt.Println("File Permission:", fileInfo.Mode())
	// fmt.Println("File Modified At:", fileInfo.ModTime())

	// For Reading the file

	// f, err := os.Open("example.txt")
	// if err != nil {
	// 	panic(err)
	// }

	// defer f.Close()

	// buf := make([]byte, 10)
	// d, err := f.Read(buf)
	// if err != nil {
	// 	panic(err)
	// }

	// for i := 0; i < len(buf); i++ {
	// 	println("data", d, string(buf[i]))
	// }

	// Another technique for reading file(loads whole file at once not suitable for large files)

	// data, err := os.ReadFile("example.txt")
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(string(data))

	//  Read Folders

	// dir, err := os.Open("../")
	// if err != nil {
	// 	panic(err)
	// }

	// defer dir.Close()

	// fileInfo, err := dir.ReadDir(1)

	// for _, fi := range fileInfo {
	// 	fmt.Println(fi.Name(), fi.IsDir())
	// }

	// Creating the file

	// f, err := os.Create("example2.txt")
	// if err != nil {
	// 	panic(err)
	// }

	// defer f.Close()

	// f.WriteString("Hello from Team Tanjavur")
	// f.WriteString("South India")

	// bytes := []byte("Hello JayNagar People")

	// f.Write(bytes)

	// read and write to another file (streaming fashion)

	// sourceFile, err := os.Open("example.txt")
	// if err != nil {
	// 	panic(err)
	// }

	// defer sourceFile.Close()

	// destFile, err := os.Create("example2.txt")
	// if err != nil {
	// 	panic(err)
	// }

	// defer destFile.Close()

	// reader := bufio.NewReader(sourceFile)
	// writer := bufio.NewWriter(destFile)

	// for {
	// 	b, err := reader.ReadByte()
	// 	if err != nil {
	// 		if err.Error() != "EOF" {
	// 			panic(err)
	// 		}
	// 		break
	// 	}

	// 	e := writer.WriteByte(b)
	// 	if err != nil {
	// 		panic(e)
	// 	}
	// }
	// writer.Flush()

	// fmt.Println("Written to new file successfully")

	// Delete a file

	// sourceFile, err := os.Open("example2.txt")
	// if err != nil {
	// 	 panic(err)
	// }

	// defer sourceFile.Close()

	// err := os.Remove("example2.txt")
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println("File Deleted successfully")

}
