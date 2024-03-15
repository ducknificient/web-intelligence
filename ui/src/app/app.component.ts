import { Component } from '@angular/core';
import { CommonModule } from '@angular/common';
import { RouterOutlet } from '@angular/router';
import { SidebarComponent } from './layout-components/sidebar/sidebar.component';
import { FooterComponent } from './layout-components/footer/footer.component';
import { Sidebar } from './entity/sidebar';
import { ContentComponent } from './layout-components/content/content.component';

@Component({
  selector: 'app-root',
  standalone: true,
  imports: [CommonModule, RouterOutlet, SidebarComponent, FooterComponent, ContentComponent],
  templateUrl: './app.component.html',
  styleUrl: './app.component.scss'
})
export class AppComponent {
  title = 'Web Intelligence';

  sidebar: Sidebar = {
    appname: "Web Intelligence",
    applogo:"../assets/logo/sim_logo.png",
    profilepicture:"../assets/logo/sim_logo.png",
    sidebarLinkList:[
      {
        title:"Crawler",
        href:"/crawler",
        icon:"bi-alexa"
      },
      {
        title:"Text Classification",
        href:"/text-classification",
        icon:"bi-alexa"
      }
    ],
    dropdownLinkList:[
      {
        title:"Settings",
        href:"",
      },
      {
        title:"Profile",
        href:"",
      },
      {
        title:"Log Out",
        href:"",
      }
    ],
    username:"Jeremy"
  }

  calculateContentHeight(): number {
    // Calculate the content height as 90% of the viewport height
    const viewportHeight = window.innerHeight;
    const contentHeightPercentage = 85; // Adjust as needed
    return (viewportHeight * contentHeightPercentage) / 100
  }
  
}
