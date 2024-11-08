#### Catatan

- Pastikan sudah menginstall NodeJS (referensi: https://nodejs.org/en)
- Pastikan sudah menginstall Go (referensi: https://go.dev/doc/install)

---

#### Menjalankan Front End

1. Jalankan `npm install`
2. Jalankan `npm run serve`

---

#### Menjalankan Back End

1. Masuk ke folder backend
2. Jalankan `npm install`
3. Jalankan `npm run serve`
4. Tambahkan di Collection `virtual-machines` data VM

```bash
{
  "ipPublic": "127.0.0.1",
  "ipPrivate": "127.0.0.1",
  "currentUsage": 0
}
```

---

#### Menjalankan WhatsApp Engine

1. Salin `_id` dari hasil simpan data di `virtual-machines`
2. Masuk ke folder engine
3. Buka file `.env`
4. Ubah value dari key ID_VM sesuai dengan `_id` tadi
5. Jalankan `go run .`
