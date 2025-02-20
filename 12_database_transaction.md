# Database Transaction

- Salah satu fitur andalan di databse adalah transaction
- Materi databse transaction sudah saya bahas dengan tuntas di materi MySQL database, jadi silahkan pelajari di course tersebut
- Di course ini kita akan fokus bagaimana menggunakan database transaction di Go


## Transaction di Golang

- Secara default, semua perintah SQL yang kita kirim menggunakan Go akan otomatis di-commit, atau istilahnya auto commit
- Namun kita bisa menggunakan fitur transaksi sehingga SQL yang kita kirim tidak secara otomatis di-commit ke database
- Untuk memulai transaksi, kita bisa menggunakan function (DB) Begin(), dimana akan menghasilkan struct Tx yang merupakan representasi Transaction
- Struct Tx ini yang kita gunakan sebagai pengganti DB untuk melakukan transaksi, dimana hampir semua function di DB ada di Tx, seperti Exec, Query, atau Prepare
- Setelah selesai proses transaksi, kita bisa gunakan function (Tx) Commit() untuk melakukan commit atau Rollback()
