import {Component} from '@angular/core';
import {MatTabChangeEvent} from '@angular/material';
import {Router} from '@angular/router';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss'],
})
export class AppComponent {
  title = 'prebuilt';

  constructor(
    private router: Router
  ) {
  }

  onTabClick($event: MatTabChangeEvent) {
    console.log($event);
    if ($event.index === 0) {
      return this.router.navigateByUrl('/code');
    }else {
      return this.router.navigateByUrl('/home');
    }
  }
}
