package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	XuatMenu()
	for {
		fmt.Print("Nhập Bài Giải Số: ")
		option := NhapOption()
		if option == 0 {
			fmt.Println(" ...Thoát Bài Tập...")
			break
		}
		XuOption(option)
	}

	fmt.Println("Bye")
}

func NhapOption() int {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	// handle
	if err != nil {
		fmt.Println("Đã xảy ra lỗi: ", err)
		return -1
	}

	// convert string to int
	option, err := strconv.Atoi(input[:len(input)-2])
	if err != nil {
		fmt.Println("Đã xảy ra lỗi: ", err)
		return -1
	}
	return option
}

func XuatMenu() {
	fmt.Println("===================25 Bài Tập Đầu Năm Con Mồn Lèo=======================")
	fmt.Println("============1. In Các Số Từ 1 Đến 1000 Tăng Dần=========================")
	fmt.Println("============2. In Các Số Từ 1000 Đến 1 Giảm Dần=========================")
	fmt.Println("============3. In Bảng Số Từ 1 Đến 200==================================")
	fmt.Println("============4. Bội Số Của n Với Các Số Từ 1 Đến 20======================")
	fmt.Println("============5. Bảng Cửu Chương Rút Gọn 1 Đến 100========================")
	fmt.Println("============6. In Số Chẵn Từ 1 Đến n ===================================")
	fmt.Println("============7. In Số Lẻ Từ 1 Đến n =====================================")
	fmt.Println("============8. Đếm Số Từ Và Ký Tự Trong Câu=============================")
	fmt.Println("============9. Đảo Ngược Số Nguyên======================================")
	fmt.Println("============10. In Số Nguyên Tố Từ 0 Đến n Bằng FOR=====================")
	fmt.Println("============11. In Số Nguyên Tố Từ 0 Đến n Bằng While===================")
	fmt.Println("============12. In n Phần Tử Đầu Tiên Của Dãy Fibbo=====================")
	fmt.Println("============13. In Phần Tử Đầu Tiên Từ 0 Của Dãy Fibbo Đến n============")
	fmt.Println("============14. Tính Giai Thừa Của x====================================")
	fmt.Println("============15. Vẽ Tam Giác Có Chiều Cao Là n Hàng======================")
	fmt.Println("============16. Chương Trình Nhập n Số Dương============================") //số âm thì dừng
	fmt.Println("============17. Vẽ Tam Giác Cân Rỗng Có Chiều Cao Là n Hàng=============")
	fmt.Println("============18. S= 1+1/23+1/33+...+1/n3=================================")
	fmt.Println("============19. Tổng Bình Phương Từ x Đến y=============================")
	fmt.Println("============20. In Từ n Đến -100. ======================================") //-1 thì nhập lại
	fmt.Println("============21. Kiểm Tra Số Hoàn Hảo====================================")
	fmt.Println("============22. Tính x^y================================================")
	fmt.Println("============23. Kiểm Tra Số Nguyên Tố===================================")
	fmt.Println("============24. Tổng Số Chẵn và Lẻ Giữ 2 Số=============================")
	fmt.Println("============25. S= 1.2 + 2.3 + 3.4 + ... + n(n+1).======================")
	fmt.Println("=============================End========================================")
}

