import { Component, Input } from '@angular/core';
import { Sidebar } from '../../entity/sidebar';
import { CommonModule } from '@angular/common';
import { SidebarLinkComponent } from './sidebar-link/sidebar-link.component';
import { SidebarLink } from '../../entity/sidebar-link';
import { SidebarDropdownlinkComponent } from './sidebar-dropdownlink/sidebar-dropdownlink.component';

@Component({
  selector: 'app-sidebar',
  standalone: true,
  imports: [
    SidebarLinkComponent,
    SidebarDropdownlinkComponent,
    CommonModule,
  ],
  templateUrl: './sidebar.component.html',
  styleUrls: ['./sidebar.component.scss','./../../../assets/css/spil-theme.scss'],
})
export class SidebarComponent {

  @Input() sidebar!: Sidebar;

  test_link: SidebarLink[] = [{
    title:"Test title",
    href:"/test_href",
    icon:"bi-alexa",
  }]
}
