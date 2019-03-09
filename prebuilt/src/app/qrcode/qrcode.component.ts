import {Component, OnInit} from '@angular/core';
import {SizeService} from '../size.service';
import {ActivatedRoute, ParamMap} from '@angular/router';

@Component({
  selector: 'app-qrcode',
  templateUrl: './qrcode.component.html',
  styleUrls: ['./qrcode.component.scss']
})
export class QrcodeComponent implements OnInit {
  public size: SizeService;
  public qrcode: string;
  private router: ActivatedRoute;
  private params: ParamMap;

  constructor(size: SizeService, router: ActivatedRoute) {
    this.size = size;
    this.router = router;
    this.qrcode = 'http://localhost:8080';
  }

  ngOnInit() {
    this.router.queryParamMap.subscribe((params: ParamMap) => {
      console.log(params);
      this.params = params;
    });
  }

  onSuccess() {
    console.log(this.qrcode);
  }
}
