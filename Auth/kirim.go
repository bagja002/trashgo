package Auth

import(
	"log"
	"net/smtp"
)

func Kirim(body, subject string) {
	// AKUN GMAIL PENGIRIM - allow less secure apps : ON (google account setting)
	surel_pengirim := "barjafaskan9@gmail.com"
	kata_sandi := "Barisan123"

	// PENERIMA SUREL
	penerima := "barjafaskan04@gmail.com"

	pesan := "From: " + surel_pengirim + "\n" +
		"To: " + penerima + "\n" +
		"Subject: " + subject + "\n" +
		body

	// FUNGSI UNTUK MENGIRIM EMAIL MELALUI SMTP
	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", surel_pengirim, kata_sandi, "smtp.gmail.com"),
		surel_pengirim, []string{penerima}, []byte(pesan),
	)

	// GALAT PROGRAM AKAN TERCETAK JIKA ADA
	if err != nil {
		log.Print(err)
	}

	log.Print("Surel berahsil terkirim!")
}

func main() {
	// FUNGSI INI MENERIMA DUA PARAMETER
	Kirim("Hello, world!", "Sapaan")
	// PARAMETER PERTAMA BERUPA ISI DARI SUREL
	// PARAMETER KEDUA ADALAH SUBJEK SUREL
}