package main

import (
	"fmt"
)

func main() {

	//nil is the zero value for pointers, interfaces, maps, slices, channels and function types,\

	//Arrays
	//Array (mảng) trong Go tương tự các ngôn ngữ khác, tuy nhiên nó có kích thước cố định (fixed size) và các phần tử bên trong phải cùng loại dữ liệu
	// Declaring
	// Var array_name[length]Type
	// var array_name[length]Typle{item1, item2, item3, ...itemN}

	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)
	primes := [6]int{2, 3, 5, 7, 11} // [6]int{2, 3, 5, 7, 11,0}
	fmt.Println(primes)

	// Slice
	// Slice là một tham chiếu đến Array, nó mô tả một phần (hoặc toàn bộ) Array. Nó có kích thước động nên thường được sử dụng nhiều hơn Array.
	// Khởi tạo Slice s bằng cách cắt từ phần tử ở vị trí 1 (low) đến phần tử ở vị trí 3 (high - 1) của Array primes
	var s []int = primes[1:4]
	// In ra giá trị của Slice s
	fmt.Println(s) // Giá trị của s là [3, 5, 7]

	// Do Slice chỉ là tham chiếu đến Array, do đó thay đổi giá trị của Slice sẽ làm thay đổi giá trị của Array mà nó tham chiếu đến.
	// Nếu có nhiều Slice cùng tham chiếu đến một Array thì khi thay đổi giá trị một Slice có thể làm thay đổi

	//khi thay đổi 1 phần tử trong slice thì sẽ thay đổi array mà nó tham chiếu tới
	s[0] = 9

	fmt.Println(s)
	fmt.Println(primes)

	// array[i:n] cắt từ i đến n-1
	new := s[:2]
	fmt.Println(new)

	// xóa phần tử thứ n của mảng
	var newprimes []int = append(primes[:2], primes[3:]...) //(thực ra là tạo mảng mới rồi đưa 2 Slice từ đầu đến n-1 và từ n+1 đến hết vào) =))
	fmt.Println(newprimes)

	mySlice := append(primes[:2], primes[3:]...)
	fmt.Printf("giá trị mySlice là %d \n", mySlice)

	// map
	//Map (bản đồ) là một kiểu dữ liệu được dựng sẵn trong Go,
	//một map là tập hợp các cặp key/value (khóa/giá trị) trong đó một value được liên kết với một key. Value chỉ được truy xuất bởi key tương ứng.
	// khởi tạo
	//make(map[type of key]type of value)
	type Vertex struct {
		Lat, Long float64
	}
	var m map[string]Vertex

	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])

	//map inside map
	var d = map[string]string{
		"w": "x",
	}

	var data = map[string]map[string]string{
		"a": d,
		"b": d,
		"c": d,
	}

	fmt.Println(data)

	//Map là kiểu tham chiếu
	//Tương tự như slice, map là kiểu tham chiếu.
	//Khi một map được gán cho một biến mới, cả hai đều trỏ đến cùng một cấu trúc dữ liệu nội bộ.
	//Do đó những thay đổi được thực hiện ở một biến sẽ ánh xạ đến biến kia.

	personSalary := map[string]int{
		"steve": 12000,
		"jamie": 15000,
	}
	personSalary["mike"] = 9000
	fmt.Println("Original person salary", personSalary)
	newPersonSalary := personSalary
	newPersonSalary["mike"] = 18000
	fmt.Println("Person salary changed", personSalary)

	//Map không thể so sánh với nhau bằng toán tử ==. Ta chỉ có thể sử dụng == để kiểm tra xem một map có phải là nil hay không.

	map1 := map[string]int{
		"one": 1,
		"two": 2,
	}

	if map1 == nil {
		fmt.Println(" map1 = nil")
	} else {
		fmt.Println(" map1 != nil")
	}
	// map2 := map1

	// if map1 == map2 {
	// 	fmt.Println("So sánh được")
	// }

	//invalid operation: map1 == map2 (map can only be compared to nil).

	//Pointers
	x := 1
	i, j := 42, 2701

	p := &i         // p trỏ tới i
	fmt.Println(p)  // in ra giá trị ngan nho của i
	fmt.Println(*p) // in ra giá trị của i
	*p = 21         // gán i  thông qua pointer
	fmt.Println(i)  // in ra giá trị của i

	p = &j         // p trỏ tới j
	*p = *p / 37   // chia j cho 37 thông qua con trỏ *p
	fmt.Println(j) // in ra giá trị của j

	o := 25
	var n *int
	if n == nil {
		fmt.Println("giá trị n là", n)
		n = &o
		fmt.Println("giá trị n sau khi đưa giá trị vào o", n)
	}
	//
	//fmt.Scanf("%d", &x)

	a1 := 58
	fmt.Println("giá trị của a trước khi gọi hàm change là ", a1)
	b1 := &a1
	change(b1)
	fmt.Println("giá trị của a sau khi gọi hàm change là", a1)
	//fmt.Scanf("%d", &x)

	//Structs
	// Structures là gì?
	// Structures (cấu trúc) được người dùng định nghĩa cho một tập hợp các trường.
	// Nó có thể được sử dụng để nhóm dữ liệu vào một cấu trúc thay vì viết từng loại riêng biệt dưới các dạng riêng biệt.

	type Vertex1 struct {
		X int
		Y int
	}

	// Định nghĩa một struct với từ khóa type
	type Student struct {
		name string
		age  int
	}
	// Khởi tạo biến s1 có giá trị là struct Student
	s1 := Student{"Robin", 30} // {"Robin", 30}

	// Khởi tạo biến s2 có giá trị là struct Student với 1 field là name
	// Field còn lại sẽ có giá trị mặc định (zero value)
	s2 := Student{name: "Robin"} // {"Robin", 0}

	// Khởi tạo biến s3 có giá trị là struct Student và không khai báo giá trị cho trường nào
	s3 := Student{} // {"", 0}

	// Thay đổi giá trị field trong struct
	s3.name = "Robert"
	s3.age = 25

	fmt.Println(s3) // s3 = {"Robert", 25}

	s1 = Student{"Steve", 30}
	s2 = Student{"Steve", 30}
	s3 = Student{"Job", 30}

	if s1 == s2 {
		fmt.Println("s1 = s2t") // s1 bằng s2
	} else {
		fmt.Println("s1 != s2")
	}

	if s2 == s3 {
		fmt.Println("s2 = s3")
	} else {
		fmt.Println("s2 != s3") // s2 khác s3
	}
	//fmt.Scanf("%d", &b)

	name := MyString("Sam Anderson")
	var v VowelsFinder
	v = name // Điều này được chấp nhận vì MyString implement interface VowelsFinder
	fmt.Printf("Vowels are %c \n", v.FindVowels())

	pemp1 := Permanent{1, 5000, 20}
	pemp2 := Permanent{2, 6000, 30}
	cemp1 := Contract{3, 3000}
	employees := []SalaryCalculator{pemp1, pemp2, cemp1}
	totalExpense(employees)

	var d1 Describer
	if d1 == nil {
		fmt.Printf("d1 is nil and has type %T value %v\n", d1, d1)
	}

	//Các biến cấu trúc không thể so sánh nếu chúng chứa các trường không thể so sánh được.

	// type image struct {
	// 	data map[int]int
	// }

	// image1 := image{data: map[int]int{
	//     0: 155,
	// }}
	// image2 := image{data: map[int]int{
	//     0: 155,
	// }}
	// if image1 == image2 {
	//     fmt.Println("image1 and image2 are equal")
	// }
	//cannot compare image1 == image2 (operator == not defined for image)

	fmt.Printf("Kết thúc phần code demo 3 \n")
	fmt.Scanf("%d", &x)

}

