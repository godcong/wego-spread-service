import { Component, OnInit } from '@angular/core';
declare var window: Window;
@Component({
  selector: 'app-favorite',
  templateUrl: './favorite.component.html',
  styleUrls: ['./favorite.component.scss']
})
export class FavoriteComponent implements OnInit {
  private height: number;

  constructor() {
    this.height = window.innerHeight;
  }

  ngOnInit() {
  }

}
