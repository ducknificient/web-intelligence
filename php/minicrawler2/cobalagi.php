<?php


	// URL halaman web yang ingin Anda ambil innerHTML-nya
	$url = "https://sinta.kemdikbud.go.id/affiliations/profile/2136?page=200&view=scopus";

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

	// Mencari semua elemen <h5>
	$elements = $xpath->query('//h5');

	// Menampilkan innerHTML dari semua elemen <h5> yang ditemukan
	foreach ($elements as $element) {
		$innerHTML = $dom->saveHTML($element);
		echo "InnerHTML dari elemen <h5>: " . $innerHTML . "<br>";
	}