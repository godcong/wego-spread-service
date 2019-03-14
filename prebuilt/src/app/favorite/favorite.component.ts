import {Component, OnInit} from '@angular/core';
import {DataService} from '../data.service';
import {SizeService} from '../size.service';

@Component({
  selector: 'app-favorite',
  templateUrl: './favorite.component.html',
  styleUrls: ['./favorite.component.scss']
})
export class FavoriteComponent implements OnInit {
  private data: DataService;
  private userActivities: any;
  public size: SizeService;

  constructor(data: DataService, size: SizeService) {
    this.data = data;
    this.size = size;
  }

  ngOnInit() {
    this.data.getUserActivityList(true).subscribe((params: any) => {
      console.log(params);
      this.userActivities = params;
    }, error => {
      console.log(error);
      alert(error.error.message);
    });
  }

  joinFavorite(id: any, status: boolean) {
    this.data.postFavoriteJoin(id, status).subscribe((params: any) => {
      console.log(params);
    }, error => {
      console.log(error);
      alert(error.error.message);
    });
  }
}
