<?php

	require 'Queue.php';
	require 'misc.php';
	
	//1. mengisi awal $Q dengan deretan URL valid yang SEHARUSNYA ditelusuri
	//   URL yang harus ditelusuri pada intinya adalah pagination di bagian bawah
	//   masalahnya kita tidak bisa menebak ada berapa halaman URL page
	//   maka sambil increment, sambil cek apakah alamat valid atau tidak
	//   cara melakukan cek VALID atau tidaknya URL bisa dengan:
	//   a. alamat masih bisa diakses atau tidak
    //      masalahnya sudah dicoba dengan mengganti page menjadi 3000 pun masih bisa diakses
	//   b. cek apa yang beda
	//      yang beda adalah ada tulisan "Publication Not Found"  
	
	$masih_valid = true;
	$c = 1;
	$Q = new Queue();
	
	while ($masih_valid){
		$alamat = "https://sinta.kemdikbud.go.id/affiliations/profile/2136?page=" . $c . "&view=scopus";
		
		$halaman = fetch($alamat);
			
		if (strpos($halaman, 'Publication Not Found') !== false) {
			//mengandung kata tersebut, maka stop menambah halaman
			$masih_valid = false;
		} else {
			$Q->enqueue($alamat);
			$c = $c + 1;
			echo $alamat . "<br>";
		}
	}
	
	// tapi proses di atas luamaaaaaaaa
	// maka opsi lain adalah di manual langsung, 
	// > cek ada berapa halaman
	// > langsung daftarkan sejumlah itu
	
	
	//2. setelah itu langsung proses dari daftar $Q kita telusuri langsung
	
	while (!$Q->isEmpty()) {
		$u = $Q->dequeue();   //dapatkan sebuah URL dari Q
		$du = fetch($u);      //ambil teks HTML-nya
		
		if (trim($du)!=""){   //kalau dokumen HTML tersebut tidak kosong
			storeD($du, $u);  //simpan ke dalam D
		
			$L = array();
			$L = extractURL($u, $du);  //ekstrak semua href "bersih" dari d(u)
			
			foreach ($L as $v) {
				storeE($u, $v);
				
				if (!$Q->contains($v) && !containsD($v)) {
					$Q->enqueue($v);
				}
			}	
		}		
	}