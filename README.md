# 🚀 Go CRUD Sederhana

Aplikasi CRUD (Create, Read, Update, Delete) sederhana menggunakan bahasa pemrograman **Go** dan database **MariaDB/MySQL** yang dijalankan melalui Docker.

## 🛠️ Tech Stack

- **Language:** [Go (Golang)](https://go.dev/)
- **Database:** [MariaDB](https://mariadb.org/) (via Docker)
- **Management Tool:** [DBeaver](https://dbeaver.io/)
- **Containerization:** [Docker Desktop](https://www.docker.com/)

## 📋 Fitur

- [ ] Koneksi Database MySQL/MariaDB
- [ ] Create: Menambah data Employee
- [ ] Read: Menampilkan daftar Employee
- [ ] Update: Mengubah data Employee
- [ ] Delete: Menghapus data Employee

## ⚙️ Persiapan Instalasi

### 1. Prasyarat

Pastikan Anda sudah menginstal:

- Go (versi 1.22 atau terbaru)
- Docker Desktop
- DBeaver (untuk manajemen database)

### 2. Jalankan Database (Docker)

Proyek ini menggunakan Docker untuk menjalankan database agar tidak mengganggu sistem lokal Anda.

```bash
docker-compose up -d
```

_Database akan berjalan di port **3307**._

### 3. Konfigurasi Database (DBeaver)

Hubungkan DBeaver ke database dengan detail berikut:

- **Host:** `localhost`
- **Port:** `3307`
- **Database:** `crud_go`
- **Username:** `user_go`
- **Password:** `password_go`

Gunakan query di file `database.db` untuk membuat tabel `employee`.

## 🚀 Cara Menjalankan Aplikasi

Setelah database siap, jalankan aplikasi Go Anda:

```bash
go run main.go
```

## 📂 Struktur Proyek

```text
.
├── mysql_data/        # Data database (diabaikan oleh Git)
├── database.db        # Query SQL untuk schema database
├── docker-compose.yml # Konfigurasi database MariaDB
├── go.mod             # Go module configuration
├── main.go            # Entry point aplikasi
└── .gitignore         # File yang diabaikan oleh Git
```

---

Dibuat dengan ❤️ untuk belajar Go.
