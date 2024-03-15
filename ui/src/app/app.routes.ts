import { Routes } from '@angular/router';
import { CrawlerComponent } from './pages/crawler/crawler.component';
import { TextClassificationComponent } from './pages/text-classification/text-classification.component';

export const routes: Routes = [
    {path: 'crawler', component: CrawlerComponent},
    {path: 'text-classification', component: TextClassificationComponent},
];
