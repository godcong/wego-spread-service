import {Component, OnInit} from '@angular/core';
import {DataService} from '../data.service';

declare var window: Window;

@Component({
  selector: 'app-favorite',
  templateUrl: './favorite.component.html',
  styleUrls: ['./favorite.component.scss']
})
export class FavoriteComponent implements OnInit {
  private height: number;
  private data: DataService;

  constructor(data: DataService) {
    this.data = data;
    this.height = window.innerHeight;
  }

  ngOnInit() {
  }

}
