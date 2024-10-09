<?php

	require 'Queue.php';
	require 'misc.php';
	
	//$S0 = "https://www.azlyrics.com/lyrics/aaroncarter/iwillbeyours.html";
	
	$Q = new Queue();
		
	for ($i=1; $i<=12; $i++){
		$url = "https://sinta.kemdikbud.go.id/affiliations/profile/2136?page=" . $i . "&view=scopus";		
		$Q->enqueue($url);
	}
	
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