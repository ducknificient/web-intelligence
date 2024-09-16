
select 
--	d.task,u,du  
	count(1) as total
from
webintelligence.tabled d
where 
--	d.du like '%accepted-answer-indicator%'
	d.u like 'https://chinese.stackexchange.com/questions%';

;


select d.task,u,du from webintelligence.tabled d
order by d.created DESC
;

select * from webintelligence.tablee e
order by e.created desc;

;


SELECT
	d.task,d.u,d.du
--	COUNT(1) as total 
FROM 
	webintelligence.tabled d
WHERE
1=1 
and task = 'SEHAT NEGERIKU'
and u like '%/baca/%'
and u not like '%/#comment%'
and u not like '%baca/author%'
--and du not like '%Just a moment%';
--order by d.created DESC
order by u desc
;

select 
	e.u, e.v, e.*
--count(1) as total 
from webintelligence.tablee e
where 
1=1
--and task = 'AYOSEHAT-XML'
order by e.created desc
;

;

SELECT
	d.task,d.u,d.du
--	COUNT(1) as TOTAL
FROM 
	webintelligence.tabled d
WHERE
1=1 
and task in ('AYOSEHAT-XML', 'AYOSEHAT')
--and task = 'AYOSEHAT-XML'
--and task = 'AYOSEHAT'
--and du not like '%Just a moment%';
--order by d.created DESC
--order by u desc

;

SELECT
--	d.task,
	COUNT(1) as TOTAL
--	d.u,
--	d.du,
--	substring(d.du FROM '<meta name="keywords" content="(.*?)">') AS category	
FROM 
	webintelligence.tabled d
WHERE
1=1 
and task = 'KEMENDAG'
--order by d.created DESC

;

select count(1) from (
	select 
		distinct d.u
	FROM 
	webintelligence.tabled d
	WHERE
	1=1 
	and task = 'AYOSEHAT-XML'
)

;



select 
	d.u
FROM 
webintelligence.tabled d
WHERE
1=1 
and task = 'AYOSEHAT-XML'

;

select e.* from webintelligence.tablee e 
order by e.created desc

;

SELECT COUNT(1) FROM webintelligence.tabled cp

;

SELECT COUNT(1) FROM webintelligence.tabled cp

;

SELECT 
	cp.du, 
	cp.u,
	cp.task
	FROM webintelligence.tabled cp
	ORDER BY cp.created DESC

;

;
select 
cp.task,
cp.pagesource,
cp.link,
cp.document
from webintelligence.crawlpage cp
order by cp.created DESC

;

SELECT 
cp.du, cp.u, cp.task
FROM webintelligence.CRAWLPAGE  cp
WHERE 
--	du ilike '%%'
--cp.u like '%https://www.kemendag.go.id/public/files/2019/07/04/sembilan-perusahaan-indonesia-pasok-sarang-burung-walet-ke-china-id0-1562217975.pdf%'
1=1
--and task = 'KEMENDAG'
--and cp.u like '%-walet-ke-china%'
--and cp.u ilike '%public/files%'
--ORDER BY cp.created DESC

--LIMIT 10 OFFSET 1
ORDER BY cp.created DESC 

;

select cp.*
from
webintelligence.crawlpage cp
where 
1=1 
--and link like '%.pdf%'
order by cp.created desc

;



select 
	cp.*
from 
webintelligence.crawlpage cp 
where 
1=1
and task <> ''
order by cp.created desc
;

update webintelligence.crawlhref  set task = 'KEMENDAG'
--
;

;

--select count(1) from (
;

select 
--count (1) 
*
from (

select 
	ch.href,
	min(cp.created) as timestamp
from 
	webintelligence.crawlhref ch
inner join webintelligence.crawlpage cp on cp.link = ch.link 
where
1=1
and ch.href not in (
	select cp.link from webintelligence.crawlpage cp 
)
group by ch.href
order by timestamp asc 
)



--order by ch.href desc


--)
;

select count(1) from webintelligence.crawlhref ch;
select * from webintelligence.crawlhref ch;

select count(1) from webintelligence.crawlpage cp;


select link from webintelligence.crawlhref ch
order by ch.created desc 
limit 1;

;

SELECT 
cp.pagesource, cp.link, cp.task,
cp.documenttype,cp.mimetype,cp.document
FROM webintelligence.crawlpage cp
WHERE 
1=1
AND cp.task = 'JATIMPROV'
ORDER BY cp.created DESC 


;

SELECT 
	ch.*
