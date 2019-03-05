import {Component, OnInit} from '@angular/core';
import {MatTabChangeEvent} from '@angular/material';
import {ActivatedRoute, Event, NavigationEnd, NavigationError, NavigationStart, Router, RoutesRecognized} from '@angular/router';
import {Location} from '@angular/common';
import {HttpParams} from '@angular/common/http';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss'],
})
export class AppComponent implements OnInit {
  constructor() {
    console.log('appInit');
  }

  onTabClick($event: MatTabChangeEvent) {
    console.log($event);
    if ($event.index === 0) {

    } else {

    }
  }


  ngOnInit(): void {
    console.log('ngOnInit');
  }

}
