import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { ShareCodeComponent } from './share-code.component';

describe('ShareCodeComponent', () => {
  let component: ShareCodeComponent;
  let fixture: ComponentFixture<ShareCodeComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ ShareCodeComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(ShareCodeComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
