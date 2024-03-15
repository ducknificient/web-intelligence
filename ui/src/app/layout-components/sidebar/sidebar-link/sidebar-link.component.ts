import { Component, Input } from '@angular/core';
import { SidebarLink } from '../../../entity/sidebar-link';
import { CommonModule } from '@angular/common';

@Component({
  selector: 'app-sidebar-link',
  standalone: true,
  imports: [CommonModule],
  templateUrl: './sidebar-link.component.html',
  styleUrl: './sidebar-link.component.scss'
})
export class SidebarLinkComponent {

  @Input() sidebarLink!: SidebarLink;
}
