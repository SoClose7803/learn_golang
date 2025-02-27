package main

import(
	"fmt"
)
type Contact struct {
	Phone string
	Email string
	Name string
}
var phoneList = make (map[string]Contact)
func addContact(phone string, email string, name string){
	phoneList[phone] = Contact{Phone: phone, Email: email, Name: name}
	fmt.Println("Contact added successfully")
}
func removeContact(name string){
	if _, ok := phoneList[name]; ok {
		delete(phoneList, name)
		fmt.Println("Contact removed successfully")
	} else {
		fmt.Println("Contact not found")
	}
}
func searchContact(name string){
	for _, contact := range phoneList {
		if contact.Name == name {
			fmt.Println("Phone:", contact.Phone)
			fmt.Println("Email:", contact.Email)
			return
		}
	}
	fmt.Println("Contact not found")
}
func DisplayContacts() {
    fmt.Println("Danh bạ liên hệ:")
    for _, contact := range phoneList {
        fmt.Printf("Tên: %s, SĐT: %s, Email: %s\n", contact.Name, contact.Phone, contact.Email)
    }
}

func main(){
	for {
        fmt.Println("\nChương trình quản lý danh bạ")
        fmt.Println("1. Thêm liên hệ")
        fmt.Println("2. Xoá liên hệ")
        fmt.Println("3. Tìm kiếm liên hệ")
        fmt.Println("4. Hiển thị danh bạ")
        fmt.Println("5. Thoát")
        fmt.Print("Chọn một tuỳ chọn: ")

        var choice int
        fmt.Scan(&choice)

        switch choice {
        case 1:
            var name, phone, email string
            fmt.Print("Nhập tên: ")
            fmt.Scan(&name)
            fmt.Print("Nhập số điện thoại: ")
            fmt.Scan(&phone)
            fmt.Print("Nhập email: ")
            fmt.Scan(&email)
            addContact(name, phone, email)
        case 2:
            var name string
            fmt.Print("Nhập tên liên hệ cần xoá: ")
            fmt.Scan(&name)
            removeContact(name)
        case 3:
            var name string
            fmt.Print("Nhập tên liên hệ cần tìm: ")
            fmt.Scan(&name)
            searchContact(name)
        case 4:
            DisplayContacts()
        case 5:
            fmt.Println("Thoát chương trình.")
            return
        default:
            fmt.Println("Lựa chọn không hợp lệ, vui lòng thử lại!")
        }
    }
}
