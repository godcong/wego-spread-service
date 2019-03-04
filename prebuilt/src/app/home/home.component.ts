import {Component, OnChanges, OnInit, SimpleChanges} from '@angular/core';
import {HomeDataService} from './home-data.service';

declare var window: Window;

@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.scss']
})
export class HomeComponent implements OnInit, OnChanges {
  private links: string[] = ['0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '9', '9', '9', '9', '9', '9', '9', '9', '9', '9'];
  private data: HomeDataService;
  public activities: any;
  public height: number;

  constructor(data: HomeDataService) {
    this.data = data;
    console.log('constructor');
    this.height = window.innerHeight;
    this.activities = this.data.getActivityList();
  }

  ngOnInit() {
    console.log('ngOnInit');
    // this.height = window.innerHeight;
    // console.log(window.innerWidth);
    // this.activities = this.data.getActivityList();
  }

  ngOnChanges(changes: SimpleChanges): void {
    console.log('ngOnChanges');
    // this.activities = this.data.getActivityList();
  }

}
