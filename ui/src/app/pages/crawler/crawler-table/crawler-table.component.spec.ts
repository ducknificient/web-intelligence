import { ComponentFixture, TestBed } from '@angular/core/testing';

import { CrawlerTableComponent } from './crawler-table.component';

describe('CrawlerTableComponent', () => {
  let component: CrawlerTableComponent;
  let fixture: ComponentFixture<CrawlerTableComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [CrawlerTableComponent]
    })
    .compileComponents();
    
    fixture = TestBed.createComponent(CrawlerTableComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
