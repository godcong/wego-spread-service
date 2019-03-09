import {Component, OnInit} from '@angular/core';
import {SizeService} from '../size.service';

@Component({
  selector: 'app-activity',
  templateUrl: './activity.component.html',
  styleUrls: ['./activity.component.scss']
})
export class ActivityComponent implements OnInit {
  public size: SizeService;

  constructor(size: SizeService) {
    this.size = size;
  }

  ngOnInit() {
  }

}
