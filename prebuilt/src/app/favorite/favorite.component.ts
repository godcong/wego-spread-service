import {Component, OnInit} from '@angular/core';
import {DataService} from '../data.service';

declare var window: Window;

@Component({
  selector: 'app-favorite',
  templateUrl: './favorite.component.html',
  styleUrls: ['./favorite.component.scss']
})
export class FavoriteComponent implements OnInit {
  private height: number;
  private data: DataService;
  private userActivities: any;

  constructor(data: DataService) {
    this.data = data;
    this.height = window.innerHeight;
  }

  ngOnInit() {
    this.data.getUserActivityList().subscribe((ret: any) => {
      this.userActivities = ret;
    }, error => {
      console.log(error);
    });
  }

}
