import {Component, OnInit} from '@angular/core';
import {SizeService} from '../size.service';

@Component({
  selector: 'app-qrcode',
  templateUrl: './qrcode.component.html',
  styleUrls: ['./qrcode.component.scss']
})
export class QrcodeComponent implements OnInit {
  public size: SizeService;
  public qrcode: string;

  constructor(size: SizeService) {
    this.size = size;
    this.qrcode = 'http://localhost:8080';
  }

  ngOnInit() {
  }

}
