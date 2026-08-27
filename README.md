# Tandur VAC API (Virtual Assessment Center Backend)

[![Go Version](https://img.shields.io/badge/Go-1.19+-00ADD8?style=for-the-badge&logo=go)](https://golang.org/)
[![Gin Framework](https://img.shields.io/badge/Framework-Gin-008ECF?style=for-the-badge&logo=gin)](https://gin-gonic.com/)
[![Database](https://img.shields.io/badge/Database-MySQL%20%2F%20MariaDB-4479A1?style=for-the-badge&logo=mysql)](https://www.mysql.com/)
[![Cache](https://img.shields.io/badge/Cache-Redis-DC382D?style=for-the-badge&logo=redis)](https://redis.io/)
[![Architecture](https://img.shields.io/badge/Architecture-Hexagonal%20%2F%20Ports%20%26%20Adapters-green?style=for-the-badge)](#arsitektur-sistem)

Backend REST API untuk sistem **Tandur Virtual Assessment Center (VAC)**. Sistem ini dirancang untuk memfasilitasi pelaksanaan asesmen kompetensi online secara komprehensif, multi-instrumen, dan real-time bagi asesi (kandidat) dan asesor.

---

## Daftar Isi
- [Gambaran Umum](#gambaran-umum)
- [Arsitektur Sistem](#arsitektur-sistem)
- [Modul Asesmen](#modul-asesmen)
- [Katalog Endpoint REST API](#katalog-endpoint-rest-api)
- [Struktur Database & View Mapping](#struktur-database--view-mapping)
- [Variabel Lingkungan (Environment)](#variabel-lingkungan-environment)
- [Panduan Instalasi & Menjalankan](#panduan-instalasi--menjalankan)
- [Optimasi & System Tuning](#optimasi--system-tuning)

---

## Gambaran Umum

**Tandur VAC API** mengelola seluruh siklus hidup asesmen daring (*online assessment session*), mulai dari autentikasi token asesi, penyajian instruksi dan soal, pelaksanaan simulasi inbox / email, studi kasus problem analysis, diskusi kelompok tanpa pemimpin (LGD), roleplay simulasi, wawancara 1-on-1, pengisian CV digital, hingga pengumpulan hasil tes dan feedback.

---

## Arsitektur Sistem

Proyek ini dibangun menggunakan pola **Hexagonal Architecture (Ports and Adapters / Clean Architecture)** untuk memisahkan domain bisnis dari framework dan database:

\\\mermaid
graph TD
    Client[Client / Frontend Web] -->|HTTP Requests| Router[Gin REST Controller Inbound Adapter]
    Router -->|Input Ports| Service[Application Service Layer]
    Service -->|Business Logic| Domain[Domain Entities & Rules]
    Service -->|Output Ports| AdapterOut[Outbound Adapters]
    AdapterOut -->|SQL Queries| MySQL[(MySQL / MariaDB Database)]
    AdapterOut -->|Cache & Pub/Sub| Redis[(Redis Cache)]
    AdapterOut -->|File Upload| BunnyCDN[(BunnyCDN Cloud Storage)]
\\\

### Struktur Direktori:
\\\	ext
tandur-vac-api-main/
├── docker-compose.yaml             # Layanan MariaDB & Redis untuk development
├── environment.local               # Contoh environment variable lokal
├── .env.example                    # Template environment variables
├── go.mod / go.sum                 # Definisi dependency Go
├── main.go                         # Entry point, router setup, connection pool, graceful shutdown
├── src/
│   ├── domain/                     # Entity bisnis murni & data contract
│   │   ├── assessor.go
│   │   ├── cv.go
│   │   ├── feedback.go
│   │   ├── inbasket.go
│   │   ├── lgd.go
│   │   ├── mailbox.go
│   │   ├── one_on_one.go
│   │   ├── problem_analysis.go
│   │   ├── qna.go
│   │   ├── roleplay.go
│   │   ├── subtes.go
│   │   └── user.go
│   ├── app/
│   │   ├── use_case/               # Interface port use case
│   │   └── service/                # Implementasi use case & business logic
│   ├── adapter/
│   │   ├── in/rest/                # Inbound HTTP REST Controllers (Gin)
│   │   └── out/
│   │       ├── mysql/              # Secondary Adapters: Repositories & Entities MySQL
│   │       └── redis/              # Secondary Adapters: Redis notification & cache
│   └── util/                       # Koneksi Database pool, Redis, Logger, Response helpers
└── logs/                           # Direktori log server harian
\\\

---

## Modul Asesmen

1. **User / Auth (/user)**: Mengambil profil dan data autentikasi asesi berdasarkan token asesmen.
2. **Subtes Management (/subtes)**: Mengelola jadwal subtes, status progres (NOT_STARTED, IN_PROGRESS, DONE), dan submit hasil.
3. **In-Basket / Inbox Simulation (/inbasket)**: Simulasi inbox email manajerial. Asesi menerima email/event terjadwal, membaca, membuat draft, membalas email (dengan CC dan lampiran), serta membuat email baru.
4. **Problem Analysis / PA (/pa)**: Asesmen analisis masalah bisnis berupa studi kasus perusahaan, dokumen profil, instruksi, dan pertanyaan analisis.
5. **Leaderless Group Discussion / LGD (/lgd)**: Diskusi kelompok simulasi tanpa pemimpin dengan link video conference dan daftar asesor pengamat.
6. **One-on-One Interview / 1on1 (/1on1)**: Wawancara terstruktur antara asesi dan asesor.
7. **Roleplay Simulation (/roleplay)**: Simulasi bermain peran interaktif antara asesi dengan role player / asesor.
8. **Question & Answer / QnA (/qna)**: Soal esai atau kuesioner terstruktur dengan penyimpanan jawaban dinamis.
9. **Curriculum Vitae / CV (/cv)**: Pengisian data portofolio dan riwayat hidup asesi (13 halaman data).
10. **Feedback (/feedback)**: Survei kepuasan dan evaluasi pelaksanaan asesmen oleh asesi.
11. **Storage Upload (/upload)**: Endpoint unggah berkas langsung ke BunnyCDN storage.

---

## Katalog Endpoint REST API

### 1. Root & Health Check
| Method | Endpoint | Deskripsi |
|---|---|---|
| \GET\ | \/\ | Cek status server aktif |
| \GET\ | \/health\ | Health check (status koneksi database, uptime, timestamp) |
| \GET\ | \/benchmark\ | Pengujian throughput transfer file |

### 2. User
| Method | Endpoint | Deskripsi |
|---|---|---|
| \GET\ | \/user/token/:token\ | Mengambil informasi akun asesi berdasarkan token |

### 3. Subtes
| Method | Endpoint | Deskripsi |
|---|---|---|
| \GET\ | \/subtes/token/:token\ | Mengambil seluruh daftar subtes dan jadwal asesi |
| \GET\ | \/subtes/result/:id\ | Mengambil hasil subtes berdasarkan ID |
| \PATCH\ | \/subtes/status\ | Mengubah status pengerjaan subtes (\id\, \status\) |
| \POST\ | \/subtes/submit\ | Mengirimkan hasil pengerjaan subtes (\id\, \esult\) |

### 4. In-Basket (Simulasi Email)
| Method | Endpoint | Deskripsi |
|---|---|---|
| \GET\ | \/inbasket/token/:token\ | Mengambil berkas instruksi, event, dan seluruh email inbox |
| \GET\ | \/inbasket/mailbox/:id\ | Mengambil detail email berdasarkan ID |
| \GET\ | \/inbasket/mailboxdraft/:id\| Mengambil draf balasan untuk email tertentu |
| \POST\ | \/inbasket/reply\ | Mengirim balasan email (disertai CC, attachment, body) |
| \POST\ | \/inbasket/replydraft\ | Menyimpan draf balasan email |
| \POST\ | \/inbasket/compose\ | Membuat email baru dari asesi |
| \PATCH\| \/inbasket/updatedraft\| Memperbarui isi draf email |
| \PATCH\| \/inbasket/status\ | Mengubah status email (\UNREAD\, \READ\, \DRAFT\, \SENT\, \REPLIED\) |
| \DELETE\| \/inbasket/:id\ | Menghapus email/draf |

### 5. Problem Analysis (PA)
| Method | Endpoint | Deskripsi |
|---|---|---|
| \GET\ | \/pa/token/:token\ | Mengambil studi kasus, instruksi, dan soal problem analysis |

### 6. Leaderless Group Discussion (LGD)
| Method | Endpoint | Deskripsi |
|---|---|---|
| \GET\ | \/lgd/token/:token\ | Mengambil instruksi diskusi, tautan meeting, dan data asesor |

### 7. One-on-One Interview
| Method | Endpoint | Deskripsi |
|---|---|---|
| \GET\ | \/1on1/token/:token\ | Mengambil jadwal, tautan video call, dan asesor 1-on-1 |

### 8. Roleplay
| Method | Endpoint | Deskripsi |
|---|---|---|
| \GET\ | \/roleplay/token/:token\ | Mengambil skenario roleplay, instruksi peserta, dan meeting URL |

### 9. Question & Answer (QnA)
| Method | Endpoint | Deskripsi |
|---|---|---|
| \GET\ | \/qna/token/:token\ | Mengambil daftar master soal QnA |
| \GET\ | \/qna/id/:id\ | Mengambil detail set soal QnA |
| \POST\ | \/qna/result\ | Menyimpan/memperbarui satu jawaban (\Upsert\) |
| \POST\ | \/qna/results\ | Menyimpan/memperbarui banyak jawaban sekaligus |
| \GET\ | \/qna/results/token/:token\| Mengambil seluruh jawaban yang tersimpan |

### 10. Curriculum Vitae (CV)
| Method | Endpoint | Deskripsi |
|---|---|---|
| \GET\ | \/cv/token/:token\ | Mengambil data CV 13 halaman yang sudah tersimpan |
| \POST\ | \/cv/submit\ | Menyimpan / memperbarui data CV |

### 11. Feedback
| Method | Endpoint | Deskripsi |
|---|---|---|
| \POST\ | \/feedback/submit\ | Mengirim kuesioner feedback asesi |

### 12. Upload
| Method | Endpoint | Deskripsi |
|---|---|---|
| \POST\ | \/upload\ | Mengunggah file (multipart/form-data) ke CDN storage |

---

## Struktur Database & View Mapping

Sistem memanfaatkan View MySQL yang sudah dioptimasi untuk agregasi data:

| Entity Domain | View / Tabel Database Utama |
|---|---|
| \Subtes\ | \V_GET_SUBTES\, \	rx_ap_class_management_detail\ |
| \Inbasket File\ | \V_GET_INBASKET_FILE\ |
| \Inbasket Event\ | \V_GET_INBASKET_EVENT\ |
| \Inbasket Email\ | \V_GET_INBASKET_EMAIL\ |
| \Inbasket Mailbox\ | \	rx_ac_inbasket_mailbox\ |
| \Problem Analysis\| \V_GET_PA\ |
| \LGD\ | \V_GET_LGD\ |
| \One-on-One\ | \V_GET_1ON1\ |
| \Roleplay\ | \V_GET_ROLEPLAY\ |
| \QnA\ | \V_GET_QNA_MASTER\, \	rx_ac_question_answer_detail\, \	rx_ac_qna_result\ |
| \CV\ | \V_GET_CV\, \	rx_ac_cv\ |
| \Feedback\ | \	rx_ac_feedback\ |
| \User / Asesi\ | \V_GET_LOGIN\ |
| \Assessor\ | \V_GET_AP_USER\ |

---

## Variabel Lingkungan (Environment)

Buat file \.env\ atau set environment variable berikut:

\\\ash
# Port Aplikasi (Default: 3030)
PORT=3030

# Database MySQL / MariaDB
DB_HOST=127.0.0.1:3306
DB_USER=devtandurusr
DB_PASSWORD=D3vp4_T4ndur#
DB_NAME=dev_tandur

# Cache Redis
CACHE_HOST=127.0.0.1
CACHE_PORT=6379
CACHE_PASSWORD=

# CDN Storage (BunnyCDN)
CDN_API_KEY=your_bunnycdn_access_key
\\\

---

## Panduan Instalasi & Menjalankan

### 1. Menjalankan Database & Redis dengan Docker
\\\ash
docker compose up -d
\\\

### 2. Menjalankan Server secara Lokal
\\\ash
# Download dependency
go mod download

# Jalankan server
go run main.go
\\\

### 3. Build Binary Production
\\\ash
# Build binary Linux/macOS
go build -o tandur-vac-api main.go

# Build binary Windows
go build -o tandur-vac-api.exe main.go
\\\

---

## Optimasi & System Tuning

Sistem ini telah melewati audit menyeluruh dan peningkatan performa:

1. **Singleton Database Connection Pooling**:
   - Menghapus pembukaan/penutupan koneksi database per-request yang sebelumnya membebani I/O dan memicu *race condition* / sql: database is closed.
   - Mengatur parameter pool optimal: MaxOpenConns(25), MaxIdleConns(10), ConnMaxLifetime(5m), ConnMaxIdleTime(1m).
2. **Eliminasi Connection Leaks**:
   - Seluruh loop db.Query kini dilengkapi defer results.Close() untuk memastikan cursor koneksi segera dikembalikan ke pool.
   - Mengganti operasi UPDATE/DELETE dari db.Query menjadi db.Exec.
3. **Parameterized Queries (Keamanan & Stabilitas)**:
   - Seluruh string concatenation SQL diganti dengan placeholder ?. Ini mencegah crash saat user menginput teks dengan tanda petik satu ' dan menutup celah SQL Injection.
4. **Perbaikan Logika Inbasket & Routing**:
   - Memperbaiki bug parsing ID pada DELETE /inbasket/:id.
   - Memperbaiki boolean logic query pencarian subjek/body email.
5. **Thread-Safe Logging & Graceful Shutdown**:
   - Direktori logs/ dibuat otomatis.
   - Server menangani sinyal SIGINT dan SIGTERM dengan srv.Shutdown() untuk menyelesaikan request yang sedang berjalan sebelum aplikasi berhenti.
