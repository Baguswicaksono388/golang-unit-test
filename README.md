# Aturan File Test
- Go-Lang memiliki aturan cara membuat file khusus untuk unit test
- Unk membuat file unit test, kita harus menggunakan akhiran _test
- Jadi kita misal membuat file hello_world.go, artinya untuk membuat unit testnya, kita harus membuat file hello_world_test.go

# Aturan Function Unit Test
- Selain aturan nama file, di Go-Lang juga sdh diatur unk nama function unit test
- Nama function untuk unit test harus diawali dengan nama Test
- Misal jika kita ingin mengetest func HelloWorld, maka kita akan membuat func unit test dgn nama TestHelloWorld
- Tdk ada aturan unk nama belakang func unit test harus sama dgn nama func yg akan di test, yg penting adlh harus diawali dgn kata Test
- Selanjutnya hrs memiliki parameter (t *testing.T) dan tidak mengembalikan return value

# Menjalankan Unit Test
- Unk menjalankan unit test kita bisa menggunakan perintah : go test
- Jika kita ingin melihat lebih detail func test apa saja yg sudah di running, kita bisa gunakan perintah : go test -v
- Dan jika kita ingin memilih func unit test mana yg ingim di running, kita bisa gunakan perintah : go test -v -run TestNamaFunction

# Menjalankan semua unit test
- Jika kita ingin menjalankan semua unit test dari top folder module nya, kita bisa menggunakan perintah : go test ./...

# Menggagalkan Unit Test
- Menggagalkan unit Test menggunakan panic bukanlah hal yg bagus
- Go-Lang sendiri sdh menyediakan cara unk menggagalkan unit test menggunakan testing.T
- Terdapat function Fail(), FailNow(), Error(), dan Fatal() jika kita ingin menggagalkan unit test
- Fail() akan menggagalkan unit test, namun tetap menjalankan eksekusi unit test. Namun diakhir ketika selesai, maka unit test tersebut dianggap gagal
- FailNow() akan menggagalkan unit test saat ini juga, tanpa melanjutkan eksekusi unit test
- Error() func lbh seperti melakukan log(print) error, namun setelah melakukan log error, dia akan secara otomatis memanggil function Fail(), sehingga mengakibatkan unit test dianggap gagal. Namun, krn hanya memanggil Fail(), artinya eksekusi unit test akan tetap berjalan sampai selesai.
- Fatal() mirip dgn Error(), hanya saja, setelah melakukan log error, dia akan memanggil FailNow(), sehingga mengakibatkan eksekusi unit test berhenti.
- Disarankan menggunakan Error() & Fatal(), dikarenakan dapat menambahkan knpa unit test kita gagal.

# Assertion
- Melakukan pengecekan di unit test secara manual menggunakan if else sangatlah menyebalkan
- Apalagi jika result data yg harus di cek itu banyak.
- Oleh karena itu, sangat disarankan untuk menggunakan Assertion unk melakukan pengecekan
- Sayangnya, di Go-Lang tdk menyediakan package unk assertion, sehingga kita butuh menambahkan library unk melakukan assertion ini
- Salah satu library Assertion yg paling populer di Go-Lang adlh Testify
- Kita bisa menggunakan library ini unk melakukan assertion terhadap result data di unit test
- Sumber : https://github.com/stretchr/testify
- Kita bisa menambahkannya di Go module kita : go get github.com/stretchr/testify

# Assert vs Require
- Testify menyediakan dua package unk assertion, yaitu assert dan require
- Saat kita menggunakan assert, jika pengecekan gagal, maka assert akan memanggil Fail(), artinya eksekusi unit test akan tetap dilanjutkan
- Sedangkan jika kita menggunakan require, jika pengecekan gagal, maka require akan memanggil FailNow(), artinya eksekusi unit test tidak akan dilanjutkan.

# Skip Test (Membatalkan Test)
- Kadang dalam keadaan tertentu, kita ingin membatalkan eksekusi unit test
- Di Go-Lang juga kita bisa membatalkan eksekusi unit test jika kita mau
- Unk membatalkan unit test kita bisa menggunakan func Skip()

# Before dan After Test
- Biasanya dlm unit test, kadang kita ingin melakukan sesuatu sblm dan setelah sebuah unit test dieksekusi
- Jikalau kode yg kita lakukan sebelum dan setelah selalu sama antar unit test func, maka membuat manual di unit test func nya adlh hal yg membosankan dan terlalu banyak kode duplikat jadinya.
- Untungnya di Go-Lang terdpt fitur yg bernama testing.M
- Fitur ini bernama Main, dimana digunakan unk mengatur eksekusi unit test, namun hal ini juga bisa kita gunakan unk melakukan Before dan After di unit test

# testing.M
- Unk mengatur ekeskusi unit test, kita cukup membuat sebuah func bernama TestMain dgn parameter testing.M
- Jika terdapat func TestMain tersebut, maka secara otomatis Go-Lang akan mengeksekusi func ini tiap kali akan menjalankan unit test di sebuah package
- Dgn ini kita bisa mengatur Before dan After unit test sesuai dgn yg kita mau
- Ingat, func TestMain itu dieksekusi hanya sekali per Go-Lang package, bukan per tiap func unit test

# Sub Test
- Go-Lang mendukung fitur pembuatan func unit test di dalam func unit test
- Fitur ini memang sedikit aneh dan jarang sekali dimiliki di unit test di bahasa pemograman yg lainnya
- Unk membuat sub test, kita bisa menggunakan func Run()

