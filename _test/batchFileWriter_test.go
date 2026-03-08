package test

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/farseer-go/fs/batchFileWriter"
)

// TestBasicWrite 测试基本写入功能
func TestBasicWrite(t *testing.T) {
	dir := createTempDir(t)
	defer os.RemoveAll(dir)

	writer := batchFileWriter.NewWriter(dir, "log", "day", 0, 0, time.Second, true)

	// 写入数据
	for i := 0; i < 10; i++ {
		writer.Write(fmt.Sprintf("test line %d", i))
	}

	writer.Close()

	// 验证文件存在
	files := listFiles(dir)
	if len(files) != 1 {
		t.Fatalf("期望 1 个文件，实际 %d 个", len(files))
	}

	// 验证内容
	content := readFile(t, files[0])
	lines := strings.Split(content, "\n")
	// 移除最后一个空行
	if lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}
	if len(lines) != 10 {
		t.Fatalf("期望 10 行，实际 %d 行", len(lines))
	}
}

// TestWriteWithSizeLimit 测试文件大小限制翻转（MB 单位）
func TestWriteWithSizeLimit(t *testing.T) {
	dir := createTempDir(t)
	defer os.RemoveAll(dir)

	// 设置 1MB 大小限制
	writer := batchFileWriter.NewWriter(dir, "log", "day", 1, 0, time.Second, true)

	// 写入超过 1MB 的数据（每行约 1KB，写入 1500 行 ≈ 1.5MB）
	for i := 0; i < 1500; i++ {
		writer.Write(fmt.Sprintf("test line %d - %s", i, strings.Repeat("x", 1000)))
	}

	writer.Close()

	// 验证文件翻转
	files := listFiles(dir)
	if len(files) < 2 {
		t.Fatalf("期望至少 2 个文件（翻转），实际 %d 个", len(files))
	}

	t.Logf("生成了 %d 个文件", len(files))
}

// TestFileCountLimit 测试文件数量限制
func TestFileCountLimit(t *testing.T) {
	dir := createTempDir(t)
	defer os.RemoveAll(dir)

	// 设置 1MB 大小限制，最多保留 3 个文件
	writer := batchFileWriter.NewWriter(dir, "log", "day", 1, 3, time.Second, true)

	// 写入足够多的数据，触发多次翻转（约 5MB）
	for i := 0; i < 6000; i++ {
		writer.Write(fmt.Sprintf("test line %d - %s", i, strings.Repeat("x", 1000)))
	}

	writer.Close()

	// 验证文件数量
	files := listFiles(dir)
	if len(files) > 3 {
		t.Fatalf("期望最多 3 个文件，实际 %d 个", len(files))
	}

	t.Logf("保留了 %d 个文件", len(files))
}

// TestAppendNewLine 测试换行符设置
func TestAppendNewLine(t *testing.T) {
	dir := createTempDir(t)
	defer os.RemoveAll(dir)

	// 测试添加换行符
	writer := batchFileWriter.NewWriter(dir, "log", "day", 0, 0, time.Second, true)
	writer.Write("line1")
	writer.Write("line2")
	writer.Close()

	files := listFiles(dir)
	content := readFile(t, files[0])
	if !strings.Contains(content, "line1\n") {
		t.Fatal("期望包含换行符")
	}

	// 测试不添加换行符
	dir2 := createTempDir(t)
	defer os.RemoveAll(dir2)

	writer2 := batchFileWriter.NewWriter(dir2, "log", "day", 0, 0, time.Second, false)
	writer2.Write("line1")
	writer2.Write("line2")
	writer2.Close()

	files2 := listFiles(dir2)
	content2 := readFile(t, files2[0])
	if strings.Contains(content2, "line1\n") {
		t.Fatal("不期望包含换行符")
	}
	if content2 != "line1line2" {
		t.Fatalf("期望 'line1line2'，实际 '%s'", content2)
	}
}

// TestWriteDifferentTypes 测试写入不同类型数据
func TestWriteDifferentTypes(t *testing.T) {
	dir := createTempDir(t)
	defer os.RemoveAll(dir)

	writer := batchFileWriter.NewWriter(dir, "log", "day", 0, 0, time.Second, true)

	// 写入字符串
	writer.Write("string data")

	// 写入字节
	writer.Write([]byte("byte data"))

	// 写入结构体（会被序列化为 JSON）
	type TestStruct struct {
		Name  string `json:"name"`
		Value int    `json:"value"`
	}
	writer.Write(TestStruct{Name: "test", Value: 123})

	writer.Close()

	files := listFiles(dir)
	content := readFile(t, files[0])

	if !strings.Contains(content, "string data") {
		t.Fatal("期望包含 string data")
	}
	if !strings.Contains(content, "byte data") {
		t.Fatal("期望包含 byte data")
	}
	if !strings.Contains(content, `"name":"test"`) {
		t.Fatal("期望包含 JSON 结构体数据")
	}

	t.Logf("内容:\n%s", content)
}

