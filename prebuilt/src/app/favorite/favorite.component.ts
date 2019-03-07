import {Component, OnInit} from '@angular/core';
import {DataService} from '../data.service';
import {SizeService} from '../size.service';

@Component({
  selector: 'app-favorite',
  templateUrl: './favorite.component.html',
  styleUrls: ['./favorite.component.scss']
})
export class FavoriteComponent implements OnInit {
  private height: number;
  private data: DataService;
  private userActivities: any;
  private size: SizeService;

  constructor(data: DataService, size: SizeService) {
    this.data = data;
    this.size = size;
    this.height = window.innerHeight;
  }

  ngOnInit() {
    this.data.getUserActivityList().subscribe((ret) => {
      console.log(ret);
      this.userActivities = ret;
    }, error => {
      console.log(error);
    });
  }

}
