import {Component, OnInit} from '@angular/core';
import {DataService} from '../../data.service';
import {SizeService} from '../../size.service';

@Component({
  selector: 'app-my-share',
  templateUrl: './my-share.component.html',
  styleUrls: ['./my-share.component.scss']
})
export class MyShareComponent implements OnInit {
  private data: DataService;
  private userActivities: any;
  public size: SizeService;

  constructor(data: DataService, size: SizeService) {
    this.data = data;
    this.size = size;
  }

  ngOnInit() {
    this.data.getUserActivityList(false).subscribe((params: any) => {
      console.log(params);
      this.userActivities = params;
    }, error => {
      console.log(error);
    });
  }

  joinFavorite(id: any) {
    this.data.postFavoriteJoin(id).subscribe((params: any) => {
      console.log(params);
      this.userActivities = params;
    }, error => {
      console.log(error);
    });
  }
}
