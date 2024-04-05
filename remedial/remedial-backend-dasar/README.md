# Proyek: Student Portal dengan Interfaces

## Deskripsi

Portal mahasiswa sederhana ini akan berfungsi sebagai pengantar untuk Golang, berfokus pada struktur data dan kontrol dasar serta interfaces. Portal ini akan memungkinkan mahasiswa untuk login, melihat detail pribadi mereka, mata kuliah, dan nilai. 

Proyek ini dirancang untuk menjadi sederhana dan ramah bagi pemula. Ini memungkinkan mahasiswa untuk berlatih menggunakan sintaks dasar Go dan struktur data seperti slices, maps, dan structs, struktur kontrol seperti loops dan conditionals, dan interfaces.

Proyek ini menampilkan beberapa fitur utama seperti:

- Student Login: Memeriksa apakah data mahasiswa yang diinput oleh pengguna sudah terdaftar.
- View Student Details: Menampilkan detail pribadi mahasiswa.
- View Course Details: Menampilkan detail mata kuliah yang diikuti oleh mahasiswa.
- View Grades: Menampilkan nilai dari setiap mata kuliah yang diambil oleh mahasiswa.

## Constraints

- ID mahasiswa hanya berupa kombinasi dari huruf dan angka, dengan panjang 5 karakter.
- Nama mahasiswa hanya berupa huruf, dengan panjang maksimal 50 karakter.
- mata kuliah yang diambil oleh mahasiswa disimpan dalam bentuk array (slices).
- Nilai dari setiap mata kuliah disimpan dalam bentuk map, dengan key berupa Course dan value berupa int.
- Struct Course hanya memiliki dua field, yaitu ID dan Name.

## Test Case Examples

Berikut adalah beberapa contoh kasus uji:

Test Case 1

Input:

```go
student := manager.Login(1)
```
Expected Output / Behavior:

Menampilkan detail mahasiswa dengan ID 1 jika mahasiswa tersebut ada dalam array Students. Jika tidak, mengembalikan nilai nil.

Test Case 2

Input:

```go
manager.ViewStudentDetails(student)
```
Expected Output / Behavior:

Menampilkan detail mahasiswa seperti "Student ID:", "First Name:", dan "Last Name:".

Test Case 3

Input:

```go
manager.ViewCourseDetails(student)
```
Expected Output / Behavior:

Menampilkan detail mata kuliah yang diambil oleh mahasiswa seperti "Course ID:" dan "Course Name:".

Test Case 4

Input:

```go
manager.ViewGrades(student)
```
Expected Output / Behavior:

Menampilkan nilai dari setiap mata kuliah yang diambil oleh mahasiswa seperti "Course:" dan "Grade:".
