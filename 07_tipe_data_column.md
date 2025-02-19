# Tipe Data Column

- Sebelumnya kita hanya membuat table dengan tipe data di kolom nya berupa VARCHAR
- Untuk VARCHAR di database, biasanya kita gunakan String di Go
- Bagaimana dengan tipe data yang lain?
- Apa representasinya di Go, misal tipe data timestamp, date, dan lain-lain


## Mapping Tipe Data

| Tipe Data Database              | Tipe Dat Golang  |
|---------------------------------|------------------|
| VARCHAR, CHAR                   | string           |
| INT, BIGINT                     | int32, int64     |
| FLOAT, DOUBLE                   | float32, float64 |
| BOOLEAN                         | bool             |
| DATE, DATETIME, TIME, TIMESTAMP | time.Time        |


## Error Tipe Data Date

- Secara default, Driver MySQL untuk Go akan melakukan query tipe data DATE, DATETIME, TIMESTAMP menjadi []byte/[]uint8. Dimana ini bisa dikonversi menjadi String, lalu diparsing menjadi time.Time
- Namun hal ini merepotkan jika dilakukan manual, kita bisa meminta Drive MySQL untuk Go secara otomatis melakukan parsing dengan menambahkan parameter parseTime=true pada dataSource ketika sql.Open('mysql', dataSource)


## Nullable Type

- Go database tidak mengerti dengan tipe data NULL di database
- Oleh karena itu, khusus untuk kolom yang bisa NULL di database, akan jadi masalah jika kita melakukan Scan secara bulat-bulat menggunakan tipe data representative di Go


## Error Data Null

- Konversi secara otomatis NULL tidak didukung oleh Driver MySQL di Go
- Oleh karena itu, khusus tipe kolom yang bisa NULL, kita perlu menggunakan tipe data yang ada dalam package sql

## Tipe Data Nullable

| Tipe Data Go | Tipe Data Nullable       |
|--------------|--------------------------|
| string       | database/sql.NullString  |
| bool         | database/sql.NullBool    |
| float64      | database/sql.NullFloat64 |
| int32        | database/sql.NullInt32   |
| int64        | database/sql.NullInt64   |
| time.Time    | database/sql.NullTime    |
