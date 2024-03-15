import { SidebarDropdownlink } from "./sidebar-dropdownlink";
import { SidebarLink } from "./sidebar-link";

export interface Sidebar {
    appname: string,
    applogo: string,
    profilepicture: string,
    sidebarLinkList: SidebarLink[],
    dropdownLinkList: SidebarDropdownlink[],
    username: string,
}
