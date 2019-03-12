import {Component, OnInit} from '@angular/core';
import {SizeService} from '../size.service';
import {DataService} from '../data.service';
import {ParamMap} from '@angular/router';

@Component({
  selector: 'app-my',
  templateUrl: './my.component.html',
  styleUrls: ['./my.component.scss']
})
export class MyComponent implements OnInit {
  public size: SizeService;
  private data: DataService;
  private info: ParamMap;

  constructor(size: SizeService, data: DataService) {
    this.size = size;
    this.data = data;
  }

  ngOnInit() {
    this.data.getMyInfo().subscribe((params: ParamMap) => {
      console.log(params);
      this.info = params;
    });
  }

}
