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
  public activities: any;

  constructor(size: SizeService, data: DataService) {
    this.data = data;
    this.size = size;
  }

  ngOnInit() {
    this.data.getActivityList('user').subscribe((params: any) => {
      console.log(params);
      this.activities = params;
    }, error => {
      console.log(error);
      alert(error.error.message);
    });
  }

}
