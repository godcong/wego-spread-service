import { TestBed } from '@angular/core/testing';

import { WebTokenService } from './web-token.service';

describe('WebTokenService', () => {
  beforeEach(() => TestBed.configureTestingModule({}));

  it('should be created', () => {
    const service: WebTokenService = TestBed.get(WebTokenService);
    expect(service).toBeTruthy();
  });
});
