import {Component, OnInit} from '@angular/core';
import {SizeService} from '../size.service';

@Component({
  selector: 'app-my',
  templateUrl: './my.component.html',
  styleUrls: ['./my.component.scss']
})
export class MyComponent implements OnInit {
  size: SizeService;

  constructor(size: SizeService) {
    this.size = size;

  }

  ngOnInit() {
  }

}
