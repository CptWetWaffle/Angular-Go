import {Component, OnInit} from '@angular/core';
import {HttpClient} from "@angular/common/http";
import {Task} from "protractor/built/taskScheduler";
import {environment} from "../../environments/environment.prod";

interface INewsfeedItem {
  title: string,
  post: string
}

@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.css']
})
export class HomeComponent implements OnInit {
  public title = '';
  public post = '';
  public newsfeedItems: INewsfeedItem[] = [];

  constructor(private httpClient: HttpClient) {

  }

  async ngOnInit() {
    await this.loadNewsItems();
  }

  async loadNewsItems() {
    this.newsfeedItems = await this.httpClient.get<INewsfeedItem[]>(environment.gateway + '/api/newsfeed').toPromise();
  }

  async addPost() {
    await this.httpClient.post(environment.gateway + '/api/newsfeed', {
      title: this.title,
      post: this.post
    }).toPromise();

    await this.loadNewsItems();

    this.title = '';
    this.post = '';
  }

}
