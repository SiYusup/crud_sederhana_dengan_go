# 🚀 Go CRUD Employee Management

Aplikasi manajemen karyawan sederhana berbasis web yang dibangun menggunakan **Go (Golang)**, **MariaDB**, dan **Bootstrap 4**. Proyek ini dirancang untuk mendemonstrasikan implementasi CRUD (Create, Read, Update, Delete) dasar dengan koneksi database MySQL/MariaDB menggunakan Docker.

## ✨ Fitur Utama & Penjelasan

Aplikasi ini mencakup seluruh siklus manajemen data (CRUD) dengan fungsionalitas sebagai berikut:

*   **📋 Manajemen Daftar Karyawan (Read)**
    Menampilkan seluruh data karyawan yang tersimpan di database MariaDB ke dalam tabel yang rapi. Fitur ini menggunakan sistem penomoran otomatis yang dinamis (dimulai dari angka 1) untuk memudahkan identifikasi data.
*   **➕ Penambahan Karyawan Baru (Create)**
    Menyediakan antarmuka form untuk memasukkan informasi karyawan meliputi Nama, NPWP, dan Alamat. Data yang dikirim akan diproses oleh backend Go dan disimpan secara permanen ke database.
*   **✏️ Pembaruan Data Karyawan (Update)**
    Memungkinkan pengguna untuk mengubah informasi karyawan yang sudah ada. Saat tombol edit diklik, sistem akan mengambil data spesifik berdasarkan ID dan menampilkannya kembali ke dalam form untuk dimodifikasi.
*   **🗑️ Penghapusan Data (Delete)**
    Fitur untuk menghapus record karyawan dari database secara permanen. Proses ini dilakukan berdasarkan ID unik karyawan untuk memastikan akurasi data yang dihapus.
*   **📱 Antarmuka Responsif (Modern UI)**
    Dibangun menggunakan **Bootstrap 4**, memastikan tampilan aplikasi tetap optimal dan mudah digunakan baik melalui perangkat desktop maupun perangkat mobile (Responsive Design).
*   **🐳 Infrastruktur Terisolasi (Dockerized DB)**
    Database MariaDB dijalankan di dalam kontainer Docker, sehingga pengguna tidak perlu melakukan instalasi database secara manual di sistem operasi mereka, menjaga lingkungan pengembangan tetap bersih.

## 🛠️ Tech Stack
*   **Backend:** Go 1.22+
*   **Database:** MariaDB (via Docker)
*   **Frontend:** HTML, CSS (Bootstrap 4)
*   **DB Driver:** `github.com/go-sql-driver/mysql`

---

## ⚙️ Cara Menjalankan Proyek

### 1. Persiapan Database (Docker)
Pastikan Docker Desktop Anda sudah berjalan, lalu buka terminal di folder proyek dan jalankan:
```bash
docker-compose up -d
```
> Database MariaDB akan berjalan di `localhost:3307`.

### 2. Inisialisasi Database (DBeaver / HeidiSQL)
Hubungkan aplikasi database manager Anda ke MariaDB Docker, lalu jalankan query berikut untuk membuat tabel:

```sql
CREATE DATABASE IF NOT EXISTS crud_employee;
USE crud_employee;

CREATE TABLE employee (
    id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    npwp VARCHAR(20),
    address TEXT
);
```

### 3. Konfigurasi Go
Pastikan dependensi sudah terinstal:
```bash
go mod tidy
```

### 4. Jalankan Aplikasi
Jalankan perintah berikut untuk memulai server:
```bash
go run main.go
```
Aplikasi akan berjalan di: [http://localhost:8080](http://localhost:8080)

---

## 📂 Struktur Folder
```text
.
├── controller/        # Logika handler untuk setiap route
├── database/          # Inisialisasi koneksi database
├── routes/            # Mapping URL ke controller
├── views/             # File HTML (Templates)
├── docker-compose.yml # Konfigurasi MariaDB Docker
├── main.go            # Entry point aplikasi
└── README.md          # Dokumentasi proyek
```

## 📝 Catatan Penting
- **Port Database:** Proyek ini menggunakan port `3307` untuk menghindari konflik dengan MySQL lokal.
- **Form Edit:** Menggunakan query parameter `?id=` untuk mengidentifikasi data yang diubah.

---
Dibuat oleh [SiYusup](https://github.com/SiYusup) untuk keperluan belajar pemrograman Go.