FROM webintelligence.crawlhref ch
WHERE ch.link ='https://jatimprov.go.id/berita?berita=setda-opt&page=2'
ORDER BY ch.created DESC
;

SELECT cp.documenttype,cp.* FROM webintelligence.crawlpage cp
WHERE cp.task ='JATIMPROV'
AND cp.documenttype IS NULL
AND link LIKE 'https://jatimprov%'
ORDER BY cp.created DESC  
;

SELECT 
count(1)
FROM webintelligence.crawlpage c 
WHERE task ='JATIMPROV'
AND DOCUMENTtype IS null
AND link LIKE 'https://jatimprov%'
;

SELECT 
count(1)
FROM webintelligence.crawlhref c 
WHERE task ='JATIMPROV'

;

SELECT 
--count(1) AS total
link,
count(sentence) AS total
FROM (
	SELECT 
		regexp_split_to_table(paragraph, E'[.!?]') AS sentence,
		link
	FROM (
		SELECT 
			COALESCE(substring(
			    cp.pagesource
			    FROM
			    '<p><p>(.*?)</p></p>'
			),'') AS paragraph,
			link
		FROM 
			webintelligence.crawlpage cp
		WHERE task ='JATIMPROV'	
		AND link LIKE 'https://jatimprov%'
		AND link LIKE '%/berita/%'
		AND cp.pagesource like '%%'
--		ORDER BY cp.created DESC
	)	
)
GROUP BY link
ORDER BY total desc
--ORDER BY total ASC 
;

SELECT 
--count(1) AS total
count(sentence) AS total
FROM (
	SELECT 
		regexp_split_to_table(paragraph, E'[.!?]') AS sentence,
		link
	FROM (
		SELECT 
			COALESCE(substring(
			    cp.pagesource
			    FROM
			    '<p><p>(.*?)</p></p>'
			),'') AS paragraph,
			link
		FROM 
			webintelligence.crawlpage cp
		WHERE task ='JATIMPROV'	
		AND link LIKE 'https://jatimprov%'
		AND link LIKE '%/berita/%'
		AND cp.pagesource like '%%'
--		ORDER BY cp.created DESC
	)	
)
--GROUP BY link
ORDER BY total desc


;

SELECT 
COALESCE(cp.link,''),
COALESCE(cp.pagesource,''),
COALESCE(substring(
        cp.pagesource
        FROM
        '<meta property="og:title" content="(.*?)">'
),'') AS document_metatitle,
COALESCE(substring(
    cp.pagesource
    FROM
    '<meta property="og:description" content="(.*?)">'
),'') AS document_metacontent,
COALESCE(substring(
    cp.pagesource
    FROM
    '<div class="pr-bg pr-bg-white"></div>
<h3>(.*?)</h3>'
),'') AS document_title,
COALESCE(substring(
    cp.pagesource
    FROM
    '<div class="parallax-header"> <a href="#">(.*?)</a>'
),'') AS document_date,
COALESCE(substring(
    cp.pagesource
    FROM
    '<span>Kategori : </span><a href="#">(.*?)</a> </div>'
),'') AS document_category,
COALESCE(substring(
    cp.pagesource
    FROM
    '<li><span><i class="fal fa-eye"></i>(.*?)</span></li>'
),'') AS document_totalview,
COALESCE(substring(
    cp.pagesource
    FROM
    '<li><span><i class="fal fa-hashtag"></i>(.*?)</span></li>'
),'') AS document_hashtag,
COALESCE(substring(
    cp.pagesource
    FROM
    '<p><p>(.*?)</p></p>'
),'') AS document_content,
COALESCE(substring(
    cp.pagesource
    FROM
    '<span>Berita Terkait</span>(.*?)</div>
</div>
</div>
</div>
</div>
</div>'
),'') AS document_relatednews
FROM webintelligence.crawlpage cp
WHERE task ='JATIMPROV'	
AND link LIKE 'https://jatimprov%'
AND link LIKE '%/berita/%'
AND cp.pagesource like '%%'
ORDER BY cp.created DESC
;

SELECT 
	count(1) AS to
FROM webintelligence.crawlpage cp
WHERE task ='JATIMPROV'
--AND DOCUMENTtype IS null
AND link LIKE 'https://jatimprov%'
AND link LIKE '%/berita/%';

;
SELECT 
	substring(
        cp.pagesource
        FROM
        '<span>Kategori : </span><a href="#">(.*?)</a> </div>'
    ) AS document_category
FROM webintelligence.crawlpage cp
WHERE task ='JATIMPROV'
--AND DOCUMENTtype IS null
AND link LIKE 'https://jatimprov%'
AND link LIKE '%/berita/%'
GROUP BY document_category

