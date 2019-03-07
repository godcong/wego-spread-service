import {Component, OnChanges, OnInit, SimpleChanges} from '@angular/core';
import {ActivatedRoute, ParamMap} from '@angular/router';
import {WebTokenService} from '../web-token.service';
import {SizeService} from '../size.service';
import {DataService} from '../data.service';

@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.scss']
})
export class HomeComponent implements OnInit, OnChanges {
  size: SizeService;
  private data: DataService;
  public activities: any;
  private router: ActivatedRoute;

  constructor(data: DataService, router: ActivatedRoute, size: SizeService) {
    this.data = data;
    this.router = router;
    this.size = size;
    console.log('constructor');
  }

  ngOnInit() {
    console.log('ngOnInit');
    this.router.queryParamMap.subscribe((params: ParamMap) => {
      if (params.has('token')) {
        WebTokenService.setToken(params.get('token'));
      }
    });
    this.data.getActivityList().subscribe((ret: any) => {
      this.activities = ret;
    }, error => {
      console.log(error);
    });
  }

  ngOnChanges(changes: SimpleChanges): void {
    console.log('ngOnChanges');
  }
}
