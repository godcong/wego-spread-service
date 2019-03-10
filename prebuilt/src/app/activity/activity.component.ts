import {Component, OnInit} from '@angular/core';
import {SizeService} from '../size.service';
import {DataService} from '../data.service';
import {ActivatedRoute, ParamMap} from '@angular/router';

@Component({
  selector: 'app-activity',
  templateUrl: './activity.component.html',
  styleUrls: ['./activity.component.scss']
})
export class ActivityComponent implements OnInit {
  public size: SizeService;
  private data: DataService;
  private router: ActivatedRoute;
  private id: string;
  private activity: ParamMap;

  constructor(size: SizeService, data: DataService, router: ActivatedRoute) {
    this.size = size;
    this.data = data;
    this.router = router;
  }

  ngOnInit() {
    this.router.paramMap.subscribe((params: ParamMap) => {
      console.log(params);
      this.id = params.get('id');
    });
    this.data.getActivityInfo(this.id).subscribe((params: ParamMap) => {
      console.log(params);
      this.activity = params;
    });
  }

}
