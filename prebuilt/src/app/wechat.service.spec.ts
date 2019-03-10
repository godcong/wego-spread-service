import { TestBed } from '@angular/core/testing';

import { WechatService } from './wechat.service';

describe('WechatService', () => {
  beforeEach(() => TestBed.configureTestingModule({}));

  it('should be created', () => {
    const service: WechatService = TestBed.get(WechatService);
    expect(service).toBeTruthy();
  });
});
