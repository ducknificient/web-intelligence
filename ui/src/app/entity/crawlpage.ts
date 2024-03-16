import { Crawlhref } from "./crawlhref";

export interface Crawlpage {
    id: string,
    pagesource: string,
    link: string,
    hreflist: Crawlhref[],
}
