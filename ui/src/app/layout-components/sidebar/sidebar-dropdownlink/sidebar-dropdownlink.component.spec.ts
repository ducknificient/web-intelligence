import { ComponentFixture, TestBed } from '@angular/core/testing';

import { SidebarDropdownlinkComponent } from './sidebar-dropdownlink.component';

describe('SidebarDropdownlinkComponent', () => {
  let component: SidebarDropdownlinkComponent;
  let fixture: ComponentFixture<SidebarDropdownlinkComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [SidebarDropdownlinkComponent]
    })
    .compileComponents();
    
    fixture = TestBed.createComponent(SidebarDropdownlinkComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