# Menjalankan Hanya Sub Test
- Kita sdh tahu jika ingin menjalankan sebuah unit test func, kita bisa gunakan perintah : go test -run TestNamaFunction
- Jika kita ingin menjelankan hanya salah satu sub test, kita bisa gunakan perintah : go test -run TestNamaFunction/NamaSubTest
- Atau unk semua test semua sub test di semua function, kita bisa gunakan perintah : go test - run /NamaSubTest

# Table Test
- Jika diperhatikan, sebenarnya dgn sub test, kita bisa membuat test secara dinamis
- Dan fitur sub test ini, biasa digunakan oleh programmer Go-Lang unk membuat test dgn konsep table test
- Table test yaitu dimana kita menyediakan data berupa slice yg berisi parameter dan ekspektasi hasil dari unit test
- Lalu slice tersebut kita iterasi menggunakan sub test

# Mock
- Mock adlh object yg sdh kita program dgn ekspektasi tertentu sehingga ketika dipanggil, dia akan menghasilkan data yg sudah kita program diawal
- Mock adlh salah satu teknik dlm unit testing, dimana kita bisa membuat mock object dari suatu object yg memang sulit unk di testing
- Misal kita ingin membuat unit test, namun ternyata ada kode program kita yg harus memanggil API Call ke third party service. Hal ini sangat sulit unk di test, karena unit testing kita hrs selalu memanggil third party service, dan blm tentu response nya sesuai dgn apa yg kita mau
- Pd kasus seperti ini, cocok sekali unk menggunakan mock object

# Testify Mock
- Unk membuat mock object, tdk ada fitur bawaan Go-Lang, namun kita bisa menggunakan library testify.
- Testify mendukung pembuatan mock object, sehingga cocok unk kita gunakan ketika ingin membuat mock object
- Namun, perlu diperhatikan, jika desain kode program kita jelek, akan sulit unk melakukan mocking, jadi pastikan kita melakukan pembuatan desain kode program kita dengan baik

# Contoh Kasus Mock (Aplikasi Querry ke Database)
- Kita akan coba contoh kasus dgn membuat contoh aplikasi golang yg melakukan query ke database
- Dimana kita akan buat layer Service sbg business logic, dan layer Repository sbg jembatan ke database
- Agar kode kita mudah unk di test, disarankan agar membuat kontrak berupa Interface.

# Benchmark
- Selain unit test, Go-Lang testing package juga mendukung melakukan benchmark
- Benchmark adalah mekanisme menghitung kecepatan performa kode aplikasi kita
- Benchmark di Go-Lang dilakukan dgn cara secara otomatis melakukan iterasi kode yg kita panggil berkali-kali sampai waktu tertentu
- Kita tdk perlu menentukan jumlah iterasi dan lamanya, karena itu sdh diatur oleh testing.B bawaan dari testing package

# testing.B
- testing.B adlh struct yg digunakan unk melakukan benchmark
- testing.B mirip dgn testing.T, terdpt func Fail(), FailNow(), Error(), Fatal(), dan lain-lain
- Yg membedakan, ada beberapa attribute dan function tambahan yg digunakan unk melakukan benchmark
- Salah satunya adlh attribute N, ini digunakan unk melakukan total iterasi sebuah benchmark.

# Cara Kerja Benchmark
- Cara kerja benchmark di Go-Lang sangat sederhana
- Gimana kita hanya perlu membuat perulangan sejumlah N attribute
- Nanti secara otomatis Go-Lang akan melakukan eksekusi sejumlah perulangan yg ditentukan secara otomatis, lalu mendeteksi berapa lama proses tersebut berjalan, dan disimpulkan performa benchmark nya dalam waktu

# Benchmark Function
- Mirip seperti unit test, unk benchmark pun, di Go-Lang sudah ditentukan nama func nya, harus diawali dgn kata Benchmark, misal BenchmarkHelloWorld, BenchmarkXxx
- Selain itu, harus memiliki parameter (b *testing.B)
- Dan tdk boleh mengembalikan retur value
- Unk nama file Benchmark, sama seperti unit test, diakhiri dengan _test, misal hello_world_test.go

# Menjalankan Benchmark
- Unk menjalankan seluruh benchmark di module, kita bisa menggunakan perintah sama seperti test, namun ditambahkan parameter bench : go test -v -bench=
- Jika kita hanya ingin menjalankan benchmark tanpa unit test, kita bisa gunakan perintah : go test -v -run=NotMathUnitTest -bench=
- Kode diatas selain menjalankan benchmark, akan menjalankan unit test juga, jika kita hanya ingin menjalankan benchmark tertentu, kita bisa gunakan perintah : go test -v -run=NotMatchUnitTest -bench=BenchmarkTest
- Jika kita menjalankan benchmark di root module dan ingin semua module dijalankan, kita bisa gunakan perintah : go test -v -bench=../...

# Menjalankan Hanya Sub Benchmark
- Saat kita menjalankan benchmark function, maka semua sub benchmark akan berjalan
- Namkun jika kita ingin menjalankan salah satu sub benchmark saja, kita bisa gunakan perintah : go test -v -bench=BenchmarkNama/NamaSub

# Table Bechmark
- Sama seperti di unit test, programmer Go-Lang terbiasa membuat table benchmark
- Ini digunakan agar kita bisa mudah melakukan performance test dengan kombinasi data berbeda-beda tanpa harus membuat banyak benchmark function.