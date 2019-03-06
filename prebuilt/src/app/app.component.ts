import {Component, OnInit} from '@angular/core';
import {ActivatedRoute, ParamMap, Router} from '@angular/router';
import {WebTokenService} from './web-token.service';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss'],
})
export class AppComponent implements OnInit {
  private router: ActivatedRoute;

  constructor(router: ActivatedRoute) {
    this.router = router;
    console.log('appInit');
  }

  ngOnInit(): void {
    this.router.queryParamMap.subscribe((params: ParamMap) => {
      console.log(params);
    });
  }

}