func XuOption(option int) {
	switch option {
	case 1:
		fmt.Println("Bài Giải Số 1: ", option)
		fmt.Println("================In Các Số Từ 1 Đến 1000 Tăng Dần=================")
		InDaySoTangDan()
		fmt.Println("-----------------------------Done--------------------------------")
		fmt.Println("================In Các Số Từ 1 Đến 1000 Tăng Dần=================")
	case 2:
		fmt.Println("Bài Giải Số 2: ", option)
		fmt.Println("================In Các Số Từ 1 Đến 1000 Giảm Dần=================")
		InDaySoGiamDan()
		fmt.Println("-----------------------------Done--------------------------------")
		fmt.Println("================In Các Số Từ 1 Đến 1000 Giảm Dần=================")
	case 3:
		fmt.Println("Bài Giải Số 3: ", option)
		fmt.Println("================In Bảng Số Từ 1 Đến 200=================")
		InBangSo()
		fmt.Println("-----------------------------Done--------------------------------")
		fmt.Println("================In Bảng Số Từ 1 Đến 200=================")
	case 4:
		fmt.Println("Bài Giải Số 4: ", option)
		fmt.Println("================Bội Số Của n Với Các Số Từ 1 Đến 20=================")
		Boisonhan()
		fmt.Println("-----------------------------Done--------------------------------")
		fmt.Println("================Bội Số Của n Với Các Số Từ 1 Đến 20=================")
	case 5:
		fmt.Println("Bài Giải Số 5: ", option)
		fmt.Println("================Bảng Cửu Chương Rút Gọn 1 Đến 100=================")
		Bangcuuchuongrutgon()
		fmt.Println("-----------------------------Done--------------------------------")
		fmt.Println("================Bảng Cửu Chương Rút Gọn 1 Đến 100=================")
	case 6:
		fmt.Println("Bài Giải Số 6: ", option)
		fmt.Println("================In Số Chẵn Từ 1 Đến n =================")
		InSoChan()
		fmt.Println("-----------------------------Done--------------------------------")
		fmt.Println("================In Số Chẵn Từ 1 Đến n =================")
	case 7:
		fmt.Println("Bài Giải Số 7: ", option)
		fmt.Println("================In Số Lẻ Từ 1 Đến n =================")
		InSoLe()
		fmt.Println("-----------------------------Done--------------------------------")
		fmt.Println("================In Số Lẻ Từ 1 Đến n =================")
	case 8:
		fmt.Println("Bài Giải Số 8: ", option)
		fmt.Println("================Đếm Số Từ Và Ký Tự Trong Câu The Gioi Di Dong=================")
		DemKyTuVaSoTu()
		fmt.Println("-----------------------------Done--------------------------------")
		fmt.Println("================Đếm Số Từ Và Ký Tự Trong Câu The Gioi Di Dong=================")
	case 9:
		fmt.Println("Bài Giải Số 9: ", option)
		fmt.Println("================Đảo Ngược Số Nguyên=================")
		DaoNguocSo()
		fmt.Println("-----------------------------Done--------------------------------")
		fmt.Println("================Đảo Ngược Số Nguyên=================")
	case 10:
		fmt.Println("Bài Giải Số 10: ", option)
		fmt.Println("================In Số Nguyên Tố Từ 0 Đến n Bằng FOR=================")
		SoNguyenDuongDenNPhanTu()
		fmt.Println("-----------------------------Done--------------------------------")
		fmt.Println("================In Số Nguyên Tố Từ 0 Đến n Bằng FOR=================")
	case 11:
		fmt.Println("Bài Giải Số 11: ", option)
		fmt.Println("================In Số Nguyên Tố Từ 0 Đến n Bằng WHILE=================")
		SoNguyenToDenNPhanTuWhile()
		fmt.Println("-----------------------------Done--------------------------------")
		fmt.Println("================In Số Nguyên Tố Từ 0 Đến n Bằng WHILE=================")
	case 12:
		fmt.Println("Bài Giải Số 12: ", option)
		fmt.Println("================In n Phần Tử Đầu Tiên Của Dãy Fibbo=================")
		FibonacciInNPhanTuDauTien()
		fmt.Println(" \n ")
		fmt.Println("-----------------------------Done--------------------------------")
		fmt.Println("================In n Phần Tử Đầu Tiên Của Dãy Fibbo=================")
	case 13:
		fmt.Println("Bài Giải Số 13: ", option)
		fmt.Println("================In Phần Tử Đầu Tiên Từ 0 Của Dãy Fibbo Đến n=================")
		FibonacciInNPhanTuDauTienTu0DenN()
		fmt.Println(" \n ")
		fmt.Println("-----------------------------Done--------------------------------")
		fmt.Println("================In Phần Tử Đầu Tiên Từ 0 Của Dãy Fibbo Đến n=================")
	case 14:
		fmt.Println("Bài Giải Số 14: ", option)
		fmt.Println("================Giai Thừa Của X=================")
		GiaithuacuaX()
		fmt.Println(" \n ")
		fmt.Println("-----------------------------Done--------------------------------")
		fmt.Println("================Giai Thừa Của X=================")
	case 15:
		fmt.Println("Bài Giải Số 15: ", option)
		fmt.Println("================Vẽ Tam Giác Có Chiều Cao Là n Hàng=================")
		Triangle()
		fmt.Println(" \n ")
		fmt.Println("-----------------------------Done--------------------------------")
		fmt.Println("================Vẽ Tam Giác Có Chiều Cao Là n Hàng=================")
	case 16:
		fmt.Println("Bài Giải Số 16: ", option)
		fmt.Println("================Chương Trình Nhập n Số Dương=================")
		DungNeuLaSoAm()
		fmt.Println(" \n ")
		fmt.Println("-----------------------------Done--------------------------------")
		fmt.Println("================Chương Trình Nhập n Số Dương=================")
	case 17:
		fmt.Println("Bài Giải Số 17: ", option)
		fmt.Println("================Vẽ Tam Giác Cân Rỗng Có Chiều Cao Là n Hàng=================")
		TamGiacCan()
		fmt.Println(" \n ")
		fmt.Println("-----------------------------Done--------------------------------")
		fmt.Println("================Vẽ Tam Giác Cân Rỗng Có Chiều Cao Là n Hàng=================")
	case 18:
		fmt.Println("Bài Giải Số 18: ", option)
		fmt.Println("================S= 1+1/23+1/33+...+1/n3=================")
		KetQuaBieuThuc()
		fmt.Println(" \n ")
		fmt.Println("-----------------------------Done--------------------------------")
		fmt.Println("================S= 1+1/23+1/33+...+1/n3=================")
	case 19:
		fmt.Println("Bài Giải Số 19: ", option)
		fmt.Println("================Tổng Bình Phương Từ x Đến y=================")
		TongBinhPhuongXY()
		fmt.Println(" \n ")
		fmt.Println("-----------------------------Done--------------------------------")
		fmt.Println("================Tổng Bình Phương Từ x Đến y=================")
	case 20:
		fmt.Println("Bài Giải Số 20: ", option)
		fmt.Println("================In Từ n Đến -100=================")
		InSoAm()
		fmt.Println(" \n ")
		fmt.Println("-----------------------------Done--------------------------------")
		fmt.Println("================In Từ n Đến -100=================")
	case 21:
		fmt.Println("Bài Giải Số 21: ", option)
		fmt.Println("================Kiểm Tra Số Hoàn Hảo=================")
		SoHoanHao()
		fmt.Println(" \n ")
		fmt.Println("-----------------------------Done--------------------------------")
		fmt.Println("================Kiểm Tra Số Hoàn Hảo=================")
	case 22:
		fmt.Println("Bài Giải Số 22: ", option)
		fmt.Println("================Tính x^y=================")
		XLuyThuaY()
		fmt.Println(" \n ")
		fmt.Println("-----------------------------Done--------------------------------")
		fmt.Println("================Tính x^y=================")
	case 23:
		fmt.Println("Bài Giải Số 23: ", option)
		fmt.Println("================Kiểm Tra Số Nguyên Tố=================")
		KtraSoNguyenTo()
		fmt.Println(" \n ")
		fmt.Println("-----------------------------Done--------------------------------")
		fmt.Println("================Kiểm Tra Số Nguyên Tố=================")
	case 24:
		fmt.Println("Bài Giải Số 24: ", option)
		fmt.Println("================In Bảng Số Từ 1 Đến 200=================")
		TongSoChanLe()
		fmt.Println(" \n ")
		fmt.Println("-----------------------------Done--------------------------------")
		fmt.Println("================In Bảng Số Từ 1 Đến 200=================")
	case 25:
		fmt.Println("Bài Giải Số 25: ", option)
		fmt.Println("================S= 1.2 + 2.3 + 3.4 + ... + n(n+1)=================")
		GiatriBieuThuc()
		fmt.Println(" \n ")
		fmt.Println("-----------------------------Done--------------------------------")
		fmt.Println("================S= 1.2 + 2.3 + 3.4 + ... + n(n+1)=================")

	default:
		fmt.Println("Option bạn chọn không có trong menu. Vui lòng chọn lại")
		fmt.Println("------------------------------------------------------")
		break
	}
}

