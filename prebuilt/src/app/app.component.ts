import {Component, OnInit} from '@angular/core';
import {MatTabChangeEvent} from '@angular/material';
import {ActivatedRoute, Router} from '@angular/router';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss'],
})
export class AppComponent {
  private router: ActivatedRoute;

  constructor(router: ActivatedRoute) {
    console.log('appInit');
    this.router = router;

  }

  onTabClick($event: MatTabChangeEvent) {
    console.log($event);
    if ($event.index === 0) {
      // this.router.navigateByUrl('/code');
    } else {
      // this.router.navigateByUrl('/home');
    }
  }


  getURL() {
    this.router.queryParamMap.subscribe(val => {
      // if ((val !== null) || (val['token'] !== null)) {
      // }
      console.log(val);
    });
  }

}
