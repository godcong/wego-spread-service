import {Component, OnInit} from '@angular/core';
import {SizeService} from '../size.service';
import {ActivatedRoute, Params} from '@angular/router';

@Component({
  selector: 'app-qrcode',
  templateUrl: './qrcode.component.html',
  styleUrls: ['./qrcode.component.scss']
})
export class QrcodeComponent implements OnInit {
  public size: SizeService;
  public qrcode: string;
  private router: ActivatedRoute;

  constructor(size: SizeService, router: ActivatedRoute) {
    this.size = size;
    this.router = router;
    this.qrcode = 'http://localhost:8080';
  }

  ngOnInit() {
    this.router.queryParamMap.subscribe((v: Params) => {
      console.log(v);
    });
  }

  onSuccess() {
    console.log(this.qrcode);
  }
}
