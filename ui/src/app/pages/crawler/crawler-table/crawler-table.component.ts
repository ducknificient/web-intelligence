import { CommonModule } from '@angular/common';
// import { Component, inject } from '@angular/core';
import { Crawlpage } from '../../../entity/crawlpage';
import { CrawlService } from '../../../service/crawl.service';
import {AfterViewInit, Component, ViewChild, inject} from '@angular/core';
import {MatPaginator, MatPaginatorModule} from '@angular/material/paginator';
import {MatSort, MatSortModule} from '@angular/material/sort';
import {MatTableDataSource, MatTableModule} from '@angular/material/table';
import {MatInputModule} from '@angular/material/input';
import {MatFormFieldModule} from '@angular/material/form-field';
@Component({
  selector: 'app-crawler-table',
  standalone: true,
  imports: [CommonModule,MatFormFieldModule, MatInputModule, MatTableModule, MatSortModule, MatPaginatorModule],
  templateUrl: './crawler-table.component.html',
  styleUrl: './crawler-table.component.scss'
})
export class CrawlerTableComponent implements AfterViewInit {

  crawlpageList: Crawlpage[] = []
  crawlService: CrawlService = inject(CrawlService);

  displayedColumns: string[] = ['no', 'link', 'pagesource'];
  dataSource: MatTableDataSource<Crawlpage>;

  @ViewChild(MatSort) sort: MatSort = new MatSort();
  // @ViewChild(MatPaginator) paginator: MatPaginator = new MatPaginator(new MatPaginatorIntl(), ChangeDetectorRef.prototype);
  @ViewChild(MatPaginator) paginator: MatPaginator | null = null;



  constructor() {
    this.crawlpageList = this.crawlService.getAllCrawlpage();

    // Create 100 users
    // const users = Array.from({length: 100}, (_, k) => createNewUser(k + 1));

    // Assign the data to the data source for the table to render
    this.dataSource = new MatTableDataSource(this.crawlpageList);
    console.log(this.dataSource)

  }

  ngAfterViewInit() {
    this.dataSource.paginator = this.paginator;
    this.dataSource.sort = this.sort;
  }

  applyFilter(event: Event) {
    const filterValue = (event.target as HTMLInputElement).value;
    this.dataSource.filter = filterValue.trim().toLowerCase();

    if (this.dataSource.paginator) {
      this.dataSource.paginator.firstPage();
    }
  }

/** Builds and returns a new User. */
// function createNewUser(id: number): UserData {
//   const name =
//     NAMES[Math.round(Math.random() * (NAMES.length - 1))] +
//     ' ' +
//     NAMES[Math.round(Math.random() * (NAMES.length - 1))].charAt(0) +
//     '.';

//   return {
//     id: id.toString(),
//     name: name,
//     progress: Math.round(Math.random() * 100).toString(),
//     fruit: FRUITS[Math.round(Math.random() * (FRUITS.length - 1))],
//   };
// }

}
