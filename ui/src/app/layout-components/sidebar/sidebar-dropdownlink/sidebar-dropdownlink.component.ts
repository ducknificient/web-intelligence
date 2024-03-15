import { Component, Input } from '@angular/core';
import { SidebarDropdownlink } from '../../../entity/sidebar-dropdownlink';

@Component({
  selector: 'app-sidebar-dropdownlink',
  standalone: true,
  imports: [],
  templateUrl: './sidebar-dropdownlink.component.html',
  styleUrl: './sidebar-dropdownlink.component.scss'
})
export class SidebarDropdownlinkComponent {
  @Input() dropdownLink!: SidebarDropdownlink
}
