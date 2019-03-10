import {Component, OnInit} from '@angular/core';
import {SizeService} from '../size.service';
import {ActivatedRoute, ParamMap} from '@angular/router';
import {DataService} from '../data.service';

@Component({
  selector: 'app-qrcode',
  templateUrl: './qrcode.component.html',
  styleUrls: ['./qrcode.component.scss']
})
export class QrcodeComponent implements OnInit {
  public size: SizeService;
  public qrcode: string;
  private router: ActivatedRoute;
  private id: string;
  private user: string;
  private data: DataService;

  constructor(size: SizeService, router: ActivatedRoute, data: DataService) {
    this.size = size;
    this.router = router;
    this.data = data;
    this.qrcode = 'http://localhost:8080';
  }

  ngOnInit() {
    this.router.queryParamMap.subscribe((params: ParamMap) => {
      console.log(params);
      this.user = params.get('user');
    });
    this.router.paramMap.subscribe((params: ParamMap) => {
      console.log(params);
      this.id = params.get('id');
    });
    this.data.getSpreadShareInfo(this.id, this.user).subscribe((params: ParamMap) => {
      console.log(params);
      this.qrcode = params['url'];
      this.data.wxInit(params);
    });
  }

  onSuccess() {
    console.log(this.qrcode);
  }
}
