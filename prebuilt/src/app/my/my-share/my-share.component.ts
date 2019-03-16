import {Component, OnInit} from '@angular/core';
import {SizeService} from '../../size.service';
import {DataService} from '../../data.service';

@Component({
  selector: 'app-my-share',
  templateUrl: './my-share.component.html',
  styleUrls: ['./my-share.component.scss']
})
export class MyShareComponent implements OnInit {
  public size: SizeService;
  private data: DataService;
  public userActivities: any;

  constructor(size: SizeService, data: DataService) {
    this.data = data;
    this.size = size;
  }

  ngOnInit() {
    this.data.getUserActivityList(false).subscribe((params: any) => {
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
