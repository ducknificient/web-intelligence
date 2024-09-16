;

SELECT 
	count(1) AS total
FROM 
	webintelligence.tabled t

;

SELECT 
	t.du, t.u, t.created 
FROM
	webintelligence.tabled t
ORDER BY t.created DESC
		
;

SELECT 
	t.du,LENGTH(t.du) AS sizee, t.u, t.created 
FROM
	webintelligence.tabled t
ORDER BY sizee DESC

;

SELECT 
	count(1) AS total
FROM 
	webintelligence.tablee t

;

SELECT 
	t.u, t.v 
FROM
	webintelligence.tablee t
WHERE 
	1=1
AND v like '%https://www.azlyrics.coma/a1.html%'
ORDER BY 
	t.created DESC
		
;

SELECT
	t.u,
	count(t.u) AS total
FROM
	webintelligence.tablee t
WHERE 
	1=1
GROUP BY t.u
ORDER BY total DESC

;

SELECT 
	t.du,
	t.u
FROM 
	webintelligence.tabled t 
WHERE 
	t.u = 'https://www.azlyrics.com/a.html'
	
;

;
SELECT t.du,t.u FROM webintelligence.tabled t WHERE u = 'https://www.azlyrics.com/a.html'; 

SELECT * FROM webintelligence.backup_tabled t ;

;
--DELETE FROM webintelligence.backup_tabled;
--DELETE FROM webintelligence.backup_tablee;
;

INSERT INTO webintelligence.backup_tabled (du,u,created,createdby,updated,updatedby,flag) 
(
SELECT du,u,created,createdby,updated,updatedby,flag FROM webintelligence.tabled
)

;

INSERT INTO webintelligence.backup_tablee (u,v,created,createdby,updated,updatedby,flag) 
(
SELECT u,v,created,createdby,updated,updatedby,flag FROM webintelligence.tablee 
)

;

SELECT 
	count(1) AS total
--	t.*
FROM webintelligence.tabled t
WHERE 
1=1
AND t.created < to_timestamp('2024-02-15 21:19:39','YYYY-MM-DD hh24:mi:ss')
--ORDER BY t.created ASC

;
	
SELECT 
	count(1) AS total
--	t.*
FROM webintelligence.tabled t
WHERE 
1=1
--AND t.u LIKE '%stack%'
AND t.created < to_timestamp('2024-02-15 21:19:39','YYYY-MM-DD hh24:mi:ss')

;

SELECT 
	count(1) AS total
--	t.*
FROM 
	webintelligence.tablee t
WHERE 
1=1
AND t.created < to_timestamp('2024-02-15 21:19:39','YYYY-MM-DD hh24:mi:ss')
--ORDER BY t.created ASC

;

create table webintelligence.crawlpage (
    pagesource text null,
    link text null,
    created timestamp without time zone default now() not null,
    createdby text default 'POSTGRES',
    updated timestamp without time zone default now() not null,
    updatedby text default 'POSTGRES',
    flag types.flag default 'Insert'::types.flag
);

create table webintelligence.crawlhref (
    link text null,
    href text null,
    created timestamp without time zone default now() not null,
    createdby text default 'POSTGRES',
    updated timestamp without time zone default now() not null,
    updatedby text default 'POSTGRES',
    flag types.flag default 'Insert'::types.flag
);

;

INSERT INTO webintelligence.crawlpage (pagesource,link,created,createdby,updated,updatedby,flag) 
(
SELECT du,u,created,createdby,updated,updatedby,flag FROM webintelligence.tabled t
WHERE t.created > to_timestamp('2024-02-15 21:19:39','YYYY-MM-DD hh24:mi:ss')
)

;

INSERT INTO webintelligence.crawlhref (link,href,created,createdby,updated,updatedby,flag) 
(
SELECT u,v,created,createdby,updated,updatedby,flag FROM webintelligence.tablee t
WHERE t.created > to_timestamp('2024-02-15 21:19:39','YYYY-MM-DD hh24:mi:ss')
)
;

SELECT 
	t.*
FROM webintelligence.crawlpage t

;

SELECT 
	t.*
FROM webintelligence.crawlhref t

;