;
SELECT 
task,
sum(1) AS total
FROM webintelligence.tabled
GROUP BY task 
UNION ALL 
SELECT 
cp.task,
sum(1) AS total
FROM webintelligence.crawlpage cp
GROUP BY task 
;
 
AYOSEHAT
PORTKESMAS
SEHAT NEGERIKU
;

SELECT 
	
FROM
webintelligence.crawlpage cp
WHERE 
cp.task = 'JATIMPROV'
--AND cp.link LIKE '%page=%'
ORDER BY cp.created DESC;

;

SELECT 
	d.u,d.du,
	substring(
        d.du
        FROM
        '<div class="content-inner">(.*?)</div>'
    ) AS document_content
FROM
	webintelligence.tabled d
WHERE 
	task = 'SEHAT NEGERIKU'
	AND d.u LIKE '%/baca/%'
	AND d.u NOT LIKE '%comment%'
	AND d.du NOT LIKE '<?xml%'
ORDER BY d.created DESC 

--
;

SELECT 
	cp.documenttype,
	cp.mimetype,
	sum(1) AS total_document,
	sum(length(cp.pagesource)) AS total_character
FROM
	webintelligence.crawlpage cp
WHERE 
	task = 'ALODOKTER'
GROUP BY cp.documenttype,cp.mimetype 
ORDER BY documenttype ASC 
--ORDER BY cp.created DESC 

;

SELECT
	cp.link,
	cp.pagesource,	
	cp.task,
	cp.documenttype,
	cp.mimetype
FROM 
	webintelligence.crawlpage cp
WHERE
	cp.task = 'ALODOKTER'
--	AND lower(link) LIKE '%zell%';
--	AND lower(cp.pagesource) LIKE '%anemia%';
--	AND cp.link = 'https://www.alodokter.com/anemia'
ORDER BY created DESC 

;

SELECT count(1) AS total FROM webintelligence.crawlpage cp;
SELECT count(1) AS total FROM webintelligence.crawlhref ch;

;

SELECT 
	ch.*
FROM
	webintelligence.crawlhref ch
WHERE 
	1=1
	AND task = 'ALODOKTER'
ORDER BY ch.created DESC

;

-- menghitung text alo dokter

SELECT 
--count(1) AS total
count(sentence) AS total
FROM (
	SELECT 
		regexp_split_to_table(paragraph, E'[.!?]') AS sentence,
		link
	FROM (
		SELECT 
			COALESCE(substring(
				    cp.pagesource
				    FROM
				    '<html><body>(.*?)</body></html>'
				),'') AS paragraph,
			link
		FROM 
			webintelligence.crawlpage cp
		WHERE 
		task ='ALODOKTER'	
--		AND link LIKE 'https://jatimprov%'
--		AND link LIKE '%/berita/%'
		AND cp.pagesource like '%%'
--		ORDER BY cp.created DESC
	)	
)
--GROUP BY link
ORDER BY total desc

--

SELECT 
	link,
	COALESCE(substring(
	    cp.pagesource
	    FROM
	    '<html><body>(.*?)</body></html>'
	),'') AS paragraph
FROM 
	webintelligence.crawlpage cp
WHERE 
	task ='ALODOKTER'	
--	AND link LIKE 'https://jatimprov%'
--	AND link LIKE '%/berita/%'
ORDER BY cp.created DESC


--

SELECT 
		regexp_split_to_table(paragraph, E'[.!?]') AS sentence,
		link
	FROM (
		SELECT 
			COALESCE(substring(
				    cp.pagesource
				    FROM
				    '<html><body>(.*?)</body></html>'
				),'') AS paragraph,
			link
		FROM 
			webintelligence.crawlpage cp
		WHERE 
		task ='ALODOKTER'	
--		AND link LIKE 'https://jatimprov%'
--		AND link LIKE '%/berita/%'
		AND cp.pagesource like '%%'
--		ORDER BY cp.created DESC
	)	

;

SELECT 
cp.*
FROM
	webintelligence.crawlpage cp 
WHERE 
1=1
AND cp.documenttype = 'image'


;

SELECT * FROM webintelligence.backup_tabled 
WHERE lower(du) LIKE '%lyric%'
ORDER BY created ASC

;

SELECT 
count(1) AS total 
FROM webintelligence.crawlpage cp
WHERE 
	cp.task = 'ALODOKTER';

;

SELECT 
count(1) AS total 
FROM webintelligence.crawlpage cp
WHERE 
	cp.task = 'ALODOKTER';
















