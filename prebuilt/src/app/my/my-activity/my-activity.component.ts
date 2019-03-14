import {Component, OnInit} from '@angular/core';
import {DataService} from '../../data.service';
import {ParamMap} from '@angular/router';
import {SizeService} from '../../size.service';

@Component({
  selector: 'app-my-activity',
  templateUrl: './my-activity.component.html',
  styleUrls: ['./my-activity.component.scss']
})
export class MyActivityComponent implements OnInit {
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