func InDaySoTangDan() {
	for i := 1; i < 1001; i++ {
		fmt.Print(i, " ")
	}
}
func InDaySoGiamDan() {
	for i := 1000; i > 0; i-- {
		fmt.Print(i, " ")
	}
}
func InBangSo() {
	for i := 1; i < 201; i++ {

		fmt.Print(i, "\t")
	}
}
func Boisonhan() {

	var num int = 0

	fmt.Print("Enter Number: ")
	fmt.Scanf("%d", &num)

	for count := 1; count <= 20; count++ {
		fmt.Printf("%d * %d = %d\n", num, count, num*count)
	}

}
func Bangcuuchuongrutgon() {
	for x := 1; x <= 100; x++ {
		for y := 1; y <= 10; y++ {
			fmt.Print(x*y, "  ")
		}
		fmt.Print("\n")
	}
}
func InSoChan() {
	var evnum, i int
	fmt.Print("Enter the Number to Print Even's=")
	fmt.Scanln(&evnum)
	fmt.Println("Even Numbers from 1 to ", evnum, " are = ")
	for i = 2; i <= evnum; i = i + 2 {
		fmt.Print(i, "\t")
	}
	fmt.Println()
}
func InSoLe() {
	var odd, i int
	fmt.Print("Enter the Number to Print Odd's=")
	fmt.Scanln(&odd)
	fmt.Println("Odd Numbers from 1 to ", odd, " are = ")
	for i = 1; i <= odd; i = i + 2 {
		fmt.Print(i, "\t")
	}
	fmt.Println()
}
func DemKyTuVaSoTu() {
	Output, Output2, _ := Giaiphap("the gioi di dong")
	fmt.Println("Số Từ: ", Output)
	fmt.Println("Số Ký Tự: ", Output2)
}
func Giaiphap(S string) (int, int, string) {
	demtu := 1
	demkytu := 0
	for _, v := range S {
		if v == 32 {
			demkytu++
		}
		if v != 32 {
			demtu++
		}
	}
	return demtu, demkytu, "hello"
}
func DaoNguocSo() {
	var revnum, remainder int

	fmt.Print("Enter the Number to Reverse = ")
	fmt.Scanln(&revnum)

	reverse := 0

	for revnum > 0 {
		remainder = revnum % 10
		reverse = reverse*10 + remainder
		revnum = revnum / 10
	}

	fmt.Println("The Reverse of the Given Number = ", reverse)
}
func SoNguyenDuongDenNPhanTu() {
	var primNum, primMin, primMax, primcount int

	fmt.Print("Đặt giới hạn cho số nguyên tố từ 0 đến n = ")
	fmt.Scanln(&primMin, &primMax)

	fmt.Println("Số Nguyên Tố Giữa ", primMin, " và ", primMax, " là ")
	for primNum = primMin; primNum <= primMax; primNum++ {
		primcount = 0
		for i := 2; i < primNum/2; i++ {
			if primNum%i == 0 {
				primcount++
				break
			}
		}
		if primcount == 0 && primNum != 1 {
			fmt.Print(primNum, "\t")
		}
	}
	fmt.Println()
}
func SoNguyenToDenNPhanTuWhile() {
	var primNum, primMin, primMax, primcount int

	fmt.Print("Enter the Minimum and Maximum Limit for Prime Numbers = ")
	fmt.Scanln(&primMin, &primMax)

	fmt.Println("Prime Numbers between ", primMin, " and ", primMax, " are ")
	for primNum = primMin; primNum <= primMax; primNum++ {
		primcount = 0
		for i := 2; i < primNum/2; i++ {
			if primNum%i == 0 {
				primcount++
				break
			}
		}
		if primcount == 0 && primNum != 1 {
			fmt.Print(primNum, "\t")
		}
	}
	fmt.Println()
}
func FibonacciInNPhanTuDauTien() {
	var n int
	t1 := 0
	t2 := 1
	nextTerm := 0

	fmt.Print("Nhập Số Nguyên Dương n : ")
	fmt.Scan(&n)
	fmt.Print("Dãy Fibonacci Đến n :")
	for i := 1; i <= n; i++ {
		if i == 1 {
			fmt.Print(" ", t1)
			continue
		}
		if i == 2 {
			fmt.Print(" ", t2)
			continue
		}
		nextTerm = t1 + t2
		t1 = t2
		t2 = nextTerm
		fmt.Print(" ", nextTerm)
	}
}
func FibonacciInNPhanTuDauTienTu0DenN() {
	var n int
	t1 := 0
	t2 := 1
	nextTerm := 0

	fmt.Print("Nhập Số Nguyên Dương n : ")
	fmt.Scan(&n)
	fmt.Print("Dãy Fibonacci Đến n :")
	for i := 1; i <= n; i++ {
		if i == 1 {
			fmt.Print(" ", t1)
			continue
		}
		if i == 2 {
			fmt.Print(" ", t2)
			continue
		}

		nextTerm = t1 + t2
		t1 = t2
		t2 = nextTerm
		if nextTerm == n {
			fmt.Print(nextTerm)
		} else {
			fmt.Print(" ", nextTerm)

		}
	}
}
func GiaithuacuaX() {

	var _ int = 1
	var n int

	fmt.Print("Nhập X : ")
	fmt.Scan(&n)
	fmt.Print("Bội Số Của X là: ", BoiSoNhan(n))
}
func BoiSoNhan(n int) uint64 {
	var factVal uint64 = 1 // uint64 is the set of all unsigned 64-bit integers.
	// Range: 0 through 18446744073709551615.
	if n < 0 {
		fmt.Print("Factorial of negative number doesn't exist.")
	} else {
		for i := 1; i <= n; i++ {
			factVal *= uint64(i) // mismatched types int64 and int
		}

	}
	return factVal /* return from function*/
}
func Triangle() {
	var i, j, row int
	var n int
	fmt.Print("Nhập Số Nguyên Dương n : ")
	fmt.Scan(&n)
	row = n
	fmt.Println("\nRight Angled Star Pattern")
	for i = 1; i <= row; i++ {
		for j = 1; j <= i; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}
func DungNeuLaSoAm() {
	var n int = 0
	fmt.Print("Nhập Số Nguyên Dương n : ")
	fmt.Scan(&n)
	for i := 0; i <= n; i++ {
		if n <= -15 {
			fmt.Print("n phải là số dương ")
		} else {
			fmt.Print(i, " ")
		}
	}
}
func TamGiacCan() {
	var n int
	fmt.Print("Nhập Số Nguyên Dương n : ")
	fmt.Scan(&n)
	if n > 0 {
		for i := 0; i < n; i++ {
			for j := 0; j < 2*n+1; j++ {
				if j == n-i || j == n+i {
					fmt.Print(" * ")
				} else {
					fmt.Print("   ")
				}
			}
			fmt.Print(("\n"))
		}
		for j := 0; j < 2*n+1; j++ {
			fmt.Print(" * ")
		}

	}
}
func KetQuaBieuThuc() {
	var n int
	fmt.Print("Nhập Số Nguyên Dương n : ")
	fmt.Scan(&n)
	var S float64 = 0
	for i := 1; i <= n; i++ {
		S += 1 / float64(i*i*i)
	}
	fmt.Print(S)

}
func TongBinhPhuongXY() {
	var x int
	var y int
	var S int
	fmt.Print("Nhập Số Nguyên Dương x : ")
	fmt.Scan(&x)
	fmt.Print("Nhập Số Nguyên Dương y : ")
	fmt.Scan(&y)
	for i := x; i <= y; i++ {
		S = S + i*i
	}
	fmt.Print(S)
}
func InSoAm() {
	var n int
	fmt.Print("Nhập Số Nguyên Dương n : ")
	fmt.Scan(&n)
	if n == -1 {
		fmt.Print("Nhập Lại n : ")
		fmt.Scan(&n)
	}
	if n > -100 {
		for i := n; i >= -100; i-- {
			fmt.Print(i, "\n")
		}
	} else {
		for i := n; i <= -100; i++ {
			fmt.Print(i, "\n")
		}
	}
}
func SoHoanHao() {
	var n int
	var i, s int

	fmt.Print("Nhập Số Nguyên Dương n : ")
	fmt.Scan(&n)

	if n <= 0 {
		fmt.Println("N phai lon hon 0")
	} else {
		for i = 1; i <= n-1; i++ {
			if n%i == 0 {
				s += i
			}
		}
		if s == n {
			fmt.Print(n, " la so hoan hao")
		} else {
			fmt.Print(n, " khong phai la so hoan hao")
		}
	}
}
func XLuyThuaY() {
	var x int
	var y int
	var luy_thua = 1
	fmt.Print("Nhập Số Nguyên x : ")
	fmt.Scan(&x)
	fmt.Print("Nhập Số Nguyên y : ")
	fmt.Scan(&y)
	for i := 1; i <= y; i++ {
		luy_thua *= x
	}
	fmt.Print(luy_thua)
}
func KtraSoNguyenTo() {
	var n int
	kiemtra := 1
	fmt.Print("Nhập Số Nguyên n : ")
	fmt.Scan(&n)
	for i := 2; i <= n-1; i++ {
		if n%i == 0 {
			kiemtra = 0
			break
		}
	}
	if kiemtra == 1 {
		fmt.Println(n, "là số nguyên tố")
	} else if kiemtra == 0 {
		fmt.Println(n, "không phải là số nguyên tố")
	}
}
func TongSoChanLe() {
	var a int
	var b int
	fmt.Print("Nhập Số Nguyên a : ")
	fmt.Scan(&a)
	fmt.Print("Nhập Số Nguyên b : ")
	fmt.Scan(&b)
	if a == b {
		fmt.Println("a khong duoc bang b")
	}
	tong_so_chan := 0
	tong_so_le := 0
	if a > b {
		for i := a; i >= b; i-- {
			if i%2 == 0 {
				tong_so_chan += i
			} else {
				tong_so_le += i
			}
		}
		fmt.Println("Tổng số lẻ", tong_so_le)
		fmt.Println("Tổng số chẵn", tong_so_chan)
	} else {
		for i := b; i >= a; i-- {
			if i%2 == 0 {
				tong_so_chan += i
			} else {
				tong_so_le += i
			}
		}
		fmt.Println("Tổng số lẻ", tong_so_le)
		fmt.Println("Tổng số chẵn", tong_so_chan)
	}
}
func GiatriBieuThuc() {
	var n int
	var s int = 0
	fmt.Print("Nhập Số Nguyên n : ")
	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		s = s + i*(i+1)
	}
	fmt.Print(s)
}