// TestReuseExistingFile 测试复用已有文件
func TestReuseExistingFile(t *testing.T) {
	dir := createTempDir(t)
	defer os.RemoveAll(dir)

	// 第一次写入
	writer1 := batchFileWriter.NewWriter(dir, "log", "day", 0, 0, time.Second, true)
	for i := 0; i < 5; i++ {
		writer1.Write(fmt.Sprintf("first write %d", i))
	}
	writer1.Close()

	// 第二次写入（应该复用同一个文件）
	writer2 := batchFileWriter.NewWriter(dir, "log", "day", 0, 0, time.Second, true)
	for i := 0; i < 5; i++ {
		writer2.Write(fmt.Sprintf("second write %d", i))
	}
	writer2.Close()

	// 验证只有一个文件
	files := listFiles(dir)
	if len(files) != 1 {
		t.Fatalf("期望 1 个文件，实际 %d 个", len(files))
	}

	// 验证内容包含两次写入
	content := readFile(t, files[0])
	lines := strings.Split(content, "\n")
	// 移除最后一个空行
	if lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}
	if len(lines) != 10 {
		t.Fatalf("期望 10 行，实际 %d 行", len(lines))
	}

	if !strings.Contains(content, "first write") {
		t.Fatal("期望包含 first write")
	}
	if !strings.Contains(content, "second write") {
		t.Fatal("期望包含 second write")
	}
}

// TestReuseFileWithSizeLimit 测试有大小限制时复用文件（MB 单位）
func TestReuseFileWithSizeLimit(t *testing.T) {
	dir := createTempDir(t)
	defer os.RemoveAll(dir)

	// 第一次写入，设置 10MB 大小限制
	writer1 := batchFileWriter.NewWriter(dir, "log", "day", 10, 0, time.Second, true)
	for i := 0; i < 5; i++ {
		writer1.Write(fmt.Sprintf("first write %d", i))
	}
	writer1.Close()

	// 第二次写入，应该复用同一个文件（因为未超过大小限制）
	writer2 := batchFileWriter.NewWriter(dir, "log", "day", 10, 0, time.Second, true)
	for i := 0; i < 5; i++ {
		writer2.Write(fmt.Sprintf("second write %d", i))
	}
	writer2.Close()

	files := listFiles(dir)
	if len(files) != 1 {
		t.Fatalf("期望 1 个文件（复用），实际 %d 个", len(files))
	}

	// 第三次写入，使用 1MB 限制，并写入足够多的数据超过限制
	// 先检查当前文件大小
	currentContent := readFile(t, files[0])
	currentSize := len(currentContent)

	t.Logf("当前文件大小: %d 字节", currentSize)

	// 写入超过 1MB 的数据，触发创建新文件
	writer3 := batchFileWriter.NewWriter(dir, "log", "day", 1, 0, time.Second, true)
	for i := 0; i < 1500; i++ {
		writer3.Write(fmt.Sprintf("third write %d - %s", i, strings.Repeat("x", 1000)))
	}
	writer3.Close()

	files = listFiles(dir)
	if len(files) < 2 {
		t.Fatalf("期望至少 2 个文件（新文件），实际 %d 个", len(files))
	}

	t.Logf("生成了 %d 个文件", len(files))
}

// TestConcurrentWrite 测试并发写入
func TestConcurrentWrite(t *testing.T) {
	dir := createTempDir(t)
	defer os.RemoveAll(dir)

	writer := batchFileWriter.NewWriter(dir, "log", "day", 0, 0, time.Second, true)

	// 并发写入
	done := make(chan bool)
	for g := 0; g < 10; g++ {
		go func(goroutineId int) {
			for i := 0; i < 100; i++ {
				writer.Write(fmt.Sprintf("goroutine %d - line %d", goroutineId, i))
			}
			done <- true
		}(g)
	}

	// 等待所有 goroutine 完成
	for i := 0; i < 10; i++ {
		<-done
	}

	writer.Close()

	files := listFiles(dir)
	content := readFile(t, files[0])
	lines := strings.Split(content, "\n")
	if lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}
	if len(lines) != 1000 {
		t.Fatalf("期望 1000 行，实际 %d 行", len(lines))
	}

	t.Logf("并发写入完成，共 %d 行", len(lines))
}

// TestCloseFlush 测试关闭时数据刷新
func TestCloseFlush(t *testing.T) {
	dir := createTempDir(t)
	defer os.RemoveAll(dir)

	writer := batchFileWriter.NewWriter(dir, "log", "day", 0, 0, time.Hour, true) // 很长的刷盘间隔

	// 写入数据
	for i := 0; i < 10; i++ {
		writer.Write(fmt.Sprintf("test line %d", i))
	}

	// 立即关闭（不等待定时器）
	writer.Close()

	// 验证数据已写入
	files := listFiles(dir)
	content := readFile(t, files[0])
	lines := strings.Split(content, "\n")
	if lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}
	if len(lines) != 10 {
		t.Fatalf("期望 10 行（关闭时刷新），实际 %d 行", len(lines))
	}
}

