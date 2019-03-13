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
  private activity: any;

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
    this.data.getActivityInfo(this.id).subscribe((params: any) => {
      console.log(params);
      this.activity = params;
    });
  }

  joinActivity(code: string) {
    console.log(code);
    this.data.postActivityJoin(code).subscribe((params: any) => {
      console.log(params);
      alert('加入申请成功');
    }, (error: any) => {
      console.log(error);
      alert(error.error.message);
    });
  }

}
