import {Component, OnInit} from '@angular/core';
import {SizeService} from '../../size.service';
import {DataService} from '../../data.service';
import {ParamMap} from '@angular/router';

@Component({
  selector: 'app-my-spread',
  templateUrl: './my-spread.component.html',
  styleUrls: ['./my-spread.component.scss']
})
export class MySpreadComponent implements OnInit {
  public size: SizeService;
  private data: DataService;
  public spreads: ParamMap;

  constructor(size: SizeService, data: DataService) {
    this.size = size;
    this.data = data;
  }

  ngOnInit() {
    this.data.getMySpread().subscribe((params: ParamMap) => {
      console.log(params);
      this.spreads = params;
    });
  }

}
