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
			
		// Membuat objek DOMDocument
		$dom = new DOMDocument();

		// Mematikan error dan warning karena HTML tidak selalu sempurna
		libxml_use_internal_errors(true);

		// Memuat konten HTML ke dalam objek DOMDocument
		$dom->loadHTML($halaman);

		// Membuat objek DOMXPath
		$xpath = new DOMXPath($dom);

		// Mencari semua elemen dengan class "ar-title"
		$elements = $xpath->query('//div[contains(@class, "ar-title")]');
	
		if ($elements->length > 0) {
			$Q->enqueue($alamat);
			$c = $c + 1;
			echo $alamat . "<br>";
		}
	}
	
	// masalahnya....sampai nilai C di ganti 2000 sekalipun, 
	// masih ada artikel yang secara random dia munculkan, maka...cek database saja
	// sudah terdaftar belum artikel tersebut
	
	
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