package mysql

import "github.com/stryn26/MiniProjectEvermosRakamin/internal/daos"

var userSeed = []daos.User{
	{
		Nama:          "John Doe",
		Kata_Sandi:    "password1",
		Notelp:        "08123456789",
		Tanggal_lahir: "1990-01-01",
		Jenis_kelamin: "Laki-laki",
		Tentang:       "Seorang pengguna aktif",
		Pekerjaan:     "Karyawan",
		Email:         "johndoe@example.com",
		Id_provinsi:   "1",
		Id_kota:       "1",
		IsAdmin:       false,
	},
	{
		Nama:          "Jane Smith",
		Kata_Sandi:    "password2",
		Notelp:        "08987654321",
		Tanggal_lahir: "1995-02-15",
		Jenis_kelamin: "Perempuan",
		Tentang:       "Seorang pengguna dengan minat fashion",
		Pekerjaan:     "Freelancer",
		Email:         "janesmith@example.com",
		Id_provinsi:   "2",
		Id_kota:       "2",
		IsAdmin:       true,
	},
}