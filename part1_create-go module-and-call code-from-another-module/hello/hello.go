package main

import (
    "fmt"

    "example.com/greetings"
)

func main() {
    // Get a greeting message and print it.
    message := greetings.Hello("Fauzi")
    fmt.Println(message)
}

//printah "go mod init example.com/hello" To enable dependency tracking for your code
//perintah pada Bash "go mod edit -replace example.com/greetings=../greetings" menentukan apa "example.com/greetings" yang harus diganti dengan ""../greetings" untuk tujuan menemukan dependensi.
//perintah "go mod tidy" untuk menyinkronkan dependensi modul example.com/hello, menambahkan dependensi yang diperlukan oleh kode, tetapi belum dilacak dalam modul
//untuk menyambungkan gunakan  "go get example.com/greetings"
//Perintah "go run ." untuk menjalankan kode. 

