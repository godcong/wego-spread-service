import {Injectable} from '@angular/core';
import {ParamMap} from '@angular/router';

declare let wx: any;

@Injectable({
  providedIn: 'root'
})
export class WechatService {

  constructor() {
  }

  init(params: any) {
    wx.config({
      debug: true, // 开启调试模式,调用的所有api的返回值会在客户端alert出来，若要查看传入的参数，可以在pc端打开，参数信息会通过log打出，仅在pc端时才会打印。
      appId: params.appID, // 必填，公众号的唯一标识
      timestamp: params.timestamp, // 必填，生成签名的时间戳
      nonceStr: params.nonceStr, // 必填，生成签名的随机串
      signature: params.signature, // 必填，签名
      jsApiList: ['onMenuShareTimeline', 'onMenuShareAppMessage', 'onMenuShareQQ', 'onMenuShareQZone'] // 必填，需要使用的JS接口列表
    });
  }

  shareAppMessage(url: string, img: string) {
    wx.onMenuShareAppMessage({
      title: '测试1', // 分享标题
      desc: '这个是测试页面', // 分享描述
      link: url, // 分享链接
      imgUrl: img, // 分享图标
      // type: '', // 分享类型,music、video或link，不填默认为link
      // dataUrl: '', // 如果type是music或video，则要提供数据链接，默认为空
      success() {
        // 用户确认分享后执行的回调函数
        alert('分享成功');
      },
      cancel() {
        // 用户取消分享后执行的回调函数
        alert('分享取消');
      }
    });
  }

}
