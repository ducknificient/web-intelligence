import { CommonModule } from '@angular/common';
import { Component, inject } from '@angular/core';
import { CrawlerTableComponent } from './crawler-table/crawler-table.component';

@Component({
  selector: 'app-crawler',
  standalone: true,
  imports: [CommonModule,CrawlerTableComponent],
  templateUrl: './crawler.component.html',
  styleUrl: './crawler.component.scss'
})
export class CrawlerComponent {

  

}
