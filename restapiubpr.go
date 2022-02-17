package main

import (
	"encoding/json"
	"net/http"
)

type Karyawan struct {
	Nik		    int    `json:"nik"`
	Nama        string `json:"nama"`
	Divisi	 	string `json:"divisi"`
	Jabatan     string  `json:"jabatan"`
	Cabang	  	string `jeson:"cabang"` 
}

var data = []Karyawan{
		{
			Nik:      001,
			Nama:    "MichaelDevenBahari",
			Divisi:  "DigitalBanking",
			Jabatan: "ITprogammer",
			Cabang:  "Bintaro",
		},
		{
			Nik:      002,
			Nama:    "Kezia",
			Divisi:  "DigitalBanking",
			Jabatan: "ITSupprot",
			Cabang:  "BSD",
		},
		{
			Nik:      003,
			Nama:    "EvelineMary",
			Divisi:  "DigitalBanking",
			Jabatan: "Head",
			Cabang:  "Bintaro",
		},

		{
			Nik:      004,
			Nama:    "Dimas",
			Divisi:  "DigitalBanking",
			Jabatan: "Komisaris",
			Cabang:  "Bogor",
		},
	}
	
func main() {
	http.HandleFunc("/karyawan", GetKaryawan)
	http.HandleFunc("/karyawan-id", GetKaryawanByNama)

	http.ListenAndServe(":8000", nil)
}

func GetKaryawan(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	if r.Method == "POST" {
		res, err := json.Marshal(data)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(res)
		return
	}

	http.Error(w, "", http.StatusBadRequest)
}

func GetKaryawanByNama(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	if r.Method == "GET" {
		Nama := r.FormValue("karyawan")

		for _, v := range data {
			if v.Nama == Nama {
				result, err := json.Marshal(v)

				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}

				w.Write(result)
				return
			}
		}
	}

	http.Error(w, "400 status bad request", http.StatusBadRequest)
}
