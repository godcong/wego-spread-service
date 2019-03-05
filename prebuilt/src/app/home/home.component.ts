import {Component, OnChanges, OnDestroy, OnInit, SimpleChanges} from '@angular/core';
import {HomeDataService} from './home-data.service';
import {ActivatedRoute, ParamMap, Params} from '@angular/router';
import {HttpParams} from '@angular/common/http';
import {WebTokenService} from '../web-token.service';
import {SizeService} from '../size.service';

@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.scss']
})
export class HomeComponent implements OnInit, OnChanges {
  private links: string[] = ['0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '9', '9', '9', '9', '9', '9', '9', '9', '9', '9'];
  size: SizeService;
  private data: HomeDataService;
  public activities: any;
  private router: ActivatedRoute;

  constructor(data: HomeDataService, router: ActivatedRoute, size: SizeService) {
    this.data = data;
    this.router = router;
    this.size = size;
    console.log('constructor');
    this.data.getActivityList().subscribe((ret: any) => {
      this.activities = ret;
    }, error => {
      console.log(error);
    });
  }

  ngOnInit() {
    console.log('ngOnInit');
    this.router.queryParamMap.subscribe((params: ParamMap) => {
      if (params.has('token')) {
        WebTokenService.setToken(params.get('token'));
      }
    });
  }

  ngOnChanges(changes: SimpleChanges): void {
    console.log('ngOnChanges');
  }
}