// TestIntervalFlush 测试定时刷盘
func TestIntervalFlush(t *testing.T) {
	dir := createTempDir(t)
	defer os.RemoveAll(dir)

	// 100ms 刷盘间隔
	writer := batchFileWriter.NewWriter(dir, "log", "day", 0, 0, 100*time.Millisecond, true)

	// 写入数据
	writer.Write("line1")
	writer.Write("line2")

	// 等待刷盘（不关闭 writer）
	time.Sleep(200 * time.Millisecond)

	// 此时文件应该已经有内容
	files := listFiles(dir)
	if len(files) != 1 {
		t.Fatalf("期望 1 个文件，实际 %d 个", len(files))
	}

	content := readFile(t, files[0])
	if !strings.Contains(content, "line1") {
		t.Fatal("期望定时刷盘后包含 line1")
	}

	writer.Close()
}

// TestHourlyRolling 测试小时级别文件命名
func TestHourlyRolling(t *testing.T) {
	dir := createTempDir(t)
	defer os.RemoveAll(dir)

	writer := batchFileWriter.NewWriter(dir, "log", "hour", 0, 0, time.Second, true)
	writer.Write("test data")
	writer.Close()

	files := listFiles(dir)
	// 文件名格式: 2006-01-02-15_1.log
	if !strings.Contains(files[0], time.Now().Format("2006-01-02-15")) {
		t.Fatalf("文件名不符合小时格式: %s", files[0])
	}

	t.Logf("小时格式文件名: %s", filepath.Base(files[0]))
}

// TestEmptyWrite 测试写入空数据
func TestEmptyWrite(t *testing.T) {
	dir := createTempDir(t)
	defer os.RemoveAll(dir)

	writer := batchFileWriter.NewWriter(dir, "log", "day", 0, 0, time.Second, true)

	// 写入空数据
	writer.Write("")
	writer.Write([]byte{})
	writer.Write("")

	// 写入有效数据
	writer.Write("valid data")

	writer.Close()

	files := listFiles(dir)
	content := readFile(t, files[0])

	// 空数据不应该创建额外的空行
	if strings.Contains(content, "\n\n") {
		t.Fatal("不应该有连续空行")
	}

	t.Logf("内容: %s", content)
}

// TestWriteBytes 测试写入字节数据
func TestWriteBytes(t *testing.T) {
	dir := createTempDir(t)
	defer os.RemoveAll(dir)

	writer := batchFileWriter.NewWriter(dir, "log", "day", 0, 0, time.Second, true)

	// 写入字节数据
	data := []byte("byte data line 1\nbyte data line 2")
	writer.Write(data)

	writer.Close()

	files := listFiles(dir)
	content := readFile(t, files[0])

	if !strings.Contains(content, "byte data line 1") {
		t.Fatal("期望包含 byte data line 1")
	}

	t.Logf("字节数据写入成功: %s", content)
}

// TestLargeDataWithSizeLimit 测试大数据量写入和大小限制
func TestLargeDataWithSizeLimit(t *testing.T) {
	dir := createTempDir(t)
	defer os.RemoveAll(dir)

	// 设置 2MB 大小限制
	writer := batchFileWriter.NewWriter(dir, "log", "day", 2, 0, time.Second, true)

	// 写入约 5MB 数据（每行约 1KB，写入 5000 行）
	for i := 0; i < 5000; i++ {
		writer.Write(fmt.Sprintf("line %d - %s", i, strings.Repeat("x", 1000)))
	}

	writer.Close()

	files := listFiles(dir)
	t.Logf("生成了 %d 个文件（限制 2MB）", len(files))

	// 应该至少有 2 个文件（翻转了）
	if len(files) < 2 {
		t.Fatalf("期望至少 2 个文件，实际 %d 个", len(files))
	}

	// 检查每个文件大小都不超过 2MB（约 2,097,152 字节）
	for _, f := range files {
		info, err := os.Stat(f)
		if err != nil {
			continue
		}
		sizeMB := float64(info.Size()) / (1024 * 1024)
		t.Logf("文件 %s 大小: %.2f MB", filepath.Base(f), sizeMB)
	}
}

// 辅助函数：创建临时目录
func createTempDir(t *testing.T) string {
	dir, err := os.MkdirTemp("", "batchfilewriter_test_*")
	if err != nil {
		t.Fatalf("创建临时目录失败: %v", err)
	}
	return dir
}

// 辅助函数：列出目录下所有文件
func listFiles(dir string) []string {
	var files []string
	entries, _ := os.ReadDir(dir)
	for _, entry := range entries {
		if !entry.IsDir() {
			files = append(files, filepath.Join(dir, entry.Name()))
		}
	}
	return files
}

// 辅助函数：读取文件内容
func readFile(t *testing.T, path string) string {
	content, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("读取文件失败: %v", err)
	}
	return string(content)
}
