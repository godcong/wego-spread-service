import {Component, OnInit} from '@angular/core';

declare var window: Window;

@Component({
  selector: 'app-my',
  templateUrl: './my.component.html',
  styleUrls: ['./my.component.scss']
})
export class MyComponent implements OnInit {
  private height: number;

  constructor() {
    this.height = window.innerHeight;

  }

  ngOnInit() {
  }

}
