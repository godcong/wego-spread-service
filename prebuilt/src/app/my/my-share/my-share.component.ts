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
    this.data.getUserActivityList(false).subscribe((ret) => {
      console.log(ret);
      this.userActivities = ret;
    }, error => {
      console.log(error);
    });
  }
}