// Describer interface có fnc Describe
type Describer interface {
	Describe()
}

//hàm này gán giá trị 55 vào cho biến được trỏ tới bởi con trỏ val
func change(val *int) {
	*val = 55
}

//VowelsFinder định nghĩa 1 interface
type VowelsFinder interface {
	FindVowels() []rune
}

//MyString  MyString được gọi là implement interface VowelsFinder.
type MyString string

//Range
//Range là một hình thức của vòng lặp for dùng để duyệt qua một slice hoặc map
//Mỗi một vòng lặp sẽ trả về 2 giá trị: Giá trị đầu tiên là chỉ số (vị trí) của phần tử, và giá trị thứ hai là bản sao của phần tử đó (cùng giá trị)
//Trong trường hợp khi lặp chỉ sử dụng 1 trong 2 giá trị trả về thì ta sẽ bỏ qua giá trị còn lại bằng cách thay tên biến
//bằng ký tự gạch dưới (vì nếu không thì khi biên dịch sẽ báo lỗi biến được định nghĩa mà không sử dụng).

// FindVowels MyString implements VowelsFinder
func (ms MyString) FindVowels() []rune {
	var vowels []rune
	for _, rune := range ms {
		if rune == 'a' || rune == 'e' || rune == 'i' || rune == 'o' || rune == 'u' {
			vowels = append(vowels, rune)
		}
	}
	return vowels
}

// SalaryCalculator interface
type SalaryCalculator interface {
	CalculateSalary() int
}

// Declaring
// type identifier struct{
// 	field1 data_type
// 	field2 data_type
// 	field3 data_type
//   }

//Permanent struct nhân viên hợp đồng dài hạn
type Permanent struct {
	empID    int
	basicpay int
	pf       int
}

//Contract struct nhân viên hợp đồng ngắn hạn
type Contract struct {
	empID    int
	basicpay int
}

//CalculateSalary hàm tính tiền lương của nhân viên permanent bằng tổng của basic pay và pf
func (p Permanent) CalculateSalary() int {
	return p.basicpay + p.pf
}

//CalculateSalary hàm tính tiền lương của nhân viên contract chỉ là basic pay
func (c Contract) CalculateSalary() int {
	return c.basicpay
}

/*
tổng chi phí được tính bằng cách duyệt qua từng phần tử của slice SalaryCalculator
và tính tổng mức lương của từng nhân viên
*/
func totalExpense(s []SalaryCalculator) {
	expense := 0
	for _, v := range s {
		expense = expense + v.CalculateSalary()
	}
	fmt.Printf("Total Expense Per Month $%d \n", expense)
}
