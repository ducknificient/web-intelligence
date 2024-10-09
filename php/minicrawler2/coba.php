<?php

	// URL halaman web yang ingin Anda ambil innerHTML-nya
	$url = "https://sinta.kemdikbud.go.id/affiliations/profile/2136?page=1&view=scopus";

	// Mendapatkan konten HTML dari URL
	$html = file_get_contents($url);

	// Membuat objek DOMDocument
	$dom = new DOMDocument();

	// Mematikan error dan warning karena HTML tidak selalu sempurna
	libxml_use_internal_errors(true);

	// Memuat konten HTML ke dalam objek DOMDocument
	$dom->loadHTML($html);

	// Membuat objek DOMXPath
	$xpath = new DOMXPath($dom);

	// Mencari semua elemen dengan class "ar-title"
	$elements = $xpath->query('//div[contains(@class, "ar-title")]');

	// Menyimpan innerHTML dari elemen yang pertama ditemukan dengan class "ar-title"
	if ($elements->length > 0) {
		foreach ($elements as $element){
			$innerHTML = $dom->saveHTML($element);
			echo $innerHTML . "<br>";
		}				
	} else {
		echo "Tidak ada elemen dengan class 'ar-title' ditemukan.";
	}