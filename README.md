# docker 

## xóa toàn bộ image 
docker rmi -f $(docker images -a -q)

## build image 
docker build -t app .

## docker run 
docker run -p 80:8081 -it app



# Go Kiểu Dữ Liệu

### Kiểu chuỗi: String string
 Kiểu Dữ Liệu String trong Golang các chuỗi trong Golang bao gồm các ký tự được đặt trong dấu ngoặc kép "":
var a string
a = "Abcdef"
hoặc
var a string = "Abcdef"

### Số nguyên: Integer int. Ngoài ra còn có int8, int16, int32 và int64 để quy định số lượng bit tối đa có thể sử dụng để lưu giá trị. Khi sử dụng int Go sẽ tự động chuyển thành int32 trên máy tính 32 bit và int64 trên máy 64 bit.

Kiểu số nguyên bao gồm các số 0 số nguyên âm và nguyên dương:

var a, b int
a = 100
b = -100


### Số nguyên dương: Unsigned integer uint
Unsigned Integer trong Golang
Unsigned Integer bao gồm các số nguyên không âm (hay chính là số tự nhiên):

var a uint
a = 100 // OK
a = -100 // Lỗi

### Kiểu logic đúng/sai: Boolean bol
Kiểu Dữ Liệu Boolean trong Golang
Kiểu boolean chỉ gồm 2 giá trị là True và False:
var a, b bool
a = True
b = 0 // Tương tự b = False

### Kiểu byte: Byte byte bao gồm các số nguyên với chỉ 8 byte dữ liệu còn được viết với cách khác là int8
Kiểu Dữ Liệu Byte Trong Golang
Kiểu dữ liệu byte tương đương với các số nguyên cần tới tối đa 8 bit bộ nhớ để lưu trữ.
var a byte
a = 2

### Kiểu số thực float : Gồm có float32 (với các số float sử dụng tối đa 32 bit để lưu trữ) và float64 (với các số float sử dụng tối đa 64 bit để lưu trữ)
Kiểu Dữ Liệu Float trong Golang
Kiểu Float dùng để biểu diễn các số thực, chúng ta sử dụng float64 với hệ điều hành 64 bit và float32 với hệ điều hành 32 bit:
var a float64
a = 1,234




