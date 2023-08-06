# Go Starter API

Go Starter API adalah proyek awal (starter) untuk membuat RESTful API menggunakan bahasa pemrograman Go. Proyek ini mencakup fitur-fitur dasar seperti registrasi, login, manajemen produk, dan manajemen pengguna. Proyek ini juga dilengkapi dengan otomatisasi reload (hot-reload) untuk memudahkan pengembangan.

## Fitur

- Registrasi pengguna
- Login pengguna dengan token JWT
- Manajemen produk (CRUD)
- Manajemen pengguna (CRUD)

## Instalasi

1. Pastikan Anda telah menginstal Go dan mengatur GOPATH dengan benar.

2. Unduh proyek ini menggunakan perintah berikut: `go get -u github.com/username/go-starter-api`

3. Pindah ke direktori proyek: `cd $GOPATH/src/github.com/username/go-starter-api`

4. Instal dependensi proyek: `go mod tidy`

## Penggunaan

1. Jalankan `cp .env.example .env`

2. Jalankan key:generate untuk membuat SECRET_KEY: `go run main.go key:generate`

3. Buat file .env dan konfigurasikan koneksi ke database MySQL:
    - DB_USER=username
    - DB_PASS=password
    - DB_HOST=localhost
    - DB_PORT=3306
    - DB_NAME=database_name
    - SECRET_KEY=your_secret_key

4. Jalankan migrasi untuk membuat tabel pada database: `go run main.go migrate`

5. Jalankan seed untuk menambahkan data awal ke database: `go run main.go db:seed`

6. Jalankan serve untuk menjalankan aplikasi: `go run main.go serve`

7. Proyek ini akan dijalankan pada http://localhost:8080

## Endpoint API

### Auth (User)

- `POST /api/auth/register` - Registrasi pengguna baru
- `POST /api/auth/login` - Login pengguna
- `POST /api/auth/logout` - Logout pengguna

### Produk (Product)

- `GET /api/products` - Mendapatkan daftar produk
- `GET /api/products/:id` - Mendapatkan detail produk berdasarkan ID
- `POST /api/products` - Menambahkan produk baru
- `PUT /api/products/:id` - Mengupdate produk berdasarkan ID
- `DELETE /api/products/:id` - Menghapus produk berdasarkan ID

### User (User)

- `GET /api/users` - Mendapatkan daftar user
- `GET /api/users/:id` - Mendapatkan detail user berdasarkan ID
- `POST /api/users` - Menambahkan user baru
- `PUT /api/users/:id` - Mengupdate user berdasarkan ID
- `DELETE /api/users/:id` - Menghapus user berdasarkan ID

## Kontribusi

Jika Anda ingin berkontribusi pada proyek ini, silakan fork proyek ini dan buat pull request dengan perubahan Anda.

## Author

Mochammad Hairullah


## Lisensi

Proyek ini dilisensikan di bawah [MIT License](LICENSE).
